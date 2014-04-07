# Functions

## DigitalRead

Returns an integer value that represents the digital read from the pin.

#### Returns

- **int** - integer value

#### API Command

**DigitalReadC**

## DigitalWrite(level byte)

Writes a ON/OFF using digital write to a pin

#### Params

- **level** - **[]byte** - on/off value

#### Returns

- **byte(level)** - on/off value

#### API Command

**DigitalWriteC**

## AnalogRead

Returns an integer value that represents the analog read from the pin.

#### Returns

- **int**- the analog read

#### API Command

**AnalogReadC**

## AnalogWrite(level byte)

Writes a value using analog write to a pin

#### Params

- **level** - **byte**

#### Returns

- **byte(level)** - pin value

#### API Command

**AnalogWriteC**

## PwmWrite(level byte)

Writes a value using Pulse Width Modulation (PWM) to a pin

#### Params

- **level** - **byte**

#### Returns

- **byte(level)** - PWM value

#### API Command

**PwmWriteC**

## ServoWrite(level byte)

Writes a value using servo write to a pin

##### Params

- **level** - **byte**

#### Returns

- **byte(level)** - pin value

#### API Command

**ServoWriteC**