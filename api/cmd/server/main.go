package main

import "github.com/FelpsCorrea/GoExpertPostgraduation/API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
