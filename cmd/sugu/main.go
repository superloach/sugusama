package main

import (
	"flag"
	"fmt"

	"github.com/superloach/minori"
	"github.com/superloach/sugusama"
)

var (
	user = flag.String("user", "", "instagram username")
	pass = flag.String("pass", "", "instagram password")
)

var Log = minori.GetLogger("sugu")

func main() {
	flag.Parse()

	if *user == "" {
		panic("please provide user")
	}
	if *pass == "" {
		panic("please provide pass")
	}

	c, err := sugusama.NewClient(nil)
	if err != nil {
		panic(err)
	}

	err = c.Login(*user, *pass)
	if err != nil {
		panic(err)
	}

	fmt.Printf("logged in as %s\n", c.Viewer.Username)

	err = c.FetchActivity()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", c.Activity)

	err = c.FetchHome()
	if err != nil {
		panic(err)
	}
}
