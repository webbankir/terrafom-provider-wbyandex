// Code generated by protoc-gen-goext. DO NOT EDIT.

package logging

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *LogGroup) SetId(v string) {
	m.Id = v
}

func (m *LogGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *LogGroup) SetCloudId(v string) {
	m.CloudId = v
}

func (m *LogGroup) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *LogGroup) SetName(v string) {
	m.Name = v
}

func (m *LogGroup) SetDescription(v string) {
	m.Description = v
}

func (m *LogGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *LogGroup) SetStatus(v LogGroup_Status) {
	m.Status = v
}

func (m *LogGroup) SetRetentionPeriod(v *duration.Duration) {
	m.RetentionPeriod = v
}
