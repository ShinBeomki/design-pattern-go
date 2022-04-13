package ProxyPattern

import "fmt"

type Car struct{}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

type Driven interface {
	Drive()
}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}
