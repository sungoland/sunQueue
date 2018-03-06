// Copyright 2018 sun. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//时间类处理信息 通过时钟调用 记录时间戳 频繁进入内核态 获取时间
package sunQueue

import (
	_ "fmt"
	"time"
)

//当前时间
var newTime = 0

var isRun = false

//如果没有运行 则开启一个协成运行
func runTick() {
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
}

func timeTick() {
	t := time.Now()
	UnixTime := t.Unix()
	timestamp := int(UnixTime)
	newTime = timestamp

}

func Gettimestamp() int {
	for newTime <= 0 && isRun == false {
		runTick()
		timeTick()
	}

	return newTime
}