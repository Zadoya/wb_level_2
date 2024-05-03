package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func scanStdin(stdin chan<- string, stop context.CancelFunc) {
	defer stop()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		stdin <- (scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func scanConn(conn net.Conn, connChan chan<- string, stop context.CancelFunc) {
	defer stop()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		connChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeData(conn net.Conn, connChan, stdin <-chan string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-connChan:
			fmt.Println(data)
		case data := <-stdin:
			conn.Write([]byte(data))
		}
	}
}

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "10s", "timeout for connecting to the server")
	flag.StringVar(&timeout, "t", "10s", "timeout for connecting to the server (shorthand)")

	flag.Parse()

	var address string

	switch len(flag.Args()) {
	case 0:
		log.Fatal("host and port have to be specified")
	case 1:
		log.Fatal("host or port is not specified")
	case 2:
		address = net.JoinHostPort(flag.Args()[0], flag.Args()[1])
	default:
		log.Fatalf("to many arguments %d, want 2: host and port", len(flag.Args()))
	}

	timeDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("wrong timeout value: %s", err.Error())
	}

	conn, err := net.DialTimeout("tcp", address, timeDuration)
	if err != nil {
		log.Fatalf("problems with connection: %s", err.Error())
	}
	defer conn.Close()

	connChan := make(chan string)
	stdin := make(chan string)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	go scanStdin(stdin, stop)
	go scanConn(conn, connChan, stop)

	wait := make(chan struct{})
	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		writeData(conn, connChan, stdin, ctx)
	}()
	<-wait
}
