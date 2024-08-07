/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// The SKU identifier of the entitlement
	Sku *string `json:"sku,omitempty"`
	// The user-facing name of the entitlement
	Name *string `json:"name,omitempty"`
	// The product type of the entitlement
	ProductType *string `json:"productType,omitempty"`
	// The product class associated with the entitlement
	ProductClass *string `json:"productClass,omitempty"`
	// The feature tier associated with the entitlement (null for add-ons)
	FeatureTier *string `json:"featureTier,omitempty"`
	// Whether or not the entitlement is an add-on
	IsAddOn *bool `json:"isAddOn,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InlineResponse2001) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InlineResponse2001) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *InlineResponse2001) SetSku(v string) {
	o.Sku = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2001) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2001) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2001) SetName(v string) {
	o.Name = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse2001) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse2001) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse2001) SetProductType(v string) {
	o.ProductType = &v
}

// GetProductClass returns the ProductClass field value if set, zero value otherwise.
func (o *InlineResponse2001) GetProductClass() string {
	if o == nil || isNil(o.ProductClass) {
		var ret string
		return ret
	}
	return *o.ProductClass
}

// GetProductClassOk returns a tuple with the ProductClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetProductClassOk() (*string, bool) {
	if o == nil || isNil(o.ProductClass) {
    return nil, false
	}
	return o.ProductClass, true
}

// HasProductClass returns a boolean if a field has been set.
func (o *InlineResponse2001) HasProductClass() bool {
	if o != nil && !isNil(o.ProductClass) {
		return true
	}

	return false
}

// SetProductClass gets a reference to the given string and assigns it to the ProductClass field.
func (o *InlineResponse2001) SetProductClass(v string) {
	o.ProductClass = &v
}

// GetFeatureTier returns the FeatureTier field value if set, zero value otherwise.
func (o *InlineResponse2001) GetFeatureTier() string {
	if o == nil || isNil(o.FeatureTier) {
		var ret string
		return ret
	}
	return *o.FeatureTier
}

// GetFeatureTierOk returns a tuple with the FeatureTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetFeatureTierOk() (*string, bool) {
	if o == nil || isNil(o.FeatureTier) {
    return nil, false
	}
	return o.FeatureTier, true
}

// HasFeatureTier returns a boolean if a field has been set.
func (o *InlineResponse2001) HasFeatureTier() bool {
	if o != nil && !isNil(o.FeatureTier) {
		return true
	}

	return false
}

// SetFeatureTier gets a reference to the given string and assigns it to the FeatureTier field.
func (o *InlineResponse2001) SetFeatureTier(v string) {
	o.FeatureTier = &v
}

// GetIsAddOn returns the IsAddOn field value if set, zero value otherwise.
func (o *InlineResponse2001) GetIsAddOn() bool {
	if o == nil || isNil(o.IsAddOn) {
		var ret bool
		return ret
	}
	return *o.IsAddOn
}

// GetIsAddOnOk returns a tuple with the IsAddOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetIsAddOnOk() (*bool, bool) {
	if o == nil || isNil(o.IsAddOn) {
    return nil, false
	}
	return o.IsAddOn, true
}

// HasIsAddOn returns a boolean if a field has been set.
func (o *InlineResponse2001) HasIsAddOn() bool {
	if o != nil && !isNil(o.IsAddOn) {
		return true
	}

	return false
}

// SetIsAddOn gets a reference to the given bool and assigns it to the IsAddOn field.
func (o *InlineResponse2001) SetIsAddOn(v bool) {
	o.IsAddOn = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.ProductClass) {
		toSerialize["productClass"] = o.ProductClass
	}
	if !isNil(o.FeatureTier) {
		toSerialize["featureTier"] = o.FeatureTier
	}
	if !isNil(o.IsAddOn) {
		toSerialize["isAddOn"] = o.IsAddOn
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


