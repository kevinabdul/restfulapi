package main

import (
	config "restfulalta/part-3-code-structuring/config"
	routes "restfulalta/part-3-code-structuring/routes"
)

func main() {
	config.InitDb()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}

