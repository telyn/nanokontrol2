package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rakyll/portmidi"
	"github.com/telyn/midi/korg/devicesearch"
	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/nanokontrol2"
	"github.com/telyn/midi/sysex"
)

var on = flag.Bool("on", false, "turn native mode on")
var off = flag.Bool("off", false, "turn native mode off")

func main() {
	flag.Parse()
	err := portmidi.Initialize()
	if err != nil {
		fmt.Println("Couldn't initialize portmidi")
		os.Exit(1)
		return
	}
	fmt.Println("Looking for Korg NanoKONTROL...")
	res, err := devicesearch.Search(korgdevices.NanoKONTROL2)
	fmt.Println("Done!")
	if err != nil {
		fmt.Printf("Couldn't find device: %s\n", err)
		os.Exit(1)
		return
	}
	fmt.Println("Found it!")
	if *on {
		fmt.Println("Enabling native mode")
		err := sysex.Write(res.Out, portmidi.Time(), nanokontrol2.SetModeRequest{
			Channel:    res.Channel,
			NativeMode: true,
		})
		if err != nil {
			fmt.Printf("Couldn't send native mode request to nanoKONTROL: %v\n", err)
			return
		}
		on := byte(0x7F)
		for {
			for i := byte(0x00); i < 0x08; i++ {
				bs := []byte{0xBF, i | 0x20, on}
				fmt.Printf("res.Out.WriteSysExBytes(%x,%x,%x)\n", bs[0], bs[1], bs[2])
				res.Out.WriteSysExBytes(portmidi.Time(), bs)
			}
			on = (on ^ 0x7F) & 0x7F
			time.Sleep(500 * time.Millisecond)
		}
	} else if *off {
		fmt.Println("Disabling native mode")
		sysex.Write(res.Out, portmidi.Time(), nanokontrol2.SetModeRequest{
			Channel:    res.Channel,
			NativeMode: false,
		})
	}

	err = portmidi.Terminate()
	if err != nil {
		fmt.Println("Couldn't terminate portmidi")
		os.Exit(1)
		return
	}
}
