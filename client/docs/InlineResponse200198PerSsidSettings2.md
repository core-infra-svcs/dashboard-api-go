# InlineResponse200198PerSsidSettings2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of SSID | [optional] 
**MinBitrate** | Pointer to **int32** | Sets min bitrate (Mbps) of this SSID. Can be one of &#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. | [optional] 
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39;, &#39;5ghz&#39;, &#39;6ghz&#39; or &#39;multi&#39;. | [optional] 
**Bands** | Pointer to [**InlineResponse200198ApBandSettingsBands**](InlineResponse200198ApBandSettingsBands.md) |  | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false. | [optional] 

## Methods

### NewInlineResponse200198PerSsidSettings2

`func NewInlineResponse200198PerSsidSettings2() *InlineResponse200198PerSsidSettings2`

NewInlineResponse200198PerSsidSettings2 instantiates a new InlineResponse200198PerSsidSettings2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200198PerSsidSettings2WithDefaults

`func NewInlineResponse200198PerSsidSettings2WithDefaults() *InlineResponse200198PerSsidSettings2`

NewInlineResponse200198PerSsidSettings2WithDefaults instantiates a new InlineResponse200198PerSsidSettings2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200198PerSsidSettings2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200198PerSsidSettings2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200198PerSsidSettings2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200198PerSsidSettings2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse200198PerSsidSettings2) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse200198PerSsidSettings2) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse200198PerSsidSettings2) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineResponse200198PerSsidSettings2) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetBandOperationMode

`func (o *InlineResponse200198PerSsidSettings2) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *InlineResponse200198PerSsidSettings2) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *InlineResponse200198PerSsidSettings2) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *InlineResponse200198PerSsidSettings2) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBands

`func (o *InlineResponse200198PerSsidSettings2) GetBands() InlineResponse200198ApBandSettingsBands`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *InlineResponse200198PerSsidSettings2) GetBandsOk() (*InlineResponse200198ApBandSettingsBands, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *InlineResponse200198PerSsidSettings2) SetBands(v InlineResponse200198ApBandSettingsBands)`

SetBands sets Bands field to given value.

### HasBands

`func (o *InlineResponse200198PerSsidSettings2) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *InlineResponse200198PerSsidSettings2) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *InlineResponse200198PerSsidSettings2) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *InlineResponse200198PerSsidSettings2) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *InlineResponse200198PerSsidSettings2) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


