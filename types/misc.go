package types

import (
	"fmt"
	"unsafe"

	. "gtasamod/complextypes"
	. "gtasamod/structoffsetvalidator"
)

type CSimpleTransform struct {
	Posn    CVector
	Heading float32
}

func (cst *CSimpleTransform) String() string {
	return fmt.Sprintf("%s -> %0.2f", cst.Posn.String(), cst.Heading)
}

func init() {
	sz := unsafe.Sizeof(CSimpleTransform{})
	if sz != 0x10 {
		panic(fmt.Sprintf("CSimpleTransform size is wrong, expected %d, got %d\n", 0x10, sz))
	}
}

type CVector struct {
	X, Y, Z float32 //do we need offset tags in this too?
}

func init() {
	sz := unsafe.Sizeof(CVector{})
	if sz != 12 {
		panic(fmt.Sprintf("CVector size is wrong, expected %d, got %d\n", 12, sz))
	}
}

func (cv *CVector) String() string {
	return fmt.Sprintf("%.2f %.2f %.2f", cv.X, cv.Y, cv.Z)
}

type CVector2D struct {
	X, Y float32
}

func init() {
	sz := unsafe.Sizeof(CVector2D{})
	if sz != 8 {
		panic(fmt.Sprintf("CVector size is wrong, expected %d, got %d\n", 8, sz))
	}
}

// (84 bytes):
type CMatrix struct {
	Right                 CVector   `offset:"0x0"`  // 0x0  // RW: Right
	flags                 uint32    `offset:"0xc"`  // 0xC
	Forward               CVector   `offset:"0x10"` // 0x10 // RW: Up
	pad1                  uint32    `offset:"0x1c"` // 0x1C
	Up                    CVector   `offset:"0x20"` // 0x20 // RW: At
	pad2                  uint32    `offset:"0x2c"` // 0x2C
	Position              CVector   `offset:"0x30"` // 0x30
	pad3                  uint32    `offset:"0x3c"` // 0x3C
	m_pAttachMatrix       uintptr   `offset:"0x40"` // 0x40 RwMatrix* (unused)
	m_bOwnsAttachedMatrix byte      `offset:"0x44"` // 0x44 - Do we need to delete attached matrix at detaching
	_                     [0xF]byte //padding
}

func init() {
	var mx CMatrix
	err := ValidateOffsets(&mx)
	if err != nil {
		panic(err)
	}

	sz := unsafe.Sizeof(CMatrix{})
	if sz != 84 {
		panic(fmt.Sprintf("CMatrix size is wrong, expected %d, got %d\n", 84, sz))
	}
}

func (cps *CMatrix) String() string {
	return cps.Position.String()
}

type InteriorState byte

const (
	InteriourOutside InteriorState = 0
	InteriourInside                = 3
)

func (is InteriorState) String() string {
	switch is {
	case InteriourOutside:
		return "InteriourOutside"
	case InteriourInside:
		return "InteriourInside"
	default:
		return "Unknown InteriorState"
	}
}

type AnimationState byte

const (
	Landing  AnimationState = 0
	Punching                = 61
	Stopped                 = 102
	Sprint                  = 154
	Run                     = 205
)

func (as AnimationState) String() string {
	switch as {
	case Landing:
		return "Landing"
	case Punching:
		return "Punching"
	case Stopped:
		return "Stopped"
	case Sprint:
		return "Sprint"
	case Run:
		return "Run"
	default:
		return "Unknown AnimationState"
	}
}

type Properties byte

const (
	None            Properties = 0
	Invisible                  = 3
	NotHeadshotable            = 12
	Drowning                   = 20
)

func (pp Properties) String() string {
	switch pp {
	case None:
		return "None"
	case Invisible:
		return "Invisible"
	case NotHeadshotable:
		return "NotHeadshotable"
	case Drowning:
		return "Drowning"
	default:
		return "Unknown Properties"
	}
}

type State uint32

const (
	ExitVehicle State = 0
	Normal            = 1
	Driving           = 50
	Wasted            = 55
	Busted            = 63
)

func (ss State) String() string {
	switch ss {
	case ExitVehicle:
		return "ExitVehicle"
	case Normal:
		return "Normal"
	case Driving:
		return "Driving"
	case Wasted:
		return "Wasted"
	case Busted:
		return "Busted"
	default:
		return "Uknown State"
	}
}

type LockInPlace byte

const (
	CanMove LockInPlace = 0
	Locked              = 1
)

func (lip LockInPlace) String() string {
	switch lip {
	case CanMove:
		return "Can Move"
	case Locked:
		return "Locked in Place"
	default:
		return "Unknown Lock State"
	}
}

func (lip *LockInPlace) Lock() {
	*lip = Locked
}

func (lip *LockInPlace) Unlock() {
	*lip = CanMove
}

func (lip *LockInPlace) Toggle() {
	if *lip == Locked {
		lip.Unlock()
	} else {
		lip.Lock()
	}
}

