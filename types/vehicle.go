package types

import (
	"fmt"
	"log"
	"strings"
	"unsafe"

	. "gtasamod/complextypes"
	. "gtasamod/madresser"
	. "gtasamod/structoffsetvalidator"
)

// go doesnt know an array of unknown size, so we pretent to have a 512 len array
// use MaxNumberOfVehicles to check how many there are actually!
type VehicleInfo struct {
	FirstVehicleInPool      *CVehicle
	VehiclesInUse           *[512]PoolObjectFlag //size -> MaxNumberOfVehicles byte array -> isEmpty
	MaxNumberOfVehicles     int32                //num
	CurrentNumberOfVehicles int32                //first free
}

var vehiclepool = TypeAtAbsolute[*VehicleInfo](0xB74494)

func GetVehiclePool() *VehicleInfo {
	return *vehiclepool
}

type PoolObjectFlag uint8

func (pof PoolObjectFlag) IsEmpty() bool {
	return (pof & 0x80) != 0
}

// there will never be a vehicle over first in pool, so
// you loop until you hit that one first
var slotempty = fmt.Errorf("Slot Empty")
var outofindex = fmt.Errorf("Slot > FirstVehicleInPool")
var invalidindex = fmt.Errorf("index < 0")

func (vi *VehicleInfo) GetAt(i int32) (*CVehicle, error) {
	if i < 0 {
		return nil, invalidindex
	}

	if (*vi.VehiclesInUse)[i].IsEmpty() {
		return nil, slotempty
	} else if i > vi.CurrentNumberOfVehicles {
		return nil, outofindex
	} else {
		//log.Printf("Bits: %08b", (*vi.VehiclesInUse)[i])
		tt := (*CVehicle)(unsafe.Add(unsafe.Pointer(vi.FirstVehicleInPool), uintptr(CVehicleSize*uintptr(i))))
		return tt, nil
	}
}

func (vi *VehicleInfo) PrintAllVehicles() {
	for i := int32(0); i < vi.MaxNumberOfVehicles; i++ {
		vehicle, err := vi.GetAt(i)
		if err != nil {
			// log.Printf("%d: %s", i, err)
			continue
		}
		log.Printf("%d: %d", i, vehicle.ModelID)
	}
}

// The Vehicle list behaves a little weird
// Is there some internal defragging going on?
func (vi *VehicleInfo) EveryVehicle() []int32 {
	var list []int32
	for i := int32(0); i < vi.MaxNumberOfVehicles; i++ {
		_, err := vi.GetAt(i)
		if err == outofindex {
			continue //break //no more vehicles
		}
		if err == slotempty {
			continue //nothing here
		}
		list = append(list, i)

	}
	return list
}

//0xB6F980 – Is the direct pointer to the pool start (CVehicle)
// Each vehicle object is 2584 (0xA18) bytes. It starts at 0xC502AA0.

// 0x0A18 2584
// size is 0x58a accprdomg to reversoids
type CVehicle struct {
	CPhysical      `offset:"0x0"`
	_              [0x2F0]byte   //padding
	VehicleFlags   VehicleFlags  `offset:"0x428"`
	CreationTime   uint32        `offset:"0x430"`
	BodyColor      [4]byte       `offset:"0x434"` //could be turned into type
	_              [0x28]byte    //padding
	Driver         *CPed         `offset:"0x460"`
	Passenger1     *CPed         `offset:"0x464"`
	Passenger2     *CPed         `offset:"0x468"`
	Passenger3     *CPed         `offset:"0x46c"`
	Passenger4     *CPed         `offset:"0x470"` // Bus
	Passenger5     *CPed         `offset:"0x474"` // Bus
	Passenger6     *CPed         `offset:"0x478"` // Bus
	Passenger7     *CPed         `offset:"0x47c"` // Bus
	Passenger8     *CPed         `offset:"0x480"` // Bus
	Passenger9     *CPed         `offset:"0x484"` // Bus
	_              [0x2]byte     //padding
	HasNitro       byte          `offset:"0x48a"` // 0=no, 1-10=yes
	_              [0x11]byte    //padding
	GasPedal       float32       `offset:"0x49c"`
	BrakePedal     float32       `offset:"0x4a0"`
	_              [0x4]byte     //padding
	CarBomb        CarBomb       `offset:"0x4a8"`
	_              [0x17]byte    //padding
	Health         CarHealth     `offset:"0x4c0"`
	_              [0x34]byte    //padding
	Locked         CarLocked     `offset:"0x4f8"`
	_              [0x94]byte    //padding
	CarType        byte          `offset:"0x590"` // options:"0=car/plane,5=boat,6=train,9=bike"
	_              [0x13]byte    //padding
	CarTyreState   CarTyreState  `offset:"0x5a4"` // options:"0=ok,1=flat,2=landing gear"
	_              [0xb4]byte    //padding
	BikeTyreState  BikeTyreState `offset:"0x65c"` // options:"0=ok,1=flat"
	CycleTyreState BikeTyreState `offset:"0x65e"`
	_              [0x68]byte    //padding
	IsBMX          byte          `offset:"0x6c8"` // options:"0=false,1=true"
	_              [0xf3]byte    //padding
	BurnTimerBike  float32       `offset:"0x7bc"`
	_              [0xe4]byte    //padding
	NitroCount     float32       `offset:"0x8a4"` //-1.0=discharged 0=empty(if hasnitro=>0) 1.0=filled
	_              [0x3c]byte    //padding
	BurnTimer      float32       `offset:"0x8e4"`
	_              [0x130]byte   //padding
}

