package ObserverPattern

import (
	"fmt"
	"math/rand"
)

type Observable interface {
	subscribe(observer Observer)
	notifyAll(data interface{})
}

type Patient struct {
	name         string
	temperature  float64
	heartRate    int
	observerList []Observer
}

func NewPatient(name string, temperature float64, heartRate int) *Patient {
	return &Patient{
		name,
		temperature,
		heartRate,
		make([]Observer, 0),
	}
}

func (p *Patient) subscribe(observer Observer) {
	p.observerList = append(p.observerList, observer)
}

func (p *Patient) notifyAll(data interface{}) {
	for _, observer := range p.observerList {
		observer.update(data)
	}
}

func (p *Patient) setTemperature(temperature float64) {
	p.temperature = temperature
	medicalInfo := NewMedicalInfo(p.name, TEMPERATURE, p.temperature)
	p.notifyAll(medicalInfo)
}

func (p *Patient) setHeartRate(heartRate int) {
	p.heartRate = heartRate
	medicalInfo := NewMedicalInfo(p.name, HEARTRATE, p.heartRate)
	p.notifyAll(medicalInfo)
}

const (
	TEMPERATURE = iota
	HEARTRATE
)

type MedicalInfo struct {
	PatientName string
	Kind        int
	Value       interface{}
}

func NewMedicalInfo(name string, kind int, value interface{}) MedicalInfo {
	return MedicalInfo{
		name,
		kind,
		value,
	}
}

type Observer interface {
	update(data interface{})
	getId() int
}

type Nurse struct {
	name string
	id   int
}

func NewNurse(name string) *Nurse {
	id := rand.Int()
	return &Nurse{
		name: name,
		id:   id,
	}
}

func (n *Nurse) update(data interface{}) {
	if medicalInfo, ok := data.(MedicalInfo); ok {
		switch medicalInfo.Kind {
		case TEMPERATURE:
			if medicalInfo.Value.(float64) > 38 {
				fmt.Printf("%s goes to %s\n", medicalInfo.PatientName, n.name)
			}
		case HEARTRATE:
			if medicalInfo.Value.(int) < 30 {
				fmt.Printf("%s calls a doctor\n", n.name)
			}
		}
	}
}

func (n *Nurse) getId() int {
	return n.id
}
