# Grand Theft Auto San Andreas Trainer
For Version 1.0 HOODLUM

[Demo Video](https://github.com/BieHDC/go-gtasa-trainer/blob/8b010e0d90c9df9ccab517fa1464010db58407cf/demo.webm)

---
## What is going on?
I wanted to test if you can combine Go with game hacking.
And it turns out, yes you can!

---
## How does it work?
It generates a dinput8.dll file that you place next to the game exe to have it load.  
After it got loaded, DLLMain returns and Go's init() funcs will run.  
After doing its confirmation dance, it will run main() in a goroutine.  
This goroutine then waits until the static pointer to IDirect3DDevice9 is not null anymore.  
After that we know the game exe has really loaded and we hook our functions.  

There are 2 main hooked functions, Key-Input and the Render call.  
Key-Input is for menu input.  
The Render call runs all that render frames function that are supposed to run.  
These are the feedback printers, quick menu inf health cheat. The rotation advancing animation in the vehicle spawner... see anything that calls `func AddFn(fn GameFunc)`.  

---
## What does it do?
#### Ingame:
It has some text events, like chaos caused, getting in vehicle etc. which exists for activity check.  
Then there are menu keys:

- `F5` has a quick menu with a few functions, like infinite health, give weapons, make car indestructible...
- `F6` is a command prompt that can set wanted level and game speed.
- `F7` is a fly hack.
- `F8` has a list of active vehicles and you can cycle through them to teleport on top of them.
- `F9` is a vehicle spawner. It has all vehicles except those that crashed when i tried the whole list.

#### Web Interface:
Thats right, once you loaded into gameplay you will see a help text message telling you ip and port of a webserver.  
Browse to the address and then you can spawn yourself Weapons and Vehicles. This is the part that made the project idea interesting.

#### Native UI:
in the cmd/native_ui folder is a fyne based simple user interface that basically does the same as the web interface, and it also just calls web queries.  
One extra feature it has is to scan your lan for existing GTA SA game instances (it really just http-gets every ip on the static port and adds it to the list if it got a 200OK) so you dont have to type the ip in.  
It is also available as Android App.

---
## Go and making dlls
The `make.bat`/`make.sh` files are your greatest friends and you can stop searching around how to do the thing.

---
## Gems
Check out `structoffsetvalidator` and the types inside `types/`.
Here you see stuff like:
```go
type CPlayer struct {
	Ped                        *CPed        `offset:"0x0"`
	PlayerData                 *CPlayerData `offset:"0x4"`
	_                          [0xb0]byte   //padding
	Money                      uint32       `offset:"0xb8"` //100% correct
	MoneyDisplayed             uint32       `offset:"0xbc"`
	//...
}

func init() {
	var player CPlayer
	err := ValidateOffsets(&player)
	if err != nil {
		panic(err)
	}
	sz := unsafe.Sizeof(player)
	if sz != 400 {
		panic(fmt.Sprintf("CPlayer size is wrong, expected %d, got %d\n", 400, sz))
	}
}
```
During startup `ValidateOffsets` checks if the compiled offset matches the expected offset in the struct tag. The Go compiler usually aligns very good with native C.  
Additionally we verify the size of the struct we expect.  
And in case something went wrong, we will get yelled at.  
You can invoke this path only through rundll32.exe and dont need to start the full game.  

---
## Credits
[The Go Project and Team](https://go.dev/) for the language.  
[gtasa-reversed Project](https://github.com/gta-reversed/gta-reversed-modern) for all the GTA SA internals so i could focus on my main task.  
[Fyne](https://fyne.io/) that allowed me to make a painless gui with great support channels.  
