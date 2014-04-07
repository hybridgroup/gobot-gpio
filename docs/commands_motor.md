# Functions

## On

Starts the motor.

#### Returns

- **CurrentSpeed** - current speed value

## Off

Stops the motor.

#### Returns

- **Speed** - speed value

## IsOn

Returns true if the motor is on

#### Returns

- **bool** 

## IsOff

Returns true if the motor is off

#### Returns

- **bool**

## Toggle

Sets the state of the motor to the oposite of the current state, if motor is on then sets it to off.

## Speed(value byte)

Sets the speed of the motor to the value provided in the speed param, speed value must be an integer between 0 and 255.

#### Params

- **value** - **byte**

#### Returns

- **SpeedPin value**

## Min

Sets the motor to it's minimum speed.

## Max

Sets the motor to it's max speed.

## Forward(speed byte) 

Sets the motor to a forward direction at the specified speed.

#### Params

- **speed** - **byte**

#### Returns

- **speed**

## Backward(speed byte)

Sets the motor to a backward direction at the specified speed.

#### Params

- **speed** - **byte**

#### Returns

- **speed**

## Direction(direction string) 

Sets the direction of the motor

#### Params

- **direction** - **string**