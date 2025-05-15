# InlineResponse200163

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**InlineResponse200163DefaultSettings**](InlineResponse200163DefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]InlineResponse200163Overrides**](InlineResponse200163Overrides.md) | Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings. | [optional] 

## Methods

### NewInlineResponse200163

`func NewInlineResponse200163() *InlineResponse200163`

NewInlineResponse200163 instantiates a new InlineResponse200163 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200163WithDefaults

`func NewInlineResponse200163WithDefaults() *InlineResponse200163`

NewInlineResponse200163WithDefaults instantiates a new InlineResponse200163 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *InlineResponse200163) GetDefaultSettings() InlineResponse200163DefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *InlineResponse200163) GetDefaultSettingsOk() (*InlineResponse200163DefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *InlineResponse200163) SetDefaultSettings(v InlineResponse200163DefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *InlineResponse200163) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse200163) GetOverrides() []InlineResponse200163Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse200163) GetOverridesOk() (*[]InlineResponse200163Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse200163) SetOverrides(v []InlineResponse200163Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse200163) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


