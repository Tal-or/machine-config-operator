// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/machine/v1"
)

// AWSResourceReferenceApplyConfiguration represents an declarative configuration of the AWSResourceReference type for use
// with apply.
type AWSResourceReferenceApplyConfiguration struct {
	Type    *v1.AWSResourceReferenceType           `json:"type,omitempty"`
	ID      *string                                `json:"id,omitempty"`
	ARN     *string                                `json:"arn,omitempty"`
	Filters *[]AWSResourceFilterApplyConfiguration `json:"filters,omitempty"`
}

// AWSResourceReferenceApplyConfiguration constructs an declarative configuration of the AWSResourceReference type for use with
// apply.
func AWSResourceReference() *AWSResourceReferenceApplyConfiguration {
	return &AWSResourceReferenceApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AWSResourceReferenceApplyConfiguration) WithType(value v1.AWSResourceReferenceType) *AWSResourceReferenceApplyConfiguration {
	b.Type = &value
	return b
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *AWSResourceReferenceApplyConfiguration) WithID(value string) *AWSResourceReferenceApplyConfiguration {
	b.ID = &value
	return b
}

// WithARN sets the ARN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ARN field is set to the value of the last call.
func (b *AWSResourceReferenceApplyConfiguration) WithARN(value string) *AWSResourceReferenceApplyConfiguration {
	b.ARN = &value
	return b
}

func (b *AWSResourceReferenceApplyConfiguration) ensureAWSResourceFilterApplyConfigurationExists() {
	if b.Filters == nil {
		b.Filters = &[]AWSResourceFilterApplyConfiguration{}
	}
}

// WithFilters adds the given value to the Filters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Filters field.
func (b *AWSResourceReferenceApplyConfiguration) WithFilters(values ...*AWSResourceFilterApplyConfiguration) *AWSResourceReferenceApplyConfiguration {
	b.ensureAWSResourceFilterApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFilters")
		}
		*b.Filters = append(*b.Filters, *values[i])
	}
	return b
}