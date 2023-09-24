package main

import "fmt"

// Copy of internal data to allow for easy cross compiling
type WeaponType uint32

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
