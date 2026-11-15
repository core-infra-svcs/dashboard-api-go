# InlineObject50

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RfProfile** | Pointer to [**DevicesSerialWirelessRadioOverridesRfProfile**](DevicesSerialWirelessRadioOverridesRfProfile.md) |  | [optional] 
**Radios** | Pointer to [**[]DevicesSerialWirelessRadioOverridesRadios**](DevicesSerialWirelessRadioOverridesRadios.md) | Radio overrides. If omitted, existing per-radio override settings remain unchanged. If provided, only the radios included in the array are updated; the array does not replace the device&#39;s full existing set of radio overrides. Read-only response fields for each radio, such as &#39;band&#39;, may be included in the request body and are ignored. | [optional] 

## Methods

### NewInlineObject50

`func NewInlineObject50() *InlineObject50`

NewInlineObject50 instantiates a new InlineObject50 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject50WithDefaults

`func NewInlineObject50WithDefaults() *InlineObject50`

NewInlineObject50WithDefaults instantiates a new InlineObject50 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRfProfile

`func (o *InlineObject50) GetRfProfile() DevicesSerialWirelessRadioOverridesRfProfile`

GetRfProfile returns the RfProfile field if non-nil, zero value otherwise.

### GetRfProfileOk

`func (o *InlineObject50) GetRfProfileOk() (*DevicesSerialWirelessRadioOverridesRfProfile, bool)`

GetRfProfileOk returns a tuple with the RfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfProfile

`func (o *InlineObject50) SetRfProfile(v DevicesSerialWirelessRadioOverridesRfProfile)`

SetRfProfile sets RfProfile field to given value.

### HasRfProfile

`func (o *InlineObject50) HasRfProfile() bool`

HasRfProfile returns a boolean if a field has been set.

### GetRadios

`func (o *InlineObject50) GetRadios() []DevicesSerialWirelessRadioOverridesRadios`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *InlineObject50) GetRadiosOk() (*[]DevicesSerialWirelessRadioOverridesRadios, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *InlineObject50) SetRadios(v []DevicesSerialWirelessRadioOverridesRadios)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *InlineObject50) HasRadios() bool`

HasRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


