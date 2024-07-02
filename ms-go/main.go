package main

import (
	_ "ms-go/db"
	"ms-go/router"
	"ms-go/app/consumers"
)

func main() {
	go func() {
		consumers.RunConsumer()
	}()

	router.Run()
}
