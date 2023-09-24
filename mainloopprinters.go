package main

import (
	"fmt"
	"log"
	"time"

	. "gtasamod/directcalls"
	. "gtasamod/types"
)

func timedprinter() GameFunc {
	lasttime := time.Now()

	return func(gi *GameInternals) bool {
		if time.Since(lasttime) > time.Second*10 {
			lasttime = time.Now()
			log.Printf("Weapons:\n")
			for weapon := range gi.player.Ped.Weapons {
				log.Printf("\t" + gi.player.Ped.Weapons[weapon].String() + "\n")
			}
			log.Printf("Money:%d\n", gi.player.Money)
			log.Printf("MoneyDisplayed:%d\n", gi.player.MoneyDisplayed)
			log.Printf("Hours not eaten:%d\n", gi.player.NumHoursNotEaten)
			log.Printf("TryingToExitVehicle:(%d)\n", gi.player.TryingToExitVehicle)
			log.Printf("LastTargetVehicle:(0x%x)\n", gi.player.LastTargetVehicle)

			log.Printf("WheelieNumCounter:(%d)\n", gi.player.WheelieNumCounter)
			log.Printf("WheelingDistanceCounter:(%.2f)\n", gi.player.WheelingDistanceCounter)

			log.Printf("TaxiTimer:%d\n", gi.player.TaxiTimer) //seems to be ok?

			log.Printf("Position:(%s)\n", gi.player.Ped.Pos.String())

			log.Printf("IsStanding:%t", gi.player.Ped.PlayerState.IsStanding())
			log.Printf("InVehicle:%t", gi.player.Ped.PlayerState.InVehicle())
			log.Printf("IsFiringWeapon:%t", gi.player.Ped.PlayerState.IsFiringWeapon())

			cw := *gi.currentvehicle
			if cw != nil {
				log.Println("Vehicle:")
				log.Printf("IsHandbrakeOn:%t", cw.VehicleFlags.IsHandbrakeOn())
				log.Printf("ComedyControls:%t", cw.VehicleFlags.ComedyControls())
				log.Printf("TakeLessDamage:%t", cw.VehicleFlags.TakeLessDamage())
				log.Printf("HasBeenOwnedByPlayer:%t", cw.VehicleFlags.HasBeenOwnedByPlayer())
				log.Printf("CanBeDamaged:%t", cw.VehicleFlags.CanBeDamaged())
				log.Printf("VehicleColProcessed:%t", cw.VehicleFlags.VehicleColProcessed())
				log.Printf("IsDrowning:%t", cw.VehicleFlags.IsDrowning())
				log.Printf("TyresDontBurst:%t", cw.VehicleFlags.TyresDontBurst())
				log.Printf("EngineBroken:%t", cw.VehicleFlags.EngineBroken())
				log.Printf("WinchCanPickMeUp:%t", cw.VehicleFlags.WinchCanPickMeUp())
				log.Printf("SirenOrAlarm:%t", cw.VehicleFlags.SirenOrAlarm())
				log.Printf("MadDriver:%t", cw.VehicleFlags.MadDriver())
				log.Printf("PetrolTankIsWeakPoint:%t", cw.VehicleFlags.PetrolTankIsWeakPoint())
				log.Printf("HasBeenResprayed:%t", cw.VehicleFlags.HasBeenResprayed())

			}

			GetVehiclePool().PrintAllVehicles()
		}
		return true
	}
}

var _ = func() struct{} { AddFn(timedprinter()); return struct{}{} }()

func havoker() GameFunc {
	var LastHavoc uint32

	return func(gi *GameInternals) bool {
		currenthavoc := gi.player.Havoc
		if currenthavoc != LastHavoc {
			Messages_AddBigMessageQ(fmt.Sprintf("You caused %d Havoc", currenthavoc))
			LastHavoc = currenthavoc
		}
		return true
	}
}

var _ = func() struct{} { AddFn(havoker()); return struct{}{} }()

func feeter() GameFunc {
	var oldVehicle *CVehicle
	return func(gi *GameInternals) bool {
		cw := *gi.currentvehicle
		if cw != oldVehicle {
			oldVehicle = cw
			if cw == nil {
				Messages_AddBigMessageQ("On Foot")
			} else {
				Messages_AddBigMessageQ(fmt.Sprintf("In Vehicle (%s)", VehicleID(cw.ModelID).String()))
				log.Println(cw.String())
			}
		}
		return true
	}
}

// we do a little trolling
var _ = func() struct{} { AddFn(feeter()); return struct{}{} }()

func healther() GameFunc {
	var OldHealth float32
	return func(gi *GameInternals) bool {
		if health := gi.player.Ped.Health; health != OldHealth {
			curhealthstring := fmt.Sprintf("Health is: %.2f \n Health was: %.2f", health, OldHealth)
			//log.Printf(curhealthstring)
			Messages_AddBigMessageQ(curhealthstring)
			OldHealth = health
		}
		return true
	}
}

var _ = func() struct{} { AddFn(healther()); return struct{}{} }()
