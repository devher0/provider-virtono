/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VirtualMachineParameters are the configurable fields of a VirtualMachine.
type VirtualMachineParameters struct {
	// The virtualization to be specified relevant to which the VPS will be created.
	Virtualization string `json:"virtualization"`
	// The server group where the VPS will be created and assigned.
	ServerGroupId *int64 `json:"serverGroupId,omitempty"`
	// The password for the root user / administrator of the VPS.
	RootPassword string `json:"rootPassword"`
	// Hostname.
	Hostname string `json:"hostname"`
	// The allowed disk space for the VPS.
	DiskSpace int64 `json:"diskSpace"`
	// The amount of RAM which the VPS will always have.
	RAM int64 `json:"ram"`
	// The maximum amount of RAM that the VPS can use.
	BurstableRAM *int64 `json:"burstableRAM,omitempty"`
	// Swap RAM.
	SwapRAM *int64 `json:"swapRAM,omitempty"`
	// Monthly bandwidth limit of the VPS.
	Bandwidth int64 `json:"bandwidth"`
	// The user ID under whom the vps will be created.
	UserId int64 `json:"userId"`
	// The CPU weight that has been assigned to the user.
	CPU *int64 `json:"cpu,omitempty"`
	// The CPU share in % that will be assigned to the VPS.
	CPUPercent *int64 `json:"cpuPercent,omitempty"`
	// Number of CPU cores that will be used by the VPS.
	CPUCores int64 `json:"cpuCores"`
	// Number of IPv4 addresses that will be assigned to the VPS.
	IPv4 *int64 `json:"numIpv4,omitempty"`
	// 	Number of internal IP addresses that will be assigned to the VPS.
	IPInternal *int64 `json:"numInternalIps,omitempty"`
	// Number of IPv6 addresses that will be assigned to the VPS.
	IPv6 *int64 `json:"numIpv6,omitempty"`
	// Number of IPv6 subnets that will be assigned to the VPS.
	Ipv6Subnet *int64 `json:"numIpv6Subnets,omitempty"`
	// Indicates if the VNC should be set of this VPS.
	EnableVNC *bool `json:"enableVNC,omitempty"`
	// Optional VNC password or automatically generated if VNC is enabled.
	VNCPassword *string `json:"vncPassword,omitempty"`
	// Used for Xen HVM. This sets the shadow memory for the VPS.
	ShadowMemory *int64 `json:"shadowMemory,omitempty"`
	// Create VPS using ISO present on the server.
	ISO *string `json:"iso,omitempty"`
	// Sets the boot order in the VPS.
	BootOrder *int64 `json:"bootOrder,omitempty"`
	// Indicates if `tuntap` should be enabled. Note: For OpenVZ only.
	EnableTuntap *int64 `json:"enableTuntap,omitempty"`
	// If set then IO priority for the vps will be enabled. Note: For OpenVZ only.
	EnableIOPriority *int64 `json:"enableIOPriority,omitempty"`
	// Indicates if the VPS should be suspended if the bandwidth limit is exceeded.
	SuspendIfBandwidthExceeded *bool `json:"suspendIfBandwidthExceeded,omitempty"`
	// The number of maximum OS re-installations allowed.
	OsReinstallLimit *int64 `json:"osReinstallLimit,omitempty"`
	// The OS (operating-system) ID.
	OSId int64 `json:"osId"`
	// Media group ID for OS (operating-system) templates.
	MediaGroupId *string `json:"mediaGroupId,omitempty"`
	// Virtual network interface type.
	NIC *string `json:"nic,omitempty"`
	// Provide one of these ssh_options: add_ssh_keys,
	// generate_keys, use_ssh_keys, we have provided detailed information about these options below
	SSHOptions *string `json:"sshOptions,omitempty"`
	// Public SSH key.
	SSHPublicKey *string `json:"sshPublicKey,omitempty"`
	// Private SSH Key.
	SSHPrivateKey *int64 `json:"sshPrivateKey,omitempty"`
	// Provide an array of public keys (compulsory if `use_ssh_keys` is passed in `ssh_options`)
	ExistingKeys *[]string `json:"existingSSHKeys,omitempty"`
	// BIOS type `seabios` or `uefi`. Default is `seabios`. Note: will be applied only while booting with iso. (KVM ONLY)
	BIOS *string `json:"bios,omitempty"`
}

// VirtualMachineObservation are the observable fields of a VirtualMachine.
type VirtualMachineObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A VirtualMachineSpec defines the desired state of a VirtualMachine.
type VirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualMachineParameters `json:"forProvider"`
}

// A VirtualMachineStatus represents the observed state of a VirtualMachine.
type VirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A VirtualMachine is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,virtono}
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

// VirtualMachine type metadata.
var (
	VirtualMachineKind             = reflect.TypeOf(VirtualMachine{}).Name()
	VirtualMachineGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualMachineKind}.String()
	VirtualMachineKindAPIVersion   = VirtualMachineKind + "." + SchemeGroupVersion.String()
	VirtualMachineGroupVersionKind = SchemeGroupVersion.WithKind(VirtualMachineKind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
