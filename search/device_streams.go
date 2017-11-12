package main

import (
	"fmt"
	"os"

	"github.com/rakyll/portmidi"
)

type deviceStream struct {
	stream *portmidi.Stream
	info   *portmidi.DeviceInfo
}

type deviceStreams []deviceStream

func (d deviceStreams) Close() error {
	errs := []error{}
	for _, device := range d {
		err := device.stream.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 1 {
		return errs[0]
	}
	return nil
}

func initialize() (in deviceStreams, out deviceStreams) {

	err := portmidi.Initialize()
	if err != nil {
		fmt.Printf("Couldn't initialize portmidi: %v\n", err)
		os.Exit(1)
		return
	}

	for i := 0; i < portmidi.CountDevices(); i++ {
		device := portmidi.DeviceID(i)
		info := portmidi.Info(device)
		io := ""
		if info.IsInputAvailable {
			io += "i"
		}
		if info.IsOutputAvailable {
			io += "o"
		}
		fmt.Printf("%d: %s %s: (%s)\n", i, info.Interface, info.Name, io)
		if info.IsInputAvailable {
			input, err := portmidi.NewInputStream(device, 32)
			if err != nil {
				fmt.Printf("Couldn't connect input for %v %v\n", info.Interface, info.Name)
			} else {
				in = append(in, deviceStream{input, info})
			}

		}
		if info.IsOutputAvailable {
			output, err := portmidi.NewOutputStream(device, 32, 32)
			if err != nil {
				fmt.Printf("Couldn't connect output for %v %v\n", info.Interface, info.Name)
			} else {
				out = append(out, deviceStream{output, info})
			}
		}

	}
	return
}
