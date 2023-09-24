package types

import (
	"fmt"
	"strings"
	"unsafe"

	. "gtasamod/structoffsetvalidator"
)

// Can Map Cleanly to the data
type CWeapon struct { //Total 28 bytes
	Type            WeaponType  `offset:"0x0"`
	State           WeaponState `offset:"0x4"`
	AmmoInClip      uint32      `offset:"0x8"`
	TotalAmmo       uint32      `offset:"0xC"`
	TimeForNextShot uint32      `offset:"0x10"` //not sure what this means? recoil time?
	pad1, pad2      uint32
}

func init() {
	var cw CWeapon
	err := ValidateOffsets(&cw)
	if err != nil {
		panic(err)
	}

	sz := unsafe.Sizeof(CWeapon{})
	if sz != 28 {
		panic(fmt.Sprintf("CWeapon size is wrong, expected %d, got %d\n", 28, sz))
	}
}

func (cw *CWeapon) String() string {
	var out strings.Builder

	out.WriteString("Weapon: ")
	out.WriteString(fmt.Sprintf("Type:(%s) ", cw.Type.String()))
	out.WriteString(fmt.Sprintf("State:(%s) ", cw.State.String())) //Useless output
	out.WriteString(fmt.Sprintf("Clip:(%d) ", cw.AmmoInClip))
	out.WriteString(fmt.Sprintf("TotalAmmo:(%d) ", cw.TotalAmmo))
	//out.WriteString(fmt.Sprintf("Recoil?:(%d)", cw.TimeForNextShot)) //Useless output

	return out.String()
}

type WeaponState uint32

const (
	WEAPONSTATE_READY WeaponState = iota
	WEAPONSTATE_FIRING
	WEAPONSTATE_RELOADING
	WEAPONSTATE_OUT_OF_AMMO
	WEAPONSTATE_MELEE_MADECONTACT
)

func (ws WeaponState) String() string {
	switch ws {
	case WEAPONSTATE_READY:
		return "WEAPONSTATE_READY"
	case WEAPONSTATE_FIRING:
		return "WEAPONSTATE_FIRING"
	case WEAPONSTATE_RELOADING:
		return "WEAPONSTATE_RELOADING"
	case WEAPONSTATE_OUT_OF_AMMO:
		return "WEAPONSTATE_OUT_OF_AMMO"
	case WEAPONSTATE_MELEE_MADECONTACT:
		return "WEAPONSTATE_MELEE_MADECONTACT"
	default:
		return "Unknwon WeaponState"
	}
}

type WeaponType uint32

func (wt WeaponType) AsInt() string {
	return fmt.Sprintf("%d", wt)
}

func (wt WeaponType) String() string {
	switch wt {
	case Fist:
		return "Fist"
	case BrassKnuckles:
		return "BrassKnuckles"
	case GoldClub:
		return "GoldClub"
	case Nitestick:
		return "Nitestick"
	case Knife:
		return "Knife"
	case BaseballBat:
		return "BaseballBat"
	case Shovel:
		return "Shovel"
	case PoolCue:
		return "PoolCue"
	case Katana:
		return "Katana"
	case Chainsaw:
		return "Chainsaw"
	case Cane:
		return "Cane"
	case Pistol:
		return "Pistol"
	case SliencePistol:
		return "SliencePistol"
	case DesertEagle:
		return "DesertEagle"
	case Shotgun:
		return "Shotgun"
	case SawnOffShotgun:
		return "SawnOffShotgun"
	case SPAS12:
		return "SPAS12"
	case MicroUzi:
		return "MicroUzi"
	case MP5:
		return "MP5"
	case TEC9:
		return "TEC9"
	case AK47:
		return "AK47"
	case M4:
		return "M4"
	case CountryRifle:
		return "CountryRifle"
	case SniperRifle:
		return "SniperRifle"
	case RocketLauncher:
		return "RocketLauncher"
	case HeatSeekingRPG:
		return "HeatSeekingRPG"
	case FlameThrower:
		return "FlameThrower"
	case Minigun:
		return "Minigun"
	case Grenade:
		return "Grenade"
	case MolotovCocktail:
		return "MolotovCocktail"
	case RemoteExplosives:
		return "RemoteExplosives"
	case FireExtinguisher:
		return "FireExtinguisher"
	case Camera:
		return "Camera"
	case Flowers:
		return "Flowers"
	case Dildo1:
		return "Dildo1"
	case Dildo2:
		return "Dildo2"
	case Vibe1:
		return "Vibe1"
	case Vibe2:
		return "Vibe2"
	case NVGoggles:
		return "NVGoggles"
	case IRGoggles:
		return "IRGoggles"
	case Parachute:
		return "Parachute"
	case Detonator:
		return "Detonator"
	case NormalRockets:
		return "NormalRockets"
	case HeatseekingRockets:
		return "HeatseekingRockets"
	case Flares:
		return "Flares"
	case Teargas:
		return "Teargas"
	case Spraycan:
		return "Spraycan"
	default:
		return "Unknown Weapon"
	}
}

