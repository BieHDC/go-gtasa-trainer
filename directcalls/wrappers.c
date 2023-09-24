#include "shared.h"

// 0x609F10
typedef void(__thiscall *fpCPlayerPed_SetWantedLevel)(void *cped,
                                                      int32_t level);
void CPlayerPed_SetWantedLevel(void *cped, int32_t level) {
  ((fpCPlayerPed_SetWantedLevel)0x609F10)(cped, level);
}

// 0x5E4110
typedef void(__thiscall *fpCPed_Teleport)(void *cped, CVector destination,
                                          bool resetRotation);
void CPed_Teleport(void *cped, CVector destination, bool resetRotation) {
  ((fpCPed_Teleport)0x5E4110)(cped, destination, resetRotation);
}

// 0x6A9CA0
typedef void(__thiscall *fpCAutomobile_Teleport)(void *vehicle,
                                                 CVector destination,
                                                 bool resetRotation);
void CAutomobile_Teleport(void *vehicle, CVector destination,
                          bool resetRotation) {
  ((fpCAutomobile_Teleport)0x6A9CA0)(vehicle, destination, resetRotation);
}

// 0x43A0B0
// CVehicle* CCheat::VehicleCheat(eModelID modelId)
typedef void *(*fpCCheat_VehicleCheat)(int32_t);
void *CCheat_VehicleCheat(int32_t modelid) {
  return ((fpCCheat_VehicleCheat)0x43A0B0)(modelid);
}

// 0x6D2250
// void CVehicle::DestroyVehicleAndDriverAndPassengers(CVehicle* vehicle)
typedef void (*fpCVehicle_DestroyVehicleAndDriverAndPassengers)(void *);
void CVehicle_DestroyVehicleAndDriverAndPassengers(void *cvehicle) {
  ((fpCVehicle_DestroyVehicleAndDriverAndPassengers)0x6D2250)(cvehicle);
}

// 0x59B020
typedef void(__thiscall *fpCMatrix_SetRotateZOnly)(void *matrix, float angle);
void CMatrix_SetRotateZOnly(void *matrix, float angle) {
  ((fpCMatrix_SetRotateZOnly)0x59B020)(matrix, angle);
}

// Go doesnt like weird enums
// enum eMessageStyle : uint16_t
const uint16_t STYLE_MIDDLE = 0;         // In The Middle
const uint16_t STYLE_BOTTOM_RIGHT = 1;   // At The Bottom Right
const uint16_t STYLE_WHITE_MIDDLE = 2;   // White Text In The Middle
const uint16_t STYLE_MIDDLE_SMALLER = 3; // In The Middle Smaller
const uint16_t STYLE_MIDDLE_SMALLER_HIGHER =
    4; // In The Middle Smaller A Bit Higher On The Screen
const uint16_t STYLE_WHITE_MIDDLE_SMALLER =
    5; // Small White Text In The Middle Of The Screen
const uint16_t STYLE_LIGHT_BLUE_TOP = 6; // Light Blue Text On Top Of The Screen

// 0x69F370
typedef void (*fpAddBigMessageQ)(const char *text, uint32_t time,
                                 uint16_t style);
void AddBigMessageQ(const char *text, uint32_t time, uint16_t style) {
  ((fpAddBigMessageQ)0x69F370)(text, time, style);
}

// 0x69F0B0
typedef void (*fpCMessages_AddMessageQ)(const char *text, uint32_t time,
                                        uint16_t flag, bool bPreviousBrief);
void CMessages_AddMessageQ(const char *text, uint32_t time, uint16_t flag,
                           bool bPreviousBrief) {
  ((fpCMessages_AddMessageQ)0x69F0B0)(text, time, flag, bPreviousBrief);
}

// 0x588BE0
typedef void (*fpCHud_SetHelpMessage)(const char *text, bool quickMessage,
                                      bool permanent, bool addToBrief);
void CHud_SetHelpMessage(const char *text, bool quickMessage, bool permanent,
                         bool addToBrief) {
  ((fpCHud_SetHelpMessage)0x588BE0)(text, quickMessage, permanent, addToBrief);
}

