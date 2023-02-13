package main

import (
	"fmt"
	route2 "github.com/VictorHugo1211/simulator/application/route"
)

func main() {
	route := route2.Route{
		ID: "4",
		ClientID: "4",
	}
	fmt.Println("Hello, playground")
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println("Hello, playground")
	fmt.Println(stringjson[0])
}
