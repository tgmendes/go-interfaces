package newcar

type Updater interface {
	UpdateVRM()
	UpdateEngine()
	UpdateWheels()
	UpdateRadio()
}
type Car struct{}

func (c Car) UpdateVRM() {}

func (c Car) UpdateEngine() {}

func (c Car) UpdateWheels() {}

func (c Car) UpdateRadio() {}