typedef void(__thiscall *CPed_GiveDelayedWeapon)(void *cped,
                                                 uint32_t weaponType,
                                                 uint32_t ammo);
void CPed_GiveWeapon(void *cped, uint32_t weaponType, uint32_t ammo) {
  ((CPed_GiveDelayedWeapon)0x5E89B0)(cped, weaponType, ammo);
  // player->SetCurrentWeapon(WEAPON_MICRO_UZI);
}

// 0x456F20
typedef int32_t (*CPickups_GenerateNewOne)(CVector coors, uint32_t modelId,
                                           uint8_t pickupType, uint32_t ammo,
                                           uint32_t moneyPerDay, bool isEmpty,
                                           char *message);
void MakePickup(CVector pos, uint16_t modelID, uint8_t pickuptype,
                uint32_t ammo, uint32_t moneyPerDay) {
  ((CPickups_GenerateNewOne)0x456F20)(pos, (uint32_t)modelID, pickuptype, ammo,
                                      moneyPerDay, false, NULL);
}
/*
// example go code
var newpos C.CVector
newpos.X = (C.float)(player.Ped.Pos.Position.X)
newpos.Y = (C.float)(player.Ped.Pos.Position.Y + 5)
newpos.Z = (C.float)(player.Ped.Pos.Position.Z)
modelID := *TypeAtAbsolute[C.uint16_t](0x8CD758)
pickuptype := C.uint8_t(10) //PICKUP_MINE_ARMED //only trigger if player in
vehicle ammo := C.uint32_t(0) //??? moneyPerDay := C.uint32_t(0) //???
C.MakePickup(newpos, modelID, pickuptype, ammo, moneyPerDay)
*/
/*
enum ePickupType : uint8 {
    PICKUP_NONE = 0,
    PICKUP_IN_SHOP = 1,
    PICKUP_ON_STREET = 2,
    PICKUP_ONCE = 3,
    PICKUP_ONCE_TIMEOUT = 4,
    PICKUP_ONCE_TIMEOUT_SLOW = 5,
    PICKUP_COLLECTABLE1 = 6,
    PICKUP_IN_SHOP_OUT_OF_STOCK = 7,
    PICKUP_MONEY = 8,
    PICKUP_MINE_INACTIVE = 9,
    PICKUP_MINE_ARMED = 10,
    PICKUP_NAUTICAL_MINE_INACTIVE = 11,
    PICKUP_NAUTICAL_MINE_ARMED = 12,
    PICKUP_FLOATINGPACKAGE = 13,
    PICKUP_FLOATINGPACKAGE_FLOATING = 14,
    PICKUP_ON_STREET_SLOW = 15,
    PICKUP_ASSET_REVENUE = 16,
    PICKUP_PROPERTY_LOCKED = 17,
    PICKUP_PROPERTY_FORSALE = 18,
    PICKUP_MONEY_DOESNTDISAPPEAR = 19,
    PICKUP_SNAPSHOT = 20,
    PICKUP_2P = 21,
    PICKUP_ONCE_FOR_MISSION = 22
};

namespace ModelIndices {
    ModelIndex& MI_TRAFFICLIGHTS = *(ModelIndex*)0x8CD4F4;
    ModelIndex& MI_TRAFFICLIGHTS_VERTICAL = *(ModelIndex*)0x8CD4F8;
    ModelIndex& MI_TRAFFICLIGHTS_MIAMI = *(ModelIndex*)0x8CD4FC;
    ModelIndex& MI_TRAFFICLIGHTS_VEGAS = *(ModelIndex*)0x8CD500;
    ModelIndex& MI_TRAFFICLIGHTS_TWOVERTICAL = *(ModelIndex*)0x8CD504;
    ModelIndex& MI_TRAFFICLIGHTS_3 = *(ModelIndex*)0x8CD508;
    ModelIndex& MI_TRAFFICLIGHTS_4 = *(ModelIndex*)0x8CD50C;
    ModelIndex& MI_TRAFFICLIGHTS_5 = *(ModelIndex*)0x8CD510;
    ModelIndex& MI_TRAFFICLIGHTS_GAY = *(ModelIndex*)0x8CD514;
    ModelIndex& MI_SINGLESTREETLIGHTS1 = *(ModelIndex*)0x8CD518;
    ModelIndex& MI_SINGLESTREETLIGHTS2 = *(ModelIndex*)0x8CD51C;
    ModelIndex& MI_SINGLESTREETLIGHTS3 = *(ModelIndex*)0x8CD520;
    ModelIndex& MI_DOUBLESTREETLIGHTS = *(ModelIndex*)0x8CD524;
    ModelIndex& MI_STREETLAMP1 = *(ModelIndex*)0x8CD528;
    ModelIndex& MI_STREETLAMP2 = *(ModelIndex*)0x8CD52C;
    ModelIndex& MODELID_CRANE_1 = *(ModelIndex*)0x8CD530;
    ModelIndex& MODELID_CRANE_2 = *(ModelIndex*)0x8CD534;
    ModelIndex& MODELID_CRANE_3 = *(ModelIndex*)0x8CD538;
    ModelIndex& MODELID_CRANE_4 = *(ModelIndex*)0x8CD53C;
    ModelIndex& MODELID_CRANE_5 = *(ModelIndex*)0x8CD540;
    ModelIndex& MODELID_CRANE_6 = *(ModelIndex*)0x8CD544;
    ModelIndex& MI_PARKINGMETER = *(ModelIndex*)0x8CD548;
    ModelIndex& MI_PARKINGMETER2 = *(ModelIndex*)0x8CD54C;
    ModelIndex& MI_MALLFAN = *(ModelIndex*)0x8CD550;
    ModelIndex& MI_HOTELFAN_NIGHT = *(ModelIndex*)0x8CD554;
    ModelIndex& MI_HOTELFAN_DAY = *(ModelIndex*)0x8CD558;
    ModelIndex& MI_HOTROOMFAN = *(ModelIndex*)0x8CD55C;
    ModelIndex& MI_PHONEBOOTH1 = *(ModelIndex*)0x8CD560;
    ModelIndex& MI_WASTEBIN = *(ModelIndex*)0x8CD564;
    ModelIndex& MI_BIN = *(ModelIndex*)0x8CD568;
    ModelIndex& MI_POSTBOX1 = *(ModelIndex*)0x8CD56C;
    ModelIndex& MI_NEWSSTAND = *(ModelIndex*)0x8CD570;
    ModelIndex& MI_TRAFFICCONE = *(ModelIndex*)0x8CD574;
    ModelIndex& MI_DUMP1 = *(ModelIndex*)0x8CD578;
    ModelIndex& MI_ROADWORKBARRIER1 = *(ModelIndex*)0x8CD57C;
    ModelIndex& MI_ROADBLOCKFUCKEDCAR1 = *(ModelIndex*)0x8CD580;
    ModelIndex& MI_ROADBLOCKFUCKEDCAR2 = *(ModelIndex*)0x8CD584;
    ModelIndex& MI_BUSSIGN1 = *(ModelIndex*)0x8CD588;
    ModelIndex& MI_NOPARKINGSIGN1 = *(ModelIndex*)0x8CD58C;
    ModelIndex& MI_PHONESIGN = *(ModelIndex*)0x8CD590;
    ModelIndex& MI_FIRE_HYDRANT = *(ModelIndex*)0x8CD594;
    ModelIndex& MI_COLLECTABLE1 = *(ModelIndex*)0x8CD598;
    ModelIndex& MI_MONEY = *(ModelIndex*)0x8CD59C;
    ModelIndex& MI_CARMINE = *(ModelIndex*)0x8CD5A0;
    ModelIndex& MI_NAUTICALMINE = *(ModelIndex*)0x8CD5A4;
    ModelIndex& MI_TELLY = *(ModelIndex*)0x8CD5A8;
    ModelIndex& MI_BRIEFCASE = *(ModelIndex*)0x8CD5AC;
    ModelIndex& MI_GLASS1 = *(ModelIndex*)0x8CD5B0;
    ModelIndex& MI_GLASS8 = *(ModelIndex*)0x8CD5B4;
    ModelIndex& MI_EXPLODINGBARREL = *(ModelIndex*)0x8CD5B8;
    ModelIndex& MI_PICKUP_ADRENALINE = *(ModelIndex*)0x8CD5BC;
    ModelIndex& MI_PICKUP_BODYARMOUR = *(ModelIndex*)0x8CD5C0;
    ModelIndex& MI_PICKUP_INFO = *(ModelIndex*)0x8CD5C4;
    ModelIndex& MI_PICKUP_HEALTH = *(ModelIndex*)0x8CD5C8;
    ModelIndex& MI_PICKUP_BONUS = *(ModelIndex*)0x8CD5CC;
    ModelIndex& MI_PICKUP_BRIBE = *(ModelIndex*)0x8CD5D0;
    ModelIndex& MI_PICKUP_KILLFRENZY = *(ModelIndex*)0x8CD5D4;
    ModelIndex& MI_PICKUP_CAMERA = *(ModelIndex*)0x8CD5D8;
    ModelIndex& MI_PICKUP_PARACHUTE = *(ModelIndex*)0x8CD5DC;
    ModelIndex& MI_PICKUP_REVENUE = *(ModelIndex*)0x8CD5E0;
    ModelIndex& MI_PICKUP_SAVEGAME = *(ModelIndex*)0x8CD5E4;
    ModelIndex& MI_PICKUP_PROPERTY = *(ModelIndex*)0x8CD5E8;
    ModelIndex& MI_PICKUP_PROPERTY_FORSALE = *(ModelIndex*)0x8CD5EC;
    ModelIndex& MI_PICKUP_CLOTHES = *(ModelIndex*)0x8CD5F0;
    ModelIndex& MI_PICKUP_2P_KILLFRENZY = *(ModelIndex*)0x8CD5F4;
    ModelIndex& MI_PICKUP_2P_COOP = *(ModelIndex*)0x8CD5F8;
    ModelIndex& MI_BOLLARDLIGHT = *(ModelIndex*)0x8CD5FC;
    ModelIndex& MI_FENCE = *(ModelIndex*)0x8CD600;
    ModelIndex& MI_FENCE2 = *(ModelIndex*)0x8CD604;
    ModelIndex& MI_BUOY = *(ModelIndex*)0x8CD608;
    ModelIndex& MI_PARKTABLE = *(ModelIndex*)0x8CD60C;
    ModelIndex& MI_LAMPPOST1 = *(ModelIndex*)0x8CD610;
    ModelIndex& MI_MLAMPPOST = *(ModelIndex*)0x8CD614;
    ModelIndex& MI_BARRIER1 = *(ModelIndex*)0x8CD618;
    ModelIndex& MI_LITTLEHA_POLICE = *(ModelIndex*)0x8CD61C;
    ModelIndex& MI_TELPOLE02 = *(ModelIndex*)0x8CD620;
    ModelIndex& MI_TRAFFICLIGHT01 = *(ModelIndex*)0x8CD624;
    ModelIndex& MI_PARKBENCH = *(ModelIndex*)0x8CD628;
    ModelIndex& MI_LIGHTBEAM = *(ModelIndex*)0x8CD62C;
    ModelIndex& MI_AIRPORTRADAR = *(ModelIndex*)0x8CD630;
    ModelIndex& MI_RCBOMB = *(ModelIndex*)0x8CD634;
    ModelIndex& MI_BEACHBALL = *(ModelIndex*)0x8CD638;
    ModelIndex& MI_SANDCASTLE1 = *(ModelIndex*)0x8CD63C;
    ModelIndex& MI_SANDCASTLE2 = *(ModelIndex*)0x8CD640;
    ModelIndex& MI_JELLYFISH = *(ModelIndex*)0x8CD644;
    ModelIndex& MI_JELLYFISH01 = *(ModelIndex*)0x8CD648;
    ModelIndex& MI_FISH1SINGLE = *(ModelIndex*)0x8CD64C;
    ModelIndex& MI_FISH1S = *(ModelIndex*)0x8CD650;
    ModelIndex& MI_FISH2SINGLE = *(ModelIndex*)0x8CD654;
    ModelIndex& MI_FISH2S = *(ModelIndex*)0x8CD658;
    ModelIndex& MI_FISH3SINGLE = *(ModelIndex*)0x8CD65C;
    ModelIndex& MI_FISH3S = *(ModelIndex*)0x8CD660;
    ModelIndex& MI_TURTLE = *(ModelIndex*)0x8CD664;
    ModelIndex& MI_DOLPHIN = *(ModelIndex*)0x8CD668;
    ModelIndex& MI_SHARK = *(ModelIndex*)0x8CD66C;
    ModelIndex& MI_SUBMARINE = *(ModelIndex*)0x8CD670;
    ModelIndex& MI_ESCALATORSTEP = *(ModelIndex*)0x8CD674;
    ModelIndex& MI_ESCALATORSTEP8 = *(ModelIndex*)0x8CD678;
    ModelIndex& MI_LOUNGE_WOOD_UP = *(ModelIndex*)0x8CD67C;
    ModelIndex& MI_LOUNGE_TOWEL_UP = *(ModelIndex*)0x8CD680;
    ModelIndex& MI_LOUNGE_WOOD_DN = *(ModelIndex*)0x8CD684;
    ModelIndex& MI_LOTION = *(ModelIndex*)0x8CD688;
    ModelIndex& MI_BEACHTOWEL01 = *(ModelIndex*)0x8CD68C;
    ModelIndex& MI_BEACHTOWEL02 = *(ModelIndex*)0x8CD690;
    ModelIndex& MI_BEACHTOWEL03 = *(ModelIndex*)0x8CD694;
    ModelIndex& MI_BEACHTOWEL04 = *(ModelIndex*)0x8CD698;
    ModelIndex& MI_BLIMP_NIGHT = *(ModelIndex*)0x8CD69C;
    ModelIndex& MI_BLIMP_DAY = *(ModelIndex*)0x8CD6A0;
    ModelIndex& MI_YT_MAIN_BODY = *(ModelIndex*)0x8CD6A4;
    ModelIndex& MI_YT_MAIN_BODY2 = *(ModelIndex*)0x8CD6A8;
    ModelIndex& MI_SAMSITE = *(ModelIndex*)0x8CD6AC;
    ModelIndex& MI_SAMSITE2 = *(ModelIndex*)0x8CD6B0;
    ModelIndex& MI_TRAINCROSSING = *(ModelIndex*)0x8CD6B4;
    ModelIndex& MI_TRAINCROSSING1 = *(ModelIndex*)0x8CD6B8;
    ModelIndex& MI_MAGNOCRANE = *(ModelIndex*)0x8CD6BC;
    ModelIndex& MI_CRANETROLLEY = *(ModelIndex*)0x8CD6C0;
    ModelIndex& MI_QUARRYCRANE_ARM = *(ModelIndex*)0x8CD6C4;
    ModelIndex& MI_OBJECTFORMAGNOCRANE1 = *(ModelIndex*)0x8CD6C8;
    ModelIndex& MI_OBJECTFORMAGNOCRANE2 = *(ModelIndex*)0x8CD6CC;
    ModelIndex& MI_OBJECTFORMAGNOCRANE3 = *(ModelIndex*)0x8CD6D0;
    ModelIndex& MI_OBJECTFORMAGNOCRANE4 = *(ModelIndex*)0x8CD6D4;
    ModelIndex& MI_OBJECTFORMAGNOCRANE5 = *(ModelIndex*)0x8CD6D8;
    ModelIndex& MI_OBJECTFORBUILDINGSITECRANE1 = *(ModelIndex*)0x8CD6DC;
    ModelIndex& MI_MAGNOCRANE_HOOK = *(ModelIndex*)0x8CD6E0;
    ModelIndex& MI_HARVESTERBODYPART1 = *(ModelIndex*)0x8CD6E4;
    ModelIndex& MI_HARVESTERBODYPART2 = *(ModelIndex*)0x8CD6E8;
    ModelIndex& MI_HARVESTERBODYPART3 = *(ModelIndex*)0x8CD6EC;
    ModelIndex& MI_HARVESTERBODYPART4 = *(ModelIndex*)0x8CD6F0;
    ModelIndex& MI_GRASSHOUSE = *(ModelIndex*)0x8CD6F4;
    ModelIndex& MI_GRASSPLANT = *(ModelIndex*)0x8CD6F8;
    ModelIndex& MI_CRANE_HARNESS = *(ModelIndex*)0x8CD6FC;
    ModelIndex& MI_CRANE_MAGNET = *(ModelIndex*)0x8CD700;
    ModelIndex& MI_QUARY_ROCK1 = *(ModelIndex*)0x8CD704;
    ModelIndex& MI_QUARY_ROCK2 = *(ModelIndex*)0x8CD708;
    ModelIndex& MI_QUARY_ROCK3 = *(ModelIndex*)0x8CD70C;
    ModelIndex& MI_ATM = *(ModelIndex*)0x8CD710;
    ModelIndex& MI_DEAD_TIED_COP = *(ModelIndex*)0x8CD714;
    ModelIndex& MI_WINDSOCK = *(ModelIndex*)0x8CD718;
    ModelIndex& MI_WRECKING_BALL = *(ModelIndex*)0x8CD71C;
    ModelIndex& MI_FREEFALL_BOMB = *(ModelIndex*)0x8CD720;
    ModelIndex& MI_WONG_DISH = *(ModelIndex*)0x8CD724;
    ModelIndex& MI_GANG_DRINK = *(ModelIndex*)0x8CD728;
    ModelIndex& MI_GANG_SMOKE = *(ModelIndex*)0x8CD72C;
    ModelIndex& MI_RHYMESBOOK = *(ModelIndex*)0x8CD730;
    ModelIndex& MI_KMB_ROCK = *(ModelIndex*)0x8CD734;
    ModelIndex& MI_KMB_PLANK = *(ModelIndex*)0x8CD738;
    ModelIndex& MI_KMB_BOMB = *(ModelIndex*)0x8CD73C;
    ModelIndex& MI_MINI_MAGNET = *(ModelIndex*)0x8CD740;
    ModelIndex& MI_HANGING_CARCASS = *(ModelIndex*)0x8CD744;
    ModelIndex& MI_IMY_SHASH_WALL = *(ModelIndex*)0x8CD748;
    ModelIndex& MI_PARACHUTE_BACKPACK = *(ModelIndex*)0x8CD74C;
    ModelIndex& MI_OYSTER = *(ModelIndex*)0x8CD750;
    ModelIndex& MI_HORSESHOE = *(ModelIndex*)0x8CD754;
    ModelIndex& MI_OFFROAD_WHEEL = *(ModelIndex*)0x8CD758;
    ModelIndex& MI_FLARE = *(ModelIndex*)0x8CD75C;
    ModelIndex& MI_NITRO_BOTTLE_SMALL = *(ModelIndex*)0x8CD760;
    ModelIndex& MI_NITRO_BOTTLE_LARGE = *(ModelIndex*)0x8CD764;
    ModelIndex& MI_NITRO_BOTTLE_DOUBLE = *(ModelIndex*)0x8CD768;
    ModelIndex& MI_HYDRAULICS = *(ModelIndex*)0x8CD76C;
    ModelIndex& MI_STEREO_UPGRADE = *(ModelIndex*)0x8CD770;
    ModelIndex& MI_BASKETBALL = *(ModelIndex*)0x8CD774;
    ModelIndex& MI_POOL_CUE_BALL = *(ModelIndex*)0x8CD778;
    ModelIndex& MI_PUNCHBAG = *(ModelIndex*)0x8CD77C;
    ModelIndex& MI_IMY_GRAY_CRATE = *(ModelIndex*)0x8CD780;
}
*/
