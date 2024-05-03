package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	var host, port, protocol string	
	flag.StringVar(&host, "host", "localhost", "Хост")
	flag.StringVar(&port, "port", "8080", "Порт")
	flag.StringVar(&protocol, "protocol", "tcp", "Протокол (tcp/udp)")
	flag.StringVar(&host, "h", "localhost", "Хост")
	flag.StringVar(&port, "p", "8080", "Порт")
	flag.StringVar(&protocol, "prot", "tcp", "Протокол (tcp/udp)")
	flag.Parse()

	// Формируем адрес для подключения
	addr := fmt.Sprintf("%s:%s", host, port)

	// Подключаемся к серверу
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	// Копируем данные из stdin в соединение
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Println("Ошибка при отправке данных:", err)
		return
	}
}