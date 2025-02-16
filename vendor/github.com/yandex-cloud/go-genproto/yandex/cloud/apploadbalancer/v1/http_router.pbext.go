// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *HttpRouter) SetId(v string) {
	m.Id = v
}

func (m *HttpRouter) SetName(v string) {
	m.Name = v
}

func (m *HttpRouter) SetDescription(v string) {
	m.Description = v
}

func (m *HttpRouter) SetFolderId(v string) {
	m.FolderId = v
}

func (m *HttpRouter) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *HttpRouter) SetVirtualHosts(v []*VirtualHost) {
	m.VirtualHosts = v
}

func (m *HttpRouter) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *HttpRouter) SetRouteOptions(v *RouteOptions) {
	m.RouteOptions = v
}
