package main

import (
	"fmt"
)

func getTest(gcx *gorineeWebContext) {
	fmt.Println("test")
	gcx.ReplyJSON("test")
}

func getId(gcx *gorineeWebContext) {
	//param, query, path
	id := InterfaceConvertString(gcx.Param("id"))
	gcx.ReplyJSON(id)
}

func main() {
	g := new(gorineeWeb)
	r := g.New()
	r.Get("/test", getTest)
	r.Get("/{id}", getId)
	r.Run(":3000")
	//t := Get
	fmt.Println(int(Continue))
}
