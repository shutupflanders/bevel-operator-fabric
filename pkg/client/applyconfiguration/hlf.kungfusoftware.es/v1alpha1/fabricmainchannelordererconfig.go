/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
)

// FabricMainChannelOrdererConfigApplyConfiguration represents an declarative configuration of the FabricMainChannelOrdererConfig type for use
// with apply.
type FabricMainChannelOrdererConfigApplyConfiguration struct {
	OrdererType  *string                                                       `json:"ordererType,omitempty"`
	Capabilities []string                                                      `json:"capabilities,omitempty"`
	Policies     *map[string]FabricMainChannelPoliciesConfigApplyConfiguration `json:"policies,omitempty"`
	BatchTimeout *string                                                       `json:"batchTimeout,omitempty"`
	BatchSize    *FabricMainChannelOrdererBatchSizeApplyConfiguration          `json:"batchSize,omitempty"`
	State        *hlfkungfusoftwareesv1alpha1.FabricMainChannelConsensusState  `json:"state,omitempty"`
	EtcdRaft     *FabricMainChannelEtcdRaftApplyConfiguration                  `json:"etcdRaft,omitempty"`
}

// FabricMainChannelOrdererConfigApplyConfiguration constructs an declarative configuration of the FabricMainChannelOrdererConfig type for use with
// apply.
func FabricMainChannelOrdererConfig() *FabricMainChannelOrdererConfigApplyConfiguration {
	return &FabricMainChannelOrdererConfigApplyConfiguration{}
}

// WithOrdererType sets the OrdererType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OrdererType field is set to the value of the last call.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithOrdererType(value string) *FabricMainChannelOrdererConfigApplyConfiguration {
	b.OrdererType = &value
	return b
}

// WithCapabilities adds the given value to the Capabilities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Capabilities field.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithCapabilities(values ...string) *FabricMainChannelOrdererConfigApplyConfiguration {
	for i := range values {
		b.Capabilities = append(b.Capabilities, values[i])
	}
	return b
}

// WithPolicies sets the Policies field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policies field is set to the value of the last call.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithPolicies(value map[string]FabricMainChannelPoliciesConfigApplyConfiguration) *FabricMainChannelOrdererConfigApplyConfiguration {
	b.Policies = &value
	return b
}

// WithBatchTimeout sets the BatchTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BatchTimeout field is set to the value of the last call.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithBatchTimeout(value string) *FabricMainChannelOrdererConfigApplyConfiguration {
	b.BatchTimeout = &value
	return b
}

// WithBatchSize sets the BatchSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BatchSize field is set to the value of the last call.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithBatchSize(value *FabricMainChannelOrdererBatchSizeApplyConfiguration) *FabricMainChannelOrdererConfigApplyConfiguration {
	b.BatchSize = value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithState(value hlfkungfusoftwareesv1alpha1.FabricMainChannelConsensusState) *FabricMainChannelOrdererConfigApplyConfiguration {
	b.State = &value
	return b
}

// WithEtcdRaft sets the EtcdRaft field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EtcdRaft field is set to the value of the last call.
func (b *FabricMainChannelOrdererConfigApplyConfiguration) WithEtcdRaft(value *FabricMainChannelEtcdRaftApplyConfiguration) *FabricMainChannelOrdererConfigApplyConfiguration {
	b.EtcdRaft = value
	return b
}
