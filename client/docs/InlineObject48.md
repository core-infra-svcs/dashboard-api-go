# InlineObject48

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RfProfile** | Pointer to [**DevicesSerialWirelessRadioOverridesRfProfile**](DevicesSerialWirelessRadioOverridesRfProfile.md) |  | [optional] 
**Radios** | Pointer to [**[]DevicesSerialWirelessRadioOverridesRadios**](DevicesSerialWirelessRadioOverridesRadios.md) | Radio overrides. If omitted, existing per-radio override settings remain unchanged. If provided, only the radios included in the array are updated; the array does not replace the device&#39;s full existing set of radio overrides. Read-only response fields for each radio, such as &#39;band&#39;, may be included in the request body and are ignored. | [optional] 

## Methods

### NewInlineObject48

`func NewInlineObject48() *InlineObject48`

NewInlineObject48 instantiates a new InlineObject48 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject48WithDefaults

`func NewInlineObject48WithDefaults() *InlineObject48`

NewInlineObject48WithDefaults instantiates a new InlineObject48 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRfProfile

`func (o *InlineObject48) GetRfProfile() DevicesSerialWirelessRadioOverridesRfProfile`

GetRfProfile returns the RfProfile field if non-nil, zero value otherwise.

### GetRfProfileOk

`func (o *InlineObject48) GetRfProfileOk() (*DevicesSerialWirelessRadioOverridesRfProfile, bool)`

GetRfProfileOk returns a tuple with the RfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfProfile

`func (o *InlineObject48) SetRfProfile(v DevicesSerialWirelessRadioOverridesRfProfile)`

SetRfProfile sets RfProfile field to given value.

### HasRfProfile

`func (o *InlineObject48) HasRfProfile() bool`

HasRfProfile returns a boolean if a field has been set.

### GetRadios

`func (o *InlineObject48) GetRadios() []DevicesSerialWirelessRadioOverridesRadios`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *InlineObject48) GetRadiosOk() (*[]DevicesSerialWirelessRadioOverridesRadios, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *InlineObject48) SetRadios(v []DevicesSerialWirelessRadioOverridesRadios)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *InlineObject48) HasRadios() bool`

HasRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


