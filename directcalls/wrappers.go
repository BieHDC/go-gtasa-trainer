package directcalls

/*
// Shared definitions
#include "shared.h"

// Functions
void CPlayerPed_SetWantedLevel(void* cped, int32_t level);

void CPed_Teleport(void* cped, CVector destination, bool resetRotation);
void CAutomobile_Teleport(void* vehicle, CVector destination, bool resetRotation);

extern const uint16_t STYLE_MIDDLE;                // In The Middle
extern const uint16_t STYLE_BOTTOM_RIGHT;          // At The Bottom Right
extern const uint16_t STYLE_WHITE_MIDDLE;          // White Text In The Middle
extern const uint16_t STYLE_MIDDLE_SMALLER;        // In The Middle Smaller
extern const uint16_t STYLE_MIDDLE_SMALLER_HIGHER; // In The Middle Smaller A Bit Higher On The Screen
extern const uint16_t STYLE_WHITE_MIDDLE_SMALLER;  // Small White Text In The Middle Of The Screen
extern const uint16_t STYLE_LIGHT_BLUE_TOP;        // Light Blue Text On Top Of The Screen
void AddBigMessageQ(const char* text, uint32_t time, uint16_t style);
void CHud_SetHelpMessage(const char* text, bool quickMessage, bool permanent, bool addToBrief);
void CMessages_AddMessageQ(const char* text, uint32_t time, uint16_t flag, bool bPreviousBrief);

void CPed_GiveWeapon(void* cped, uint32_t weaponType, uint32_t ammo);

void MakePickup(CVector pos, uint16_t modelID, uint8_t pickuptype, uint32_t ammo, uint32_t moneyPerDay);

void* CCheat_VehicleCheat(int32_t modelid);
void CVehicle_DestroyVehicleAndDriverAndPassengers(void* cvehicle);
void CMatrix_SetRotateZOnly(void* matrix, float angle);
*/
import "C"

import (
	"strings"
	"time"
	"unsafe"

	. "gtasamod/types"
)

// Helper function
// Since we copy the contents, we can always safely pass the pointer
func MakeCCVector(orig *CVector) C.CVector {
	return C.CVector{
		X: C.float(orig.X),
		Y: C.float(orig.Y),
		Z: C.float(orig.Z),
	}
}

// fixme find the right encoder or make one yourself
// "golang.org/x/text/encoding/charmap"
// var gtaencoder = charmap.Windows1252.NewEncoder()
// Right now only repaces \n with ~n~
// map:
// ß = 150 0x96
// ä = 154 0x9a
// ö = 168 0xa8
// ü = 172 0xac
func ToGtaSaString(s string) string {
	//var err error
	var result string
	result = strings.ReplaceAll(s, "\n", "~n~")
	//result, err = gtaencoder.String(result)
	//if err != nil {
	//	result = s //just ignore
	//}
	return result
}

func Cheat_NewVehicle(modelid int32) *CVehicle {
	if modelid < 400 || modelid > 611 {
		//not a valid car id
		return nil
	}
	return (*CVehicle)(C.CCheat_VehicleCheat(C.int32_t(modelid)))
}

func Vehicle_Destroy(cvehicle *CVehicle) {
	C.CVehicle_DestroyVehicleAndDriverAndPassengers(unsafe.Pointer(cvehicle))
}

func Matrix_SetRotateZOnly(matrix *CMatrix, angle float32) {
	C.CMatrix_SetRotateZOnly(unsafe.Pointer(matrix), C.float(angle))
}

// Here we keep all functions we dont hook and just call directly
func PlayerPed_SetWantedLevel(playerped *CPed, level int) {
	C.CPlayerPed_SetWantedLevel(unsafe.Pointer(playerped), C.int32_t(level))
}

func Ped_Teleport(cped *CPed, pos *CVector, resetRotation bool, Zoff float32) {
	cvec := MakeCCVector(pos)
	cvec.Z += C.float(Zoff)
	C.CPed_Teleport(unsafe.Pointer(cped), cvec, C.bool(resetRotation))
}

func Automobile_Teleport(vehicle *CVehicle, pos *CVector, resetRotation bool) {
	C.CAutomobile_Teleport(unsafe.Pointer(vehicle), MakeCCVector(pos), C.bool(resetRotation))
}

func Messages_AddBigMessageQWithDuration(str string, duration uint32) {
	go func(str string, duration uint32) {
		cstr := C.CString(str)
		C.AddBigMessageQ(cstr, (C.uint32_t)(duration), C.STYLE_LIGHT_BLUE_TOP)
		//we need to keep the memory around until its done, with extra timeout
		<-time.After(time.Millisecond * time.Duration(duration*2))
		C.free(unsafe.Pointer(cstr))
	}(str, duration)
}

func Messages_AddBigMessageQ(str string) {
	Messages_AddBigMessageQWithDuration(str, 1500)
}

// fixme there is some kind of limit of max tokens?
func Hud_SetHelpMessage(message string, quickMessage, permanent, addToBrief bool) {
	str := C.CString(strings.TrimSuffix(ToGtaSaString(message), " ")) //there is a bug if the messages last char is a space
	C.CHud_SetHelpMessage(str, C.bool(quickMessage), C.bool(permanent), C.bool(addToBrief))
	C.free(unsafe.Pointer(str))
}

func Messages_AddMessageQ(msg string, time uint32, flag uint16, bPreviousBrief bool) {
	// unlike bigmessage, this does have a copy function
	text := C.CString(msg)
	C.CMessages_AddMessageQ(text, C.uint32_t(time), C.uint16_t(flag), C.bool(bPreviousBrief))
	C.free(unsafe.Pointer(text))
}

func Ped_GiveWeapon(cped *CPed, wt WeaponType, ammo uint32) {
	C.CPed_GiveWeapon(unsafe.Pointer(cped), (C.uint32_t)(wt), (C.uint32_t)(ammo))
}

func CPickups_MakePickup(pos *CVector, modelID uint16, pickuptype uint8, ammo uint32, moneyPerDay uint32) {
	C.MakePickup(MakeCCVector(pos), C.uint16_t(modelID), C.uint8_t(pickuptype), C.uint32_t(ammo), C.uint32_t(moneyPerDay))
}