var CVehicleSize = unsafe.Sizeof(CVehicle{})

func init() {
	var cvehicle CVehicle
	err := ValidateOffsets(&cvehicle)
	if err != nil {
		panic(err)
	}
	sz := unsafe.Sizeof(cvehicle)
	if sz != 2584 {
		panic(fmt.Sprintf("CVehicle size is wrong, expected %d, got %d\n", 2584, sz))
	}
}

// Depending on some curcumstances (i guess full vs limited physics)
// the position is ether in Pos or Posn, return the valid one
// Logically nothing should ever be at pos 0,0,0
func (cv *CVehicle) GetPosition() *CVector {
	if cv.Pos.Position.X == 0 && cv.Pos.Position.Y == 0 && cv.Pos.Position.Z == 0 {
		return &cv.Posn.Posn
	}
	return &cv.Pos.Position
}

type VehicleFlags [8]byte

// Flags for byte[0]
const (
	isLawEnforcer uint8 = 1 << (1 * iota)
	isAmbulanceOnDuty
	isFireTruckOnDuty
	isLocked      // Is this guy locked by the script (cannot be removed)
	engineOn      // For sound purposes. Parked cars have their engines switched off (so do destroyed cars)
	isHandbrakeOn // How's the handbrake doing ?
	lightsOn      // Are the lights switched on ?
	freebies
)

func (vf *VehicleFlags) IsHandbrakeOn() bool {
	return GetBit(vf[0], isHandbrakeOn)
}

func (vf *VehicleFlags) LightsOn() bool {
	return GetBit(vf[0], lightsOn)
}

// Flags for byte[1]
const (
	isVan uint8 = 1 << (1 * iota) // Is this vehicle a van (doors at back of vehicle)
	isBus
	isBig
	lowVehicle
	comedyControls // Will make the car hard to control (hopefully in a funny way)
	warnedPeds
	craneMessageDone
	takeLessDamage // This vehicle is stronger (takes about 1/4 of damage)
)

func (vf *VehicleFlags) ComedyControls() bool {
	return GetBit(vf[1], comedyControls)
}

func (vf *VehicleFlags) ComedyControlsToggle() bool {
	if vf.ComedyControls() {
		vf[1] = ClearBit(vf[1], comedyControls)
		return false
	} else {
		vf[1] = SetBit(vf[1], comedyControls)
		return true
	}
}

func (vf *VehicleFlags) TakeLessDamage() bool {
	return GetBit(vf[1], takeLessDamage)
}

// Flags for byte[2]
const (
	isDamaged            uint8 = 1 << (1 * iota)
	hasBeenOwnedByPlayer       // To work out whether stealing it is a crime
	fadeOutVehicle
	isBeingCarJacked
	createRoadBlockPeds
	canBeDamaged // Set to FALSE during cut scenes to avoid explosions
	occupantsHaveBeenGenerated
	gunSwitchedOff
)

func (vf *VehicleFlags) HasBeenOwnedByPlayer() bool {
	return GetBit(vf[2], hasBeenOwnedByPlayer)
}

func (vf *VehicleFlags) CanBeDamaged() bool {
	return GetBit(vf[2], canBeDamaged)
}

func (vf *VehicleFlags) CanBeDamagedToggle() bool {
	if vf.CanBeDamaged() {
		vf[2] = ClearBit(vf[2], canBeDamaged)
		return false
	} else {
		vf[2] = SetBit(vf[2], canBeDamaged)
		return true
	}
}

// Flags for byte[3]
const (
	vehicleColProcessed uint8 = 1 << (1 * iota) // Has ProcessEntityCollision been processed for this car?
	isCarParked
	hasAlreadyBeenRecordedVehicle
	partOfConvoy
	heliMinimumTilt
	audioChangingGear
	isDrowningVehicle // is vehicle occupants taking damage in water (i.e. vehicle is dead in water)
	tyresDontBurst    // If this is set the tyres are invincible
)

func (vf *VehicleFlags) VehicleColProcessed() bool {
	return GetBit(vf[3], vehicleColProcessed)
}

func (vf *VehicleFlags) IsDrowning() bool {
	return GetBit(vf[3], isDrowning)
}

func (vf *VehicleFlags) TyresDontBurst() bool {
	return GetBit(vf[3], tyresDontBurst)
}

func (vf *VehicleFlags) TyresDontBurstToggle() bool {
	if vf.TyresDontBurst() {
		vf[3] = ClearBit(vf[3], tyresDontBurst)
		return false
	} else {
		vf[3] = SetBit(vf[3], tyresDontBurst)
		return true
	}
}

// Flags for byte[4]
const (
	createdAsPoliceVehicle uint8 = 1 << (1 * iota)
	restingOnPhysical
	parking
	canPark
	fireGun
	driverLastFrame // Was there a driver present last frame ?
	neverUseSmallerRemovalRange
	isRCVehicle
)

