package main

import "douyin/initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Redis()
	initialize.Router()
}
