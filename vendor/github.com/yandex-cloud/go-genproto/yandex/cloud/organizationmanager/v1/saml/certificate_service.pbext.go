// Code generated by protoc-gen-goext. DO NOT EDIT.

package saml

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetCertificateRequest) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *ListCertificatesRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *ListCertificatesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListCertificatesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListCertificatesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListCertificatesResponse) SetCertificates(v []*Certificate) {
	m.Certificates = v
}

func (m *ListCertificatesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateCertificateRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *CreateCertificateRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateCertificateRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateCertificateRequest) SetData(v string) {
	m.Data = v
}

func (m *CreateCertificateMetadata) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *UpdateCertificateRequest) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *UpdateCertificateRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateCertificateRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateCertificateRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateCertificateRequest) SetData(v string) {
	m.Data = v
}

func (m *UpdateCertificateMetadata) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *DeleteCertificateRequest) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *DeleteCertificateMetadata) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *ListCertificateOperationsRequest) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *ListCertificateOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListCertificateOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListCertificateOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListCertificateOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
