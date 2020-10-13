package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tidwall/gjson"
)

const (
	SIG_EXIT        = 1
	PREFIX_KEYSPACE = "__keyspace@0__:"
	PREFIX_TAG      = "sysman:hardware:"
	RedisHost       = "localhost"
	RedisPort       = "6379"
)

type DBClient struct {
	wgCancel  sync.WaitGroup
	rdsDB     *redis.Client
	rdsPubSub *redis.PubSub
	apiToken  string
	workers   int
	l         *log.Logger
	ctx       context.Context
}

var ctx = context.Background()

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	val3, err := rdb.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("key3 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key3", val3)
	}
}

// DbClientNew : device database client.
func DbClientNew(origin string) *DBClient {
	var db DBClient
	db.l = log.New(os.Stdout, "["+origin+"]", log.Lshortfile)
	db.wgCancel = sync.WaitGroup{}
	db.rdsDB = redis.NewClient(&redis.Options{Addr: "127.0.0.1:" + RedisPort})
	db.ctx = db.rdsDB.Context()
	db.rdsPubSub = db.rdsDB.PSubscribe(db.ctx, PREFIX_KEYSPACE+PREFIX_TAG+"*")

	// connect db
	retry := 0
	retryMax := 3
	if os.Getenv("TP_DEVICE_DEBUG") == "1" {
		retryMax = 0
	}
	for retry < retryMax {
		if _, err := db.rdsPubSub.Receive(db.ctx); err != nil {
			db.l.Printf("Create pubsub receive failed(%v)\n", err)
			time.Sleep(5 * time.Second)
			retry++
			continue
		}
		db.l.Printf("Create pubsub receive success\n")
		return &db
	}
	db.l.Printf("Create pubsub receive exit because timeout.\n")
	db.rdsPubSub = nil
	return &db
}

func (self *DBClient) GetDBInfoAll() string {
	// call appman resource api
	url := "http://localhost:6379"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		self.l.Printf("Get device db failed[0], err: %v\n", err)
		return "[]"
	}
	res, err := client.Do(req)
	if err != nil {
		self.l.Printf("Get device db failed[1], err: %v\n", err)
		return "[]"
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		self.l.Printf("Get device db failed[2], err: %v\n", err)
		return "[]"
	}
	return gjson.Get(string(response), "data").String()
}

func (self *DBClient) GetDBInfo(item_type string, id string) string {
	// call appman resource api
	url := "http://localhost:6379"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		self.l.Printf("Get device db failed[0], err: %v\n", err)
		return "{}"
	}
	req.Header.Add("mx-api-token", self.apiToken)

	res, err := client.Do(req)
	if err != nil {
		self.l.Printf("Get device db failed[1], err: %v\n", err)
		return "{}"
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		self.l.Printf("Get device db failed[2], err: %v\n", err)
		return "{}"
	}

	// get item information from api response.
	ifaceId, _ := strconv.Atoi(id)

	// get all
	if ifaceId < 0 {
		return gjson.Get(string(response), "data").String()
	}

	result := gjson.Get(string(response), "data")
	target := ""
	result.ForEach(func(key, value gjson.Result) bool {
		if value.Get("id").Int() == int64(ifaceId) {
			target = value.String()
			return false
		}
		return true
	})

	if len(target) > 0 {
		return target
	}
	return "{}"
}

func (self *DBClient) Subscribe(hook func(msg string)) {
	if self.rdsPubSub == nil {
		self.l.Printf("device db pubsub channel is null\n")
		return
	}
	ch := self.rdsPubSub.Channel()
	for true {
		select {
		case msg, ok := <-ch:
			if ok {
				re, _ := regexp.Compile("__keyspace@[0-9]__:" + PREFIX_TAG + "(.*:.*)")
				key := re.ReplaceAllString(msg.Channel, "$1")
				arr := strings.Split(key, ":")
				if len(arr) < 2 {
					self.l.Printf("device db key is incorrect[%s]\n", key)
					continue
				}
				data := self.GetDBInfo(arr[0], arr[1])
				hook(data)
			}
		case <-self.ctx.Done():
			return
		}
	}
}

func (self *DBClient) Close() {
	if self.rdsPubSub != nil {
		err := self.rdsPubSub.Close()
		if err != nil {
			self.l.Printf("DBClient rdsPubSub.Close() error: %s", err)
		}
		self.l.Println("DBClient rdsPubSub.Close()")
	}
	if self.rdsDB != nil {
		err := self.rdsDB.Close()
		if err != nil {
			self.l.Printf("DBClient rdsDB.Close() error: %s", err)
		}
		self.l.Println("DBClient rdsDB.Close()")
	}
	self.ctxClose()
}

// func (w *WifiEntry) SubscribeDB(db *platform.DBClient) {
// 	// subscribe DB
// 	dbinfo := db.GetDBInfoAll()
// 	w.db = db
// 	w.Run(dbinfo)
// }

// func (w *WifiEntry) Run(dbinfo string) error {
// 	w.mutex.Lock()
// 	defer w.mutex.Unlock()
// 	go w.db.Subscribe(w.OnEthernetChange) // subscribe wifi change callback
// 	result := gjson.Parse(dbinfo)
// 	result.ForEach(func(key, value gjson.Result) bool {
// 		w.UpdateEthernetServiceStatus(value.String(), false)
// 		return true
// 	})

// 	return nil
// }

func DBupdate(message string) {
	fmt.Println(message)
}
