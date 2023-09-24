package main

/*
// Internal Use Stuff
void GoRuntimeHasFullyLoaded(void);
*/
import "C"

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	. "gtasamod/madresser"
	. "gtasamod/types"
)

// Used for rundll32 dry testing
//
//export Ccalls
func Ccalls() {
	fmt.Println("go called from c")
}

func init() {
	//fmt.Println("go init called")

	C.GoRuntimeHasFullyLoaded()
}

//export GODLLmain
func GODLLmain() {
	// After going from init() -> C.GoRuntimeHasFullyLoaded -> GODLLmain,
	// we kick off our main "thread" now that we know 100% everything is
	// taken care of in terms of loading.
	go main()
}

// Called by GODLLmain, or normally when not being a dll
func main() {
	/*
		// Useful for analysation
		go func() {
			for {
				time.Sleep(30 * time.Second)
				dumpStackTraces()
			}
		}()
	*/
	setupLogging()

	// we need to wait here so windows can fully load the main exe,
	// otherwise we will silently fail and exit. This variable will
	// be not zero once the game starts starting, after which we can
	// do our hooks.
	dxptr := TypeAtAbsolute[uint32](0xC97C28) //pointer to IDirect3DDevice9
	for {
		if *dxptr != 0 {
			break
		}
		runtime.Gosched()
	}

	SetupFunctionsHooks() // gtasa_functionhooks.go

	select {
	//Block until we get unloaded. Maybe do something useful here.
	}
}

type GameInternals struct {
	player         *CPlayer
	wanted         *CWanted
	currentvehicle **CVehicle
}

// For GameFunc:
// return true to run again next iteration
// return false to get removed from the run queue
// also just return false if you want to run it once
type GameFunc func(*GameInternals) bool

var extfuncs []GameFunc
var extfuncsmutex sync.Mutex

func AddFn(fn GameFunc) {
	extfuncsmutex.Lock()
	extfuncs = append(extfuncs, fn)
	extfuncsmutex.Unlock()
}

//fixme todo:
// Might be intersting to do something with it
// 0xB73458 – Start of controls block.
//    +0x20 = [word] Accelerate:
//        0 = off
//        255 = on
//    +0x22 = [word] Brake
// https://gtamods.com/wiki/Memory_Addresses_(SA)

var gameInternals GameInternals
var gamehasinit bool

// We use this to initialise everything we need once on the first
// ingame frame, and then rewire the call to the bare function.
//
//export RenderSceneFirstGo
func RenderSceneFirstGo() {
	//log.Println("Ingame Render")
	if gamehasinit {
		log.Println("if you see this every frame it didnt rehook correctly to the more efficient function")
	}

	gameInternals = GameInternals{
		// the player struct is store there
		player: TypeAtAbsolute[CPlayer](0xB7CD98),
		wanted: *TypeAtAbsolute[*CWanted](0xB7CD9C),
		//0xBA18FC – Current vehicle pointer:
		//    0 = on-foot
		//    >0 = in-car
		currentvehicle: TypeAtAbsolute[*CVehicle](0xBA18FC),
	}
	gamehasinit = true

	go launchWeb(&gameInternals)

	// rewire this call to RenderSceneGo
	SetRenderFuncLooping()

	// run the custom functions here, afterwards
	// that function will be fun directly
	RenderSceneGo()
}

// This is the bare function the other comment talked about.
//
//export RenderSceneGo
func RenderSceneGo() {
	i := 0
	extfuncsmutex.Lock()
	for _, fn := range extfuncs {
		if fn(&gameInternals) {
			extfuncs[i] = fn
			i++
		}
	}
	extfuncs = extfuncs[:i]
	extfuncsmutex.Unlock()
}

var quitfuncs []func()
var quitfuncsmutex sync.Mutex

func AddQuitFn(fn func()) {
	quitfuncsmutex.Lock()
	quitfuncs = append(quitfuncs, fn)
	quitfuncsmutex.Unlock()
}

//export GameIsQuitting
func GameIsQuitting() {
	log.Println("Game is quitting")
	for _, fn := range quitfuncs {
		fn()
	}
	time.Sleep(1 * time.Second)
	dumpStackTraces()
}

//export TypedChars
func TypedChars(c uint32, pinputting *bool) bool {
	// We only care about input after we are fully loaded
	if !gamehasinit {
		return false
	}

	return Menuhandler(c, pinputting)
}
