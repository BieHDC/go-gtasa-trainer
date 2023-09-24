package main

import (
	"fmt"

	. "gtasamod/directcalls"
	. "gtasamod/types"
)

const mincar = 400
const maxcar = 611

var lastcar VehicleID
var currcar VehicleID = MODEL_HYDRA //initialise

var currspawn *CVehicle
var angle float32
var keeprotating bool

func menu_vehiclespawn_setup() {
	currspawn = nil
	lastcar = -1
	keeprotating = true
	AddFn(func(gi *GameInternals) bool {
		if currspawn != nil {
			Matrix_SetRotateZOnly(currspawn.Pos, angle)
			angle += 0.07
		}
		return keeprotating
	})
}

func menu_vehiclespawn(c uint32) bool {
	switch c {
	//currcar is zero because no button is pressed
	case ArrLeft:
		currcar = clampcar(mincar, currcar-1, maxcar, false)
	case ArrRight:
		currcar = clampcar(mincar, currcar+1, maxcar, true)
	case Esc:
		//remove spawned car
		if currspawn != nil {
			//delete car
			Vehicle_Destroy(currspawn)
			currspawn = nil
		}
		fallthrough
	case Enter:
		//keep spawned car
		if currspawn != nil {
			//turn indestructible off
			currspawn.VehicleFlags.CanBeDamagedToggle()
			currspawn.PhysFlags.CollidableToggle(true)
		}
		keeprotating = false
		Hud_SetHelpMessage("\x00", false, false, false) //clear
		return false
	}

	if currcar != lastcar {
		if currspawn != nil {
			//delete car
			Vehicle_Destroy(currspawn)
		}
		//spawn new car
		currspawn = Cheat_NewVehicle((int32)(currcar))
		if currspawn != nil {
			//disable destructability
			currspawn.VehicleFlags.CanBeDamagedToggle()
			currspawn.PhysFlags.CollidableToggle(false)
		}
		lastcar = currcar
	}

	Hud_SetHelpMessage(fmt.Sprintf("CarID: %d\nName: %s", currcar, currcar.String()), false, true, false)

	return true
}

func init() {
	NewMenu(F9, menu_vehiclespawn, menu_vehiclespawn_setup)
}

func clampcar(min, current, max VehicleID, up bool) VehicleID {
	switch current {
	case 449, 537, 538, 569, 570, 590:
		//crashy id, skip
		if up {
			current++
		} else {
			current--
		}
		return clampcar(min, current, max, up)
	default:
		break
	}
	//nether the min-est nor the max-est is crashy, so no issue
	if max < min {
		//sanity
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
