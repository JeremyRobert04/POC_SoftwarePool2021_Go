package main

import "SofwareGoDay2/server"

func main() {
	i := server.NewServer()
	i.Run(":8080")
}