package main

import "github.com/WilfredDube/ginny/internal"

func main() {
	router := internal.Routes()
	router.Run()
}
