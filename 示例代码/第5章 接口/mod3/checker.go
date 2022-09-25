//第5章/mod3/checker.go
package main

import "fengyoulin.com/mod1"

type uc struct{}

func (*uc) IsUser(a interface{}) bool {
	_, ok := a.(*mod1.User)
	return ok
}

var UserChecker uc
