###Что выведет программа? Объяснить вывод программы.

```go
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
 
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
   	go func() {
       	for {
           	select {
            case v := <-a:
               	c <- v
            case v := <-b:
            	c <- v
        	}
    	}
	}()
	return c
}
 
func main() {
 
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

###Ответ: программа не завершиться, и после вывода чисел, которые мы подали на горутины asChan, программа начнет выдавать нули, так как при чтении из закрытого канала всегда будет читаться дефолтное значение для типа данных, которое содержит канал.

Поэтому в данном случае, когда канал А или В закрыты, чтение из закрытого канала С будет продолжаться, и в каждой итерации мы получим нулевое значение для типа int.

Чтобы избежать бесконечного чтения из закрытого канала, нужно использовать второе возвращаемое значение из операции чтения (ok), чтобы проверить, был ли канал закрыт.

Исправвить можно так:

```go
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)	//закрываем канал, чтобы основная горутина не ждала ответа и не случился deadlock
		for {
			select {
			case v, ok := <-a: 	//получаем значение ok
				if !ok {		//false - означает, что канал закрыт и мы присваиваем a значение nil
					a = nil
					if b == nil { // если и канал В закрыт, то завершить горутину и закрыть канал С
						return
					}
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					if a == nil { // если и канал А закрыт, то завершить горутину и закрыть канал С
						return
					}
				}
				c <- v
			}
		}
	}()

	return c
}
```




