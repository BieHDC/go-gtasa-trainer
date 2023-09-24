package main

import (
	"fmt"
	"log"

	. "gtasamod/directcalls"
	. "gtasamod/types"
)

var selectedcar int

var keepgoing2 bool
var currveh *CVehicle

func contteleport(ped *CPed) {
	if currveh != nil {
		Ped_Teleport(ped, currveh.GetPosition(), false, 2)
	}
}

func menu_vehiclemap_setup() {
	currveh = nil
	selectedcar = 0
	keepgoing2 = true
	AddFn(func(gi *GameInternals) bool {
		contteleport(gi.player.Ped)
		return keepgoing2
	})
}

func menu_vehiclemap(c uint32) bool {
	switch c {
	case ArrUp:
		selectedcar -= 1
	case ArrDown:
		selectedcar += 1
	case Esc, Enter:
		keepgoing2 = false
		Hud_SetHelpMessage("\x00", false, false, false) //clear

		if currveh != nil {
			log.Println(currveh.String())
		}
		return false
	}

	var display string
	vehiclepool := GetVehiclePool()
	vehicles := vehiclepool.EveryVehicle()
	selectedcar = clamp(0, selectedcar, len(vehicles)-1)
	carid := int32(-1)
	for i, vhid := range vehicles {
		var sel string
		if i == selectedcar {
			sel = "> "
			carid = vhid
		} else {
			sel = "  "
		}

		vh, _ := vehiclepool.GetAt(vhid)
		display += sel + fmt.Sprintf("%d: ID:%d\n", i, vh.ModelID)
	}
	if display == "" {
		display = "No Vehicles"
	}

	currveh, _ = vehiclepool.GetAt(carid)

	Hud_SetHelpMessage(display, false, true, false)

	return true
}

func init() {
	NewMenu(F8, menu_vehiclemap, menu_vehiclemap_setup)
}
