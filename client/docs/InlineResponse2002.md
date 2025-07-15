# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The SKU identifier of the entitlement | [optional] 
**Name** | Pointer to **string** | The user-facing name of the entitlement | [optional] 
**ProductType** | Pointer to **string** | The product type of the entitlement | [optional] 
**ProductClass** | Pointer to **string** | The product class associated with the entitlement | [optional] 
**FeatureTier** | Pointer to **string** | The feature tier associated with the entitlement (null for add-ons) | [optional] 
**IsAddOn** | Pointer to **bool** | Whether or not the entitlement is an add-on | [optional] 
**IsFree** | Pointer to **bool** | Whether or not the entitlement is granted for free | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *InlineResponse2002) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *InlineResponse2002) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *InlineResponse2002) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *InlineResponse2002) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2002) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2002) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2002) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2002) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse2002) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse2002) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse2002) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse2002) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductClass

`func (o *InlineResponse2002) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *InlineResponse2002) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *InlineResponse2002) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *InlineResponse2002) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### GetFeatureTier

`func (o *InlineResponse2002) GetFeatureTier() string`

GetFeatureTier returns the FeatureTier field if non-nil, zero value otherwise.

### GetFeatureTierOk

`func (o *InlineResponse2002) GetFeatureTierOk() (*string, bool)`

GetFeatureTierOk returns a tuple with the FeatureTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTier

`func (o *InlineResponse2002) SetFeatureTier(v string)`

SetFeatureTier sets FeatureTier field to given value.

### HasFeatureTier

`func (o *InlineResponse2002) HasFeatureTier() bool`

HasFeatureTier returns a boolean if a field has been set.

### GetIsAddOn

`func (o *InlineResponse2002) GetIsAddOn() bool`

GetIsAddOn returns the IsAddOn field if non-nil, zero value otherwise.

### GetIsAddOnOk

`func (o *InlineResponse2002) GetIsAddOnOk() (*bool, bool)`

GetIsAddOnOk returns a tuple with the IsAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAddOn

`func (o *InlineResponse2002) SetIsAddOn(v bool)`

SetIsAddOn sets IsAddOn field to given value.

### HasIsAddOn

`func (o *InlineResponse2002) HasIsAddOn() bool`

HasIsAddOn returns a boolean if a field has been set.

### GetIsFree

`func (o *InlineResponse2002) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *InlineResponse2002) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *InlineResponse2002) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *InlineResponse2002) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


