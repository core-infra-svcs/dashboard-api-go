# InlineResponse20084

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**InlineResponse20084DefaultSettings**](InlineResponse20084DefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]InlineResponse20084Overrides**](InlineResponse20084Overrides.md) | Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings. | [optional] 

## Methods

### NewInlineResponse20084

`func NewInlineResponse20084() *InlineResponse20084`

NewInlineResponse20084 instantiates a new InlineResponse20084 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20084WithDefaults

`func NewInlineResponse20084WithDefaults() *InlineResponse20084`

NewInlineResponse20084WithDefaults instantiates a new InlineResponse20084 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *InlineResponse20084) GetDefaultSettings() InlineResponse20084DefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *InlineResponse20084) GetDefaultSettingsOk() (*InlineResponse20084DefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *InlineResponse20084) SetDefaultSettings(v InlineResponse20084DefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *InlineResponse20084) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse20084) GetOverrides() []InlineResponse20084Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse20084) GetOverridesOk() (*[]InlineResponse20084Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse20084) SetOverrides(v []InlineResponse20084Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse20084) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


