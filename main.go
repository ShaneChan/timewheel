package main

import (
	"log"
	"strconv"
	"time"
	"timewheel/timer"
)

func main() {
	for i := 0; i < 100000; i++ {
		timer.Add("test0"+strconv.Itoa(i), 5, func() {
			time.Sleep(time.Second * 2)
			log.Println("你好世界a" + strconv.Itoa(i))
		})
	}
	for i := 0; i < 100000; i++ {
		timer.Add("test1"+strconv.Itoa(i), 10, func() {
			time.Sleep(time.Second * 2)
			log.Println("你好世界b" + strconv.Itoa(i))
		})
	}
	for i := 0; i < 100000; i++ {
		timer.Add("test2"+strconv.Itoa(i), 15, func() {
			time.Sleep(time.Second * 2)
			log.Println("你好世界c" + strconv.Itoa(i))
		})
	}
	for i := 0; i < 100000; i++ {
		timer.Add("test3"+strconv.Itoa(i), 20, func() {
			time.Sleep(time.Second * 2)
			log.Println("你好世界d" + strconv.Itoa(i))
		})
	}
	for i := 0; i < 100000; i++ {
		timer.Add("test4"+strconv.Itoa(i), 25, func() {
			time.Sleep(time.Second * 2)
			log.Println("你好世界e" + strconv.Itoa(i))
		})
	}

	timer.Start()

	time.Sleep(time.Hour)
}
