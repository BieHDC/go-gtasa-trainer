package types

import (
	"fmt"
	"strings"
	"unsafe"

	. "gtasamod/structoffsetvalidator"
)

// Wanted pool start (CWanted). Each slot has 668 bytes of data.
// 668 0x29C
type CWanted struct {
	Level                       uint32      `offset:"0x0"`
	LevelBeforeParole           uint32      `offset:"0x4"`
	LastDecreased               uint32      `offset:"0x8"`
	LastChanged                 uint32      `offset:"0xc"`
	TimeOfParole                uint32      `offset:"0x10"`
	ScoreMultiplier             float32     `offset:"0x14"`
	NumCopsInPursuit            byte        `offset:"0x18"`
	MaxCopsOnFootShooting       byte        `offset:"0x19"`
	MaxCopsInCarInPursuit       byte        `offset:"0x1a"`
	NumCopsBeatingSuspect       byte        `offset:"0x1b"`
	ChanceOfRoadblock           uint16      `offset:"0x1C"`
	IsPlayerIgnoredByCopsScript byte        `offset:"0x1e" options:"0=false,1=true"`
	IsPlayerIgnoredByCopsGarage byte        `offset:"0x1f" options:"0=false,1=true"`
	IsPlayerIgnoredByEveryone   byte        `offset:"0x20" options:"0=false,1=true"`
	StreamerShouldLoadSWAT      byte        `offset:"0x21"`
	StreamerShouldLoadFBI       byte        `offset:"0x22"`
	StreamerShouldLoadArmy      byte        `offset:"0x23"`
	ChaseTime                   uint32      `offset:"0x24"`
	ChaseTimeCounter            uint32      `offset:"0x28"`
	LevelAsStars                uint32      `offset:"0x2c"`
	LevelAsStarsBeforeParole    uint32      `offset:"0x30"`
	_                           [0x258]byte //padding/stuff i dont care about
	LeavePlayerAlone            uint32      `offset:"0x28c"`
	MaximumWantedLevel          *uint32     `offset:"0x290"`
	MaximumChaosLevel           *uint32     `offset:"0x294"`
	_                           [0x4]byte   //padding
}

func init() {
	var cwanted CWanted
	err := ValidateOffsets(&cwanted)
	if err != nil {
		panic(err)
	}
	sz := unsafe.Sizeof(cwanted)
	if sz != 668 {
		panic(fmt.Sprintf("CWanted size is wrong, expected %d, got %d\n", 668, sz))
	}
}

func (cw *CWanted) String() string {
	var out strings.Builder

	out.WriteString("Wanted: ")
	out.WriteString(fmt.Sprintf("Level:(%d) ", cw.Level))
	out.WriteString(fmt.Sprintf("Level before Parole:(%d) ", cw.LevelBeforeParole))
	//ignore boring ones
	out.WriteString(fmt.Sprintf("Time Of Parole:(%d) ", cw.TimeOfParole))
	out.WriteString(fmt.Sprintf("Num Cops in Pursuit:(%d) ", cw.NumCopsInPursuit))
	out.WriteString(fmt.Sprintf("Stars:(%d)\n", cw.LevelAsStars))

	return out.String()
}
