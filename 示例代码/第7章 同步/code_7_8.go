//第7章/code_7_8.go
select {
case v := <-ch1:
	println(v)
case ch2 <- 10:
default:
}
