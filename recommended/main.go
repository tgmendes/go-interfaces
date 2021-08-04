package recommended

import (
	"github.com/tgmendes/go-interfaces/recommended/car"
	"github.com/tgmendes/go-interfaces/recommended/changerequest"
	"github.com/tgmendes/go-interfaces/recommended/motorcycle"
	"github.com/tgmendes/go-interfaces/recommended/newcar"
	"github.com/tgmendes/go-interfaces/recommended/profile"
)

func main() {
	cUpdater := car.Car{}
	mUpdater := motorcycle.Motorcycle{}
	pUpdater := profile.Profile{}

	changerequest.UpdateVRM(cUpdater)
	changerequest.UpdateVRM(mUpdater)
	changerequest.UpdateProfileName(pUpdater)

	newcUpdater := newcar.Car{}
	changerequest.UpdateVRM(newcUpdater)
}
