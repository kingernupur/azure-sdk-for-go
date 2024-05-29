//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arminformaticadatamgmt

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ApplicationType - Various application types of the Serverless Runtime environments
type ApplicationType string

const (
	// ApplicationTypeCDI - Data Integration
	ApplicationTypeCDI ApplicationType = "CDI"
	// ApplicationTypeCDIE - Advanced Data Integration
	ApplicationTypeCDIE ApplicationType = "CDIE"
)

// PossibleApplicationTypeValues returns the possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{
		ApplicationTypeCDI,
		ApplicationTypeCDIE,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PlatformType - Various types of the Platform types.
type PlatformType string

const (
	// PlatformTypeAZURE - Azure platform type
	PlatformTypeAZURE PlatformType = "AZURE"
)

// PossiblePlatformTypeValues returns the possible values for the PlatformType const type.
func PossiblePlatformTypeValues() []PlatformType {
	return []PlatformType{
		PlatformTypeAZURE,
	}
}

// ProvisioningState - Provisioning State of the Organization resource.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Organization resource creation request accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Organization resource creation canceled
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - Organization resource creation started
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted - Organization resource is deleted
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting - Organization resource deletion started
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Organization resource creation failed
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateNotSpecified - Organization resource state is unknown
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	// ProvisioningStateSucceeded - Organization resource creation successful
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Organization resource is being updated
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RuntimeType - Various types of the runtime types.
type RuntimeType string

const (
	// RuntimeTypeSERVERLESS - Serverless Runtime type
	RuntimeTypeSERVERLESS RuntimeType = "SERVERLESS"
)

// PossibleRuntimeTypeValues returns the possible values for the RuntimeType const type.
func PossibleRuntimeTypeValues() []RuntimeType {
	return []RuntimeType{
		RuntimeTypeSERVERLESS,
	}
}
