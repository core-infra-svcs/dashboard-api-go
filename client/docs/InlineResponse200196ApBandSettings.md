# InlineResponse200196ApBandSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39;, &#39;5ghz&#39;, &#39;6ghz&#39; or &#39;multi&#39;. Defaults to dual. | [optional] 
**Bands** | Pointer to [**InlineResponse200196ApBandSettingsBands**](InlineResponse200196ApBandSettingsBands.md) |  | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band. Can be either true or false. Defaults to true. | [optional] 

## Methods

### NewInlineResponse200196ApBandSettings

`func NewInlineResponse200196ApBandSettings() *InlineResponse200196ApBandSettings`

NewInlineResponse200196ApBandSettings instantiates a new InlineResponse200196ApBandSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200196ApBandSettingsWithDefaults

`func NewInlineResponse200196ApBandSettingsWithDefaults() *InlineResponse200196ApBandSettings`

NewInlineResponse200196ApBandSettingsWithDefaults instantiates a new InlineResponse200196ApBandSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *InlineResponse200196ApBandSettings) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *InlineResponse200196ApBandSettings) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *InlineResponse200196ApBandSettings) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *InlineResponse200196ApBandSettings) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBands

`func (o *InlineResponse200196ApBandSettings) GetBands() InlineResponse200196ApBandSettingsBands`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *InlineResponse200196ApBandSettings) GetBandsOk() (*InlineResponse200196ApBandSettingsBands, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *InlineResponse200196ApBandSettings) SetBands(v InlineResponse200196ApBandSettingsBands)`

SetBands sets Bands field to given value.

### HasBands

`func (o *InlineResponse200196ApBandSettings) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *InlineResponse200196ApBandSettings) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *InlineResponse200196ApBandSettings) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *InlineResponse200196ApBandSettings) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *InlineResponse200196ApBandSettings) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


