package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние - паттерн поведения объектов.
Он позволяет объекту варьировать свое поведение в зависимости от внутреннего состояния.
Извне создается впечатление, что изменился класс объекта.

Применимость паттерна:
- когда поведение объекта зависит от его состояния и должно изменяться во время выполнения;
- когда в коде операций встречаются состоящие из многих ветвей условные операторы,
в которых выбор ветви зависит от состояния.

Плюсы и минусы:
- локализует зависящее от состояния поведение и делит его на части, соответствующие состояниям;
- делает явными переходы между состояниями;
- объекты состояния можно разделять.

Примеры использования:
- В графических редакторах позволяет клиентам легко определять новые виды инструментов.
*/

type FreezingMode interface {
	Freeze() string
}

type Fridge struct {
	state FreezingMode
}

func (a *Fridge) Freeze() string {
	return a.state.Freeze()
}

func (a *Fridge) SetState(state FreezingMode) {
	a.state = state
}

func NewFridge() *Fridge {
	return &Fridge{state: &DefaultFreezingMode{}}
}

type DefaultFreezingMode struct {
}

func (a *DefaultFreezingMode) Freeze() string {
	return "Freeze"
}

type SuperFreezingMode struct {
}

func (a *SuperFreezingMode) Freeze() string {
	return "Super freeze"
}

func main() {
	fridge := NewFridge()

	fmt.Println(fridge.Freeze())

	fridge.SetState(&SuperFreezingMode{})

	fmt.Println(fridge.Freeze())
}
