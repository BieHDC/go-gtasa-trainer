package main

import (
	. "gtasamod/directcalls"
	. "gtasamod/types"
)

func pedorvehiclepos(gi *GameInternals) *CVector {
	if *gi.currentvehicle != nil {
		return &(*gi.currentvehicle).Pos.Position
	} else {
		return &gi.player.Ped.Pos.Position
	}
}

func pedorvehiclephysicsflags(gi *GameInternals) *PhysicsFlags {
	if *gi.currentvehicle != nil {
		return &(*gi.currentvehicle).PhysFlags
	} else {
		return &gi.player.Ped.PhysFlags
	}
}

func pedorvehiclespeed(gi *GameInternals) *CVector {
	if *gi.currentvehicle != nil {
		return &(*gi.currentvehicle).MoveSpeed
	} else {
		return &gi.player.Ped.MoveSpeed
	}
}

func menu_flyhack_setup() {
	AddFn(func(gi *GameInternals) bool {
		pedorvehiclephysicsflags(gi).ApplyGravityToggle(false)
		return false
	})
	Hud_SetHelpMessage("Flymode on.\nPress Esc to quit.", true, false, false)
}

func menu_flyhack(c uint32) bool {
	var t CVector
	switch c {
	case Numpad5:
		t.Z += 2
	case Numpad0:
		t.Z -= 2
	case Numpad8:
		t.Y += 2
	case Numpad2:
		t.Y -= 2
	case Numpad6:
		t.X += 2
	case Numpad4:
		t.X -= 2
	case Esc:
		AddFn(func(gi *GameInternals) bool {
			pedorvehiclephysicsflags(gi).ApplyGravityToggle(true)
			return false
		})
		Hud_SetHelpMessage("Flymode off", true, false, false)
		return false
	}

	AddFn(func(gi *GameInternals) bool {
		*pedorvehiclespeed(gi) = CVector{} //zero the speed every time since we can he hit by something
		dstpos := pedorvehiclepos(gi)
		dstpos.X += t.X
		dstpos.Y += t.Y
		dstpos.Z += t.Z
		return false
	})

	return true
}

func init() {
	NewMenu(F7, menu_flyhack, menu_flyhack_setup)
}
