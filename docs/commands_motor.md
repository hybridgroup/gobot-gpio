# Commands

## OffC

Stops the motor.

##### Returns

`true or nil`

## OnC

Starts the motor.

##### Returns

`true or nil`

## IsOnC

Starts the motor if is not started.

##### Returns

`true or nil`

## IsOffC

Stops the motor if is not stoped.

##### Returns

`true or nil`

## ToggleC

Sets the state of the motor to the oposite of the current state, if motor is on then sets it to off.

##### Returns

`true or nil`

## SpeedC(value byte)

Sets the speed of the motor to the value provided in the speed param, speed value must be an integer between 0 and 255.

##### Params

- **value** - params
- **byte** - params

##### Returns

`integer`

## MinC

Sets the min speed value.

##### Returns

`nil`

## MaxC

Sets the max speed value.

##### Returns

`nil`

## ForwardC(speed byte) 

Sets the motor to go forward.

##### Params

- **speed** - params
- **byte** - params

##### Returns

`nil`

## BackwardC(speed byte)

Sets the motor to go backward.

##### Params

- **speed** - params
- **byte** - params

##### Returns

`nil`

## CurrentSpeedC

Returns the current speed of the motor as an integer between 0 and 255.

##### Returns

`integer`