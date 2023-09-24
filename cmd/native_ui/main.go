package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// fixme:
// do some nice release style packaging, with icon and stuff
// add images to weapons and cars?
func main() {
	//go makeTempServer() //debugging

	myApp := app.NewWithID("biehdc.trainer.gtasahax")
	myWindow := myApp.NewWindow("GTA SA Trainer")

	if !fyne.CurrentDevice().IsMobile() {
		//myWindow.Resize(fyne.NewSize(360, 590))
		myWindow.Resize(fyne.NewSize(294, 190)) //bare minimum for desktop
	}

	scale := float32(myApp.Preferences().FloatWithFallback("appscale", platformDefaultScale[float64]()))
	applyTheme(scale)

	myWindow.CenterOnScreen()

	tabs := container.NewAppTabs(
		container.NewTabItem("Weapons", makeWeapons()),
		container.NewTabItem("Vehicles", makeVehicles()),
		container.NewTabItem("Settings", makeSettings(myWindow)),
	)

	tabs.SetTabLocation(container.TabLocationBottom)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func makeSettings(w fyne.Window) fyne.CanvasObject {
	// load the last host from the preference store
	setHost(fyne.CurrentApp().Preferences().StringWithFallback("lasthost", "127.0.0.1"))

	settings := widget.NewButton("Appearance", func() {
		w := fyne.CurrentApp().NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(440, 520))
		w.Show()
	})

	scaleslider := widget.NewSlider(0.5, 3.0)
	scaleslider.Step = 0.1
	scaleslider.Value = fyne.CurrentApp().Preferences().FloatWithFallback("appscale", platformDefaultScale[float64]())
	scaleslider.OnChangeEnded = func(newvalue float64) {
		fyne.CurrentApp().Preferences().SetFloat("appscale", newvalue)
		applyTheme(float32(newvalue))
		scaleslider.Value = newvalue
		//dialog.ShowInformation("scale", fmt.Sprintf("%f", newvalue), w)
	}

	address := widget.NewEntry()
	address.SetText(getHost())

	sethostclo := func(h string) {
		address.SetText(h)
		setHost(h)
		fyne.CurrentApp().Preferences().SetString("lasthost", h)
		dialog.ShowInformation("New Host Set", "You set "+h+" as new host", w)
	}

	exclude := "Click Scan to find hosts"
	nonefound := "No Hosts found!"
	hostlistdata := binding.BindStringList(&[]string{exclude})
	hostlist := widget.NewListWithData(
		hostlistdata,
		func() fyne.CanvasObject {
			return widget.NewLabel("255.255.255.255")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)
	hostlist.OnSelected = func(id widget.ListItemID) {
		str, err := hostlistdata.GetValue(id)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		//ugly ignore no hosts
		if str == exclude || str == nonefound {
			hostlist.Unselect(id)
			return
		}
		sethostclo(str)
	}

	scan := widget.NewButton("Scan", func() {
		newlist := FindInstances()
		if len(newlist) <= 0 {
			hostlistdata.Set([]string{nonefound})
		} else {
			hostlistdata.Set(newlist)
		}
	})

	return container.NewBorder(
		container.NewVBox(
			settings,
			container.New(layout.NewFormLayout(),
				widget.NewButton("Scale:", func() {
					//we do a little trolling
					scaleslider.OnChangeEnded(platformDefaultScale[float64]())
				}),
				scaleslider,
			),
			widget.NewForm(
				widget.NewFormItem("Client Address", address),
			),
			widget.NewButton("Set", func() {
				sethostclo(address.Text)
			}),
		),
		scan,
		nil,
		nil,
		hostlist,
	)
}

func makeWeapons() fyne.CanvasObject {
	var weaponnames []string
	var weaponids []WeaponType

	for _, weapon := range ValidWeaponIDs {
		weaponnames = append(weaponnames, weapon.String())
		weaponids = append(weaponids, weapon)
	}

	weapon := widget.NewSelect(weaponnames, nil)
	weapon.SetSelected(weaponnames[0])

	ammo := widget.NewSlider(1, 9999)
	ammo.Step = 10
	ammo.Value = 1000

	return NewBottomUpWidthStretched(
		widget.NewLabel("Select a Weapon:"),
		weapon,
		ammo,
		widget.NewButton("Get", func() {
			index := weapon.SelectedIndex()
			if index >= 0 && index < len(weaponids) {
				makeRequest(fmt.Sprintf("/weapon?weaponid=%d&ammo=%d", weaponids[index], int(ammo.Value)))
			}

		}),
	)
}

func makeVehicles() fyne.CanvasObject {
	var cars []string
	var caridsstr []string

	for _, car := range ValidVehicleIDs {
		cars = append(cars, car.String())
		caridsstr = append(caridsstr, fmt.Sprintf("%d", car))
	}

	vehicleasname := widget.NewSelect(cars, nil)
	vehicleasname.SetSelected(cars[0])

	vehicleasid := widget.NewSelect(caridsstr, nil)
	vehicleasid.SetSelected(caridsstr[0])

	// Syncronise the 2 dropdown boxes
	// maybe there is a cleaner way
	var selectedindex int
	var inprogress bool //calling SetSelectedIndex triggers OnChanged
	syncdropdown1 := func(_ string) {
		if inprogress {
			inprogress = false
			return
		}
		inprogress = true
		selectedindex = vehicleasname.SelectedIndex()
		vehicleasid.SetSelectedIndex(selectedindex)
	}
	vehicleasname.OnChanged = syncdropdown1
	syncdropdown2 := func(_ string) {
		if inprogress {
			inprogress = false
			return
		}
		inprogress = true
		selectedindex = vehicleasid.SelectedIndex()
		vehicleasname.SetSelectedIndex(selectedindex)
	}
	vehicleasid.OnChanged = syncdropdown2

	return NewBottomUpWidthStretched(
		widget.NewLabel("Select a Vehicle:"),
		vehicleasname,
		vehicleasid,
		widget.NewButton("Spawn", func() {
			if selectedindex >= 0 && selectedindex < len(caridsstr) {
				//fmt.Println(caridsstr[index])
				makeRequest(fmt.Sprintf("/spawn?vehicleid=%s", caridsstr[selectedindex]))
			}
		}),
	)
}
