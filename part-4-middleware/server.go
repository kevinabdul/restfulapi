package main

import (
	config "restfulalta/part-4-middleware/config"
	routes "restfulalta/part-4-middleware/routes"
)

func main() {
	config.InitDb()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}

