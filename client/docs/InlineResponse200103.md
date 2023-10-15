# InlineResponse200103

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | SSID Number | [optional] 
**Vlan** | Pointer to **int32** | LAN | [optional] 
**ClientMac** | Pointer to **string** | Client Mac | [optional] 
**Serial** | Pointer to **string** | Serial Number | [optional] 
**FailureStep** | Pointer to **string** | The failed onboarding step. One of: assoc, auth, dhcp, dns. | [optional] 
**Type** | Pointer to **string** | The failure type in the onboarding step | [optional] 
**Ts** | Pointer to **time.Time** | The timestamp when the client mac failed | [optional] 

## Methods

### NewInlineResponse200103

`func NewInlineResponse200103() *InlineResponse200103`

NewInlineResponse200103 instantiates a new InlineResponse200103 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103WithDefaults

`func NewInlineResponse200103WithDefaults() *InlineResponse200103`

NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse200103) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200103) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200103) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200103) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200103) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200103) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200103) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200103) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetClientMac

`func (o *InlineResponse200103) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *InlineResponse200103) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *InlineResponse200103) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *InlineResponse200103) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200103) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200103) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200103) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200103) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetFailureStep

`func (o *InlineResponse200103) GetFailureStep() string`

GetFailureStep returns the FailureStep field if non-nil, zero value otherwise.

### GetFailureStepOk

`func (o *InlineResponse200103) GetFailureStepOk() (*string, bool)`

GetFailureStepOk returns a tuple with the FailureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureStep

`func (o *InlineResponse200103) SetFailureStep(v string)`

SetFailureStep sets FailureStep field to given value.

### HasFailureStep

`func (o *InlineResponse200103) HasFailureStep() bool`

HasFailureStep returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200103) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200103) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200103) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200103) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse200103) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200103) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200103) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200103) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


