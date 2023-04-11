package analyzers

import (
	"slyfox-payload-controller/devices"
	"time"

	"go.einride.tech/pid"
)

const (
	DefaultProportionalGain float64 = 2.0
	DefaultIntegralGain     float64 = 1.0
	DefaultDerivativeGain   float64 = 1.0
)

type IPIDMotors interface {
	Update(actualValues []float64, samplingIntervals []time.Duration) []pid.ControllerState
}

type PIDMotors struct {
	motors      []devices.IServoMotor
	controllers []pid.Controller

	IPIDMotors
}

func NewPIDMotors(motors []devices.IServoMotor, ProportionalGain, IntegralGain, DerivativeGain float64) *PIDMotors {
	controllers := make([]pid.Controller, len(motors))

	for i, _ := range motors {
		c := pid.Controller{
			Config: pid.ControllerConfig{
				ProportionalGain: ProportionalGain,
				IntegralGain:     IntegralGain,
				DerivativeGain:   DerivativeGain,
			},
		}

		controllers[i] = c
	}

	return &PIDMotors{
		motors:      motors,
		controllers: controllers,
	}
}

func (pidMotors *PIDMotors) Update(actualValues []float64, samplingIntervals time.Duration) []pid.ControllerState {
	states := make([]pid.ControllerState, len(actualValues))
	for i, value := range actualValues {
		pidMotors.controllers[i].Update(pid.ControllerInput{
			ReferenceSignal:  10,
			ActualSignal:     value,
			SamplingInterval: samplingIntervals,
		})

		states[i] = pidMotors.controllers[i].State
	}

	return states
}
