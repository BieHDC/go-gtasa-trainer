package main

import (
	"fmt"
	"strconv"
	"strings"

	. "gtasamod/directcalls"
	. "gtasamod/madresser"
)

var typedtext strings.Builder

func TrimBuilderNumRight(sb *strings.Builder, amount int) {
	if amount < 0 {
		return //nopeium
	}
	ll := sb.Len()
	if ll < amount {
		sb.Reset() //delete everything
		return
	}
	current := sb.String()[:ll-amount]
	sb.Reset()
	sb.WriteString(current)
}

func menu_command(c uint32) bool {
	switch c {
	case Enter:
		processCommand(typedtext.String())
		fallthrough
	case Esc:
		//cancel inputting
		Hud_SetHelpMessage("\x00", false, false, false) //clear
		typedtext.Reset()
		return false

	case Backspace:
		TrimBuilderNumRight(&typedtext, 1)

	default:
		// Printable rune
		if c >= 32 && c <= 255 {
			typedtext.WriteRune(rune(c))
		}
	}

	//log.Println(typedtext.String())
	Hud_SetHelpMessage(">"+typedtext.String(), false, true, false)

	return true
}

var gamespeed = TypeAtAbsolute[float32](0xB7CB64) //0xB7CB64 – [float] Game speed in percent
func processCommand(cmd string) {
	tokens := strings.Split(strings.TrimSpace(cmd), " ")
	tokenlen := len(tokens)

	if tokens[0] == "" {
		return
	}

	if tokens[0] == "wanted" {
		if tokenlen <= 1 {
			AddFn(func(gi *GameInternals) bool {
				Messages_AddMessageQ(fmt.Sprintf("Current Wanted Level: %d", gi.wanted.LevelAsStars), 3000, 2, false)
				return false
			})
			return
		}
		newlevel, err := strconv.Atoi(tokens[1])
		if err != nil {
			AddFn(func(_ *GameInternals) bool {
				Messages_AddMessageQ("Not a number", 3000, 2, false)
				return false
			})
			return
		}
		if newlevel < 0 || newlevel > 6 {
			AddFn(func(_ *GameInternals) bool {
				Messages_AddMessageQ(fmt.Sprintf("Level out of range: %d", newlevel), 3000, 2, false)
				return false
			})
			return
		}

		AddFn(func(gi *GameInternals) bool {
			PlayerPed_SetWantedLevel(gi.player.Ped, newlevel)
			Messages_AddMessageQ(fmt.Sprintf("Set wanted level to %d", newlevel), 3000, 2, false)
			return false
		})
		return
	}

	if tokens[0] == "speed" {
		if tokenlen <= 1 {
			AddFn(func(_ *GameInternals) bool {
				Messages_AddMessageQ(fmt.Sprintf("Current Gamespeed is %d", int(*gamespeed*100)), 3000, 2, false)
				return false
			})
			return
		}
		if tokens[1] == "reset" {
			*gamespeed = 1.0
			AddFn(func(gi *GameInternals) bool {
				Messages_AddMessageQ(fmt.Sprintf("Gamespeed set to %d", int(*gamespeed*100)), 3000, 2, false)
				return false
			})
			return
		}
		newspeed, err := strconv.Atoi(tokens[1])
		if err != nil {
			AddFn(func(_ *GameInternals) bool {
				Messages_AddMessageQ("Not a number", 3000, 2, false)
				return false
			})
			return
		}

		*gamespeed = float32(newspeed) / 100

		AddFn(func(gi *GameInternals) bool {
			Messages_AddMessageQ(fmt.Sprintf("Gamespeed set to %d", int(*gamespeed*100)), 3000, 2, false)
			return false
		})
		return
	}

	if tokens[0] == "list" {
		AddFn(func(gi *GameInternals) bool {
			Hud_SetHelpMessage("Available Commands:\nwanted\nspeed", false, false, false)
			return false
		})
		return
	}

	AddFn(func(_ *GameInternals) bool {
		Messages_AddMessageQ("Unknown Command", 3000, 2, false)
		return false
	})
}

func init() {
	NewMenu(F6, menu_command, nil)
}
