package main

import "github.com/ppanissa-dev/go-api-restful/configs"

func main() {
	config, _ := configs.LoadConfig(".")

	println(config.DBDriver)

}
