Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
```

В первом случае печатается nil, так как функция возвращает 
интерфейсный тип error, в котором динамическое значение равно nil.
Во втором случае сравнение интерфейса error и nil возвращает false,
поскольку интерфейсный тип хоть и имеет нулевое значение, 
но обладает конкретным типом *os.PathError. А как известно, 
интерфейс равен nil только когда имеет нулевой тип и нулевое значение.