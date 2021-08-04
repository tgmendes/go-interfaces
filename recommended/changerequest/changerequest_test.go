package changerequest_test

import (
	"github.com/tgmendes/go-interfaces/recommended/changerequest"
	"testing"
)

type VehicleUpdaterMock struct {
	VRMCalled bool
}

func (v *VehicleUpdaterMock) UpdateVRM() {
	v.VRMCalled = true
}

type ProfileUpdaterMock struct {
	NameCalled bool
}

func (p *ProfileUpdaterMock) UpdateName() {
	p.NameCalled = true
}

func TestUpdateVRM(t *testing.T) {
	m := VehicleUpdaterMock{}

	changerequest.UpdateVRM(&m)

	if !m.VRMCalled {
		t.Fatal("Didn't call vrm update")
	}
}

func TestUpdateProfile(t *testing.T) {
	m := ProfileUpdaterMock{}

	changerequest.UpdateProfileName(&m)

	if !m.NameCalled {
		t.Fatal("Didn't call name")
	}
}
