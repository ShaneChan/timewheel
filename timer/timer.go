package timer

import (
	"log"
	"time"
)

// taskEntity 时间轮任务实体
type taskEntity struct {
	expireTime int64  // 触发时间戳
	handle     func() // 注册回调函数
}

// Timer 时间轮定时器
type Timer struct {
	tasks map[string]*taskEntity // 任务 任务名字->任务实体

	stopC chan struct{} // 定时器停止channel
}

var newTimer Timer // 全局唯一的定时器变量

// 全局变量初始化
func init() {
	newTimer = Timer{
		tasks: make(map[string]*taskEntity),
		stopC: make(chan struct{}),
	}
}

func Add(taskName string, duration int64, f func()) bool {
	_, ok := newTimer.tasks[taskName]
	if ok {
		return false
	}

	newTimer.tasks[taskName] = &taskEntity{
		expireTime: time.Now().Unix() + duration,
		handle:     f,
	}

	return true
}

// DeleteTask 删除任务
func DeleteTask(taskName string) bool {
	_, ok := newTimer.tasks[taskName]
	if !ok {
		return false
	}

	delete(newTimer.tasks, taskName)
	return true
}

// QueryTask 查询单个任务的剩余时间
func QueryTask(taskName string) map[string]int64 {
	task, ok := newTimer.tasks[taskName]
	if !ok {
		return nil
	}

	return map[string]int64{taskName: task.expireTime - time.Now().Unix()}
}

// QueryAllTasks 查询所有任务的剩余时间
func QueryAllTasks() map[string]int64 {
	returnMap := make(map[string]int64)

	for taskName, taskEntity := range newTimer.tasks {
		returnMap[taskName] = taskEntity.expireTime - time.Now().Unix()
	}

	return returnMap
}

// Start 定时器启动
func Start() {
	go func() {
		ticker := time.NewTicker(time.Second)
	loop:
		for {
			select {
			case <-ticker.C:
				for taskName, task := range newTimer.tasks {
					if time.Now().Unix() >= task.expireTime {
						go task.handle()
						delete(newTimer.tasks, taskName)
					}
				}
			case <-newTimer.stopC:
				log.Println("timer is ending")
				break loop
			}
		}
	}()
}

// Stop 定时器停止
func Stop() {
	newTimer.stopC <- struct{}{}
}
