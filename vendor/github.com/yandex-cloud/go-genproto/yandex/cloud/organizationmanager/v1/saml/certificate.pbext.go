// Code generated by protoc-gen-goext. DO NOT EDIT.

package saml

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Certificate) SetId(v string) {
	m.Id = v
}

func (m *Certificate) SetFederationId(v string) {
	m.FederationId = v
}

func (m *Certificate) SetName(v string) {
	m.Name = v
}

func (m *Certificate) SetDescription(v string) {
	m.Description = v
}

func (m *Certificate) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Certificate) SetData(v string) {
	m.Data = v
}
