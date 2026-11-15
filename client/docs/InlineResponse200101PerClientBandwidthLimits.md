# InlineResponse200101PerClientBandwidthLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** |     How bandwidth limits are applied by your rule. Can be one of &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;.     Secure Router networks support only the &#39;ignore&#39; setting.  | [optional] 
**BandwidthLimits** | Pointer to [**InlineResponse200101PerClientBandwidthLimitsBandwidthLimits**](InlineResponse200101PerClientBandwidthLimitsBandwidthLimits.md) |  | [optional] 

## Methods

### NewInlineResponse200101PerClientBandwidthLimits

`func NewInlineResponse200101PerClientBandwidthLimits() *InlineResponse200101PerClientBandwidthLimits`

NewInlineResponse200101PerClientBandwidthLimits instantiates a new InlineResponse200101PerClientBandwidthLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101PerClientBandwidthLimitsWithDefaults

`func NewInlineResponse200101PerClientBandwidthLimitsWithDefaults() *InlineResponse200101PerClientBandwidthLimits`

NewInlineResponse200101PerClientBandwidthLimitsWithDefaults instantiates a new InlineResponse200101PerClientBandwidthLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *InlineResponse200101PerClientBandwidthLimits) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *InlineResponse200101PerClientBandwidthLimits) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *InlineResponse200101PerClientBandwidthLimits) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *InlineResponse200101PerClientBandwidthLimits) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *InlineResponse200101PerClientBandwidthLimits) GetBandwidthLimits() InlineResponse200101PerClientBandwidthLimitsBandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *InlineResponse200101PerClientBandwidthLimits) GetBandwidthLimitsOk() (*InlineResponse200101PerClientBandwidthLimitsBandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *InlineResponse200101PerClientBandwidthLimits) SetBandwidthLimits(v InlineResponse200101PerClientBandwidthLimitsBandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *InlineResponse200101PerClientBandwidthLimits) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


