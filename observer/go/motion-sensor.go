package main

import "github.com/thoas/go-funk"

type MotionSensor struct {
	observerList   []Observer
	placeName      string
	motionDetected bool
}

func NewMotionSensor(placeName string) *MotionSensor {
	return &MotionSensor{
		placeName: placeName,
	}
}

func (m *MotionSensor) MotionDetected() {
	m.motionDetected = true
}

func (m *MotionSensor) Register(observer Observer) {
	m.observerList = append(m.observerList, observer)
}

func (m *MotionSensor) Deregister(observer Observer) {
	m.observerList = funk.Filter(m.observerList, func(o Observer) bool {
		return o.GetId() == observer.GetId()
	}).([]Observer)
}

func (m *MotionSensor) NotifyAll() {
	for _, observer := range m.observerList {
		observer.Update(m.placeName)
	}
}
