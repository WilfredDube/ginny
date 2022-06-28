package main

import "github.com/WilfredDube/ginny/internal"

func main() {
	router := internal.Routes()

	// TODO: Use http.Pusher() -> HTTP/2 & https required
	router.Static("/assets", "ui/assets")
	router.LoadHTMLGlob("ui/html/*")
	router.Run()
}
