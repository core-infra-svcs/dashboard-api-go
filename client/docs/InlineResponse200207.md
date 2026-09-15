# InlineResponse200207

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the switch stack | [optional] 
**Name** | Pointer to **string** | Name of the switch stack | [optional] 
**Serials** | Pointer to **[]string** | Serial numbers of the switches in the switch stack. Minimum 2 and maximum 8 serials. | [optional] 
**IsMonitorOnly** | Pointer to **bool** | Whether the switch stack contains monitor-only switches. | [optional] 
**VirtualMac** | Pointer to **string** | The virtual MAC address of the switch stack | [optional] 
**Members** | Pointer to [**[]NetworksNetworkIdSwitchStacksMembers**](NetworksNetworkIdSwitchStacksMembers.md) | Members of the switch stack. Minimum 2 and maximum 8 members. | [optional] 

## Methods

### NewInlineResponse200207

`func NewInlineResponse200207() *InlineResponse200207`

NewInlineResponse200207 instantiates a new InlineResponse200207 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200207WithDefaults

`func NewInlineResponse200207WithDefaults() *InlineResponse200207`

NewInlineResponse200207WithDefaults instantiates a new InlineResponse200207 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200207) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200207) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200207) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200207) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200207) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200207) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200207) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200207) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200207) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200207) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200207) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200207) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetIsMonitorOnly

`func (o *InlineResponse200207) GetIsMonitorOnly() bool`

GetIsMonitorOnly returns the IsMonitorOnly field if non-nil, zero value otherwise.

### GetIsMonitorOnlyOk

`func (o *InlineResponse200207) GetIsMonitorOnlyOk() (*bool, bool)`

GetIsMonitorOnlyOk returns a tuple with the IsMonitorOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitorOnly

`func (o *InlineResponse200207) SetIsMonitorOnly(v bool)`

SetIsMonitorOnly sets IsMonitorOnly field to given value.

### HasIsMonitorOnly

`func (o *InlineResponse200207) HasIsMonitorOnly() bool`

HasIsMonitorOnly returns a boolean if a field has been set.

### GetVirtualMac

`func (o *InlineResponse200207) GetVirtualMac() string`

GetVirtualMac returns the VirtualMac field if non-nil, zero value otherwise.

### GetVirtualMacOk

`func (o *InlineResponse200207) GetVirtualMacOk() (*string, bool)`

GetVirtualMacOk returns a tuple with the VirtualMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMac

`func (o *InlineResponse200207) SetVirtualMac(v string)`

SetVirtualMac sets VirtualMac field to given value.

### HasVirtualMac

`func (o *InlineResponse200207) HasVirtualMac() bool`

HasVirtualMac returns a boolean if a field has been set.

### GetMembers

`func (o *InlineResponse200207) GetMembers() []NetworksNetworkIdSwitchStacksMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineResponse200207) GetMembersOk() (*[]NetworksNetworkIdSwitchStacksMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineResponse200207) SetMembers(v []NetworksNetworkIdSwitchStacksMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineResponse200207) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


