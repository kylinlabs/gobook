//第7章/code_7_7.go
func (l *Spin) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(l), 0, 1) {
		pause()
	}
}
