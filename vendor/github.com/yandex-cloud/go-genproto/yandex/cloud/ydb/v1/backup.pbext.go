// Code generated by protoc-gen-goext. DO NOT EDIT.

package ydb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
)

type BackupSchedule_Policy = isBackupSchedule_Policy

func (m *BackupSchedule) SetPolicy(v BackupSchedule_Policy) {
	m.Policy = v
}

func (m *BackupSchedule) SetDailyBackupSchedule(v *DailyBackupSchedule) {
	m.Policy = &BackupSchedule_DailyBackupSchedule{
		DailyBackupSchedule: v,
	}
}

func (m *BackupSchedule) SetWeeklyBackupSchedule(v *WeeklyBackupSchedule) {
	m.Policy = &BackupSchedule_WeeklyBackupSchedule{
		WeeklyBackupSchedule: v,
	}
}

func (m *BackupSchedule) SetRecurringBackupSchedule(v *RecurringBackupSchedule) {
	m.Policy = &BackupSchedule_RecurringBackupSchedule{
		RecurringBackupSchedule: v,
	}
}

func (m *BackupSchedule) SetNextExecuteTime(v *timestamp.Timestamp) {
	m.NextExecuteTime = v
}

func (m *RecurringBackupSchedule) SetStartTime(v *timestamp.Timestamp) {
	m.StartTime = v
}

func (m *RecurringBackupSchedule) SetRecurrence(v string) {
	m.Recurrence = v
}

func (m *DaysOfWeekBackupSchedule) SetDays(v []dayofweek.DayOfWeek) {
	m.Days = v
}

func (m *DaysOfWeekBackupSchedule) SetExecuteTime(v *timeofday.TimeOfDay) {
	m.ExecuteTime = v
}

func (m *WeeklyBackupSchedule) SetDaysOfWeek(v []*DaysOfWeekBackupSchedule) {
	m.DaysOfWeek = v
}

func (m *DailyBackupSchedule) SetExecuteTime(v *timeofday.TimeOfDay) {
	m.ExecuteTime = v
}

func (m *BackupSettings) SetName(v string) {
	m.Name = v
}

func (m *BackupSettings) SetDescription(v string) {
	m.Description = v
}

func (m *BackupSettings) SetBackupSchedule(v *BackupSchedule) {
	m.BackupSchedule = v
}

func (m *BackupSettings) SetBackupTimeToLive(v *duration.Duration) {
	m.BackupTimeToLive = v
}

func (m *BackupSettings) SetSourcePaths(v []string) {
	m.SourcePaths = v
}

func (m *BackupSettings) SetSourcePathsToExclude(v []string) {
	m.SourcePathsToExclude = v
}

func (m *BackupSettings) SetType(v BackupSettings_Type) {
	m.Type = v
}

func (m *BackupSettings) SetStorageClass(v BackupSettings_StorageClass) {
	m.StorageClass = v
}

func (m *BackupConfig) SetBackupSettings(v []*BackupSettings) {
	m.BackupSettings = v
}

func (m *Backup) SetId(v string) {
	m.Id = v
}

func (m *Backup) SetName(v string) {
	m.Name = v
}

func (m *Backup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Backup) SetDatabaseId(v string) {
	m.DatabaseId = v
}

func (m *Backup) SetDescription(v string) {
	m.Description = v
}

func (m *Backup) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Backup) SetStartedAt(v *timestamp.Timestamp) {
	m.StartedAt = v
}

func (m *Backup) SetCompletedAt(v *timestamp.Timestamp) {
	m.CompletedAt = v
}

func (m *Backup) SetStatus(v Backup_Status) {
	m.Status = v
}

func (m *Backup) SetBackupSettings(v *BackupSettings) {
	m.BackupSettings = v
}

func (m *Backup) SetType(v Backup_Type) {
	m.Type = v
}

func (m *Backup) SetSize(v int64) {
	m.Size = v
}
