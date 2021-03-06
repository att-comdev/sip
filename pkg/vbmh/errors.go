package vbmh

import (
	"fmt"

	metal3 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"

	airshipv1 "sipcluster/pkg/api/v1"
)

// ErrorConstraintNotFound is returned when wrong AuthType is provided
type ErrorConstraintNotFound struct {
}

func (e ErrorConstraintNotFound) Error() string {
	return "Invalid or Not found Schedulign Constraint"
}

type ErrorUnableToFullySchedule struct {
	TargetNode   airshipv1.VMRoles
	TargetFlavor string
}

func (e ErrorUnableToFullySchedule) Error() string {
	return fmt.Sprintf("Unable to complete a schedule with a target of %v nodes, with a flavor of %v",
		e.TargetNode, e.TargetFlavor)
}

type ErrorHostIPNotFound struct {
	HostName    string
	ServiceName airshipv1.InfraService
	IPInterface string
	Message     string
}

func (e ErrorHostIPNotFound) Error() string {
	return fmt.Sprintf("Unable to identify the vBMH Host %v IP address on interface %v required by "+
		"Infrastructure Service %v %s ", e.HostName, e.IPInterface, e.ServiceName, e.Message)
}

// ErrorUknownSpreadTopology is returned when wrong AuthType is provided
type ErrorUknownSpreadTopology struct {
	Topology airshipv1.SpreadTopology
}

func (e ErrorUknownSpreadTopology) Error() string {
	return fmt.Sprintf("Uknown spread topology '%s'", e.Topology)
}

// ErrorNetworkDataNotFound is returned when NetworkData metadata is missing from BMH
type ErrorNetworkDataNotFound struct {
	BMH metal3.BareMetalHost
}

func (e ErrorNetworkDataNotFound) Error() string {
	return fmt.Sprintf("vBMH Host %v does not define NetworkData, but is required for scheduling.", e.BMH)
}
