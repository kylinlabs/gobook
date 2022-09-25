//第7章/code_7_6.go
import "sync/atomic"

type Spin int32

func (l *Spin) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(l), 0, 1) {
	}
}

func (l *Spin) Unlock() {
	atomic.StoreInt32((*int32)(l), 0)
}
