package sunQueue

import (
	_ "fmt"
	"time"
)

//当前时间
var newTime = 0

var isRun = false

func GetTime() int {
	//如果没有运行 则开启一个协成运行
	if isRun == false {
		isRun = true
		go func() {
			for {

				tick := time.Tick(1 * time.Second)
				select {
				case <-tick:
					timeTick()
				}
			}
		}()
	}
	return newTime
}

func timeTick() {
	t := time.Now()
	UnixTime := t.Unix()
	timestamp := int(UnixTime)
	newTime = timestamp

}

func Gettimestamp() int {
	for newTime <= 0 && isRun == false {
		_ = GetTime()
		timeTick()
	}

	return newTime
}
