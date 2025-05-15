# InlineResponse200159

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMtuSize** | Pointer to **int32** | MTU size for the entire network. Default value is 9578. | [optional] 
**Overrides** | Pointer to [**[]InlineResponse200159Overrides**](InlineResponse200159Overrides.md) | Override MTU size for individual switches or switch templates.       An empty array will clear overrides. | [optional] 

## Methods

### NewInlineResponse200159

`func NewInlineResponse200159() *InlineResponse200159`

NewInlineResponse200159 instantiates a new InlineResponse200159 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200159WithDefaults

`func NewInlineResponse200159WithDefaults() *InlineResponse200159`

NewInlineResponse200159WithDefaults instantiates a new InlineResponse200159 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMtuSize

`func (o *InlineResponse200159) GetDefaultMtuSize() int32`

GetDefaultMtuSize returns the DefaultMtuSize field if non-nil, zero value otherwise.

### GetDefaultMtuSizeOk

`func (o *InlineResponse200159) GetDefaultMtuSizeOk() (*int32, bool)`

GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtuSize

`func (o *InlineResponse200159) SetDefaultMtuSize(v int32)`

SetDefaultMtuSize sets DefaultMtuSize field to given value.

### HasDefaultMtuSize

`func (o *InlineResponse200159) HasDefaultMtuSize() bool`

HasDefaultMtuSize returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse200159) GetOverrides() []InlineResponse200159Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse200159) GetOverridesOk() (*[]InlineResponse200159Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse200159) SetOverrides(v []InlineResponse200159Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse200159) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


