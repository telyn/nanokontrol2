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

// DispatchConfig is a simplified struct to
type DispatchConfig struct {
	NKSysExHandler       format4.SingleDeviceHandler
	SearchHandler        search.Handler
	ControlChangeHandler msgs.Handler
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
					Search: dc.SearchHandler,
				},
			},
			msgs.ControlChange: dc.ControlChangeHandler,
		},
	}
}
