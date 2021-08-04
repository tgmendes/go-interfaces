package changerequest_test

import (
	"github.com/tgmendes/go-interfaces/antipattern/changerequest"
	"testing"
)

type CarMock struct {
	VRMCalled bool
}

func (c *CarMock) UpdateVRM() {
	c.VRMCalled = true
}
func (c *CarMock) UpdateColor()        {}
func (c *CarMock) UpdateVIN()          {}
func (c *CarMock) UpdateVehicleClass() {}

type ProfileMock struct {
	NameCalled bool
}

func (p *ProfileMock) UpdateName() {
	p.NameCalled = true
}
func (p *ProfileMock) UpdateMaritalStatus() {}
func (p *ProfileMock) UpdateAddress()       {}
func (p *ProfileMock) UpdatePreferredName() {}
func (p *ProfileMock) UpdateOccupation()    {}
func (p *ProfileMock) UpdatePhoneNumber()   {}
func (p *ProfileMock) UpdateDLN()           {}
func (p *ProfileMock) UpdateEmail()         {}

type MotorcycleMock struct {
	VRMCalled bool
}

func (m *MotorcycleMock) UpdateVRM() {
	m.VRMCalled = true
}
func (m *MotorcycleMock) UpdateColor() {}

func TestUpdateCar(t *testing.T) {
	m := CarMock{}

	changerequest.UpdateCarVRM(&m)

	if !m.VRMCalled {
		t.Fatal("Didn't call vrm update")
	}
}

func TestUpdateMotorcycle(t *testing.T) {
	m := MotorcycleMock{}

	changerequest.UpdateMotorcycleVRM(&m)

	if !m.VRMCalled {
		t.Fatal("Didn't call vrm update")
	}
}

func TestUpdateProfile(t *testing.T) {
	m := ProfileMock{}

	changerequest.UpdateProfileName(&m)

	if !m.NameCalled {
		t.Fatal("Didn't call name")
	}
}
