package nkutil

import (
	"github.com/telyn/midi"
	"github.com/telyn/midi/korg/korgdevices"
	"github.com/telyn/midi/korg/korgsysex"
	"github.com/telyn/midi/korg/korgsysex/format4"
	"github.com/telyn/midi/korg/korgsysex/search"
	"github.com/telyn/midi/msgs"
	"github.com/telyn/midi/sysex"
)

// DispatchConfig is a simplified struct to make writing programs to work with NanoKONTROLs in Native Mode easier
type DispatchConfig struct {
	NKSysExHandler        format4.SingleDeviceHandler
	SearchResponseHandler func(search.Response) error
	ControlChangeHandler  msgs.Handler
}

func (dc DispatchConfig) Dispatcher() midi.Dispatcher {
	return midi.Dispatcher{
		Handlers: map[msgs.Kind]msgs.Handler{
			msgs.SystemExclusive: sysex.Handler{
				sysex.Korg: korgsysex.MultiFormatHandler{
					Format4: format4.MultiDeviceHandler{
						Handlers: map[uint8]map[korgdevices.Device]format4.SingleDeviceHandler{
							format4.AllChannels: map[korgdevices.Device]format4.SingleDeviceHandler{
								korgdevices.NanoKONTROL2: dc.NKSysExHandler,
							},
						},
					},
					Search: search.Handler{
						Response: dc.SearchResponseHandler,
					},
				},
			},
			msgs.ControlChange: dc.ControlChangeHandler,
		},
	}
}
