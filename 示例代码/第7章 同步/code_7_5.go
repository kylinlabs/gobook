//第7章/code_7_5.go
type Spin int32

func (l *Spin) Lock() {
	lock((*int32)(l), 0, 1)
}

func (l *Spin) Unlock() {
	unlock((*int32)(l), 0)
}

func lock(ptr *int32, o, n int32)
func unlock(ptr *int32, n int32)
