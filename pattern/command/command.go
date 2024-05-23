// Команда (Command) — поведенческий паттерн проектирования, который превращает запросы в объекты,
// позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
// а также поддерживать отмену операций.

// Паттерн команда применяется:
// Когда надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие действия.
// Когда необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.

// Преимущества паттерна команда:
// Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
// Позволяет реализовать простую отмену и повтор операций.
// Позволяет реализовать отложенный запуск операций.
// Позволяет собирать сложные команды из простых.

// Недостаток паттерна команда это усложнение кода программы из-за введения множества дополнительных классов.

package main

import "fmt"

// Рассмотреть паттерн Команда можно примере телевизора. Его можно включить двумя способами:
// с помощью кнопки на самом телевизоре или пультом.

type Button struct {
    command Command
}

func (b *Button) press() {
    b.command.execute()
}

// интерфейс команды
type Command interface {
    execute()
}

// команда включения
type OnCommand struct {
    device Device
}

func (c *OnCommand) execute() {
    c.device.on()
}

// команда выключения
type OffCommand struct {
    device Device
}

func (c *OffCommand) execute() {
    c.device.off()
}

// интерфейс устройства
type Device interface {
    on()
    off()
}

// телевизор
type Tv struct {
    isRunning bool
}

func (t *Tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *Tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}

// клиентский код
func main() {
    tv := &Tv{}

    onCommand := &OnCommand{
        device: tv,
    }

    offCommand := &OffCommand{
        device: tv,
    }

    onButton := &Button{
        command: onCommand,
    }
    onButton.press()

    offButton := &Button{
        command: offCommand,
    }
    offButton.press()
}