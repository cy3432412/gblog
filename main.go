package main

import (
	"gblog/model"
	"gblog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()

}
