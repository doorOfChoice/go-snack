package main

import "github.com/nsf/termbox-go"

const (
	MOVE = iota + 1000
	END
)

type Command struct {
	command int
	key termbox.Key
}

func ListenKeyEvent(c chan Command) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch(ev.Key) {
			case termbox.KeyArrowUp: c <- Command{command : MOVE, key : ev.Key}
			case termbox.KeyArrowDown: c <- Command{command : MOVE, key : ev.Key}
			case termbox.KeyArrowRight: c <- Command{command : MOVE, key : ev.Key}
			case termbox.KeyArrowLeft: c <- Command{command : MOVE, key : ev.Key}
			case termbox.KeyCtrlC : c <- Command{command : END, key : ev.Key}
			}
		}
	}
}