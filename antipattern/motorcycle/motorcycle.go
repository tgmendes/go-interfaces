package motorcycle

type Updater interface {
	UpdateVRM()
	UpdateColor()
}

type Motorcycle struct{}

func (m Motorcycle) UpdateVRM() {}

func (m Motorcycle) UpdateColor() {}
