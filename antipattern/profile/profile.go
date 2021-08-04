package profile

type Updater interface {
	UpdateName()
	UpdateMaritalStatus()
	UpdateAddress()
	UpdatePreferredName()
	UpdateOccupation()
	UpdatePhoneNumber()
	UpdateDLN()
	UpdateEmail()
}

type Profile struct{}

func (p Profile) UpdateName()          {}
func (p Profile) UpdateMaritalStatus() {}
func (p Profile) UpdateAddress()       {}
func (p Profile) UpdatePreferredName() {}
func (p Profile) UpdateOccupation()    {}
func (p Profile) UpdatePhoneNumber()   {}
func (p Profile) UpdateDLN()           {}
func (p Profile) UpdateEmail()         {}
