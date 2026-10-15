# InlineObject110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mask** | Pointer to **int32** | Mask used for the subnet of all MGs in  this network. | [optional] 
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool. | [optional] 

## Methods

### NewInlineObject110

`func NewInlineObject110() *InlineObject110`

NewInlineObject110 instantiates a new InlineObject110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject110WithDefaults

`func NewInlineObject110WithDefaults() *InlineObject110`

NewInlineObject110WithDefaults instantiates a new InlineObject110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMask

`func (o *InlineObject110) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineObject110) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineObject110) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineObject110) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetCidr

`func (o *InlineObject110) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineObject110) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineObject110) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineObject110) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


