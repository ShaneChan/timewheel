package main

import (
	"log"
	"strconv"
	"time"
	"timewheel/timer"
)

func main() {
	for i := 0; i < 100000; i++ {
		timer.Add("test"+strconv.Itoa(i), 5, func() {
			time.Sleep(time.Second * 2)
			log.Println("你好世界" + strconv.Itoa(i))
		})
	}

	timer.Start()

	time.Sleep(time.Hour)
}
