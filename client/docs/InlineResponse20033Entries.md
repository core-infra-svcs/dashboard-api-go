# InlineResponse20033Entries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address of the ARP table entry | [optional] 
**Mac** | Pointer to **string** | The MAC address of the ARP table entry | [optional] 
**VlanId** | Pointer to **NullableInt32** | The VLAN ID of the ARP table entry. This is null for Meraki wireless devices, or for ARP table entries that do not belong to a VLAN. | [optional] 
**Interface** | Pointer to **NullableString** | The interface name of the ARP table entry, such as Vlan1, Port-channel2, or GigabitEthernet1/0/1. This is null for non-Catalyst switches as well as Meraki wireless devices. | [optional] 
**LastUpdatedAt** | Pointer to **NullableTime** | Time of the last update of the ARP table entry. This is null for Meraki wireless devices. | [optional] 

## Methods

### NewInlineResponse20033Entries

`func NewInlineResponse20033Entries() *InlineResponse20033Entries`

NewInlineResponse20033Entries instantiates a new InlineResponse20033Entries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033EntriesWithDefaults

`func NewInlineResponse20033EntriesWithDefaults() *InlineResponse20033Entries`

NewInlineResponse20033EntriesWithDefaults instantiates a new InlineResponse20033Entries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse20033Entries) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20033Entries) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20033Entries) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20033Entries) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse20033Entries) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20033Entries) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20033Entries) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20033Entries) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse20033Entries) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse20033Entries) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse20033Entries) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse20033Entries) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *InlineResponse20033Entries) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *InlineResponse20033Entries) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetInterface

`func (o *InlineResponse20033Entries) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20033Entries) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20033Entries) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20033Entries) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *InlineResponse20033Entries) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *InlineResponse20033Entries) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetLastUpdatedAt

`func (o *InlineResponse20033Entries) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse20033Entries) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse20033Entries) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse20033Entries) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### SetLastUpdatedAtNil

`func (o *InlineResponse20033Entries) SetLastUpdatedAtNil(b bool)`

 SetLastUpdatedAtNil sets the value for LastUpdatedAt to be an explicit nil

### UnsetLastUpdatedAt
`func (o *InlineResponse20033Entries) UnsetLastUpdatedAt()`

UnsetLastUpdatedAt ensures that no value is present for LastUpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


