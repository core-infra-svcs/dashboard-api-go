# InlineResponse200149

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMtuSize** | Pointer to **int32** | MTU size for the entire network. Default value is 9578. | [optional] 
**Overrides** | Pointer to [**[]InlineResponse200149Overrides**](InlineResponse200149Overrides.md) | Override MTU size for individual switches or switch templates.       An empty array will clear overrides. | [optional] 

## Methods

### NewInlineResponse200149

`func NewInlineResponse200149() *InlineResponse200149`

NewInlineResponse200149 instantiates a new InlineResponse200149 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200149WithDefaults

`func NewInlineResponse200149WithDefaults() *InlineResponse200149`

NewInlineResponse200149WithDefaults instantiates a new InlineResponse200149 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMtuSize

`func (o *InlineResponse200149) GetDefaultMtuSize() int32`

GetDefaultMtuSize returns the DefaultMtuSize field if non-nil, zero value otherwise.

### GetDefaultMtuSizeOk

`func (o *InlineResponse200149) GetDefaultMtuSizeOk() (*int32, bool)`

GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtuSize

`func (o *InlineResponse200149) SetDefaultMtuSize(v int32)`

SetDefaultMtuSize sets DefaultMtuSize field to given value.

### HasDefaultMtuSize

`func (o *InlineResponse200149) HasDefaultMtuSize() bool`

HasDefaultMtuSize returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse200149) GetOverrides() []InlineResponse200149Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse200149) GetOverridesOk() (*[]InlineResponse200149Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse200149) SetOverrides(v []InlineResponse200149Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse200149) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


