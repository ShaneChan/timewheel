package main

import (
	"log"
	"time"
	"timewheel/timer"
)

func main() {
	timer.Add("test", 10, func() {
		log.Println("你好世界")
	})

	timer.Start()
	timer.Add("test1", 20, func() {
		log.Println("再见世界")
	})

	time.Sleep(time.Hour)
}
