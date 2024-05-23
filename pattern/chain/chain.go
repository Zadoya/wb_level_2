// Цепочка вызовов (Chain of responsibility) - поведенческий шаблон проектирования, который позволяет избежать
// жесткой привязки отправителя запроса к получателю. Все возможные обработчики запроса образуют цепочку,
// а сам запрос перемещается по этой цепочке. Каждый объект в этой цепочке при получении запроса выбирает,
// либо закончить обработку запроса, либо передать запрос на обработку следующему по цепочке объекту.

// Паттерн цепочка вызовов применяется:
// Когда имеется более одного объекта, который может обработать определенный запрос.
// Когда надо передать запрос на выполнение одному из нескольких объект, точно не определяя, какому именно объекту.
// Когда набор объектов задается динамически.

// Преимущества паттерна цепочка вызовов:
// Ослабление связанности между объектами. Отправителю и получателю запроса ничего не известно друг о друге.
// Клиенту неизветна цепочка объектов, какие именно объекты составляют ее, как запрос в ней передается.
// В цепочку с легкостью можно добавлять новые типы объектов, которые реализуют общий интерфейс.

// Недостатки паттерна цепочка вызовов:
// никто не гарантирует, что запрос в конечном счете будет обработан. Если необходимого обработчика в цепочки
// не оказалось, то запрос просто выходит из цепочки и остается необработанным.

package main

import "fmt"

// Рассмотреть паттерн Цепочка обязанностей можно на примере больницы.
// Госпиталь может иметь разные помещения, например:
//  -- Приемное отделение
//  -- Доктор
//  -- Комната медикаментов
//  -- Кассир
// Когда пациент прибывает в больницу, первым делом он попадает в Приемное отделение, 
// оттуда – к Доктору, затем в Комнату медикаментов, после этого – к Кассиру, и так далее. 
// Пациент проходит по цепочке помещений, в которой каждое отправляет его 
// по ней дальше сразу после выполнения своей функции.

type Patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

type Department interface {
    execute(*Patient)
    setNext(Department)
}

type Reception struct {
    next Department
}

func (r *Reception) execute(p *Patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
    r.next = next
}

type Doctor struct {
    next Department
}

func (d *Doctor) execute(p *Patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
    d.next = next
}

type Medical struct {
    next Department
}

func (m *Medical) execute(p *Patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to patient")
        m.next.execute(p)
        return
    }
    fmt.Println("Medical giving medicine to patient")
    p.medicineDone = true
    m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
    m.next = next
}

type Cashier struct {
    next Department
}

func (c *Cashier) execute(p *Patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
    c.next = next
}

// клментский код
func main() {

    cashier := &Cashier{}

    // уставливаем значение Next для комнаты медикаментов
    medical := &Medical{}
    medical.setNext(cashier)

    //уставливаем значение Next для доктора
    doctor := &Doctor{}
    doctor.setNext(medical)

    //уставливаем значение Next для ресепшена
    reception := &Reception{}
    reception.setNext(doctor)

    patient := &Patient{name: "John Doy"}
    //пачиент пришёл в больницу
    reception.execute(patient)
}