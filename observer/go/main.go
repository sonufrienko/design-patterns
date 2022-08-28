package main

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

type Observer interface {
	Update(string)
	GetId() string
}

func main() {
	garage := NewMotionSensor("Garage")

	homeAlert := NewClient("Home")
	myPhone := NewClient("My Phone")

	garage.Register(homeAlert)
	garage.Register(myPhone)

	garage.NotifyAll()
}
