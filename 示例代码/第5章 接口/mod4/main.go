//第5章/mod4/main.go
package main

import (
	"log"
	"plugin"
	"reflect"
)

type UserFactory interface {
	NewUser(id int, nick string) interface{}
}

type UserChecker interface {
	IsUser(a interface{}) bool
}

func factory() UserFactory {
	p, err := plugin.Open("./mod2.so")
	if err != nil {
		log.Fatalln(err)
	}
	a, err := p.Lookup("UserFactory")
	if err != nil {
		log.Fatalln(err)
	}
	uf, ok := a.(UserFactory)
	if !ok {
		log.Fatalln("not a UserFactory")
	}
	return uf
}

func checker() UserChecker {
	p, err := plugin.Open("./mod3.so")
	if err != nil {
		log.Fatalln(err)
	}
	a, err := p.Lookup("UserChecker")
	if err != nil {
		log.Fatalln(err)
	}
	uc, ok := a.(UserChecker)
	if !ok {
		log.Fatalln("not a UserChecker")
	}
	return uc
}

func main() {
	uf := factory()
	uc := checker()
	u := uf.NewUser(1, "Jack")
	if !uc.IsUser(u) {
		log.Println("not a User")
	}
	t := reflect.TypeOf(u)
	println(u, t.String())
	select {}
}