// Flags for byte[5]
const (
	alwaysSkidMarks uint8 = 1 << (1 * iota)
	engineBroken          // Engine doesn't work. Player can get in but the vehicle won't drive
	vehicleCanBeTargetted
	partOfAttackWave
	winchCanPickMeUp // This car cannot be picked up by any ropes.
	impounded
	vehicleCanBeTargettedByHS
	sirenOrAlarm // Set to TRUE if siren or alarm active, else FALSE
)

func (vf *VehicleFlags) EngineBroken() bool {
	return GetBit(vf[5], engineBroken)
}

func (vf *VehicleFlags) EngineBrokenToggle() bool {
	if vf.EngineBroken() {
		vf[5] = ClearBit(vf[5], engineBroken)
		return false
	} else {
		vf[5] = SetBit(vf[5], engineBroken)
		return true
	}
}

func (vf *VehicleFlags) WinchCanPickMeUp() bool {
	return GetBit(vf[5], winchCanPickMeUp)
}

func (vf *VehicleFlags) SirenOrAlarm() bool {
	return GetBit(vf[5], sirenOrAlarm)
}

// Flags for byte[6]
const (
	hasGangLeaningOn uint8 = 1 << (1 * iota)
	gangMembersForRoadBlock
	doesProvideCover
	madDriver // This vehicle is driving like a lunatic
	UpgradedStereo
	consideredByPlayer
	petrolTankIsWeakPoint // If false shooting the petrol tank will NOT Blow up the car
	disableParticles
)

func (vf *VehicleFlags) MadDriver() bool {
	return GetBit(vf[6], madDriver)
}

func (vf *VehicleFlags) PetrolTankIsWeakPoint() bool {
	return GetBit(vf[6], petrolTankIsWeakPoint)
}

func (vf *VehicleFlags) PetrolTankIsWeakPointToggle() bool {
	if vf.PetrolTankIsWeakPoint() {
		vf[6] = ClearBit(vf[6], petrolTankIsWeakPoint)
		return false
	} else {
		vf[6] = SetBit(vf[6], petrolTankIsWeakPoint)
		return true
	}
}

// Flags for byte[7]
const (
	hasBeenResprayed uint8 = 1 << (1 * iota) // Has been resprayed in a respray garage. Reset after it has been checked.
	useCarCheats
	dontSetColourWhenRemapping
	usedForReplayVehicle
)

func (vf *VehicleFlags) HasBeenResprayed() bool {
	return GetBit(vf[7], hasBeenResprayed)
}

type EntityType uint16

const (
	PlayerAsDriver   EntityType = 2
	QuietDriver                 = 18
	SuspiciousDriver            = 26
	NoDriver                    = 34
	Destroyed                   = 42
	PlayerGettingOut            = 74
)

func (et EntityType) String() string {
	switch et {
	case PlayerAsDriver:
		return "PlayerAsDriver"
	case QuietDriver:
		return "QuietDriver"
	case SuspiciousDriver:
		return "SuspiciousDriver"
	case NoDriver:
		return "NoDriver"
	case Destroyed:
		return "Destroyed"
	case PlayerGettingOut:
		return "PlayerGettingOut"
	default:
		return "Unknown EntityType"
	}
}

type CarBomb byte

const (
	NoBomb              CarBomb = 0
	TimeBombPresent             = 1
	IgnitionBombPresent         = 2
	SetRemotely                 = 3 //idk
	TimeBombArmed               = 4
	IgnitionBombArmed           = 5
)

func (cb CarBomb) String() string {
	switch cb {
	case NoBomb:
		return "NoBomb"
	case TimeBombPresent:
		return "TimeBombPresent"
	case IgnitionBombPresent:
		return "IgnitionBombPresent"
	case SetRemotely:
		return "SetRemotely"
	case TimeBombArmed:
		return "TimeBombArmed"
	case IgnitionBombArmed:
		return "IgnitionBombArmed"
	default:
		return "Unkown CarBomb State"
	}
}

type CarHealth float32

func (ch CarHealth) String() string {
	if ch == 0 {
		return fmt.Sprintf("Dead (%0.2f)", ch)
	}
	if ch > 0 && ch <= 250 {
		return fmt.Sprintf("Burning (%0.2f)", ch)
	}
	if ch > 250 && ch < 1000 {
		return fmt.Sprintf("Damaged (%0.2f)", ch)
	}
	if ch >= 1000 {
		return fmt.Sprintf("FullHealth (%0.2f)", ch)
	}
	return fmt.Sprintf("Weird Health Value (%0.2f)", ch)
}

type CarLocked uint32

const (
	CarIsOpen   CarLocked = 1
	CarIsLocked           = 2
)

func (cl CarLocked) String() string {
	switch cl {
	case CarIsOpen:
		return "Car is Open"
	case CarIsLocked:
		return "Car is Locked"
	default:
		return "Bad Lock State"
	}
}

type CarTyreState struct {
	LeftFront  byte
	LeftRear   byte
	RightFront byte
	RightRear  byte
}

type BikeTyreState struct {
	Front byte
	Rear  byte
}

