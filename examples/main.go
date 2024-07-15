package main

import (
	"fmt"
	"time"

	"github.com/lei006/godbtt"
)

func main() {
	fmt.Println("=================")
	dbtt := godbtt.NewDBTT(godbtt.Option{})
	dbtt.Start()
	time.Sleep(time.Second * 1)
	dbtt.Stop()
	fmt.Println("=================")
}
