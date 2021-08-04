package changerequest

import (
	"github.com/tgmendes/go-interfaces/antipattern/car"
	"github.com/tgmendes/go-interfaces/antipattern/motorcycle"
	"github.com/tgmendes/go-interfaces/antipattern/profile"
)

func PassMeAnything(i interface{}) {}

func UpdateCarVRM(c car.Updater) {
	c.UpdateVRM()
}

func UpdateMotorcycleVRM(m motorcycle.Updater) {
	m.UpdateVRM()
}

func UpdateProfileName(p profile.Updater) {
	p.UpdateName()
}
