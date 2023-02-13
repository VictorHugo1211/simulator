package main

import (
	"fmt"
	route "github.com/VictorHugo1211/simulator/application/route"
)

func main() {
	route := route.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[1])
}