type PlayerState [16]byte

// Flags for byte[0]
const (
	isStanding uint8 = 1 << (1 * iota)
	wasStanding
	isLooking
	isRestoringLook
	isAimingGun
	isRestoringGun
	canPointGunAtTarget
	isTalking
)

func (p *PlayerState) IsStanding() bool { // is Standing on ground
	return GetBit(p[0], isStanding)
}

// Flags for byte[1]
const (
	//inVehicle uint8 = 128 >> (1 * iota)
	inVehicle uint8 = 1 << (1 * iota)
	isInTheAir
	isLanding
	hitSomethingLastFrame
	isNearCar
	renderPedInCar
	updateAnimHeading
	removeHead
)

func (p *PlayerState) InVehicle() bool {
	return GetBit(p[1], inVehicle)
}

// Flags for byte[2]
const (
	firingWeapon uint8 = 1 << (1 * iota)
	hasACamera
	pedIsBleeding
	stopAndShoot
	isPedDieAnimPlaying
	stayInSamePlace
	kindaStayInSamePlace
	beingChasedByPolice
)

func (p *PlayerState) IsFiringWeapon() bool {
	return GetBit(p[2], firingWeapon)
}

// Flags for byte[3]
const (
	notAllowedToDuck uint8 = 1 << (1 * iota)
	crouchWhenShooting
	isDucking
	getUpAnimStarted
	doBloodyFootprints
	dontDragMeOutCar
	stillOnValidPoly
	allowMedicsToReviveMe
)

// Flags for byte[4]
const (
	resetWalkAnims uint8 = 1 << (1 * iota)
	onBoat
	busJacked
	fadeOutPlayerState
	knockedUpIntoAir
	hitSteepSlope
	cullExtraFarAway
	tryingToReachDryLand
)

// Flags for byte[5]
const (
	collidedWithMyVehicle uint8 = 1 << (1 * iota)
	richFromMugging
	chrisCriminal
	shakeFist
	noCriticalHits
	hasAlreadyBeenRecordedPlayerState
	updateMatricesRequired
	fleeWhenStanding
)

// Flags for byte[6]
const (
	miamiViceCop uint8 = 1 << (1 * iota)
	moneyHasBeenGivenByScript
	hasBeenPhotographed
	isDrowning
	drownsInWater
	headStuckInCollision
	deadPedInFrontOfCar
	stayInCarOnJack
)

// Flags for byte[7]
const (
	dontFight uint8 = 1 << (1 * iota)
	doomAim
	canBeShotInVehicle
	pushedAlongByCar
	neverEverTargetThisPed
	thisPedIsATargetPriority
	crouchWhenScared
	knockedOffBike
)

// Flags for byte[8]
const (
	donePositionOutOfCollision uint8 = 1 << (1 * iota)
	dontRender
	hasBeenAddedToPopulation
	hasJustLeftCar
	isInDisguise
	doesntListenToPlayerGroupCommands
	isBeingArrested
	hasJustSoughtCover
)

// Flags for byte[9]
const (
	KilledByStealth uint8 = 1 << (1 * iota)
	DoesntDropWeaponsWhenDead
	CalledPreRender
	BloodPuddleCreated
	PartOfAttackWave
	ClearRadarBlipOnDeath
	NeverLeavesGroup
	TestForBlockedPositions
)

// Flags for byte[10]
const (
	rightArmBlocked uint8 = 1 << (1 * iota)
	leftArmBlocked
	duckRightArmBlocked
	midriffBlockedForJump
	fallenDown
	useAttractorInstantly
	dontAcceptIKLookAts
	hasAScriptBrain
)

// Flags for byte[11]
const (
	waitingForScriptBrainToLoad uint8 = 1 << (1 * iota)
	hasGroupDriveTask
	canExitCar
	cantBeKnockedOffBike1
	cantBeKnockedOffBike2
	hasBeenRendered
	isCached
	pushOtherPeds
)

// Flags for byte[12]
const (
	hasBulletProofVest uint8 = 1 << (1 * iota)
	usingMobilePhone
	upperBodyDamageAnimsOnly
	stuckUnderCar
	keepTasksAfterCleanUp
	isDyingStuck
	ignoreHeightCheckOnGotoPointTask
	forceDieInCar
)

// Flags for byte[13]
const (
	checkColAboveHead uint8 = 1 << (1 * iota)
	ignoreWeaponRange
	druggedUp
	wantedByPolice
	signalAfterKill
	canClimbOntoBoat
	pedHitWallLastFrame
	ignoreHeightDifferenceFollowingNodes
)

// Flags for byte[14]
const (
	moveAnimSpeedHasBeenSetByTask uint8 = 1 << (1 * iota)
	getOutUpsideDownCar
	justGotOffTrain
	deathPickupsPersist
	testForShotInVehicle
	usedForReplayPlayerState
)
