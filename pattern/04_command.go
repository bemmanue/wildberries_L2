package main

import "fmt"

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Команда - паттерн поведения объектов.
Он инкапсулирует запрос как объект, позволяя тем самым задавать параметры клиентов
для обработки соответствующих запросов, ставить запросы в очередь или протоколировать их,
а также поддерживать отмену операций.

Применимость паттерна:
- когда необходимо параметризовать объекты выполняемым действием;
- когда необходимо определять, ставить в очередь и выполнять запросы в разное время;
- когда необходимо поддержать отмену операций;
- когда необходимо поддержать протоколирование изменений,
чтобы их можно было выполнить повторно после аварийной остановки системы;
- когда необходимо структурировать систему на основе высокоуровневых операций, построенных из примитивных

Плюсы и минусы:
- позволяет добиться высокой гибкости при проектировании пользовательского интерфейса.

Примеры использования:
- в системе МасАрр [Арр89] команды широко применяются для реализации допускающих отмену операций.
*/

type Command interface {
	Execute() string
}

type Wash struct {
	receiver *Laundry
}

func (c *Wash) Execute() string {
	return c.receiver.Wash()
}

type WringOut struct {
	receiver *Laundry
}

func (c *WringOut) Execute() string {
	return c.receiver.WringOut()
}

type Laundry struct {
}

func (r *Laundry) Wash() string {
	return "Washing"
}

func (r *Laundry) WringOut() string {
	return "Wringing out"
}

type Machine struct {
	commands []Command
}

func (i *Machine) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Machine) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Machine) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func main() {
	washingMachine := &Machine{}
	laundry := &Laundry{}

	washingMachine.StoreCommand(&Wash{receiver: laundry})
	washingMachine.StoreCommand(&WringOut{receiver: laundry})

	fmt.Println(washingMachine.Execute())
}
