package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	. "gtasamod/directcalls"
	. "gtasamod/types"
)

func getIPs() []string {
	var addrlist []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			addrlist = append(addrlist, ipnet.IP.String())
		}
	}
	return addrlist
}

func launchWeb(gi *GameInternals) {
	mux := http.NewServeMux()

	mux.HandleFunc("/spawn", spawn(gi))
	mux.HandleFunc("/weapon", gun(gi))
	mux.HandleFunc("/hello", hello())
	mux.HandleFunc("/", index(gi))

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		AddQuitFn(func() {
			srv.Shutdown(nil)
		})
		msg := "Webserver on :8080"
		addrs := getIPs()
		if len(addrs) > 0 {
			msg = "Webserver on:\n"
			for _, addr := range addrs {
				msg += addr+":8080\n"
			}
		}
		<-time.After(5 * time.Second)
		AddFn(func(_ *GameInternals) bool {
			Hud_SetHelpMessage(msg, false, false, true)
			return false
		})
	}()

	srv.ListenAndServe()
}

func hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "this is cj from grove street")
	}
}

func index(gi *GameInternals) http.HandlerFunc {
	var buildpage = &bytes.Buffer{}
	vehicles := template.Must(template.New("tpl").Parse(vehicleSpawnTemplate))
	vehicles.Execute(buildpage, weblate)

	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewBuffer(buildpage.Bytes()))
	}
}

func spawn(gi *GameInternals) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cliurl, err := url.Parse(req.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		params := cliurl.Query()
		vehicleidstr := params.Get("vehicleid")

		//fmt.Println("requested vehicleid", vehicleidstr)
		vehicleidpp, err := strconv.Atoi(vehicleidstr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		vehicleid := clampcar(400, VehicleID(vehicleidpp), 611, true)
		func(vehicleid VehicleID) { //explicitly capture the vehicleid
			AddFn(func(_ *GameInternals) bool {
				Cheat_NewVehicle((int32)(vehicleid))
				return false
			})
		}(vehicleid)

		io.WriteString(w, "<head><meta http-equiv=\"refresh\" content=\"0; URL=/\"/></head>")
	}
}

func gun(gi *GameInternals) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cliurl, err := url.Parse(req.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		params := cliurl.Query()

		weaponidstr := params.Get("weaponid")
		if weaponidstr == "" {
			http.Error(w, "empty weaponid", http.StatusInternalServerError)
			return
		}
		weaponidpp, err := strconv.Atoi(weaponidstr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		weaponid := WeaponType(weaponidpp)
		if !weaponid.Valid() {
			http.Error(w, "unknown or invalid weaponid", http.StatusInternalServerError)
			return
		}

		ammostr := params.Get("ammo")
		if ammostr == "" {
			http.Error(w, "empty ammostr", http.StatusInternalServerError)
			return
		}
		ammo, err := strconv.Atoi(ammostr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// no reason to fail this
		if ammo < 1 || ammo > 9999 {
			ammo = 1000
		}

		func(weaponid WeaponType, ammo uint32) {
			AddFn(func(gi *GameInternals) bool {
				Ped_GiveWeapon(gi.player.Ped, weaponid, ammo)
				return false
			})
		}(weaponid, uint32(ammo))

		io.WriteString(w, "<head><meta http-equiv=\"refresh\" content=\"0; URL=/\"/></head>")
	}
}

type Weblate struct {
	Vehicles []VehicleID
	Weapons  []WeaponType
}

var weblate = Weblate{
	Vehicles: ValidVehicleIDs,
	Weapons:  ValidWeaponIDs,
}

//go:embed webinterface_template.html
var vehicleSpawnTemplate string
