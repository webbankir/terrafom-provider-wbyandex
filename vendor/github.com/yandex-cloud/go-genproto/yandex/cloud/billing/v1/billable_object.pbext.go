// Code generated by protoc-gen-goext. DO NOT EDIT.

package billing

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *BillableObject) SetId(v string) {
	m.Id = v
}

func (m *BillableObject) SetType(v string) {
	m.Type = v
}

func (m *BillableObjectBinding) SetEffectiveTime(v *timestamp.Timestamp) {
	m.EffectiveTime = v
}

func (m *BillableObjectBinding) SetBillableObject(v *BillableObject) {
	m.BillableObject = v
}