func (cv *CVehicle) String() string {
	var out strings.Builder

	out.WriteString("Vehicle: ")
	out.WriteString(fmt.Sprintf("Health:(%s) ", cv.Health.String()))
	out.WriteString(fmt.Sprintf("Position:(%s) ", cv.GetPosition().String()))
	out.WriteString(fmt.Sprintf("VehicleID:(%d) ", cv.ModelID))
	out.WriteString(fmt.Sprintf("VehicleName:(%s) ", VehicleID(cv.ModelID).String()))
	out.WriteString(fmt.Sprintf("Type:(%s) ", cv.EntityType.String()))
	//out.WriteString(fmt.Sprintf("PhysFlags:(%s) ", cv.PhysFlags.String())) //crashes on plane
	out.WriteString(fmt.Sprintf("Mass:(%0.2f) ", cv.Mass))
	out.WriteString(fmt.Sprintf("Driver:(%p) ", cv.Driver))
	out.WriteString(fmt.Sprintf("Passenger1:(%p) ", cv.Passenger1))
	out.WriteString(fmt.Sprintf("HasNitro:(%d) ", cv.HasNitro))
	out.WriteString(fmt.Sprintf("NitroAmount:(%0.2f) ", cv.NitroCount))
	out.WriteString(fmt.Sprintf("Gas Pedal:(%0.2f) ", cv.GasPedal))
	out.WriteString(fmt.Sprintf("Brake Pedal:(%0.2f) ", cv.BrakePedal))
	out.WriteString(fmt.Sprintf("Car Bomb:(%s) ", cv.CarBomb.String()))
	out.WriteString(fmt.Sprintf("Locked:(%s) ", cv.Locked.String()))
	out.WriteString(fmt.Sprintf("BurnTimer:(%0.2f) ", cv.BurnTimer))
	out.WriteString(fmt.Sprintf("BurnTimerBike:(%0.2f) ", cv.BurnTimerBike))

	return out.String()
}

type VehicleID int32

func (car VehicleID) AsInt() string {
	return fmt.Sprintf("%d", car)
}

