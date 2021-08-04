package antipattern

import (
	"github.com/tgmendes/go-interfaces/antipattern/car"
	"github.com/tgmendes/go-interfaces/antipattern/changerequest"
	"github.com/tgmendes/go-interfaces/antipattern/motorcycle"
	"github.com/tgmendes/go-interfaces/antipattern/profile"
)

func main() {
	cUpdater := car.Car{}
	mUpdater := motorcycle.Motorcycle{}
	profUpdater := profile.Profile{}

	changerequest.UpdateCarVRM(cUpdater)
	changerequest.UpdateMotorcycleVRM(mUpdater)
	changerequest.UpdateProfileName(profUpdater)

	// newcUpdater := newcar.Car{}
	// changerequest. ??
}
