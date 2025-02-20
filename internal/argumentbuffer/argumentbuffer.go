package argumentbuffer

import (
	"fmt"
	"sync"
)

type ArgumentBuff struct {
	command   string
	arguments []string
	mu        sync.Mutex
}

func NewArgumentBuff() *ArgumentBuff {
	return &ArgumentBuff{
		command:   "",
		arguments: []string{},
		mu:        sync.Mutex{},
	}
}

func (b *ArgumentBuff) Set(args []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(args) == 1 {
		b.command = args[0]
	} else if len(args) > 1 {
		b.command = args[0]
		b.arguments = args[1 : ]
	}
}

func (b *ArgumentBuff) Get() (command string, args []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.command, b.arguments
}

func (b *ArgumentBuff) GetArgs() (args []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.arguments
}

func (b *ArgumentBuff) PrintArgs() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.arguments) == 0 {
		fmt.Println("No arguments")
	}
	for _, arg := range b.arguments {
		fmt.Println(arg)
	}
}
