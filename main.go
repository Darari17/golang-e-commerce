package main

import "github.com/Darari17/golang-e-commerce/router"

func main() {
	router.NewServer().Run(":3000")
}
