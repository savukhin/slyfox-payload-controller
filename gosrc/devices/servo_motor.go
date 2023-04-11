package devices

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type IServoMotor interface {
	Move(value int)
}

type ServoMotor struct {
	driver gpio.ServoDriver

	IServoMotor
}

type ServoMotorPair struct {
	motor1 ServoMotor
	motor2 ServoMotor

	IServoMotor
}

func NewServoMotor(adaptor *raspi.Adaptor, pin string) *ServoMotor {
	return &ServoMotor{
		driver: *gpio.NewServoDriver(adaptor, pin),
	}
}

func (servo *ServoMotor) Move(value uint8) {
	servo.driver.Move(value)
}

func NewServoMotorPair(adaptor *raspi.Adaptor, pin1 string, pin2 string) *ServoMotorPair {
	return &ServoMotorPair{
		motor1: *NewServoMotor(adaptor, pin1),
		motor2: *NewServoMotor(adaptor, pin2),
	}
}

func (servo *ServoMotorPair) Move(value uint8) {
	servo.motor1.Move(value)
	servo.motor2.Move(value)
}
