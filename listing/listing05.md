### Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```


### Ответ: 

"error"

Задача по тематике похожа на третью. Смысл в том, что test вернет указатель на тип указатель на структуру со значением nil, который удовлетворяет интерфейсу error. Следовательно в err будет инициализироан *customError, а значит будет сам интерфейс уже будет не нулевой.

