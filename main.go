package main

import "expense-bucket-api/server"

func main() {
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
