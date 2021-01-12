package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	// /* standard parser: the smallest unit is minute */
	// c := cron.New()

	// c.AddFunc("* * * * * ", func() {
	// 	fmt.Printf("job, %s\n", time.Now())
	// })

	// c.Start()
	// time.Sleep(time.Second * 300)

	// /*
	// 	output : job, 2020-12-24 10:44:00.000168991 +0800 CST m=+52.003917302
	// 	output : job, 2020-12-24 10:45:00.000315456 +0800 CST m=+112.004063911
	// */

	c := cron.New()

	c.AddFunc("* * * * 5 ", func() {
		fmt.Printf("job, %s\n", time.Now())
	})

	c.Start()
	time.Sleep(time.Second * 60)
	<-c.Stop().Done()
	fmt.Println("go home")
	// time.Sleep(time.Second * 300)
	/*
		output : job, 2020-12-24 10:44:00.000168991 +0800 CST m=+52.003917302
		output : job, 2020-12-24 10:45:00.000315456 +0800 CST m=+112.004063911
	*/

	//----------------------------------------------------------------

	// // second parser: the smallest unit is second

	// c := cron.New(cron.WithSeconds())
	// id, _ := c.AddFunc("* * * * * *", func() {
	// 	fmt.Printf("job, %s\n", time.Now())
	// })
	// c.AddFunc("5 * * * * *", func() {
	// 	fmt.Printf("the 5th second\n")
	// })
	// c.Start()
	// time.Sleep(time.Second * 300)
	// c.Stop()
	// c.Remove(id)
	// /*
	// 	output : job, 2020-12-24 10:50:00.000424071 +0800 CST m=+38.788063264
	// 	output : job, 2020-12-24 10:50:01.000439498 +0800 CST m=+39.788077603
	// 	output : job, 2020-12-24 10:50:02.000294165 +0800 CST m=+40.787932191
	// 	output : job, 2020-12-24 10:50:03.000349644 +0800 CST m=+41.787987680
	// 	output : job, 2020-12-24 10:50:04.000282624 +0800 CST m=+42.787920664
	// 	output : the 5th second
	// 	output : job, 2020-12-24 10:50:05.000447426 +0800 CST m=+43.788085456
	// 	output : job, 2020-12-24 10:50:06.000274796 +0800 CST m=+44.787912760
	// 	output : job, 2020-12-24 10:50:07.00021668 +0800 CST m=+45.787854610
	// 	output : job, 2020-12-24 10:50:08.000369732 +0800 CST m=+46.788007768
	// */

	// // c.AddFunc("@every 1s", func() { fmt.Println("tick every 1 second") })
}
