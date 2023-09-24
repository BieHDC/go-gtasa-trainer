package types

import (
	"fmt"
	"unsafe"

	. "gtasamod/structoffsetvalidator"
)

//ref
// CurrentVehicle **CVehicle `address:"0xBA18FC"` //0=on foot, not 0 -> vehicle ptr

// size 0x190 400
type CPlayer struct {
	Ped                        *CPed        `offset:"0x0"`
	PlayerData                 *CPlayerData `offset:"0x4"`
	_                          [0xb0]byte   //padding
	Money                      uint32       `offset:"0xb8"` //100% correct
	MoneyDisplayed             uint32       `offset:"0xbc"`
	Collectables               uint32       `offset:"0xc0"` //+0xC0 = [dword] Collectables picked up
	TotalNumCollectales        uint32       `offset:"0xc4"` //+0xC4 = [dword] Total number collectables
	VehicleBumpTimer           uint32       `offset:"0xc8"` //+0xC8 = [dword] Last bump player vehicle timer
	TaxiTimer                  uint32       `offset:"0xcc"` //0xCC = [dword] Taxi timer
	_                          [0x5]byte    //padding
	TryingToExitVehicle        byte         `offset:"0xd5" options:"0=false,1=true"`
	_                          [0x2]byte    //padding
	LastTargetVehicle          *CVehicle    `offset:"0xd8"`
	PlayerState                byte         `offset:"0xdc"` //not specified what it is
	_                          [0x27]byte   //padding
	WheelieNumCounter          uint32       `offset:"0x104"`
	WheelingDistanceCounter    float32      `offset:"0x108"`
	_                          [0x34]byte   //padding
	Havoc                      uint32       `offset:"0x140"` //broken
	NumHoursNotEaten           uint16       `offset:"0x144"`
	_                          [0x6]byte    //padding
	InfiniteRun                byte         `offset:"0x14c" options:"0=false,1=true"`
	FastReload                 byte         `offset:"0x14d" options:"0=false,1=true"`
	Fireproof                  byte         `offset:"0x14e" options:"0=false,1=true"`
	MaxHealth                  byte         `offset:"0x14f" options:"0=false,1=true"`
	MaxArmour                  byte         `offset:"0x150" options:"0=false,1=true"`
	JailFreeAndKeepWeapons     byte         `offset:"0x151" options:"0=false,1=true"`
	HospitalFreeAndKeepWeapons byte         `offset:"0x152" options:"0=false,1=true"`
	CanDriveBy                 byte         `offset:"0x153" options:"0=false,1=true"`
	_                          [0x4]byte    //padding
	CrosshairActivated         uint32       `offset:"0x158"`
	CrosshairTarget            CVector2D    `offset:"0x15c"`
	SkinName                   [32]byte     `offset:"0x164"`
	_                          [0xc]byte    //padding
}

// Validate CPlayer offsets
func init() {
	var player CPlayer
	err := ValidateOffsets(&player)
	if err != nil {
		panic(err)
	}
	sz := unsafe.Sizeof(player)
	if sz != 400 {
		panic(fmt.Sprintf("CPlayer size is wrong, expected %d, got %d\n", 400, sz))
	}
}

