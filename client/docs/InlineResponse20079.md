# InlineResponse20079

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**InlineResponse20079DefaultSettings**](InlineResponse20079DefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]InlineResponse20079Overrides**](InlineResponse20079Overrides.md) | Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings. | [optional] 

## Methods

### NewInlineResponse20079

`func NewInlineResponse20079() *InlineResponse20079`

NewInlineResponse20079 instantiates a new InlineResponse20079 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079WithDefaults

`func NewInlineResponse20079WithDefaults() *InlineResponse20079`

NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *InlineResponse20079) GetDefaultSettings() InlineResponse20079DefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *InlineResponse20079) GetDefaultSettingsOk() (*InlineResponse20079DefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *InlineResponse20079) SetDefaultSettings(v InlineResponse20079DefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *InlineResponse20079) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse20079) GetOverrides() []InlineResponse20079Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse20079) GetOverridesOk() (*[]InlineResponse20079Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse20079) SetOverrides(v []InlineResponse20079Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse20079) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


