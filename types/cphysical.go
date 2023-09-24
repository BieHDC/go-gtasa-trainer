package types

import (
	"fmt"
	"unsafe"

	. "gtasamod/complextypes"
	. "gtasamod/structoffsetvalidator"
)

type CPhysical struct {
	_           [0x4]byte        //padding
	Posn        CSimpleTransform `offset:"0x4"` //used when Pos == 0,0,0
	Pos         *CMatrix         `offset:"0x14"`
	_           [0xa]byte        //padding
	ModelID     int16            `offset:"0x22"`
	_           [0xb]byte        //padding
	InInteriour InteriorState    `offset:"0x2F"`
	_           [0x6]byte        //padding
	EntityType  EntityType       `offset:"0x36"`
	_           [0x8]byte        //padding
	PhysFlags   PhysicsFlags     `offset:"0x40"`
	MoveSpeed   CVector          `offset:"0x44"`
	_           [0x3c]byte       //padding
	Mass        float32          `offset:"0x8c"`
	_           [0xA8]byte       //padding
}

func init() {
	var cphys CPhysical
	err := ValidateOffsets(&cphys)
	if err != nil {
		panic(err)
	}
	sz := unsafe.Sizeof(cphys)
	if sz != 312 {
		panic(fmt.Sprintf("CPed size is wrong, expected %d, got %d\n", 312, sz))
	}
}

// Flags for PhysicsFlags.Movement
const (
	makeMassTwiceAsBig uint8 = 1 << (1 * iota)
	applyGravity
	disableCollisionForce
	collidable
	disableTurnForce
	disableMoveForce
	infiniteMass
	disableZ
)

func (pf *PhysicsFlags) ApplyGravity() bool {
	return GetBit(pf.Movement, applyGravity)
}

func (pf *PhysicsFlags) ApplyGravityToggle(b bool) {
	if b {
		pf.Movement = SetBit(pf.Movement, applyGravity)
	} else {
		pf.Movement = ClearBit(pf.Movement, applyGravity)
	}
}

func (pf *PhysicsFlags) CollidableToggle(b bool) {
	if b {
		pf.Movement = SetBit(pf.Movement, collidable)
	} else {
		pf.Movement = ClearBit(pf.Movement, collidable)
	}
}

type PhysicsFlags struct {
	/*
	   +0x0 = Movement physics flags
	       Bit 1 = Unknown
	       Bit 2 = Apply Gravity
	       Bit 3 = Disable Collision Force
	       Bit 4 = Collidable
	       Bit 5 = Disable Turn Force
	       Bit 6 = Disable Move Force
	       Bit 7 = Infinite Mass
	       Bit 8 = Disable Z
	*/
	Movement byte

	/*
	   +0x1 = Surface physics flags
	       1 = Submerged In Water
	       2 = On Solid Surface
	       4 = Broken
	       8 = bProcessCollisionEvenIfStationary
	       16 = bSkipLineCol
	       32 = Don't Apply Speed
	       64 = b15
	       128 = bProcessingShift
	*/
	Surface byte

	/*
	   +0x2 = Special physics flags
	       1 = Soft (in other words noclip)
	       2 = Freeze
	       4 = Bullet-Proof
	       8 = Fire-Proof
	       16 = Collision-Proof (prevent fall damage)
	       32 = Melee-Proof
	       64 = Melee-Proof, Bullet-Proof, Collision-Proof
	       128 = Explosion-Proof
	*/
	Special byte

	/*
	   +0x3 = Collision physics flag
	       1 = bDontCollideWithFlyers
	       2 = Attached To Entity
	       4 = bAddMovingCollisionSpeed
	       8 = Touching Water
	       16 = Can Be Collided with
	       32 = Destroyed
	       64 = b31
	       128 = b32
	*/
	Collision byte
}

func init() {
	sz := unsafe.Sizeof(PhysicsFlags{})
	if sz != 4 {
		panic(fmt.Sprintf("PhysicsFlags size is wrong, expected %d, got %d\n", 4, sz))
	}
}

func (pf *PhysicsFlags) String() string {
	return fmt.Sprintf("%d %d %d %d", pf.Movement, pf.Surface, pf.Special, pf.Collision)
}
