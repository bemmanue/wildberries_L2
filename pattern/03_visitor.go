package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Посетитель - паттерн поведения объектов.
Он описывает операцию, выполняемую с каждым объектом из некоторой структуры.
Паттерн посетитель позволяет определить новую операцию, не изменяя классы этих объектов.

Применимость паттерна:
- в структуре присутствуют объекты многих классов с различными интерфейсами
и есть необходимость выполнять над ними операции, зависящие от конкретных классов;
- над объектами, входящими в состав структуры, надо выполнять разнообразные,
не связанные между собой операции и вы не хотите «засорять» классы такими операциями;
- классы, устанавливающие структуру объектов, изменяются редко,
но новые операции над этой структурой добавляются часто.

Плюсы и минусы:
- упрощает добавление новых операций;
- объединяет родственные операции и отсекает те, которые не имеют к ним отношения;
- добавление новых классов ConcreteElement затруднено;
- посещение различных иерархий классов;
- аккумулирование состояния;
- нарушение инкапсуляции.

Примеры использования:
- В компиляторе Smalltalk-80 имеется класс посетителя, который называется ProgramNodeEnumerator.
В основном он применяется в алгоритмах анализа исходного текста программы
и не используется ни для генерации кода, ни для красивой печати, хотя мог бы.
*/

type Visitor interface {
	VisitRestaurant(p *Restaurant) string
	VisitBar(p *Bar) string
}

type Place interface {
	Accept(v Visitor) string
}

type People struct {
}

func (v *People) VisitRestaurant(p *Restaurant) string {
	return p.Dinner()
}

func (v *People) VisitBar(p *Bar) string {
	return p.Drink()
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v) + "\n"
	}
	return result
}

type Restaurant struct {
}

func (s *Restaurant) Accept(v Visitor) string {
	return v.VisitRestaurant(s)
}

func (s *Restaurant) Dinner() string {
	return "Have dinner"
}

type Bar struct {
}

func (p *Bar) Accept(v Visitor) string {
	return v.VisitBar(p)
}

func (p *Bar) Drink() string {
	return "Drink cocktail"
}

func main() {
	city := new(City)

	city.Add(&Restaurant{})
	city.Add(&Bar{})

	fmt.Println(city.Accept(&People{}))
}
