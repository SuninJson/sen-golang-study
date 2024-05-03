package main

import (
	"runtime"
	"testing"
)

func TestGoroutineAnts(t *testing.T) {
	GoroutineAnts()
}

func TestCPU(t *testing.T) {
	// 最多利用一半的CPU
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	// 获取当前CPU的核数
	println(runtime.NumCPU())

}
