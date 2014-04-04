# Functions

## OffC

Stops the motor.

## OnC

Starts the motor.

## IsOnC

Starts the motor if is not started.

## IsOffC

Stops the motor if is not stoped.

## ToggleC

Sets the state of the motor to the oposite of the current state, if motor is on then sets it to off.

## SpeedC(value, byte)

Sets the speed of the motor to the value provided in the speed param, speed value must be an integer between 0 and 255.

#### Params

- **value** - params
- **byte** - params

#### Returns

- **integer**

## MinC

Sets the min speed value.

## MaxC

Sets the max speed value.

## ForwardC(speed, byte) 

Sets the motor to go forward.

#### Params

- **speed** - params
- **byte** - params

## BackwardC(speed byte)

Sets the motor to go backward.

#### Params

- **speed** - params
- **byte** - params

## CurrentSpeedC

Returns the current speed of the motor as an integer between 0 and 255.

#### Returns

- **integer**