package main

func main() {
	DB := DbClientNew("EeikTest")
	// fmt.Println(DB.GetDBInfoAll())
	go DB.Subscribe(DBupdate)
	DB.Close()

}
