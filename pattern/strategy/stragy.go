// Стратегия (Strategy) — поведенческий шаблон проектирования, предназначенный для определения семейства алгоритмов,
// инкапсуляции каждого из них и обеспечения их взаимозаменяемости. Это позволяет выбирать алгоритм путём определения
// соответствующего класса. Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов,
// которые его используют.

// Паттерн стратегия применяется:
// Когда есть несколько родственных классов, которые отличаются поведением. Можно задать один основной класс,
// а разные варианты поведения вынести в отдельные классы и при необходимости их применять.
// Когда необходимо обеспечить выбор из нескольких вариантов алгоритмов, которые можно легко менять
// в зависимости от условий.
// Когда необходимо менять поведение объектов на стадии выполнения программы.
// Когда класс, применяющий определенную функциональность, ничего не должен знать о ее реализации.

// Преимущества паттерна стратегия:
// Изолирует код и данные алгоритмов от остальных классов.
// Уход от наследования к делегированию.
// Горячая замена алгоритмов на лету.

// Недостатки паттерна стратегия:
// Усложняет программу за счёт дополнительных классов.
// Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

package main

import "fmt"

// При разработке «In-Memory-Cache». Необхоимо очистить кэш.
// Эту функцию можно реализовать с помощью нескольких алгоритмов, 
// самые популярные среди них:
//    - LRU: убирает запись, которая использовалась наиболее давно
//    - FIFO: убирает запись, которая была создана раньше остальных
//    - LFU: убирает запись, которая использовалась наименее часто.

// Проблема заключается в том, чтобы отделить кэш от этих алгоритмов для возможности их замены «на ходу». 
// Помимо этого, класс кэша не должен меняться при добавлении нового алгоритма.

// Паттерн "Стратегия" предполагает создание семейства алгоритмов, каждый из которых имеет свой класс.
// Все классы применяют одинаковый интерфейс, что делает алгоритмы взаимозаменяемыми внутри этого семейства. 

type EvictionAlgo interface {
    evict(c *Cache)
}

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
    fmt.Println("Evicting by fifo strtegy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
    fmt.Println("Evicting by lru strtegy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
    fmt.Println("Evicting by lfu strtegy")
}

type Cache struct {
    storage      map[string]string
    evictionAlgo EvictionAlgo
    capacity     int
    maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
    storage := make(map[string]string)
    return &Cache{
        storage:      storage,
        evictionAlgo: e,
        capacity:     0,
        maxCapacity:  2,
    }
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
    c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
    if c.capacity == c.maxCapacity {
        c.evict()
    }
    c.capacity++
    c.storage[key] = value
}

func (c *Cache) get(key string) {
    delete(c.storage, key)
}

func (c *Cache) evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}

// клиентский код

func main() {
    lfu := &Lfu{}
    cache := initCache(lfu)

    cache.add("a", "1")
    cache.add("b", "2")

    cache.add("c", "3")

    lru := &Lru{}
    cache.setEvictionAlgo(lru)

    cache.add("d", "4")

    fifo := &Fifo{}
    cache.setEvictionAlgo(fifo)

    cache.add("e", "5")

}