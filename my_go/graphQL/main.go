package main

import (
	"graphQL/router"
)

func main() {
	r := router.Router

	router.SetRouter()

	r.Run(":1234")
}