var ValidWeaponIDs = []WeaponType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 22, 23, 24, 25, 26, 27, 28, 29, 32, 30, 31, 33, 34, 35, 36, 37, 38, 16, 17, 18, 39, 41, 42, 43, 10, 11, 12, 13, 14, 15, 44, 45, 46}

func (wt WeaponType) Valid() bool {
	if wt > 46 {
		return false
	}
	for _, w := range ValidWeaponIDs {
		if wt == w {
			return true
		}
	}
	return false
}

const (
	Fist          WeaponType = 0 // Slot 0 -> No Weapon
	BrassKnuckles WeaponType = 1 // Slot 0 -> No Weapon

	GoldClub    WeaponType = 2 // Slot 1 -> Melee
	Nitestick   WeaponType = 3 // Slot 1 -> Melee
	Knife       WeaponType = 4 // Slot 1 -> Melee
	BaseballBat WeaponType = 5 // Slot 1 -> Melee
	Shovel      WeaponType = 6 // Slot 1 -> Melee
	PoolCue     WeaponType = 7 // Slot 1 -> Melee
	Katana      WeaponType = 8 // Slot 1 -> Melee
	Chainsaw    WeaponType = 9 // Slot 1 -> Melee

	Pistol        WeaponType = 22 // Slot 2 -> Handguns
	SliencePistol WeaponType = 23 // Slot 2 -> Handguns
	DesertEagle   WeaponType = 24 // Slot 2 -> Handguns

	Shotgun        WeaponType = 25 // Slot 3 -> Shotguns
	SawnOffShotgun WeaponType = 26 // Slot 3 -> Shotguns
	SPAS12         WeaponType = 27 // Slot 3 -> Shotguns

	MicroUzi WeaponType = 28 // Slot 4 -> Sub Machineguns
	MP5      WeaponType = 29 // Slot 4 -> Sub Machineguns
	TEC9     WeaponType = 32 // Slot 4 -> Sub Machineguns

	AK47 WeaponType = 30 // Slot 5 -> Machineguns
	M4   WeaponType = 31 // Slot 5 -> Machineguns

	CountryRifle WeaponType = 33 // Slot 6 -> Rifles
	SniperRifle  WeaponType = 34 // Slot 6 -> Rifles

	RocketLauncher WeaponType = 35 // Slot 7 -> Heavy Weapons
	HeatSeekingRPG WeaponType = 36 // Slot 7 -> Heavy Weapons
	FlameThrower   WeaponType = 37 // Slot 7 -> Heavy Weapons
	Minigun        WeaponType = 38 // Slot 7 -> Heavy Weapons

	Grenade          WeaponType = 16 // Slot 8 -> Projectiles
	Teargas          WeaponType = 17 // Slot 8 -> Projectiles
	MolotovCocktail  WeaponType = 18 // Slot 8 -> Projectiles
	RemoteExplosives WeaponType = 39 // Slot 8 -> Projectiles

	Spraycan         WeaponType = 41 // Slot 9 -> Special
	FireExtinguisher WeaponType = 42 // Slot 9 -> Special
	Camera           WeaponType = 43 // Slot 9 -> Special

	Dildo1  WeaponType = 10 // Slot 10 -> Gifts
	Dildo2  WeaponType = 11 // Slot 10 -> Gifts
	Vibe1   WeaponType = 12 // Slot 10 -> Gifts
	Vibe2   WeaponType = 13 // Slot 10 -> Gifts
	Flowers WeaponType = 14 // Slot 10 -> Gifts
	Cane    WeaponType = 15 // Slot 10 -> Gifts

	NVGoggles WeaponType = 44 // Slot 11 -> Special 2
	IRGoggles WeaponType = 45 // Slot 11 -> Special 2
	Parachute WeaponType = 46 // Slot 11 -> Special 2

	Detonator WeaponType = 40 // Slot 12 -> Detonators

	NormalRockets      WeaponType = 19 // Slot Else - Fired from hunter / hydra / missile launcher
	HeatseekingRockets WeaponType = 20 // Slot Else - Fired from hunter / hydra / missile launcher
	Flares             WeaponType = 58 // Slot Else - Fired from hunter / hydra / missile launcher
)

type WeaponSlot int32

func (w WeaponType) Slot() WeaponSlot {
	switch w {
	case Fist, BrassKnuckles:
		return 0
	case GoldClub, Nitestick, Knife, BaseballBat, Shovel, PoolCue, Katana, Chainsaw:
		return 1
	case Pistol, SliencePistol, DesertEagle:
		return 2
	case Shotgun, SawnOffShotgun, SPAS12:
		return 3
	case MicroUzi, MP5, TEC9:
		return 4
	case AK47, M4:
		return 5
	case CountryRifle, SniperRifle:
		return 6
	case RocketLauncher, HeatSeekingRPG, FlameThrower, Minigun:
		return 7
	case Grenade, MolotovCocktail, RemoteExplosives, Teargas:
		return 8
	case FireExtinguisher, Camera, Spraycan:
		return 9
	case Dildo1, Dildo2, Vibe1, Vibe2, Flowers, Cane:
		return 10
	case NVGoggles, IRGoggles, Parachute:
		return 11
	case Detonator:
		return 12
	case NormalRockets, HeatseekingRockets, Flares:
		return 13

	default:
		return -1
	}
}
