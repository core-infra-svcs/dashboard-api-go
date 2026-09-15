# DevicesSerialWirelessRadioOverridesRfProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The ID of an RF profile to assign to the device. If this field is omitted, the device&#39;s existing RF profile assignment will remain unchanged. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power). | [optional] 

## Methods

### NewDevicesSerialWirelessRadioOverridesRfProfile

`func NewDevicesSerialWirelessRadioOverridesRfProfile() *DevicesSerialWirelessRadioOverridesRfProfile`

NewDevicesSerialWirelessRadioOverridesRfProfile instantiates a new DevicesSerialWirelessRadioOverridesRfProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialWirelessRadioOverridesRfProfileWithDefaults

`func NewDevicesSerialWirelessRadioOverridesRfProfileWithDefaults() *DevicesSerialWirelessRadioOverridesRfProfile`

NewDevicesSerialWirelessRadioOverridesRfProfileWithDefaults instantiates a new DevicesSerialWirelessRadioOverridesRfProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DevicesSerialWirelessRadioOverridesRfProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevicesSerialWirelessRadioOverridesRfProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevicesSerialWirelessRadioOverridesRfProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevicesSerialWirelessRadioOverridesRfProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DevicesSerialWirelessRadioOverridesRfProfile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DevicesSerialWirelessRadioOverridesRfProfile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


