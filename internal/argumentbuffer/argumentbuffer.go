package argumentbuffer

import (
	"sync"
)

type ArgumentBuff struct {
	command string
	arguments []string
	mu sync.Mutex
}

func NewArgumentBuff() *ArgumentBuff {
	return &ArgumentBuff{
		command: "",
		arguments: []string{},
		mu: sync.Mutex{},
	}
}


func (b *ArgumentBuff) Set(args []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.command = args[0]
	b.arguments = args[1:len(args)-1]
}

func (b *ArgumentBuff) Get() (command string, args []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.command, b.arguments
}