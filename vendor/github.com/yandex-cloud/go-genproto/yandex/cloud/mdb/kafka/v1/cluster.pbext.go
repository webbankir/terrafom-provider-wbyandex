// Code generated by protoc-gen-goext. DO NOT EDIT.

package kafka

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func (m *Cluster) SetId(v string) {
	m.Id = v
}

func (m *Cluster) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Cluster) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Cluster) SetName(v string) {
	m.Name = v
}

func (m *Cluster) SetDescription(v string) {
	m.Description = v
}

func (m *Cluster) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Cluster) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *Cluster) SetMonitoring(v []*Monitoring) {
	m.Monitoring = v
}

func (m *Cluster) SetConfig(v *ConfigSpec) {
	m.Config = v
}

func (m *Cluster) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Cluster) SetHealth(v Cluster_Health) {
	m.Health = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *Cluster) SetHostGroupIds(v []string) {
	m.HostGroupIds = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Monitoring) SetName(v string) {
	m.Name = v
}

func (m *Monitoring) SetDescription(v string) {
	m.Description = v
}

func (m *Monitoring) SetLink(v string) {
	m.Link = v
}

func (m *ConfigSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigSpec) SetKafka(v *ConfigSpec_Kafka) {
	m.Kafka = v
}

func (m *ConfigSpec) SetZookeeper(v *ConfigSpec_Zookeeper) {
	m.Zookeeper = v
}

func (m *ConfigSpec) SetZoneId(v []string) {
	m.ZoneId = v
}

func (m *ConfigSpec) SetBrokersCount(v *wrappers.Int64Value) {
	m.BrokersCount = v
}

func (m *ConfigSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *ConfigSpec) SetUnmanagedTopics(v bool) {
	m.UnmanagedTopics = v
}

func (m *ConfigSpec) SetSchemaRegistry(v bool) {
	m.SchemaRegistry = v
}

type ConfigSpec_Kafka_KafkaConfig = isConfigSpec_Kafka_KafkaConfig

func (m *ConfigSpec_Kafka) SetKafkaConfig(v ConfigSpec_Kafka_KafkaConfig) {
	m.KafkaConfig = v
}

func (m *ConfigSpec_Kafka) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ConfigSpec_Kafka) SetKafkaConfig_2_1(v *KafkaConfig2_1) {
	m.KafkaConfig = &ConfigSpec_Kafka_KafkaConfig_2_1{
		KafkaConfig_2_1: v,
	}
}

func (m *ConfigSpec_Kafka) SetKafkaConfig_2_6(v *KafkaConfig2_6) {
	m.KafkaConfig = &ConfigSpec_Kafka_KafkaConfig_2_6{
		KafkaConfig_2_6: v,
	}
}

func (m *ConfigSpec_Kafka) SetKafkaConfig_2_8(v *KafkaConfig2_8) {
	m.KafkaConfig = &ConfigSpec_Kafka_KafkaConfig_2_8{
		KafkaConfig_2_8: v,
	}
}

func (m *ConfigSpec_Zookeeper) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Resources) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *Resources) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *Resources) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *KafkaConfig2_1) SetCompressionType(v CompressionType) {
	m.CompressionType = v
}

func (m *KafkaConfig2_1) SetLogFlushIntervalMessages(v *wrappers.Int64Value) {
	m.LogFlushIntervalMessages = v
}

func (m *KafkaConfig2_1) SetLogFlushIntervalMs(v *wrappers.Int64Value) {
	m.LogFlushIntervalMs = v
}

func (m *KafkaConfig2_1) SetLogFlushSchedulerIntervalMs(v *wrappers.Int64Value) {
	m.LogFlushSchedulerIntervalMs = v
}

func (m *KafkaConfig2_1) SetLogRetentionBytes(v *wrappers.Int64Value) {
	m.LogRetentionBytes = v
}

func (m *KafkaConfig2_1) SetLogRetentionHours(v *wrappers.Int64Value) {
	m.LogRetentionHours = v
}

func (m *KafkaConfig2_1) SetLogRetentionMinutes(v *wrappers.Int64Value) {
	m.LogRetentionMinutes = v
}

func (m *KafkaConfig2_1) SetLogRetentionMs(v *wrappers.Int64Value) {
	m.LogRetentionMs = v
}

func (m *KafkaConfig2_1) SetLogSegmentBytes(v *wrappers.Int64Value) {
	m.LogSegmentBytes = v
}

func (m *KafkaConfig2_1) SetLogPreallocate(v *wrappers.BoolValue) {
	m.LogPreallocate = v
}

func (m *KafkaConfig2_1) SetSocketSendBufferBytes(v *wrappers.Int64Value) {
	m.SocketSendBufferBytes = v
}

func (m *KafkaConfig2_1) SetSocketReceiveBufferBytes(v *wrappers.Int64Value) {
	m.SocketReceiveBufferBytes = v
}

func (m *KafkaConfig2_1) SetAutoCreateTopicsEnable(v *wrappers.BoolValue) {
	m.AutoCreateTopicsEnable = v
}

