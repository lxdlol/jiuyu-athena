package main

import "athena/rest"

func main() {
	rest.StartServer()
	select {}
}
