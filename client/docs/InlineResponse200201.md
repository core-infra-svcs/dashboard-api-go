# InlineResponse200201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**InlineResponse200201DefaultSettings**](InlineResponse200201DefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]InlineResponse200201Overrides**](InlineResponse200201Overrides.md) | Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings. | [optional] 

## Methods

### NewInlineResponse200201

`func NewInlineResponse200201() *InlineResponse200201`

NewInlineResponse200201 instantiates a new InlineResponse200201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200201WithDefaults

`func NewInlineResponse200201WithDefaults() *InlineResponse200201`

NewInlineResponse200201WithDefaults instantiates a new InlineResponse200201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *InlineResponse200201) GetDefaultSettings() InlineResponse200201DefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *InlineResponse200201) GetDefaultSettingsOk() (*InlineResponse200201DefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *InlineResponse200201) SetDefaultSettings(v InlineResponse200201DefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *InlineResponse200201) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse200201) GetOverrides() []InlineResponse200201Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse200201) GetOverridesOk() (*[]InlineResponse200201Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse200201) SetOverrides(v []InlineResponse200201Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse200201) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


