/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelOrdererBatchSizeApplyConfiguration represents an declarative configuration of the FabricMainChannelOrdererBatchSize type for use
// with apply.
type FabricMainChannelOrdererBatchSizeApplyConfiguration struct {
	MaxMessageCount   *int `json:"maxMessageCount,omitempty"`
	AbsoluteMaxBytes  *int `json:"absoluteMaxBytes,omitempty"`
	PreferredMaxBytes *int `json:"preferredMaxBytes,omitempty"`
}

// FabricMainChannelOrdererBatchSizeApplyConfiguration constructs an declarative configuration of the FabricMainChannelOrdererBatchSize type for use with
// apply.
func FabricMainChannelOrdererBatchSize() *FabricMainChannelOrdererBatchSizeApplyConfiguration {
	return &FabricMainChannelOrdererBatchSizeApplyConfiguration{}
}

// WithMaxMessageCount sets the MaxMessageCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxMessageCount field is set to the value of the last call.
func (b *FabricMainChannelOrdererBatchSizeApplyConfiguration) WithMaxMessageCount(value int) *FabricMainChannelOrdererBatchSizeApplyConfiguration {
	b.MaxMessageCount = &value
	return b
}

// WithAbsoluteMaxBytes sets the AbsoluteMaxBytes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AbsoluteMaxBytes field is set to the value of the last call.
func (b *FabricMainChannelOrdererBatchSizeApplyConfiguration) WithAbsoluteMaxBytes(value int) *FabricMainChannelOrdererBatchSizeApplyConfiguration {
	b.AbsoluteMaxBytes = &value
	return b
}

// WithPreferredMaxBytes sets the PreferredMaxBytes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreferredMaxBytes field is set to the value of the last call.
func (b *FabricMainChannelOrdererBatchSizeApplyConfiguration) WithPreferredMaxBytes(value int) *FabricMainChannelOrdererBatchSizeApplyConfiguration {
	b.PreferredMaxBytes = &value
	return b
}
