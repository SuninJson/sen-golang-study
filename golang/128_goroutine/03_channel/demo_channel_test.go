package main

import "testing"

func TestChannelSync(t *testing.T) {
	ChannelSync()
}

func TestChannelASync(t *testing.T) {
	ChannelASync()
}

func TestGoroutineNumCtl(t *testing.T) {
	GoroutineNumCtl()
}

func TestChannelDirectional(t *testing.T) {
	ChannelDirectional()
}

func TestSelectStmt(t *testing.T) {
	for i := 0; i < 10; i++ {
		println(i, ":")
		SelectStmt()
	}
}

func TestSelectFor(t *testing.T) {
	SelectFor()
}

func TestSelectEmptyBlock(t *testing.T) {
	SelectEmptyBlock()
}

func TestSelectNilChannelBlock(t *testing.T) {
	SelectNilChannelBlock()
}

func TestSelectNilChannel(t *testing.T) {
	SelectNilChannel()
}

func TestSelectNonBlock(t *testing.T) {
	SelectNonBlock()
}

func TestSelectRace(t *testing.T) {
	SelectRace()
}

func TestSelectAll(t *testing.T) {
	SelectAll()
}

func TestSelectChannelCloseSignal(t *testing.T) {
	SelectChannelCloseSignal()
}

func TestSelectSignal(t *testing.T) {
	SelectSignal()
}

func TestTimerA(t *testing.T) {
	TimerA()
}

func TestTimerB(t *testing.T) {
	TimerB()
}

func TestTickerA(t *testing.T) {
	TickerA()
}
