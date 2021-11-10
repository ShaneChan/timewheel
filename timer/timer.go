package timer

import (
	"time"
)

// taskEntity 时间轮任务实体
type taskEntity struct {
	duration int64  // 触发时间戳
	handle   func() // 注册回调函数
}

// Timer 时间轮定时器
type Timer struct {
	tasks map[string]taskEntity // 任务 任务名字->任务实体

	stopC chan struct{} // 定时器停止channel
}

var newTimer Timer

func init() {
	newTimer = Timer{
		tasks: make(map[string]taskEntity),
		stopC: make(chan struct{}),
	}
}

func Add(taskName string, duration int64, f func()) bool {
	_, ok := newTimer.tasks[taskName]
	if ok {
		return false
	}

	newTimer.tasks[taskName] = taskEntity{
		duration: time.Now().Unix() + duration,
		handle:   f,
	}

	return true
}

func Delete(taskName string) bool {
	_, ok := newTimer.tasks[taskName]
	if !ok {
		return false
	}

	delete(newTimer.tasks, taskName)
	return true
}

func Start() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				for taskName, task := range newTimer.tasks {
					if time.Now().Unix() >= task.duration {
						go task.handle()
						delete(newTimer.tasks, taskName)
					}
				}
			case <-newTimer.stopC:
				break
			}
		}
	}()
}

func (newTimer *Timer) Stop() {

}
