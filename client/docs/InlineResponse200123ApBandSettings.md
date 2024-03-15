# InlineResponse200123ApBandSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39;, &#39;5ghz&#39;, &#39;6ghz&#39; or &#39;multi&#39;. Defaults to dual. | [optional] 
**Bands** | Pointer to [**InlineResponse200123ApBandSettingsBands**](InlineResponse200123ApBandSettingsBands.md) |  | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band. Can be either true or false. Defaults to true. | [optional] 

## Methods

### NewInlineResponse200123ApBandSettings

`func NewInlineResponse200123ApBandSettings() *InlineResponse200123ApBandSettings`

NewInlineResponse200123ApBandSettings instantiates a new InlineResponse200123ApBandSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200123ApBandSettingsWithDefaults

`func NewInlineResponse200123ApBandSettingsWithDefaults() *InlineResponse200123ApBandSettings`

NewInlineResponse200123ApBandSettingsWithDefaults instantiates a new InlineResponse200123ApBandSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandOperationMode

`func (o *InlineResponse200123ApBandSettings) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *InlineResponse200123ApBandSettings) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *InlineResponse200123ApBandSettings) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *InlineResponse200123ApBandSettings) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBands

`func (o *InlineResponse200123ApBandSettings) GetBands() InlineResponse200123ApBandSettingsBands`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *InlineResponse200123ApBandSettings) GetBandsOk() (*InlineResponse200123ApBandSettingsBands, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *InlineResponse200123ApBandSettings) SetBands(v InlineResponse200123ApBandSettingsBands)`

SetBands sets Bands field to given value.

### HasBands

`func (o *InlineResponse200123ApBandSettings) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *InlineResponse200123ApBandSettings) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *InlineResponse200123ApBandSettings) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *InlineResponse200123ApBandSettings) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *InlineResponse200123ApBandSettings) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


