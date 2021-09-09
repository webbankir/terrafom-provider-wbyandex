// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1/config"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
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

func (m *Cluster) SetConfig(v *ClusterConfig) {
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

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
}

func (m *Cluster) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
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

type ClusterConfig_MysqlConfig = isClusterConfig_MysqlConfig

func (m *ClusterConfig) SetMysqlConfig(v ClusterConfig_MysqlConfig) {
	m.MysqlConfig = v
}

func (m *ClusterConfig) SetVersion(v string) {
	m.Version = v
}

func (m *ClusterConfig) SetMysqlConfig_5_7(v *config.MysqlConfigSet5_7) {
	m.MysqlConfig = &ClusterConfig_MysqlConfig_5_7{
		MysqlConfig_5_7: v,
	}
}

func (m *ClusterConfig) SetMysqlConfig_8_0(v *config.MysqlConfigSet8_0) {
	m.MysqlConfig = &ClusterConfig_MysqlConfig_8_0{
		MysqlConfig_8_0: v,
	}
}

func (m *ClusterConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ClusterConfig) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ClusterConfig) SetAccess(v *Access) {
	m.Access = v
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

func (m *Host) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Host) SetRole(v Host_Role) {
	m.Role = v
}

func (m *Host) SetHealth(v Host_Health) {
	m.Health = v
}

func (m *Host) SetServices(v []*Service) {
	m.Services = v
}

func (m *Host) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Host) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *Host) SetReplicationSource(v string) {
	m.ReplicationSource = v
}

func (m *Service) SetType(v Service_Type) {
	m.Type = v
}

func (m *Service) SetHealth(v Service_Health) {
	m.Health = v
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

func (m *Access) SetDataLens(v bool) {
	m.DataLens = v
}

func (m *Access) SetWebSql(v bool) {
	m.WebSql = v
}

func (m *PerformanceDiagnostics) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *PerformanceDiagnostics) SetSessionsSamplingInterval(v int64) {
	m.SessionsSamplingInterval = v
}

func (m *PerformanceDiagnostics) SetStatementsSamplingInterval(v int64) {
	m.StatementsSamplingInterval = v
}
