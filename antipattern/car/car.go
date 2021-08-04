package car

type Updater interface {
	UpdateVRM()
	UpdateColor()
	UpdateVIN()
	UpdateVehicleClass()
}
type Car struct{}

func (c Car) UpdateVRM() {}

func (c Car) UpdateColor() {}

func (c Car) UpdateVIN() {}

func (c Car) UpdateVehicleClass() {}