// WantedPtr **CWanted `address:"0xB7CD9C"` //but this works fine
// todo: stringer
type CPlayerData struct {
	_                   [0x4]byte //padding //Wanted: *CWanted: `offset:"0x0"` //broken and weird value, like it directly references the wanted level
	_                   [0x4]byte //padding
	CopArrestingYou     *CPed     `offset:"0x8"`
	_                   [0xC]byte //padding
	TimeCanRun          float32   `offset:"0x18"`
	MoveSpeed           float32   `offset:"0x1c"`
	_                   [0x4]byte //padding
	StandStillTimer     uint32    `offset:"0x24"`
	_                   [0x4]byte //padding
	AttackButtonCounter float32   `offset:"0x2c"`
	_                   [0x4]byte //padding
	PlayerFlags         uint32    `offset:"0x34"`
	/* PlayerFlags
	Bit 0 = Stopped Moving
	Bit 1 = Adrenaline
	Bit 2 = Have Target Selected
	Bit 3 = Free Aiming
	Bit 4 = Can Be Damaged
	Bit 5 = All Melee Attack Pts Blocked
	Bit 6 = Just Been Snacking
	Bit 7 = Require Handle Breath
	Bit 8 = Group Stuff Disabled
	Bit 9 = Group Always Follow
	Bit 10 = Group Never Follow
	Bit 11 = In Vehicle Dont Allow Weapon Change
	Bit 12 = Render Weapon
	*/
	_                            [0x8]byte  //padding
	Drunkness                    byte       `offset:"0x40"`
	_                            [0x1]byte  //padding
	DrugLevel                    byte       `offset:"0x42"`
	_                            [0x1]byte  //padding
	Breath                       float32    `offset:"0x44"`
	_                            [0x24]byte //padding
	LastTimeFiring               uint32     `offset:"0x6c"`
	_                            [0x14]byte //padding
	SprintDisabled               byte       `offset:"0x84" options:"0=false,1=true"`
	WeaponChangeDisabled         byte       `offset:"0x85" options:"0=false,1=true"`
	ForceInteriourLighting       byte       `offset:"0x86" options:"0=false,1=true"`
	_                            [0x7]byte  //padding
	WaterCoverPercent            byte       `offset:"0x8e"`
	_                            [0x9]byte  //padding
	LastHeatSeekingTarget        *CVehicle  `offset:"0x98"` //type is just a guess!
	ModelIndexOfLastBuildingShot uint32     `offset:"0x9c"`
}

// Validate CPlayerData offsets
func init() {
	var playerdata CPlayerData
	err := ValidateOffsets(&playerdata)
	if err != nil {
		panic(err)
	}
}

// 0x7C4 1988
type CPed struct {
	/*
		_                 [0x14]byte     //padding
		Pos               *CMatrix       `offset:"0x14"`
		_                 [0x17]byte     //padding
		InInteriour       InteriorState  `offset:"0x2F"`
		_                 [0x10]byte     //padding
		PhysFlags         PhysicsFlags  `offset:"0x40"`
		MoveSpeed         CVector        `offset:"0x44"`
	*/
	CPhysical         `offset:"0x0"`
	_                 [0x24]byte     //padding
	AnimState         AnimationState `offset:"0x15c"`
	_                 [0x30f]byte    //padding
	PlayerState       PlayerState    `offset:"0x46c"` //[16]byte
	_                 [0xb4]byte     //padding
	State             State          `offset:"0x530"` //ePedState           m_nPedState;
	RunningState      byte           `offset:"0x534"` //eMoveState          m_nMoveState;
	_                 [0xb]byte      //padding 	//int32               m_nSwimmingMoveState; // type is eMoveState and used for swimming in CTaskSimpleSwim::ProcessPed
	Health            float32        `offset:"0x540"`
	HealthMax         float32        `offset:"0x544"`
	Armour            float32        `offset:"0x548"`
	_                 [0xc]byte      //padding
	RotationCurrent   float32        `offset:"0x558"`
	RotationTarget    float32        `offset:"0x55C"`
	RotationSpeed     float32        `offset:"0x560"`
	_                 [0x4]byte      //padding
	TouchingCar       *CVehicle      `offset:"0x568"`
	_                 [0x20]byte     //padding
	LastOrCurrentCar  *CVehicle      `offset:"0x58c"`
	_                 [0x8]byte      //padding
	LockInPlace       LockInPlace    `offset:"0x598"`
	_                 [0x7]byte      //padding
	Weapons           [13]CWeapon    `offset:"0x5a0"` //13 Weapon Slots as array //size 0x16c
	_                 [0xc]byte      //padding
	CurrentWeaponSlot byte           `offset:"0x718"`
	_                 [0x27]byte     //padding
	CurrentWeaponID   WeaponType     `offset:"0x740"`
	_                 [0x20]byte     //padding
	Attacker          *CPed          `offset:"0x764"`
	_                 [0x34]byte     //padding
	TargettedPed      *CPed          `offset:"0x79c"`
	_                 [0x24]byte     //padding
}

// Validate CPed offsets
func init() {
	var cped CPed
	err := ValidateOffsets(&cped)
	if err != nil {
		panic(err)
	}
	sz := unsafe.Sizeof(cped)
	if sz != 1988 {
		panic(fmt.Sprintf("CPed size is wrong, expected %d, got %d\n", 1988, sz))
	}
}
