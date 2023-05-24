package main

import "log"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод - паттерн, порождающий классы.
Он определяет интерфейс для создания объекта, но оставляет подклассам решение о том,
какой класс инстанцировать. Фабричный метод позволяет классу делегировать инстанцирование подклассам.

Применимость паттерна:
- когда классу заранее неизвестно, объекты каких классов ему нужно создавать;
- когда класс спроектирован так, чтобы объекты, которые он создает, специфицировались подклассами;
- когда класс делегирует свои обязанности одному из нескольких вспомогательных
подклассов, и вы планируете локализовать знание о том, какой класс принимает эти обязанности на себя.

Плюсы и минусы:
- предоставляет подклассам операции-зацепки (hooks);
- соединяет параллельные иерархии.

Примеры использования:
- Фабричные методы в изобилии встречаются в инструментальных библиотеках и каркасах.
*/

type action string

const (
	Bread  action = "Bread"
	Cookie action = "B"
)

type Creator interface {
	Cook(action action) Product // Factory Method
}

type Product interface {
	Use() string // Every action should be usable
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) Cook(action action) Product {
	var product Product

	switch action {
	case Bread:
		product = &ConcreteBread{string(action)}
	case Cookie:
		product = &ConcreteCookie{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

type ConcreteBread struct {
	action string
}

func (p *ConcreteBread) Use() string {
	return p.action
}

type ConcreteCookie struct {
	action string
}

func (p *ConcreteCookie) Use() string {
	return p.action
}

func main() {
	bakery := NewCreator()
	products := []Product{
		bakery.Cook(Bread),
		bakery.Cook(Cookie),
	}

	for _, product := range products {
		product.Use()
	}
}
