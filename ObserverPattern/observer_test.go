package ObserverPattern

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("ObserverPattern", ObserverPatternTest)
}

func ObserverPatternTest(t *testing.T) {
	p1 := NewPatient("bob", 36.5, 60)
	p2 := NewPatient("tony", 36, 80)
	n1 := NewNurse("sarah")
	n2 := NewNurse("christina")

	p1.subscribe(n1)
	p1.subscribe(n2)
	fmt.Println(p1.observerList)

	p2.subscribe(n2)

	p1.setTemperature(40)
	p2.setHeartRate(20)
}
