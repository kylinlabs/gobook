//第5章/mod2/factory.go
package main

import "fengyoulin.com/mod1"

type uf struct{}

func (*uf) NewUser(id int, nick string) interface{} {
	return &mod1.User{
		ID:   id,
		Nick: nick,
	}
}

var UserFactory uf
