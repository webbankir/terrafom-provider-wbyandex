// Code generated by protoc-gen-goext. DO NOT EDIT.

package redis

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func (m *RedisConfig6_2) SetMaxmemoryPolicy(v RedisConfig6_2_MaxmemoryPolicy) {
	m.MaxmemoryPolicy = v
}

func (m *RedisConfig6_2) SetTimeout(v *wrappers.Int64Value) {
	m.Timeout = v
}

func (m *RedisConfig6_2) SetPassword(v string) {
	m.Password = v
}

func (m *RedisConfig6_2) SetDatabases(v *wrappers.Int64Value) {
	m.Databases = v
}

func (m *RedisConfig6_2) SetSlowlogLogSlowerThan(v *wrappers.Int64Value) {
	m.SlowlogLogSlowerThan = v
}

func (m *RedisConfig6_2) SetSlowlogMaxLen(v *wrappers.Int64Value) {
	m.SlowlogMaxLen = v
}

func (m *RedisConfig6_2) SetNotifyKeyspaceEvents(v string) {
	m.NotifyKeyspaceEvents = v
}

func (m *RedisConfigSet6_2) SetEffectiveConfig(v *RedisConfig6_2) {
	m.EffectiveConfig = v
}

func (m *RedisConfigSet6_2) SetUserConfig(v *RedisConfig6_2) {
	m.UserConfig = v
}

func (m *RedisConfigSet6_2) SetDefaultConfig(v *RedisConfig6_2) {
	m.DefaultConfig = v
}
