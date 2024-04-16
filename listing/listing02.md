###Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

```go
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

###Ответ:


Деферы добавляют функцию, следующую за ним, в стек приложения(фунции высшего порядка).
Все вызовы в стеке вызываются при возврате функции, в которой они добавлены. 
Поскольку вызовы помещаются в стек, они производятся в порядке от последнего к первому.

defer вызывается в трёх случаях:
 - Закончено выполнение функции, в которой вызывается defer (окружающая функция);
 - Окружающая функция выполнила оператор return (например, в теле цикла);
 - Программа паникует

defer не вызывается, если:
 - При обработке ошибок, или в любом другом случае, когда мы вызываем os.Exit();
 - Вызываем log.Fatal(), т.к. внутри неё также прячется os.Exit();
 - Паника возникает до вызова defer.

Их часто используют:
 - для очистки ресурсов;
 - для востановления паник;