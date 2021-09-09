// Code generated by protoc-gen-goext. DO NOT EDIT.

package devices

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetDeviceRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *GetDeviceRequest) SetDeviceView(v DeviceView) {
	m.DeviceView = v
}

func (m *GetByNameDeviceRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *GetByNameDeviceRequest) SetDeviceName(v string) {
	m.DeviceName = v
}

func (m *GetByNameDeviceRequest) SetDeviceView(v DeviceView) {
	m.DeviceView = v
}

type ListDevicesRequest_Id = isListDevicesRequest_Id

func (m *ListDevicesRequest) SetId(v ListDevicesRequest_Id) {
	m.Id = v
}

func (m *ListDevicesRequest) SetRegistryId(v string) {
	m.Id = &ListDevicesRequest_RegistryId{
		RegistryId: v,
	}
}

func (m *ListDevicesRequest) SetFolderId(v string) {
	m.Id = &ListDevicesRequest_FolderId{
		FolderId: v,
	}
}

func (m *ListDevicesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDevicesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDevicesRequest) SetDeviceView(v DeviceView) {
	m.DeviceView = v
}

func (m *ListDevicesResponse) SetDevices(v []*Device) {
	m.Devices = v
}

func (m *ListDevicesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateDeviceRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *CreateDeviceRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateDeviceRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateDeviceRequest) SetCertificates(v []*CreateDeviceRequest_Certificate) {
	m.Certificates = v
}

func (m *CreateDeviceRequest) SetTopicAliases(v map[string]string) {
	m.TopicAliases = v
}

func (m *CreateDeviceRequest) SetPassword(v string) {
	m.Password = v
}

func (m *CreateDeviceRequest_Certificate) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *CreateDeviceMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *UpdateDeviceRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *UpdateDeviceRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateDeviceRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateDeviceRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateDeviceRequest) SetTopicAliases(v map[string]string) {
	m.TopicAliases = v
}

func (m *UpdateDeviceMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeleteDeviceRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeleteDeviceMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *ListDeviceCertificatesRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *ListDeviceCertificatesResponse) SetCertificates(v []*DeviceCertificate) {
	m.Certificates = v
}

func (m *AddDeviceCertificateRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *AddDeviceCertificateRequest) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *AddDeviceCertificateMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *AddDeviceCertificateMetadata) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeleteDeviceCertificateRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeleteDeviceCertificateRequest) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeleteDeviceCertificateMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeleteDeviceCertificateMetadata) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *ListDevicePasswordsRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *ListDevicePasswordsResponse) SetPasswords(v []*DevicePassword) {
	m.Passwords = v
}

func (m *AddDevicePasswordRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *AddDevicePasswordRequest) SetPassword(v string) {
	m.Password = v
}

func (m *AddDevicePasswordMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *AddDevicePasswordMetadata) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *DeleteDevicePasswordRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeleteDevicePasswordRequest) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *DeleteDevicePasswordMetadata) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *DeleteDevicePasswordMetadata) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *ListDeviceOperationsRequest) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *ListDeviceOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDeviceOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDeviceOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDeviceOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListDeviceOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
