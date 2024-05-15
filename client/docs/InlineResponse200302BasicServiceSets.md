# InlineResponse200302BasicServiceSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bssid** | Pointer to **string** | Unique identifier for wireless access point. | [optional] 
**Ssid** | Pointer to [**InlineResponse200302Ssid**](InlineResponse200302Ssid.md) |  | [optional] 
**Radio** | Pointer to [**InlineResponse200302Radio**](InlineResponse200302Radio.md) |  | [optional] 

## Methods

### NewInlineResponse200302BasicServiceSets

`func NewInlineResponse200302BasicServiceSets() *InlineResponse200302BasicServiceSets`

NewInlineResponse200302BasicServiceSets instantiates a new InlineResponse200302BasicServiceSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302BasicServiceSetsWithDefaults

`func NewInlineResponse200302BasicServiceSetsWithDefaults() *InlineResponse200302BasicServiceSets`

NewInlineResponse200302BasicServiceSetsWithDefaults instantiates a new InlineResponse200302BasicServiceSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBssid

`func (o *InlineResponse200302BasicServiceSets) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *InlineResponse200302BasicServiceSets) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *InlineResponse200302BasicServiceSets) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *InlineResponse200302BasicServiceSets) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse200302BasicServiceSets) GetSsid() InlineResponse200302Ssid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200302BasicServiceSets) GetSsidOk() (*InlineResponse200302Ssid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200302BasicServiceSets) SetSsid(v InlineResponse200302Ssid)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200302BasicServiceSets) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetRadio

`func (o *InlineResponse200302BasicServiceSets) GetRadio() InlineResponse200302Radio`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *InlineResponse200302BasicServiceSets) GetRadioOk() (*InlineResponse200302Radio, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *InlineResponse200302BasicServiceSets) SetRadio(v InlineResponse200302Radio)`

SetRadio sets Radio field to given value.

### HasRadio

`func (o *InlineResponse200302BasicServiceSets) HasRadio() bool`

HasRadio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


