package main

import "fmt"

type Cake struct {
	biscuit    string
	cream      string
	decoration string
}

func (c Cake) show() {
	fmt.Printf("%s with %s decorated with %s\n",
		c.biscuit,
		c.cream,
		c.decoration)
}

type CakeBuilder interface {
	makeBiscuit()
	makeCream()
	makeDecoration()
}

type StrawberryCakeBuilder struct {
	cake *Cake
}

func (b StrawberryCakeBuilder) makeBiscuit() {
	b.cake.biscuit = "vanilla biscuit"
}

func (b StrawberryCakeBuilder) makeCream() {
	b.cake.cream = "strawberry cream"
}

func (b StrawberryCakeBuilder) makeDecoration() {
	b.cake.decoration = "fresh strawberries"
}

type CakeDirector struct {
	builder CakeBuilder
}

func (d CakeDirector) Cook() {
	d.builder.makeBiscuit()
	d.builder.makeCream()
	d.builder.makeDecoration()
}

func main() {
	var cake Cake

	builder := StrawberryCakeBuilder{cake: &cake}
	director := CakeDirector{builder: builder}

	director.Cook()

	cake.show()
}
