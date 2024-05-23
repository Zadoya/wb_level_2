// Посетитель (visitor) — паттерн поведения объектов. Описывает операцию, выполняемую с каждым объектом
//  из некоторой структуры. Позволяет определить новую операцию, не изменяя классы этих объектов.

// Паттерн посетитель применяется:
// Когда классам необходимо добавить одинаковый набор операций без изменения этих классов.
// Когда часто добавляются новые операции к классам, при этом общая структура классов стабильна
//  и практически не изменяется.
// Когда имеется много объектов разнородных классов с разными интерфейсами, и требуется выполнить
// ряд операций над каждым из этих объектов.

// Приемущества паттерна посетитель:
// Упрощает добавление операций, работающих со сложными структурами объектов.
// Объединяет родственные операции в одном классе.
// Посетитель может накапливать состояние при обходе структуры элементов.

// Недостатки паттерна посетитель:
// Паттерн не оправдан, если иерархия элементов часто меняется.
// Может привести к нарушению инкапсуляции элементов.

package main

import "fmt"

// Паттерн Посетитель позволяет вам добавлять поведение в структуру без ее изменения. 
// Представим, что есть библиотека, которая содержит структуры разных фигур:
//		- Квадрат
// 		- Круг
//		- Треугольник
// Структуры каждой из вышеназванных фигур реализуют общий интерфейс фигуры.
// Нам необходимо добавить в структуру функцию getArea, возвращающую площадь фигуры.

type Visitor interface {
    visitForSquare(*Square)
    visitForCircle(*Circle)
    visitForrectangle(*Rectangle)
}

type Shape interface {
    getType() string
    accept(Visitor)
}

// фигура квадрат
type Square struct {
    side int
}

func (s *Square) accept(v Visitor) {
    v.visitForSquare(s)
}

func (s *Square) getType() string {
    return "Square"
}

// фигура круг
type Circle struct {
    radius int
}

func (c *Circle) accept(v Visitor) {
    v.visitForCircle(c)
}

func (c *Circle) getType() string {
    return "Circle"
}

// фигура треугольник
type Rectangle struct {
    l int
    b int
}

func (t *Rectangle) accept(v Visitor) {
    v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
    return "rectangle"
}

type AreaCalculator struct {
    area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
    fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
    fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
    fmt.Println("Calculating area for rectangle")
}

// В случае добавления другого функционала, например getMiddleCoordinates, 
// будет использоваться все тот же метод accept(v visitor) без новых изменений структур фигур.

type MiddleCoordinates struct {
    x int
    y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
    fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
    fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
    fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() {
    square := &Square{side: 2}
    circle := &Circle{radius: 3}
    rectangle := &Rectangle{l: 2, b: 3}

    areaCalculator := &AreaCalculator{}

    square.accept(areaCalculator)
    circle.accept(areaCalculator)
    rectangle.accept(areaCalculator)

    fmt.Println()
    middleCoordinates := &MiddleCoordinates{}
    square.accept(middleCoordinates)
    circle.accept(middleCoordinates)
    rectangle.accept(middleCoordinates)
}