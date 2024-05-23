// Фабричный метод (Factory Method) — это порождающий паттерн проектирования, который решает проблему создания
// различных продуктов, без указания конкретных классов продуктов.

// Паттерн фабричный метод применяется:
// Когда заранее неизвестно, объекты каких типов необходимо создавать.
// Когда система должна быть независимой от процесса создания новых объектов и расширяемой: в нее можно
// легко вводить новые классы, объекты которых система должна создавать.
// Когда создание новых объектов необходимо делегировать из базового класса классам наследникам.

// Преимущества паттерна фабричный метод:
// Избавляет класс от привязки к конкретным классам продуктов.
// Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// Упрощает добавление новых продуктов в программу.

// Недостатки паттерна фабричный метод:
// Для каждого нового продукта необходимо создавать свой класс создателя.

package main

import "fmt"

// В Go невозможно реализовать классический вариант паттерна Фабричный метод, 
// поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность. 
// Но можно реализовать паттерн - Простая фабрика

// В этом примере будут создаваться разные типы оружия при помощи структуры фабрики.

// интерфейс Guner определяет все методы будущих пушек.
type Guner interface {
    setName(name string)
    setPower(power int)
    getName() string
    getPower() int
}

// структура Gun, которая реализует интерфейс Guner
type Gun struct {
    name  string
    power int
}

func (g *Gun) setName(name string) {
    g.name = name
}

func (g *Gun) getName() string {
    return g.name
}

func (g *Gun) setPower(power int) {
    g.power = power
}

func (g *Gun) getPower() int {
    return g.power
}

// две конкретных пушки — ak47 и musket — обе включают в себя структуру Gun
// и не напрямую реализуют все методы от Guner

type Ak47 struct {
    Gun
}

func newAk47() Guner {
    return &Ak47{
        Gun: Gun{
            name:  "AK47 gun",
            power: 4,
        },
    }
}

type Musket struct {
	Gun
}

func newMusket() Guner {
	return &Musket{
		Gun: Gun{
			name: "Musket gun",
			power: 1,
		},
	}
}

// gunFactory служит фабрикой, которая создает пушку нужного типа в зависимости 
// от аргумента на входе. Клиентом служит main.go . Вместо прямого взаимодействия 
// с объектами ak47 или musket, он создает экземпляры различного оружия при помощи 
// gunFactory, используя для контроля изготовления только параметры в виде строк.

func gunFactory(gunType string) (Guner, error) {
    if gunType == "ak47" {
        return newAk47(), nil
    }
    if gunType == "musket" {
        return newMusket(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
    ak47, _ := gunFactory("ak47")
    musket, _ := gunFactory("musket")

    printDetails(ak47)
    printDetails(musket)
}

func printDetails(g Guner) {
    fmt.Printf("Gun: %s", g.getName())
    fmt.Println()
    fmt.Printf("Power: %d", g.getPower())
    fmt.Println()
}