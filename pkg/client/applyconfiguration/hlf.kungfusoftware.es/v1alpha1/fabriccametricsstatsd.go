/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCAMetricsStatsdApplyConfiguration represents an declarative configuration of the FabricCAMetricsStatsd type for use
// with apply.
type FabricCAMetricsStatsdApplyConfiguration struct {
	Network       *string `json:"network,omitempty"`
	Address       *string `json:"address,omitempty"`
	WriteInterval *string `json:"writeInterval,omitempty"`
	Prefix        *string `json:"prefix,omitempty"`
}

// FabricCAMetricsStatsdApplyConfiguration constructs an declarative configuration of the FabricCAMetricsStatsd type for use with
// apply.
func FabricCAMetricsStatsd() *FabricCAMetricsStatsdApplyConfiguration {
	return &FabricCAMetricsStatsdApplyConfiguration{}
}

// WithNetwork sets the Network field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Network field is set to the value of the last call.
func (b *FabricCAMetricsStatsdApplyConfiguration) WithNetwork(value string) *FabricCAMetricsStatsdApplyConfiguration {
	b.Network = &value
	return b
}

// WithAddress sets the Address field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Address field is set to the value of the last call.
func (b *FabricCAMetricsStatsdApplyConfiguration) WithAddress(value string) *FabricCAMetricsStatsdApplyConfiguration {
	b.Address = &value
	return b
}

// WithWriteInterval sets the WriteInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WriteInterval field is set to the value of the last call.
func (b *FabricCAMetricsStatsdApplyConfiguration) WithWriteInterval(value string) *FabricCAMetricsStatsdApplyConfiguration {
	b.WriteInterval = &value
	return b
}

// WithPrefix sets the Prefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Prefix field is set to the value of the last call.
func (b *FabricCAMetricsStatsdApplyConfiguration) WithPrefix(value string) *FabricCAMetricsStatsdApplyConfiguration {
	b.Prefix = &value
	return b
}
