package main

import "strconv"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия - паттерн поведения объектов.
Он определяет семейство алгоритмов, инкапсулирует каждый из них и делает их взаимозаменяемыми.
Стратегия позволяет изменять алгоритмы независимо от клиентов, которые ими пользуются.

Применимость паттерна:
- когда имеется много родственных классов, отличающихся только поведением;
- когда нужно иметь несколько разных вариантов алгоритма;
- когда в алгоритме содержатся данные, о которых клиент не должен "знать";
- когда в классе определено много поведений, что представлено разветвленными условными операторами.

Плюсы и минусы:
- семейства родственных алгоритмов;
- альтернатива порождению подклассов;
- с помощью стратегий можно избавиться от условных операторов;
- выбор реализации;
- клиенты должны "знать" о различных стратегиях;
- обмен информацией между стратегией и контекстом;
- увеличение числа объектов.

Примеры использования:
- Библиотеки ЕТ++ [WGM88] и Interviews используют стратегии для инкапсуляции алгоритмов разбиения на строки.
*/

type StrategySort interface {
	Sort([]int)
}

type BubbleSort struct {
}

func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type InsertionSort struct {
}

func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

type Context struct {
	strategy StrategySort
}

func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

func main() {
	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	ctx := new(Context)
	ctx.Algorithm(&BubbleSort{})
	ctx.Sort(data1)

	var result1 string
	for _, val := range data1 {
		result1 += strconv.Itoa(val) + ","
	}

	ctx.Algorithm(&InsertionSort{})
	ctx.Sort(data2)

	var result2 string
	for _, val := range data2 {
		result2 += strconv.Itoa(val) + ","
	}
}
