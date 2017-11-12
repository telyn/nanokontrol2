package search

import (
	"fmt"
	"time"

	"github.com/rakyll/portmidi"
	"github.com/telyn/midi"
	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/korgsysex/search"
	"github.com/telyn/midi/portbidi"
	"github.com/telyn/midi/stream"
	"github.com/telyn/midi/sysex"
	"github.com/telyn/nanokontrol/nkutil"
)

func Search(device korgdevices.Device) portbidi.Stream {
	inStreams, outStreams := initialize()
	defer portmidi.Terminate()
	defer inStreams.Close()
	defer outStreams.Close()

	streamStreams := make([]stream.Stream, len(inStreams))

	dispatchers := make([]midi.Dispatcher, len(inStreams))

	for i, input := range inStreams {
		dispatch := nkutil.DispatchConfig{
			SearchHandler: search.Handler{
				ResponseHandler: func(r search.Response) error {
					outInfo := outStreams[r.EchoBackID].info
					fmt.Printf("input %d (%v %v) is connected to output %d (%v %v) and has a Korg %v\n", i, input.info.Interface, input.info.Name, r.EchoBackID, outInfo.Interface, outInfo.Name, r.Device())
					return nil
				},
			},
		}.Dispatcher()
		dispatchers[i] = dispatch

	}

	for i, output := range outStreams {
		if i > 127 { // out of data bytes
			break
		}
		sysex.Write(output.stream, portmidi.Time(), search.Request{byte(i)})
	}
	devices
	for {
		for i, inStream := range inStreams {
			bytes, err := inStream.stream.ReadSysExBytes(32)
			if err != nil {
				fmt.Printf("Couldn't read from stream %d\n")
				continue
			}

			messages := streamStreams[i].ConsumeBytes(bytes)
			for _, msg := range messages {
				dispatchers[i].HandleMessage(msg)
			}
		}
		time.Sleep(10 * time.Microsecond)
	}

}
