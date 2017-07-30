package main

import (
	"github.com/nsf/termbox-go"

)

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
	//上一次按键记录
	var record termbox.Key 
	//本次按键命令
	var command Command
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch(ev.Key) {
			case termbox.KeyArrowUp: command = Command{command : MOVE, key : ev.Key}
			case termbox.KeyArrowDown: command = Command{command : MOVE, key : ev.Key}
			case termbox.KeyArrowRight: command = Command{command : MOVE, key : ev.Key}
			case termbox.KeyArrowLeft: command = Command{command : MOVE, key : ev.Key}
			case termbox.KeyCtrlC : command = Command{command : END, key : ev.Key}
			}
		}

		//判断是否一直重复按一个键
		if(command.key != record) {
			record = command.key
			c <- command
		}
	}
}