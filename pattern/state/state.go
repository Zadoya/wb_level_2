// Состояние (State) — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.

// Паттерн состояние применяется:
// Когда поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения.
// Когда в коде методов объекта используются многочисленные условные конструкции, выбор которых зависит от
// текущего состояния объекта.

// Преимущества паттерна состояние:
// Избавляет от множества больших условных операторов машины состояний.
// Концентрирует в одном месте код, связанный с определённым состоянием.
// Упрощает код контекста.

// Паттерн состояние может неоправданно усложнить код, если состояний мало и они редко меняются.

package main

import (
	"fmt"
	"log"
)

// Рассмотрим паттерн проектирования Состояние в контексте торговых автоматов. 
// Для упрощения задачи представим, что торговый автомат может выдавать только один товар. 
// Также представим, что автомат может пребывать только в одном из четырех состояний:
// 		- hasItem (имеет товар)
//		- noItem (не имеет товаров)
//		- itemRequested (выдаёт товар)
//		- hasMoney (получил деньги)
//
// Торговый автомат может иметь различные действия. Опять-таки, для простоты оставим только четыре из них:
//		- Выбрать товар
//		- Добавить товар
//		- Ввести деньги
//		- Выдать товар
// Паттерн Состояние нужно использовать в случаях, когда объект может иметь много различных состояний, 
// которые он должен менять в зависимости от конкретного поступившего запроса.

type VendingMachine struct {
    hasItem       State
    itemRequested State
    hasMoney      State
    noItem        State

    currentState State

    itemCount int
    itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
    v := &VendingMachine{
        itemCount: itemCount,
        itemPrice: itemPrice,
    }
    hasItemState := &HasItemState{
        vendingMachine: v,
    }
    itemRequestedState := &ItemRequestedState{
        vendingMachine: v,
    }
    hasMoneyState := &HasMoneyState{
        vendingMachine: v,
    }
    noItemState := &NoItemState{
        vendingMachine: v,
    }

    v.setState(hasItemState)
    v.hasItem = hasItemState
    v.itemRequested = itemRequestedState
    v.hasMoney = hasMoneyState
    v.noItem = noItemState
    return v
}

func (v *VendingMachine) requestItem() error {
    return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
    return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
    return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
    return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
    v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
    fmt.Printf("Adding %d items\n", count)
    v.itemCount = v.itemCount + count
}

type State interface {
    addItem(int) error
    requestItem() error
    insertMoney(money int) error
    dispenseItem() error
}

// состояние, когда нет товаров
type NoItemState struct {
    vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
    return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
    i.vendingMachine.incrementItemCount(count)
    i.vendingMachine.setState(i.vendingMachine.hasItem)
    return nil
}

func (i *NoItemState) insertMoney(money int) error {
    return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) dispenseItem() error {
    return fmt.Errorf("Item out of stock")
}

// состояние, когда есть товар
type HasItemState struct {
    vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
    if i.vendingMachine.itemCount == 0 {
        i.vendingMachine.setState(i.vendingMachine.noItem)
        return fmt.Errorf("No item present")
    }
    fmt.Printf("Item requestd\n")
    i.vendingMachine.setState(i.vendingMachine.itemRequested)
    return nil
}

func (i *HasItemState) addItem(count int) error {
    fmt.Printf("%d items added\n", count)
    i.vendingMachine.incrementItemCount(count)
    return nil
}

func (i *HasItemState) insertMoney(money int) error {
    return fmt.Errorf("Please select item first")
}
func (i *HasItemState) dispenseItem() error {
    return fmt.Errorf("Please select item first")
}

// состояние, когда автомат выдает товар
type ItemRequestedState struct {
    vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
    return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
    return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
    if money < i.vendingMachine.itemPrice {
        return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
    }
    fmt.Println("Money entered is ok")
    i.vendingMachine.setState(i.vendingMachine.hasMoney)
    return nil
}
func (i *ItemRequestedState) dispenseItem() error {
    return fmt.Errorf("Please insert money first")
}

// состояние, когда автомат получил деньги
type HasMoneyState struct {
    vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
    return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
    return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
    return fmt.Errorf("Item out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
    fmt.Println("Dispensing Item")
    i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
    if i.vendingMachine.itemCount == 0 {
        i.vendingMachine.setState(i.vendingMachine.noItem)
    } else {
        i.vendingMachine.setState(i.vendingMachine.hasItem)
    }
    return nil
}

// клиентский код

func main() {
    vendingMachine := newVendingMachine(1, 10)

    err := vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()

    err = vendingMachine.addItem(2)
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()

    err = vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
}