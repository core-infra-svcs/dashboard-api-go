# InlineResponse200197Overrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Switches** | Pointer to **[]string** | List of switch serials. Applicable only for switch network. | [optional] 
**SwitchProfiles** | Pointer to **[]string** | List of switch template IDs. Applicable only for template network. | [optional] 
**MtuSize** | **int32** | MTU size for the switches or switch templates. | 

## Methods

### NewInlineResponse200197Overrides

`func NewInlineResponse200197Overrides(mtuSize int32, ) *InlineResponse200197Overrides`

NewInlineResponse200197Overrides instantiates a new InlineResponse200197Overrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200197OverridesWithDefaults

`func NewInlineResponse200197OverridesWithDefaults() *InlineResponse200197Overrides`

NewInlineResponse200197OverridesWithDefaults instantiates a new InlineResponse200197Overrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitches

`func (o *InlineResponse200197Overrides) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *InlineResponse200197Overrides) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *InlineResponse200197Overrides) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *InlineResponse200197Overrides) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *InlineResponse200197Overrides) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *InlineResponse200197Overrides) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *InlineResponse200197Overrides) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *InlineResponse200197Overrides) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetMtuSize

`func (o *InlineResponse200197Overrides) GetMtuSize() int32`

GetMtuSize returns the MtuSize field if non-nil, zero value otherwise.

### GetMtuSizeOk

`func (o *InlineResponse200197Overrides) GetMtuSizeOk() (*int32, bool)`

GetMtuSizeOk returns a tuple with the MtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuSize

`func (o *InlineResponse200197Overrides) SetMtuSize(v int32)`

SetMtuSize sets MtuSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