func (m *KafkaConfig2_1) SetNumPartitions(v *wrappers.Int64Value) {
	m.NumPartitions = v
}

func (m *KafkaConfig2_1) SetDefaultReplicationFactor(v *wrappers.Int64Value) {
	m.DefaultReplicationFactor = v
}

func (m *KafkaConfig2_6) SetCompressionType(v CompressionType) {
	m.CompressionType = v
}

func (m *KafkaConfig2_6) SetLogFlushIntervalMessages(v *wrappers.Int64Value) {
	m.LogFlushIntervalMessages = v
}

func (m *KafkaConfig2_6) SetLogFlushIntervalMs(v *wrappers.Int64Value) {
	m.LogFlushIntervalMs = v
}

func (m *KafkaConfig2_6) SetLogFlushSchedulerIntervalMs(v *wrappers.Int64Value) {
	m.LogFlushSchedulerIntervalMs = v
}

func (m *KafkaConfig2_6) SetLogRetentionBytes(v *wrappers.Int64Value) {
	m.LogRetentionBytes = v
}

func (m *KafkaConfig2_6) SetLogRetentionHours(v *wrappers.Int64Value) {
	m.LogRetentionHours = v
}

func (m *KafkaConfig2_6) SetLogRetentionMinutes(v *wrappers.Int64Value) {
	m.LogRetentionMinutes = v
}

func (m *KafkaConfig2_6) SetLogRetentionMs(v *wrappers.Int64Value) {
	m.LogRetentionMs = v
}

func (m *KafkaConfig2_6) SetLogSegmentBytes(v *wrappers.Int64Value) {
	m.LogSegmentBytes = v
}

func (m *KafkaConfig2_6) SetLogPreallocate(v *wrappers.BoolValue) {
	m.LogPreallocate = v
}

func (m *KafkaConfig2_6) SetSocketSendBufferBytes(v *wrappers.Int64Value) {
	m.SocketSendBufferBytes = v
}

func (m *KafkaConfig2_6) SetSocketReceiveBufferBytes(v *wrappers.Int64Value) {
	m.SocketReceiveBufferBytes = v
}

func (m *KafkaConfig2_6) SetAutoCreateTopicsEnable(v *wrappers.BoolValue) {
	m.AutoCreateTopicsEnable = v
}

func (m *KafkaConfig2_6) SetNumPartitions(v *wrappers.Int64Value) {
	m.NumPartitions = v
}

func (m *KafkaConfig2_6) SetDefaultReplicationFactor(v *wrappers.Int64Value) {
	m.DefaultReplicationFactor = v
}

func (m *KafkaConfig2_8) SetCompressionType(v CompressionType) {
	m.CompressionType = v
}

func (m *KafkaConfig2_8) SetLogFlushIntervalMessages(v *wrappers.Int64Value) {
	m.LogFlushIntervalMessages = v
}

func (m *KafkaConfig2_8) SetLogFlushIntervalMs(v *wrappers.Int64Value) {
	m.LogFlushIntervalMs = v
}

func (m *KafkaConfig2_8) SetLogFlushSchedulerIntervalMs(v *wrappers.Int64Value) {
	m.LogFlushSchedulerIntervalMs = v
}

func (m *KafkaConfig2_8) SetLogRetentionBytes(v *wrappers.Int64Value) {
	m.LogRetentionBytes = v
}

func (m *KafkaConfig2_8) SetLogRetentionHours(v *wrappers.Int64Value) {
	m.LogRetentionHours = v
}

func (m *KafkaConfig2_8) SetLogRetentionMinutes(v *wrappers.Int64Value) {
	m.LogRetentionMinutes = v
}

func (m *KafkaConfig2_8) SetLogRetentionMs(v *wrappers.Int64Value) {
	m.LogRetentionMs = v
}

func (m *KafkaConfig2_8) SetLogSegmentBytes(v *wrappers.Int64Value) {
	m.LogSegmentBytes = v
}

func (m *KafkaConfig2_8) SetLogPreallocate(v *wrappers.BoolValue) {
	m.LogPreallocate = v
}

func (m *KafkaConfig2_8) SetSocketSendBufferBytes(v *wrappers.Int64Value) {
	m.SocketSendBufferBytes = v
}

func (m *KafkaConfig2_8) SetSocketReceiveBufferBytes(v *wrappers.Int64Value) {
	m.SocketReceiveBufferBytes = v
}

func (m *KafkaConfig2_8) SetAutoCreateTopicsEnable(v *wrappers.BoolValue) {
	m.AutoCreateTopicsEnable = v
}

func (m *KafkaConfig2_8) SetNumPartitions(v *wrappers.Int64Value) {
	m.NumPartitions = v
}

func (m *KafkaConfig2_8) SetDefaultReplicationFactor(v *wrappers.Int64Value) {
	m.DefaultReplicationFactor = v
}

func (m *Host) SetName(v string) {
	m.Name = v
}

func (m *Host) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Host) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *Host) SetRole(v Host_Role) {
	m.Role = v
}

func (m *Host) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Host) SetHealth(v Host_Health) {
	m.Health = v
}

func (m *Host) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Host) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}
