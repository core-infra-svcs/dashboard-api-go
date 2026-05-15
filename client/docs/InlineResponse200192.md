# InlineResponse200192

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**InlineResponse200192DefaultSettings**](InlineResponse200192DefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]InlineResponse200192Overrides**](InlineResponse200192Overrides.md) | Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings. | [optional] 

## Methods

### NewInlineResponse200192

`func NewInlineResponse200192() *InlineResponse200192`

NewInlineResponse200192 instantiates a new InlineResponse200192 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200192WithDefaults

`func NewInlineResponse200192WithDefaults() *InlineResponse200192`

NewInlineResponse200192WithDefaults instantiates a new InlineResponse200192 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *InlineResponse200192) GetDefaultSettings() InlineResponse200192DefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *InlineResponse200192) GetDefaultSettingsOk() (*InlineResponse200192DefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *InlineResponse200192) SetDefaultSettings(v InlineResponse200192DefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *InlineResponse200192) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse200192) GetOverrides() []InlineResponse200192Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse200192) GetOverridesOk() (*[]InlineResponse200192Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse200192) SetOverrides(v []InlineResponse200192Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse200192) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


