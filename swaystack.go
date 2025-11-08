/************************************************************************************************100
A sway client example that connects to sway, retrieves the layout tree, and prints information about the focused node.

If calling from an ssh connection then you may need to call:

 export SWAYSOCK=/run/user/1000/sway-ipc.1000.3567.sock

 You can also call:
 swaymsg -t get_tree

swaymsg '[con_id=38]' mark managed
swaymsg '[con_mark=managed]' focus
swaymsg rename workspace "5" to "5:managed"
swaymsg rename workspace "5:managed" to "5"

swaymsg '[app_id="org.gnome.clocks"] floating enable'
swaymsg '[app_id="org.gnome.clocks"] move absolute position 100 200'
swaymsg '[app_id="org.gnome.clocks"] resize set 600 400'
swaymsg '[con_id=38] move absolute position 10 20'

listen to window events, listen to key events

***********************************************************************************************100*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/jceaser/sway-stack/lib"
	"github.com/tlinden/swayipc"
)

/***************************************************************************80*/

func ProcessTick(event *swayipc.RawResponse) error {
	var err error
	switch event.PayloadType {
	case swayipc.EV_Tick:
		ev := &swayipc.EventTick{}
		err = json.Unmarshal(event.Payload, &ev)
		//repr.Println("tick:", string(ev.Payload))
		//fmt.Printf("tick %v\n", ev)
		if strings.HasPrefix(string(ev.Payload), "layout.column-stack") {
			parts := strings.Split(string(ev.Payload), ":")
			if len(parts) == 2 {
				lib.Log.Debug.Printf("command: %s\n", parts[1])
				switch parts[1] {
				case "on":
					lib.FindWorkSpace()
				case "off":
					fmt.Println("off")
				case "up":
					lib.RotateUp()
				case "down":
					lib.RotateDown()
				case "left":
					lib.SwapLeft()
				case "right":
					lib.SwapRight()
				}
			}
		}

	case swayipc.EV_Window:
		ev := &swayipc.EventWindow{}
		err = json.Unmarshal(event.Payload, &ev)
		//repr.Println(ev.Container.Name)
		fmt.Printf("window event: %v\n", ev.Container.Name)
	case swayipc.EV_Binding:
		//ev := &swayipc.Event{}
		fmt.Println("binding")
		//repr.Println(ev)
		fmt.Println(string(event.Payload))
		//default:
		//fmt.Printf("received unsubscribed event %d\n", event.PayloadType)
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ipc := swayipc.NewSwayIPC()
	lib.ActiveConnection = ipc

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ipc.Close()

	//config, _ := ipc.GetConfig()
	//fmt.Printf("c=%v\n", config)

	tree, _ := ipc.GetTree()
	//repr.Println(tree)
	fmt.Printf("rect: %v\n", tree.Rect)

	/*
		bars, _ := ipc.GetBars()
		bar, _ := ipc.GetBar(bars[0])
		repr.Println(bar)

		bind, _ := ipc.GetBindingState()
		fmt.Printf("b=%v\n", bind)
	*/

	_, err = ipc.Subscribe(&swayipc.Event{
		Tick:    true,
		Window:  true,
		Binding: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	ipc.EventLoop(ProcessTick)
}
