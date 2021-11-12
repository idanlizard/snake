package game

import (
	term "github.com/nsf/termbox-go"
	"sync"
)

type KeyBoardEventListener struct {
	output chan term.Key
	stop   chan struct{}
	once   sync.Once
}

func NewKeyboardListener() *KeyBoardEventListener {
	return &KeyBoardEventListener{
		output: make(chan term.Key, 10),
		stop:   make(chan struct{}, 1),
	}
}

func (this *KeyBoardEventListener) listen() {
	if err := term.Init(); err != nil {
		panic(err)
	}

	term.SetInputMode(term.InputEsc)
	for {
		select {
		case <-this.stop:
			close(this.output)
			return
		default:
			ev := term.PollEvent()
			if ev.Type != term.EventKey {
				continue
			}

			this.output <- ev.Key
		}
	}
}

func (this *KeyBoardEventListener) Listen() <-chan term.Key {
	this.once.Do(func() {
		go this.listen()
	})

	return this.output
}

func (this *KeyBoardEventListener) Stop() {
	this.stop <- struct{}{}
	close(this.stop)
	term.Close()
}
