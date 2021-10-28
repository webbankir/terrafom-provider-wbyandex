// Code generated by protoc-gen-goext. DO NOT EDIT.

package saml

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Federation) SetId(v string) {
	m.Id = v
}

func (m *Federation) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *Federation) SetName(v string) {
	m.Name = v
}

func (m *Federation) SetDescription(v string) {
	m.Description = v
}

func (m *Federation) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Federation) SetCookieMaxAge(v *duration.Duration) {
	m.CookieMaxAge = v
}

func (m *Federation) SetAutoCreateAccountOnLogin(v bool) {
	m.AutoCreateAccountOnLogin = v
}

func (m *Federation) SetIssuer(v string) {
	m.Issuer = v
}

func (m *Federation) SetSsoBinding(v BindingType) {
	m.SsoBinding = v
}

func (m *Federation) SetSsoUrl(v string) {
	m.SsoUrl = v
}

func (m *Federation) SetSecuritySettings(v *FederationSecuritySettings) {
	m.SecuritySettings = v
}

func (m *Federation) SetCaseInsensitiveNameIds(v bool) {
	m.CaseInsensitiveNameIds = v
}

func (m *Federation) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *FederationSecuritySettings) SetEncryptedAssertions(v bool) {
	m.EncryptedAssertions = v
}
