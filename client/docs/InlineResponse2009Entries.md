# InlineResponse2009Entries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address of the ARP table entry | [optional] 
**Mac** | Pointer to **string** | The MAC address of the ARP table entry | [optional] 
**VlanId** | Pointer to **int32** | The VLAN ID of the ARP table entry | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Time of the last update of the ARP table entry | [optional] 

## Methods

### NewInlineResponse2009Entries

`func NewInlineResponse2009Entries() *InlineResponse2009Entries`

NewInlineResponse2009Entries instantiates a new InlineResponse2009Entries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2009EntriesWithDefaults

`func NewInlineResponse2009EntriesWithDefaults() *InlineResponse2009Entries`

NewInlineResponse2009EntriesWithDefaults instantiates a new InlineResponse2009Entries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse2009Entries) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse2009Entries) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse2009Entries) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse2009Entries) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse2009Entries) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse2009Entries) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse2009Entries) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse2009Entries) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse2009Entries) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse2009Entries) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse2009Entries) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse2009Entries) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse2009Entries) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse2009Entries) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse2009Entries) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse2009Entries) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


