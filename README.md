# Gobot Drivers for GPIO

Gobot (http://gobot.io/) is a framework and set of libraries for robotics, physical computing, and the Internet of Thing written in the Go programming language (http://golang.org/).

This library provides drivers for General Purpose Input/Output (GPIO) devices (https://en.wikipedia.org/wiki/General_Purpose_Input/Output). It is normally not used directly, but instead is registered by adaptor library such as gobot-firmata (https://github.com/hybridgroup/gobot-firmata) that supports the needed interfaces for GPIO devices.

[![Build Status](https://travis-ci.org/hybridgroup/gobot-gpio.svg?branch=master)](https://travis-ci.org/hybridgroup/gobot-gpio) [![Coverage Status](https://coveralls.io/repos/hybridgroup/gobot-gpio/badge.png?branch=master)](https://coveralls.io/r/hybridgroup/gobot-gpio?branch=master)

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

  - Analog Sensor
  - Button
  - LED
  - Motor
  - Servo

More drivers are coming soon...

## Documentation
We're busy adding documentation to our web site at http://gobot.io/  please check there as we continue to work on Gobot

Thank you!

## Contributing
* All patches must be provided under the Apache 2.0 License
* Please use the -s option in git to "sign off" that the commit is your work and you are providing it under the Apache 2.0 License
* Submit a Github Pull Request to the appropriate branch and ideally discuss the changes with us in IRC #gobotio on Freenode.
* We will look at the patch, test it out, and give you feedback.
* Avoid doing minor whitespace changes, renamings, etc. along with merged content. These will be done by the maintainers from time to time but they can complicate merges and should be done seperately.
* Take care to maintain the existing coding style.
* Add unit tests for any new or changed functionality.
* All pull requests should be "fast forward"
  * If there are commits after yours use “git rebase -i <new_head_branch>”
  * If you have local changes you may need to use “git stash”
  * For git help see [progit](http://git-scm.com/book) which is an awesome (and free) book on git


## License
Copyright (c) 2013-2014 The Hybrid Group. Licensed under the Apache 2.0 license.