const (
	MODEL_LANDSTAL VehicleID = 400
	MODEL_BRAVURA  VehicleID = 401
	MODEL_BUFFALO  VehicleID = 402
	MODEL_LINERUN  VehicleID = 403
	MODEL_PEREN    VehicleID = 404
	MODEL_SENTINEL VehicleID = 405
	MODEL_DUMPER   VehicleID = 406
	MODEL_FIRETRUK VehicleID = 407
	MODEL_TRASH    VehicleID = 408
	MODEL_STRETCH  VehicleID = 409
	MODEL_MANANA   VehicleID = 410
	MODEL_INFERNUS VehicleID = 411
	MODEL_VOODOO   VehicleID = 412
	MODEL_PONY     VehicleID = 413
	MODEL_MULE     VehicleID = 414
	MODEL_CHEETAH  VehicleID = 415
	MODEL_AMBULAN  VehicleID = 416
	MODEL_LEVIATHN VehicleID = 417
	MODEL_MOONBEAM VehicleID = 418
	MODEL_ESPERANT VehicleID = 419
	MODEL_TAXI     VehicleID = 420
	MODEL_WASHING  VehicleID = 421
	MODEL_BOBCAT   VehicleID = 422
	MODEL_MRWHOOP  VehicleID = 423
	MODEL_BFINJECT VehicleID = 424
	MODEL_HUNTER   VehicleID = 425
	MODEL_PREMIER  VehicleID = 426
	MODEL_ENFORCER VehicleID = 427
	MODEL_SECURICA VehicleID = 428
	MODEL_BANSHEE  VehicleID = 429
	MODEL_PREDATOR VehicleID = 430
	MODEL_BUS      VehicleID = 431
	MODEL_RHINO    VehicleID = 432
	MODEL_BARRACKS VehicleID = 433
	MODEL_HOTKNIFE VehicleID = 434
	MODEL_ARTICT1  VehicleID = 435
	MODEL_PREVION  VehicleID = 436
	MODEL_COACH    VehicleID = 437
	MODEL_CABBIE   VehicleID = 438
	MODEL_STALLION VehicleID = 439
	MODEL_RUMPO    VehicleID = 440
	MODEL_RCBANDIT VehicleID = 441
	MODEL_ROMERO   VehicleID = 442
	MODEL_PACKER   VehicleID = 443
	MODEL_MONSTER  VehicleID = 444
	MODEL_ADMIRAL  VehicleID = 445
	MODEL_SQUALO   VehicleID = 446
	MODEL_SEASPAR  VehicleID = 447
	MODEL_PIZZABOY VehicleID = 448
	MODEL_TRAM     VehicleID = 449 //crash
	MODEL_ARTICT2  VehicleID = 450
	MODEL_TURISMO  VehicleID = 451
	MODEL_SPEEDER  VehicleID = 452
	MODEL_REEFER   VehicleID = 453
	MODEL_TROPIC   VehicleID = 454
	MODEL_FLATBED  VehicleID = 455
	MODEL_YANKEE   VehicleID = 456
	MODEL_CADDY    VehicleID = 457
	MODEL_SOLAIR   VehicleID = 458
	MODEL_TOPFUN   VehicleID = 459
	MODEL_SKIMMER  VehicleID = 460
	MODEL_PCJ600   VehicleID = 461
	MODEL_FAGGIO   VehicleID = 462
	MODEL_FREEWAY  VehicleID = 463
	MODEL_RCBARON  VehicleID = 464
	MODEL_RCRAIDER VehicleID = 465
	MODEL_GLENDALE VehicleID = 466
	MODEL_OCEANIC  VehicleID = 467
	MODEL_SANCHEZ  VehicleID = 468
	MODEL_SPARROW  VehicleID = 469
	MODEL_PATRIOT  VehicleID = 470
	MODEL_QUAD     VehicleID = 471
	MODEL_COASTG   VehicleID = 472
	MODEL_DINGHY   VehicleID = 473
	MODEL_HERMES   VehicleID = 474
	MODEL_SABRE    VehicleID = 475
	MODEL_RUSTLER  VehicleID = 476
	MODEL_ZR350    VehicleID = 477
	MODEL_WALTON   VehicleID = 478
	MODEL_REGINA   VehicleID = 479
	MODEL_COMET    VehicleID = 480
	MODEL_BMX      VehicleID = 481
	MODEL_BURRITO  VehicleID = 482
	MODEL_CAMPER   VehicleID = 483
	MODEL_MARQUIS  VehicleID = 484
	MODEL_BAGGAGE  VehicleID = 485
	MODEL_DOZER    VehicleID = 486
	MODEL_MAVERICK VehicleID = 487
	MODEL_VCNMAV   VehicleID = 488
	MODEL_RANCHER  VehicleID = 489
	MODEL_FBIRANCH VehicleID = 490
	MODEL_VIRGO    VehicleID = 491
	MODEL_GREENWOO VehicleID = 492
	MODEL_JETMAX   VehicleID = 493
	MODEL_HOTRING  VehicleID = 494
	MODEL_SANDKING VehicleID = 495
	MODEL_BLISTAC  VehicleID = 496
	MODEL_POLMAV   VehicleID = 497
	MODEL_BOXVILLE VehicleID = 498
	MODEL_BENSON   VehicleID = 499
	MODEL_MESA     VehicleID = 500
	MODEL_RCGOBLIN VehicleID = 501
	MODEL_HOTRINA  VehicleID = 502
	MODEL_HOTRINB  VehicleID = 503
	MODEL_BLOODRA  VehicleID = 504
	MODEL_RNCHLURE VehicleID = 505
	MODEL_SUPERGT  VehicleID = 506
	MODEL_ELEGANT  VehicleID = 507
	MODEL_JOURNEY  VehicleID = 508
	MODEL_BIKE     VehicleID = 509
	MODEL_MTBIKE   VehicleID = 510
	MODEL_BEAGLE   VehicleID = 511
	MODEL_CROPDUST VehicleID = 512
	MODEL_STUNT    VehicleID = 513
	MODEL_PETRO    VehicleID = 514
	MODEL_RDTRAIN  VehicleID = 515
	MODEL_NEBULA   VehicleID = 516
	MODEL_MAJESTIC VehicleID = 517
	MODEL_BUCCANEE VehicleID = 518
	MODEL_SHAMAL   VehicleID = 519
	MODEL_HYDRA    VehicleID = 520
	MODEL_FCR900   VehicleID = 521
	MODEL_NRG500   VehicleID = 522
	MODEL_COPBIKE  VehicleID = 523
	MODEL_CEMENT   VehicleID = 524
	MODEL_TOWTRUCK VehicleID = 525
	MODEL_FORTUNE  VehicleID = 526
	MODEL_CADRONA  VehicleID = 527
	MODEL_FBITRUCK VehicleID = 528
	MODEL_WILLARD  VehicleID = 529
	MODEL_FORKLIFT VehicleID = 530
	MODEL_TRACTOR  VehicleID = 531
	MODEL_COMBINE  VehicleID = 532
	MODEL_FELTZER  VehicleID = 533
	MODEL_REMINGTN VehicleID = 534
	MODEL_SLAMVAN  VehicleID = 535
	MODEL_BLADE    VehicleID = 536
	MODEL_FREIGHT  VehicleID = 537 //crash
	MODEL_STREAK   VehicleID = 538 //crash
	MODEL_VORTEX   VehicleID = 539
	MODEL_VINCENT  VehicleID = 540
	MODEL_BULLET   VehicleID = 541
	MODEL_CLOVER   VehicleID = 542
	MODEL_SADLER   VehicleID = 543
	MODEL_FIRELA   VehicleID = 544
	MODEL_HUSTLER  VehicleID = 545
	MODEL_INTRUDER VehicleID = 546
	MODEL_PRIMO    VehicleID = 547
	MODEL_CARGOBOB VehicleID = 548
	MODEL_TAMPA    VehicleID = 549
	MODEL_SUNRISE  VehicleID = 550
	MODEL_MERIT    VehicleID = 551
	MODEL_UTILITY  VehicleID = 552
	MODEL_NEVADA   VehicleID = 553
	MODEL_YOSEMITE VehicleID = 554
	MODEL_WINDSOR  VehicleID = 555
	MODEL_MONSTERA VehicleID = 556
	MODEL_MONSTERB VehicleID = 557
	MODEL_URANUS   VehicleID = 558
	MODEL_JESTER   VehicleID = 559
	MODEL_SULTAN   VehicleID = 560
	MODEL_STRATUM  VehicleID = 561
	MODEL_ELEGY    VehicleID = 562
	MODEL_RAINDANC VehicleID = 563
	MODEL_RCTIGER  VehicleID = 564
	MODEL_FLASH    VehicleID = 565
	MODEL_TAHOMA   VehicleID = 566
	MODEL_SAVANNA  VehicleID = 567
	MODEL_BANDITO  VehicleID = 568
	MODEL_FREIFLAT VehicleID = 569 //crash
	MODEL_STREAKC  VehicleID = 570 //crash
	MODEL_KART     VehicleID = 571
	MODEL_MOWER    VehicleID = 572
	MODEL_DUNERIDE VehicleID = 573
	MODEL_SWEEPER  VehicleID = 574
	MODEL_BROADWAY VehicleID = 575
	MODEL_TORNADO  VehicleID = 576
	MODEL_AT400    VehicleID = 577
	MODEL_DFT30    VehicleID = 578
	MODEL_HUNTLEY  VehicleID = 579
	MODEL_STAFFORD VehicleID = 580
	MODEL_BF400    VehicleID = 581
	MODEL_NEWSVAN  VehicleID = 582
	MODEL_TUG      VehicleID = 583
	MODEL_PETROTR  VehicleID = 584
	MODEL_EMPEROR  VehicleID = 585
	MODEL_WAYFARER VehicleID = 586
	MODEL_EUROS    VehicleID = 587
	MODEL_HOTDOG   VehicleID = 588
	MODEL_CLUB     VehicleID = 589
	MODEL_FREIBOX  VehicleID = 590 //crash
	MODEL_ARTICT3  VehicleID = 591
	MODEL_ANDROM   VehicleID = 592
	MODEL_DODO     VehicleID = 593
	MODEL_RCCAM    VehicleID = 594
	MODEL_LAUNCH   VehicleID = 595
	MODEL_COPCARLA VehicleID = 596
	MODEL_COPCARSF VehicleID = 597
	MODEL_COPCARVG VehicleID = 598
	MODEL_COPCARRU VehicleID = 599
	MODEL_PICADOR  VehicleID = 600
	MODEL_SWATVAN  VehicleID = 601
	MODEL_ALPHA    VehicleID = 602
	MODEL_PHOENIX  VehicleID = 603
	MODEL_GLENSHIT VehicleID = 604
	MODEL_SADLSHIT VehicleID = 605
	MODEL_BAGBOXA  VehicleID = 606
	MODEL_BAGBOXB  VehicleID = 607
	MODEL_TUGSTAIR VehicleID = 608
	MODEL_BOXBURG  VehicleID = 609
	MODEL_FARMTR1  VehicleID = 610
	MODEL_UTILTR1  VehicleID = 611
)

