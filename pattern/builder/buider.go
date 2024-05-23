// Строитель (Builder) - шаблон проектирования, который инкапсулирует создание объекта
// и позволяет разделить его на различные этапы.

// Шаблон проектирования строитель состоить из 4 компонентов:
// Product: представляет объект, который должен быть создан.
// В данном случае все части объекта заключены в списке parts.
// Builder: определяет интерфейс для создания различных частей объекта Product.
// ConcreteBuilder: конкретная реализация Buildera.
// Создает объект Product и определяет интерфейс для доступа к нему.
// Director: распорядитель - создает объект, используя объекты Builder.

// Паттер строитель применяется когда процесс создания нового объекта не должен зависеть от того,
// из каких частей этот объект состоит и как эти части связаны между собой.
// Строитель применяется когда необходимо обеспечить получение различных вариаций объекта в процессе его создания.

// Паттер строитель позволяет изменять внутреннее представление продукта.
// Объект Builder предоставляет распорядителю абстрактный интерфейс для конструирования продукта,
// за которым он может скрыть представление и внутреннюю структуру продукта, а также процесс его сборки.
// Поскольку продукт конструируется через абстрактный интерфейс, то для изменения внутреннего представления
// достаточно всего лишь определить новый вид строителя.

// Строитель изолирует код, реализующий конструирование и представление.
// Улучшается модульность, инкапсулируя способ конструирования и представления сложного объекта.

// Строитель предоставляет более точный контроль над процессом конструирования.
// В отличие от порождающих паттернов, которые сразу конструируют весь объект целиком,
// builder делает это шаг за шагом под управлением director. Когда продукт завершен, director забирает его у builder.

// Строители конструируют свои продукты шаг за шагом, поэтому интерфейс класса Builder должен быть достаточно общим,
// чтобы обеспечить конструирование при любом виде конкретного строителя.

package main

import "fmt"

// Нижеприведенный код использует разные типы домов (igloo и normalHouse), 
// которые конструируются с помощью строителей iglooBuilder и normalBuilder. 
// При создании каждого дома используются одинаковые шаги. 

type House struct {
    windowType string
    doorType   string
    floor      int
}

type IBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() House
}

func getBuilder(builderType string) IBuilder {
    if builderType == "normal" {
        return newNormalBuilder()
    }

    if builderType == "igloo" {
        return newIglooBuilder()
    }
    return nil
}
type NormalBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newNormalBuilder() *NormalBuilder {
    return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
    b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
    b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
    b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
    return House{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}

type IglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newIglooBuilder() *IglooBuilder {
    return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
    b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
    b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
    b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
    return House{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}

type Director struct {
    builder IBuilder
}

func newDirector(b IBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) setBuilder(b IBuilder) {
    d.builder = b
}

func (d *Director) buildHouse() House {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}

// клиентский код
func main() {
    normalBuilder := getBuilder("normal")
    iglooBuilder := getBuilder("igloo")

    director := newDirector(normalBuilder)
    normalHouse := director.buildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

    director.setBuilder(iglooBuilder)
    iglooHouse := director.buildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}