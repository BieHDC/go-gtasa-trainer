package main

import (
	. "gtasamod/complextypes"
	. "gtasamod/directcalls"
	. "gtasamod/madresser"
	. "gtasamod/types"
)

var infhealh bool

func setInfHealth(gi *GameInternals) bool {
	gi.player.Ped.Health = 5000
	return infhealh
}

type menuentry struct {
	name    string
	command func()
}

var entries = []menuentry{
	{name: "Give Rocket Launcher", command: func() {
		AddFn(func(gi *GameInternals) bool {
			Ped_GiveWeapon(gi.player.Ped, RocketLauncher, 400)
			return false
		})
	}},
	{name: "Give Minigun", command: func() {
		AddFn(func(gi *GameInternals) bool {
			Ped_GiveWeapon(gi.player.Ped, Minigun, 8000)
			return false
		})

	}},
	{name: "Give Molotov Cocktail", command: func() {
		AddFn(func(gi *GameInternals) bool {
			Ped_GiveWeapon(gi.player.Ped, MolotovCocktail, 100)
			return false
		})
	}},
	{name: "Inf Health", command: func() {
		infhealh = !infhealh
		if infhealh {
			AddFn(setInfHealth)
		}
		Hud_SetHelpMessage("~r~Inf Health~s~~n~"+booltoonoff(infhealh), true, false, false)
	}},
	{name: "Toggle Infinite Run", command: func() {
		AddFn(func(gi *GameInternals) bool {
			FlipBoolInplace(&gi.player.InfiniteRun)
			return false
		})
	}},
	{name: "Toggle Fast Reload", command: func() {
		AddFn(func(gi *GameInternals) bool {
			FlipBoolInplace(&gi.player.FastReload)
			return false
		})
	}},
	{name: "Toggle Fireproof", command: func() {
		AddFn(func(gi *GameInternals) bool {
			FlipBoolInplace(&gi.player.Fireproof)
			return false
		})
	}},
	{name: "Toggle OnMission", command: func() {
		omptr := TypeAtAbsolute[byte](0xA49FC4)
		FlipBoolInplace(omptr)
		Hud_SetHelpMessage("~r~On Mission:~s~ "+booltoonoff(*omptr == 1), true, false, false)
	}},

	// Vehicle
	{name: "Toggle CanBeDamaged", command: func() {
		AddFn(func(gi *GameInternals) bool {
			cv := *gi.currentvehicle
			if cv != nil {
				Hud_SetHelpMessage("~r~CanBeDamaged~s~~n~"+booltoonoff((*gi.currentvehicle).VehicleFlags.CanBeDamagedToggle()), true, false, false)
			} else {
				Hud_SetHelpMessage("~r~CanBeDamaged~s~~n~Not in Vehicle", true, false, false)
			}
			return false
		})
	}},
	{name: "Toggle TyresDontBurst", command: func() {
		AddFn(func(gi *GameInternals) bool {
			cv := *gi.currentvehicle
			if cv != nil {
				Hud_SetHelpMessage("~r~TyresDontBurst~s~~n~"+booltoonoff((*gi.currentvehicle).VehicleFlags.TyresDontBurstToggle()), true, false, false)
			} else {
				Hud_SetHelpMessage("~r~TyresDontBurst~s~~n~Not in Vehicle", true, false, false)
			}
			return false
		})
	}},
	{name: "Toggle PetrolTankIsWeakPoint", command: func() {
		AddFn(func(gi *GameInternals) bool {
			cv := *gi.currentvehicle
			if cv != nil {
				Hud_SetHelpMessage("~r~PetrolTankIsWeakPoint~s~~n~"+booltoonoff((*gi.currentvehicle).VehicleFlags.PetrolTankIsWeakPointToggle()), true, false, false)
			} else {
				Hud_SetHelpMessage("~r~PetrolTankIsWeakPoint~s~~n~Not in Vehicle", true, false, false)
			}
			return false
		})
	}},
}

var currententry int
var maxentry = len(entries) - 1

func menu_quick(c uint32) bool {
	switch c {
	case Esc:
		// we cancelled our menu
		Hud_SetHelpMessage("\x00", false, false, false) //clear
		return false
	case ArrUp:
		currententry = clamp(0, currententry-1, maxentry)
	case ArrDown:
		currententry = clamp(0, currententry+1, maxentry)
	case Enter:
		Hud_SetHelpMessage("\x00", false, false, false) //clear
		entries[currententry].command()
		return false
	}

	//draw
	var display string
	for i, e := range entries {
		if i == currententry {
			display += ">~p~"
		} else {
			display += " "
		}
		display += e.name
		display += "~s~~n~" //https://gtamods.com/wiki/GXT#Tokens
		//23 is around the width limit
	}
	display = display[:len(display)-6] //cut the last newline and text reset
	Hud_SetHelpMessage(display, false, true, false)

	return true
}

func clamp(min, current, max int) int {
	if max < min {
		// this happens when the len(list) is empty and it "underflows"
		// since this is a custom thing, we return the min and the other
		// code will do the rest since we have nothing to display
		return min
	}
	if current < min {
		return min
	}
	if current > max {
		return max
	}
	return current
}

func booltoonoff(b bool) string {
	if b {
		return "On"
	} else {
		return "Off"
	}
}

func init() {
	NewMenu(F5, menu_quick, nil)
}
