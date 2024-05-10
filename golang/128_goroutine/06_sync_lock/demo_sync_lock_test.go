package _6_sync_lock

import "testing"

func TestDemoMutex(t *testing.T) {
	DemoMutex()
}

func TestDemoRWMutex(t *testing.T) {
	DemoRWMutex()
}

func TestSyncAtomicAdd(t *testing.T) {
	SyncAtomicAdd()
}

func TestSyncAtomicValue(t *testing.T) {
	SyncAtomicValue()
}

func TestSyncPool(t *testing.T) {
	SyncPool()
}

func TestSyncOnce(t *testing.T) {
	SyncOnce()
}

func TestSyncCondition(t *testing.T) {
	SyncCondition()
}
