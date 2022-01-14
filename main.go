package main

func main() {
	g := new(gorineeWeb)
	r := g.New()
	r.Run(":3000")
}