func (car VehicleID) String() string {
	switch car {
	case MODEL_LANDSTAL:
		return "LANDSTAL"
	case MODEL_BRAVURA:
		return "BRAVURA"
	case MODEL_BUFFALO:
		return "BUFFALO"
	case MODEL_LINERUN:
		return "LINERUN"
	case MODEL_PEREN:
		return "PEREN"
	case MODEL_SENTINEL:
		return "SENTINEL"
	case MODEL_DUMPER:
		return "DUMPER"
	case MODEL_FIRETRUK:
		return "FIRETRUK"
	case MODEL_TRASH:
		return "TRASH"
	case MODEL_STRETCH:
		return "STRETCH"
	case MODEL_MANANA:
		return "MANANA"
	case MODEL_INFERNUS:
		return "INFERNUS"
	case MODEL_VOODOO:
		return "VOODOO"
	case MODEL_PONY:
		return "PONY"
	case MODEL_MULE:
		return "MULE"
	case MODEL_CHEETAH:
		return "CHEETAH"
	case MODEL_AMBULAN:
		return "AMBULAN"
	case MODEL_LEVIATHN:
		return "LEVIATHN"
	case MODEL_MOONBEAM:
		return "MOONBEAM"
	case MODEL_ESPERANT:
		return "ESPERANT"
	case MODEL_TAXI:
		return "TAXI"
	case MODEL_WASHING:
		return "WASHING"
	case MODEL_BOBCAT:
		return "BOBCAT"
	case MODEL_MRWHOOP:
		return "MRWHOOP"
	case MODEL_BFINJECT:
		return "BFINJECT"
	case MODEL_HUNTER:
		return "HUNTER"
	case MODEL_PREMIER:
		return "PREMIER"
	case MODEL_ENFORCER:
		return "ENFORCER"
	case MODEL_SECURICA:
		return "SECURICA"
	case MODEL_BANSHEE:
		return "BANSHEE"
	case MODEL_PREDATOR:
		return "PREDATOR"
	case MODEL_BUS:
		return "BUS"
	case MODEL_RHINO:
		return "RHINO"
	case MODEL_BARRACKS:
		return "BARRACKS"
	case MODEL_HOTKNIFE:
		return "HOTKNIFE"
	case MODEL_ARTICT1:
		return "ARTICT1"
	case MODEL_PREVION:
		return "PREVION"
	case MODEL_COACH:
		return "COACH"
	case MODEL_CABBIE:
		return "CABBIE"
	case MODEL_STALLION:
		return "STALLION"
	case MODEL_RUMPO:
		return "RUMPO"
	case MODEL_RCBANDIT:
		return "RCBANDIT"
	case MODEL_ROMERO:
		return "ROMERO"
	case MODEL_PACKER:
		return "PACKER"
	case MODEL_MONSTER:
		return "MONSTER"
	case MODEL_ADMIRAL:
		return "ADMIRAL"
	case MODEL_SQUALO:
		return "SQUALO"
	case MODEL_SEASPAR:
		return "SEASPAR"
	case MODEL_PIZZABOY:
		return "PIZZABOY"
	case MODEL_TRAM:
		return "TRAM"
	case MODEL_ARTICT2:
		return "ARTICT2"
	case MODEL_TURISMO:
		return "TURISMO"
	case MODEL_SPEEDER:
		return "SPEEDER"
	case MODEL_REEFER:
		return "REEFER"
	case MODEL_TROPIC:
		return "TROPIC"
	case MODEL_FLATBED:
		return "FLATBED"
	case MODEL_YANKEE:
		return "YANKEE"
	case MODEL_CADDY:
		return "CADDY"
	case MODEL_SOLAIR:
		return "SOLAIR"
	case MODEL_TOPFUN:
		return "TOPFUN"
	case MODEL_SKIMMER:
		return "SKIMMER"
	case MODEL_PCJ600:
		return "PCJ600"
	case MODEL_FAGGIO:
		return "FAGGIO"
	case MODEL_FREEWAY:
		return "FREEWAY"
	case MODEL_RCBARON:
		return "RCBARON"
	case MODEL_RCRAIDER:
		return "RCRAIDER"
	case MODEL_GLENDALE:
		return "GLENDALE"
	case MODEL_OCEANIC:
		return "OCEANIC"
	case MODEL_SANCHEZ:
		return "SANCHEZ"
	case MODEL_SPARROW:
		return "SPARROW"
	case MODEL_PATRIOT:
		return "PATRIOT"
	case MODEL_QUAD:
		return "QUAD"
	case MODEL_COASTG:
		return "COASTG"
	case MODEL_DINGHY:
		return "DINGHY"
	case MODEL_HERMES:
		return "HERMES"
	case MODEL_SABRE:
		return "SABRE"
	case MODEL_RUSTLER:
		return "RUSTLER"
	case MODEL_ZR350:
		return "ZR350"
	case MODEL_WALTON:
		return "WALTON"
	case MODEL_REGINA:
		return "REGINA"
	case MODEL_COMET:
		return "COMET"
	case MODEL_BMX:
		return "BMX"
	case MODEL_BURRITO:
		return "BURRITO"
	case MODEL_CAMPER:
		return "CAMPER"
	case MODEL_MARQUIS:
		return "MARQUIS"
	case MODEL_BAGGAGE:
		return "BAGGAGE"
	case MODEL_DOZER:
		return "DOZER"
	case MODEL_MAVERICK:
		return "MAVERICK"
	case MODEL_VCNMAV:
		return "VCNMAV"
	case MODEL_RANCHER:
		return "RANCHER"
	case MODEL_FBIRANCH:
		return "FBIRANCH"
	case MODEL_VIRGO:
		return "VIRGO"
	case MODEL_GREENWOO:
		return "GREENWOO"
	case MODEL_JETMAX:
		return "JETMAX"
	case MODEL_HOTRING:
		return "HOTRING"
	case MODEL_SANDKING:
		return "SANDKING"
	case MODEL_BLISTAC:
		return "BLISTAC"
	case MODEL_POLMAV:
		return "POLMAV"
	case MODEL_BOXVILLE:
		return "BOXVILLE"
	case MODEL_BENSON:
		return "BENSON"
	case MODEL_MESA:
		return "MESA"
	case MODEL_RCGOBLIN:
		return "RCGOBLIN"
	case MODEL_HOTRINA:
		return "HOTRINA"
	case MODEL_HOTRINB:
		return "HOTRINB"
	case MODEL_BLOODRA:
		return "BLOODRA"
	case MODEL_RNCHLURE:
		return "RNCHLURE"
	case MODEL_SUPERGT:
		return "SUPERGT"
	case MODEL_ELEGANT:
		return "ELEGANT"
	case MODEL_JOURNEY:
		return "JOURNEY"
	case MODEL_BIKE:
		return "BIKE"
	case MODEL_MTBIKE:
		return "MTBIKE"
	case MODEL_BEAGLE:
		return "BEAGLE"
	case MODEL_CROPDUST:
		return "CROPDUST"
	case MODEL_STUNT:
		return "STUNT"
	case MODEL_PETRO:
		return "PETRO"
	case MODEL_RDTRAIN:
		return "RDTRAIN"
	case MODEL_NEBULA:
		return "NEBULA"
	case MODEL_MAJESTIC:
		return "MAJESTIC"
	case MODEL_BUCCANEE:
		return "BUCCANEE"
	case MODEL_SHAMAL:
		return "SHAMAL"
	case MODEL_HYDRA:
		return "HYDRA"
	case MODEL_FCR900:
		return "FCR900"
	case MODEL_NRG500:
		return "NRG500"
	case MODEL_COPBIKE:
		return "COPBIKE"
	case MODEL_CEMENT:
		return "CEMENT"
	case MODEL_TOWTRUCK:
		return "TOWTRUCK"
	case MODEL_FORTUNE:
		return "FORTUNE"
	case MODEL_CADRONA:
		return "CADRONA"
	case MODEL_FBITRUCK:
		return "FBITRUCK"
	case MODEL_WILLARD:
		return "WILLARD"
	case MODEL_FORKLIFT:
		return "FORKLIFT"
	case MODEL_TRACTOR:
		return "TRACTOR"
	case MODEL_COMBINE:
		return "COMBINE"
	case MODEL_FELTZER:
		return "FELTZER"
	case MODEL_REMINGTN:
		return "REMINGTN"
	case MODEL_SLAMVAN:
		return "SLAMVAN"
	case MODEL_BLADE:
		return "BLADE"
	case MODEL_FREIGHT:
		return "FREIGHT"
	case MODEL_STREAK:
		return "STREAK"
	case MODEL_VORTEX:
		return "VORTEX"
	case MODEL_VINCENT:
		return "VINCENT"
	case MODEL_BULLET:
		return "BULLET"
	case MODEL_CLOVER:
		return "CLOVER"
	case MODEL_SADLER:
		return "SADLER"
	case MODEL_FIRELA:
		return "FIRELA"
	case MODEL_HUSTLER:
		return "HUSTLER"
	case MODEL_INTRUDER:
		return "INTRUDER"
	case MODEL_PRIMO:
		return "PRIMO"
	case MODEL_CARGOBOB:
		return "CARGOBOB"
	case MODEL_TAMPA:
		return "TAMPA"
	case MODEL_SUNRISE:
		return "SUNRISE"
	case MODEL_MERIT:
		return "MERIT"
	case MODEL_UTILITY:
		return "UTILITY"
	case MODEL_NEVADA:
		return "NEVADA"
	case MODEL_YOSEMITE:
		return "YOSEMITE"
	case MODEL_WINDSOR:
		return "WINDSOR"
	case MODEL_MONSTERA:
		return "MONSTERA"
	case MODEL_MONSTERB:
		return "MONSTERB"
	case MODEL_URANUS:
		return "URANUS"
	case MODEL_JESTER:
		return "JESTER"
	case MODEL_SULTAN:
		return "SULTAN"
	case MODEL_STRATUM:
		return "STRATUM"
	case MODEL_ELEGY:
		return "ELEGY"
	case MODEL_RAINDANC:
		return "RAINDANC"
	case MODEL_RCTIGER:
		return "RCTIGER"
	case MODEL_FLASH:
		return "FLASH"
	case MODEL_TAHOMA:
		return "TAHOMA"
	case MODEL_SAVANNA:
		return "SAVANNA"
	case MODEL_BANDITO:
		return "BANDITO"
	case MODEL_FREIFLAT:
		return "FREIFLAT"
	case MODEL_STREAKC:
		return "STREAKC"
	case MODEL_KART:
		return "KART"
	case MODEL_MOWER:
		return "MOWER"
	case MODEL_DUNERIDE:
		return "DUNERIDE"
	case MODEL_SWEEPER:
		return "SWEEPER"
	case MODEL_BROADWAY:
		return "BROADWAY"
	case MODEL_TORNADO:
		return "TORNADO"
	case MODEL_AT400:
		return "AT400"
	case MODEL_DFT30:
		return "DFT30"
	case MODEL_HUNTLEY:
		return "HUNTLEY"
	case MODEL_STAFFORD:
		return "STAFFORD"
	case MODEL_BF400:
		return "BF400"
	case MODEL_NEWSVAN:
		return "NEWSVAN"
	case MODEL_TUG:
		return "TUG"
	case MODEL_PETROTR:
		return "PETROTR"
	case MODEL_EMPEROR:
		return "EMPEROR"
	case MODEL_WAYFARER:
		return "WAYFARER"
	case MODEL_EUROS:
		return "EUROS"
	case MODEL_HOTDOG:
		return "HOTDOG"
	case MODEL_CLUB:
		return "CLUB"
	case MODEL_FREIBOX:
		return "FREIBOX"
	case MODEL_ARTICT3:
		return "ARTICT3"
	case MODEL_ANDROM:
		return "ANDROM"
	case MODEL_DODO:
		return "DODO"
	case MODEL_RCCAM:
		return "RCCAM"
	case MODEL_LAUNCH:
		return "LAUNCH"
	case MODEL_COPCARLA:
		return "COPCARLA"
	case MODEL_COPCARSF:
		return "COPCARSF"
	case MODEL_COPCARVG:
		return "COPCARVG"
	case MODEL_COPCARRU:
		return "COPCARRU"
	case MODEL_PICADOR:
		return "PICADOR"
	case MODEL_SWATVAN:
		return "SWATVAN"
	case MODEL_ALPHA:
		return "ALPHA"
	case MODEL_PHOENIX:
		return "PHOENIX"
	case MODEL_GLENSHIT:
		return "GLENSHIT"
	case MODEL_SADLSHIT:
		return "SADLSHIT"
	case MODEL_BAGBOXA:
		return "BAGBOXA"
	case MODEL_BAGBOXB:
		return "BAGBOXB"
	case MODEL_TUGSTAIR:
		return "TUGSTAIR"
	case MODEL_BOXBURG:
		return "BOXBURG"
	case MODEL_FARMTR1:
		return "FARMTR1"
	case MODEL_UTILTR1:
		return "UTILTR1"
	default:
		return fmt.Sprintf("Unknown ID %d", car)
	}
}

var ValidVehicleIDs = []VehicleID{400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 450, 451, 452, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 525, 526, 527, 528, 529, 530, 531, 532, 533, 534, 535, 536, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 551, 552, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611}
