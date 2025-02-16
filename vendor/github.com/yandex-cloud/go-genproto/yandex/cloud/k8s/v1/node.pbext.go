// Code generated by protoc-gen-goext. DO NOT EDIT.

package k8s

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Node) SetStatus(v Node_Status) {
	m.Status = v
}

func (m *Node) SetSpec(v *Node_Spec) {
	m.Spec = v
}

func (m *Node) SetCloudStatus(v *Node_CloudStatus) {
	m.CloudStatus = v
}

func (m *Node) SetKubernetesStatus(v *Node_KubernetesStatus) {
	m.KubernetesStatus = v
}

func (m *Node_KubernetesStatus) SetId(v string) {
	m.Id = v
}

func (m *Node_KubernetesStatus) SetConditions(v []*Condition) {
	m.Conditions = v
}

func (m *Node_KubernetesStatus) SetTaints(v []*Taint) {
	m.Taints = v
}

func (m *Node_KubernetesStatus) SetAttachedVolumes(v []*AttachedVolume) {
	m.AttachedVolumes = v
}

func (m *Node_CloudStatus) SetId(v string) {
	m.Id = v
}

func (m *Node_CloudStatus) SetStatus(v string) {
	m.Status = v
}

func (m *Node_CloudStatus) SetStatusMessage(v string) {
	m.StatusMessage = v
}

func (m *Node_Spec) SetResources(v *ResourcesSpec) {
	m.Resources = v
}

func (m *Node_Spec) SetDisk(v *DiskSpec) {
	m.Disk = v
}

func (m *Condition) SetType(v string) {
	m.Type = v
}

func (m *Condition) SetStatus(v string) {
	m.Status = v
}

func (m *Condition) SetMessage(v string) {
	m.Message = v
}

func (m *Condition) SetLastHeartbeatTime(v *timestamp.Timestamp) {
	m.LastHeartbeatTime = v
}

func (m *Condition) SetLastTransitionTime(v *timestamp.Timestamp) {
	m.LastTransitionTime = v
}

func (m *Taint) SetKey(v string) {
	m.Key = v
}

func (m *Taint) SetValue(v string) {
	m.Value = v
}

func (m *Taint) SetEffect(v Taint_Effect) {
	m.Effect = v
}

func (m *AttachedVolume) SetDriverName(v string) {
	m.DriverName = v
}

func (m *AttachedVolume) SetVolumeHandle(v string) {
	m.VolumeHandle = v
}

func (m *NodeTemplate) SetPlatformId(v string) {
	m.PlatformId = v
}

func (m *NodeTemplate) SetResourcesSpec(v *ResourcesSpec) {
	m.ResourcesSpec = v
}

func (m *NodeTemplate) SetBootDiskSpec(v *DiskSpec) {
	m.BootDiskSpec = v
}

func (m *NodeTemplate) SetMetadata(v map[string]string) {
	m.Metadata = v
}

func (m *NodeTemplate) SetV4AddressSpec(v *NodeAddressSpec) {
	m.V4AddressSpec = v
}

func (m *NodeTemplate) SetSchedulingPolicy(v *SchedulingPolicy) {
	m.SchedulingPolicy = v
}

func (m *NodeTemplate) SetNetworkInterfaceSpecs(v []*NetworkInterfaceSpec) {
	m.NetworkInterfaceSpecs = v
}

func (m *NodeTemplate) SetPlacementPolicy(v *PlacementPolicy) {
	m.PlacementPolicy = v
}

func (m *NodeTemplate) SetNetworkSettings(v *NodeTemplate_NetworkSettings) {
	m.NetworkSettings = v
}

func (m *NodeTemplate) SetContainerRuntimeSettings(v *NodeTemplate_ContainerRuntimeSettings) {
	m.ContainerRuntimeSettings = v
}

func (m *NodeTemplate_NetworkSettings) SetType(v NodeTemplate_NetworkSettings_Type) {
	m.Type = v
}

func (m *NodeTemplate_ContainerRuntimeSettings) SetType(v NodeTemplate_ContainerRuntimeSettings_Type) {
	m.Type = v
}

func (m *NetworkInterfaceSpec) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *NetworkInterfaceSpec) SetPrimaryV4AddressSpec(v *NodeAddressSpec) {
	m.PrimaryV4AddressSpec = v
}

func (m *NetworkInterfaceSpec) SetPrimaryV6AddressSpec(v *NodeAddressSpec) {
	m.PrimaryV6AddressSpec = v
}

func (m *NetworkInterfaceSpec) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *NodeAddressSpec) SetOneToOneNatSpec(v *OneToOneNatSpec) {
	m.OneToOneNatSpec = v
}

func (m *OneToOneNatSpec) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

func (m *ResourcesSpec) SetMemory(v int64) {
	m.Memory = v
}

func (m *ResourcesSpec) SetCores(v int64) {
	m.Cores = v
}

func (m *ResourcesSpec) SetCoreFraction(v int64) {
	m.CoreFraction = v
}

func (m *ResourcesSpec) SetGpus(v int64) {
	m.Gpus = v
}

func (m *DiskSpec) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *DiskSpec) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *SchedulingPolicy) SetPreemptible(v bool) {
	m.Preemptible = v
}

func (m *PlacementPolicy) SetPlacementGroupId(v string) {
	m.PlacementGroupId = v
}
