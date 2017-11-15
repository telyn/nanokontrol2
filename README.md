Collection of command-line utilities to communicate with your Korg NanoKONTROL2

At the moment the only program is toggle-nativemode, which enables/disables Korg Native Mode on the first NanoKONTROL it can find (it tries every MIDI interface attached to your computer), then flashes a set of LEDs.

In the future there will be a tabs/pages interface for the NanoKONTROL which will dynamically remap the controls of the NanoKONTROL based on which buttons are pressed.

For example, the transport controls (forward, reverse, start, stop, record) could each be a separate 'page' of controls. When forward is pushed, the knobs could affect MIDI CCs 0-7. Then when reverse is pushed, they could affect CCs 8-15, and so on for the other buttons and sliders. The program will be very configurable using a fairly simple text file.

I've been working on this since the beginning of November 2017. Please file feature requests and issues! 😊
