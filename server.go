package main

import "card-utils-rest-server/routers"

func main() {
	router := routers.InitRouter()
	err := router.Run()// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
