package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель - паттерн, порождающий объекты.
Он отделяет конструирование сложного объекта от его представления,
так что в результате одного и того же процесса конструирования могут получаться разные представления.

Применимость паттерна:
- алгоритм создания сложного объекта не должен зависеть от того,
из каких частей состоит объект и как они стыкуются между собой;
- процесс конструирования должен обеспечивать различные представления конструируемого объекта.

Плюсы и минусы:
- позволяет изменять внутреннее представление продукта;
- изолирует код, реализующий конструирование и представление;
- дает более тонкий контроль над процессом конструирования.

Примеры использования:
- приложение для конвертирования из формата RTF из библиотеки ЕТ++ [WGM88]
*/

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
