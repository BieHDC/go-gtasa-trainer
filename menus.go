package main

import (
	"fmt"
	"sync"
)

type menu struct {
	initfunc func()
	handler  func(uint32) bool
}

var menus = map[uint32]menu{}
var menusmutex sync.Mutex

// key: the key to open the menu
// handler: the function executed on input.
//
//	return "true" if the menu is still active/open
//	return "false" if the menu has been closed
//
// initfunc: called every time the mneu is opened
//
//	used to set initial state
func NewMenu(key uint32, handler func(uint32) bool, initfunc func()) {
	menusmutex.Lock()
	defer menusmutex.Unlock()
	_, exists := menus[key]
	if exists {
		panic(fmt.Sprintf("Key already in use: %d", key))
	}

	if handler == nil {
		panic(fmt.Sprintf("menu with key %d has no handler declared", key))
	}

	menus[key] = menu{initfunc: initfunc, handler: handler}
}

var activemenu uint32

// returns wether we handled it (true), or if it should be passed on to the game (false)
// todo:
//
//	ptr to special keys pressed?
//	we would need that for key combos like shift+f7
func Menuhandler(c uint32, pinputting *bool) bool {
	//log.Println("rune:", c)

	// we are not in our menu but we maybe want to be
	if !*pinputting {
		menu, exists := menus[c]
		if !exists {
			// we are not in any menu and we didnt activate any ether
			return false //returning false is only ever relevant here
		}
		// the menu exists and we just opened it
		activemenu = c
		// and if there is some initial state to init, do it now
		if menu.initfunc != nil {
			menu.initfunc()
		}
		// we are now inputting
		*pinputting = true
	}

	menu, exists := menus[activemenu]
	if !exists {
		panic("should not happen")
	}
	*pinputting = menu.handler(c)

	return true //we have handled the input
}

// Same as enum in C, keep in sync!
const (
	Backspace uint32 = iota + 1000
	Tab
	Enter
	Shift
	ShiftL
	ShiftR
	Ctrl
	CtrlL
	CtrlR
	Alt
	AltL
	AltR
	Pause
	Capslock
	Esc
	PageUp
	PageDown
	End
	Home
	ArrLeft
	ArrUp
	ArrRight
	ArrDown
	Insert
	Delete
	Numpad0
	Numpad1
	Numpad2
	Numpad3
	Numpad4
	Numpad5
	Numpad6
	Numpad7
	Numpad8
	Numpad9
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
)
