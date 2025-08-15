# InlineResponse200160

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMtuSize** | Pointer to **int32** | MTU size for the entire network. Default value is 9578. | [optional] 
**Overrides** | Pointer to [**[]InlineResponse200160Overrides**](InlineResponse200160Overrides.md) | Override MTU size for individual switches or switch templates.       An empty array will clear overrides. | [optional] 

## Methods

### NewInlineResponse200160

`func NewInlineResponse200160() *InlineResponse200160`

NewInlineResponse200160 instantiates a new InlineResponse200160 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200160WithDefaults

`func NewInlineResponse200160WithDefaults() *InlineResponse200160`

NewInlineResponse200160WithDefaults instantiates a new InlineResponse200160 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMtuSize

`func (o *InlineResponse200160) GetDefaultMtuSize() int32`

GetDefaultMtuSize returns the DefaultMtuSize field if non-nil, zero value otherwise.

### GetDefaultMtuSizeOk

`func (o *InlineResponse200160) GetDefaultMtuSizeOk() (*int32, bool)`

GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtuSize

`func (o *InlineResponse200160) SetDefaultMtuSize(v int32)`

SetDefaultMtuSize sets DefaultMtuSize field to given value.

### HasDefaultMtuSize

`func (o *InlineResponse200160) HasDefaultMtuSize() bool`

HasDefaultMtuSize returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse200160) GetOverrides() []InlineResponse200160Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse200160) GetOverridesOk() (*[]InlineResponse200160Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse200160) SetOverrides(v []InlineResponse200160Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse200160) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


