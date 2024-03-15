# InlineResponse2008

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the client | [optional] 
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**Description** | Pointer to **string** | Short description of the client | [optional] 
**MdnsName** | Pointer to **string** | The client&#39;s MDNS name | [optional] 
**DhcpHostname** | Pointer to **string** | The client&#39;s DHCP hostname | [optional] 
**User** | Pointer to **string** | The client user&#39;s name | [optional] 
**Ip** | Pointer to **string** | The IP address of the client | [optional] 
**Vlan** | Pointer to **string** | The client-assigned name of the VLAN the client is connected to | [optional] 
**NamedVlan** | Pointer to **string** | The owner-assigned name of the VLAN the client is connected to | [optional] 
**Switchport** | Pointer to **string** | The name of the switchport with clients on it, if the device is a switch | [optional] 
**AdaptivePolicyGroup** | Pointer to **string** | A description of the adaptive policy group | [optional] 
**Usage** | Pointer to [**DevicesSerialClientsUsage**](DevicesSerialClientsUsage.md) |  | [optional] 

## Methods

### NewInlineResponse2008

`func NewInlineResponse2008() *InlineResponse2008`

NewInlineResponse2008 instantiates a new InlineResponse2008 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008WithDefaults

`func NewInlineResponse2008WithDefaults() *InlineResponse2008`

NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2008) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2008) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2008) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2008) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse2008) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse2008) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse2008) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse2008) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2008) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2008) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2008) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2008) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMdnsName

`func (o *InlineResponse2008) GetMdnsName() string`

GetMdnsName returns the MdnsName field if non-nil, zero value otherwise.

### GetMdnsNameOk

`func (o *InlineResponse2008) GetMdnsNameOk() (*string, bool)`

GetMdnsNameOk returns a tuple with the MdnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdnsName

`func (o *InlineResponse2008) SetMdnsName(v string)`

SetMdnsName sets MdnsName field to given value.

### HasMdnsName

`func (o *InlineResponse2008) HasMdnsName() bool`

HasMdnsName returns a boolean if a field has been set.

### GetDhcpHostname

`func (o *InlineResponse2008) GetDhcpHostname() string`

GetDhcpHostname returns the DhcpHostname field if non-nil, zero value otherwise.

### GetDhcpHostnameOk

`func (o *InlineResponse2008) GetDhcpHostnameOk() (*string, bool)`

GetDhcpHostnameOk returns a tuple with the DhcpHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHostname

`func (o *InlineResponse2008) SetDhcpHostname(v string)`

SetDhcpHostname sets DhcpHostname field to given value.

### HasDhcpHostname

`func (o *InlineResponse2008) HasDhcpHostname() bool`

HasDhcpHostname returns a boolean if a field has been set.

### GetUser

`func (o *InlineResponse2008) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse2008) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse2008) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse2008) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse2008) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse2008) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse2008) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse2008) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse2008) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse2008) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse2008) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse2008) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetNamedVlan

`func (o *InlineResponse2008) GetNamedVlan() string`

GetNamedVlan returns the NamedVlan field if non-nil, zero value otherwise.

### GetNamedVlanOk

`func (o *InlineResponse2008) GetNamedVlanOk() (*string, bool)`

GetNamedVlanOk returns a tuple with the NamedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlan

`func (o *InlineResponse2008) SetNamedVlan(v string)`

SetNamedVlan sets NamedVlan field to given value.

### HasNamedVlan

`func (o *InlineResponse2008) HasNamedVlan() bool`

HasNamedVlan returns a boolean if a field has been set.

### GetSwitchport

`func (o *InlineResponse2008) GetSwitchport() string`

GetSwitchport returns the Switchport field if non-nil, zero value otherwise.

### GetSwitchportOk

`func (o *InlineResponse2008) GetSwitchportOk() (*string, bool)`

GetSwitchportOk returns a tuple with the Switchport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchport

`func (o *InlineResponse2008) SetSwitchport(v string)`

SetSwitchport sets Switchport field to given value.

### HasSwitchport

`func (o *InlineResponse2008) HasSwitchport() bool`

HasSwitchport returns a boolean if a field has been set.

### GetAdaptivePolicyGroup

`func (o *InlineResponse2008) GetAdaptivePolicyGroup() string`

GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupOk

`func (o *InlineResponse2008) GetAdaptivePolicyGroupOk() (*string, bool)`

GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroup

`func (o *InlineResponse2008) SetAdaptivePolicyGroup(v string)`

SetAdaptivePolicyGroup sets AdaptivePolicyGroup field to given value.

### HasAdaptivePolicyGroup

`func (o *InlineResponse2008) HasAdaptivePolicyGroup() bool`

HasAdaptivePolicyGroup returns a boolean if a field has been set.

### GetUsage

`func (o *InlineResponse2008) GetUsage() DevicesSerialClientsUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineResponse2008) GetUsageOk() (*DevicesSerialClientsUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineResponse2008) SetUsage(v DevicesSerialClientsUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineResponse2008) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


