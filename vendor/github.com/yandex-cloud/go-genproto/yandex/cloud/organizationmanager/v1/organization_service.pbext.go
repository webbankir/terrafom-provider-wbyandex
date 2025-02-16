// Code generated by protoc-gen-goext. DO NOT EDIT.

package organizationmanager

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetOrganizationRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *ListOrganizationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListOrganizationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListOrganizationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListOrganizationsResponse) SetOrganizations(v []*Organization) {
	m.Organizations = v
}

func (m *ListOrganizationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateOrganizationRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *UpdateOrganizationRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateOrganizationRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateOrganizationRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateOrganizationRequest) SetTitle(v string) {
	m.Title = v
}

func (m *UpdateOrganizationRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateOrganizationMetadata) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *ListOrganizationOperationsRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *ListOrganizationOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListOrganizationOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListOrganizationOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListOrganizationOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
