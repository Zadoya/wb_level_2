package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Shell struct {
	In	io.Reader
	Out	io.Writer
}

func NewShell() {
	shell := &Shell{
		In: os.Stdin,
		Out: os.Stdout,
	}
	shell.Start()
}

func (s *Shell) Start() {
	scanner := bufio.NewScanner(s.In)
	for scanner.Scan() {
		text := scanner.Text()
		if cmds := strings.Split(text, "|"); len(cmds) > 1 {
			if err := s.ProcessPipiline(cmds); err != nil {
				fmt.Fprintln(s.Out, err.Error())
			}
		} else {
			if err := s.ProcessCommand(cmds[0]); err != nil {
				fmt.Fprintln(s.Out, err.Error())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(s.Out, "Ошибка при считывании: ", err.Error())
	}
}

func (s *Shell) ProcessPipiline(cmds []string) error {
	return nil
}

func (s *Shell) ProcessCommand(cmd string) error {
	args := strings.Split(cmd, " ")
	switch args[0] {
	case "cd":
		s.cd(args[1:])
	case "pwd":
		s.pwd()
	case "echo":
	case "kill":
	case "ps":
	case "quit":
		fmt.Fprintln(s.Out, "quitting")
		os.Exit(0)
	}
	return nil
}

func (s *Shell) cd(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("cd: string not in pwd: %s\n", args)
	} else if len(args) == 1 {
		if err := os.Chdir(args[0]); err != nil {
			return err
		}
	} else {
		if err := os.Chdir(""); err != nil {
			return err 
		}
	}
	return nil
}

func (s *Shell) pwd() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err = fmt.Fprintln(s.Out, path); err != nil {
		return err
	}
	return nil	
}

func (s *Shell) echo(args []string) error {
	for i := range args {
		
	}
}

func main() {
	NewShell()
}
