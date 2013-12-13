# gobot-gpio

Gobot (http://gobot.io/) is a library for robotics and physical computing using Go

This library provides drivers for General Purpose Input/Output (GPIO) devices (https://en.wikipedia.org/wiki/General_Purpose_Input/Output). It is normally not used directly, but instead is registered by adaptor library such as gobot-firmata (https://github.com/hybridgroup/gobot-firmata) that support the needed interfaces for GPIO devices.

## Getting Started
Install the library with: `go get -u github.com/hybridgroup/gobot-gpio`

## Examples
```go
package main

import (
        "fmt"
        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot-firmata"
        "github.com/hybridgroup/gobot-gpio"
)

func main() {

        firmata := new(gobotFirmata.FirmataAdaptor)
        firmata.Name = "firmata"
        firmata.Port = "/dev/ttyACM0"

        led := gobotGPIO.NewLed(firmata)
        led.Name = "led"
        led.Pin = "13"

        work := func() {
                gobot.Every("1s", func() {
                        led.Toggle()
                })
        }

        robot := gobot.Robot{
                Connections: []interface{}{firmata},
                Devices:     []interface{led},
                Work:        work,
        }

        robot.Start()
}
```
## Hardware Support
Gobot has a extensible system for connecting to hardware devices. The following GPIO devices are currently supported:

  - Button
  - LED
  - Servo

More drivers are coming soon...

## Documentation
We're busy adding documentation to our web site at http://gobot.io/  please check there as we continue to work on Gobot

Thank you!

## Contributing
In lieu of a formal styleguide, take care to maintain the existing coding style. Add unit tests for any new or changed functionality.


## License
Copyright (c) 2013 The Hybrid Group. Licensed under the Apache 2.0 license.
