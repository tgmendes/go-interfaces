package changerequest

type VehicleUpdater interface {
	UpdateVRM()
}

type ProfileUpdater interface {
	UpdateName()
}

func UpdateVRM(c VehicleUpdater) {
	c.UpdateVRM()
}

func UpdateProfileName(p ProfileUpdater) {
	p.UpdateName()
}
