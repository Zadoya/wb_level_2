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
	"strconv"
	"strings"
	"os/exec"

	ps "github.com/mitchellh/go-ps"
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
		fmt.Fprintln(s.Out, "Error of reading: ", err.Error())
	}
}

func (s *Shell) ProcessPipiline(cmds []string) error {
	return nil
}

func (s *Shell) ProcessCommand(cmd string) error {
	command, args, _ := strings.Cut(cmd, " ")

	switch command {
	case "cd":
		if err := s.cd(strings.Fields(args)); err != nil {
			fmt.Fprintln(s.Out, err)
		}
	case "pwd":
		if err := s.pwd(); err != nil {
			fmt.Fprintln(s.Out, err)
		}
	case "echo":
		if err := s.echo(strings.Split(args, "\"")); err != nil {
			fmt.Fprintln(s.Out, err)
		}
	case "kill":
		if err := s.kill(strings.Fields(args)); err != nil {
			fmt.Fprintln(s.Out, err)
		}
	case "ps":
		if err := s.ps(); err != nil {
			fmt.Fprintln(s.Out, err)
		}
	case "quit":
		fmt.Fprintln(s.Out, "quitting")
		os.Exit(0)
	default:
		// создание нового процесса и выполния команды в нем
		cmd := exec.Command(command, strings.Fields(args)...)
		cmd.Stdout = s.Out
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(s.Out, err)
		}
	}
	return nil
}

func (s *Shell) cd(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("cd: string not in pwd: %s", args)
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
	if len(args) % 2 == 0 {
		return fmt.Errorf("incorrect arguments")
	}
	for i := range args {
		if i % 2 == 1 {
			fmt.Fprintf(s.Out, "%s", args[i])
		} else if i == 0 {
			fmt.Fprint(s.Out, strings.Join(strings.Fields(args[i]), " "))
		} else {
			fmt.Fprint(s.Out," ", strings.Join(strings.Fields(args[i]), " "))
		}
	}
	return nil
}

func (s *Shell) kill(pid []string) []error {
	errs := make([]error, 0, len(pid))
	if len(pid) == 0 {
		return []error{fmt.Errorf("kill: not enough arguments")}
	} 
	for i := range pid {
		if pidInt, err := strconv.Atoi(pid[i]); err != nil {
			errs = append(errs, err)
		} else if proc, err := os.FindProcess(pidInt); err != nil {
			errs = append(errs, err)
		} else if err = proc.Kill(); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (s *Shell) ps() error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}
	fmt.Fprintf(s.Out, "%10s\t%10s\t%10s\n", "PID", "PPID", "EXEC")
	for i := range processList {
		if processList[i].PPid() > 1 {
			fmt.Fprintf(s.Out, "%10v\t%10v\t%10v\n", processList[i].Pid(), processList[i].PPid(), processList[i].Executable())
		}
	}
	return  nil
}
func main() {
	NewShell()
}
