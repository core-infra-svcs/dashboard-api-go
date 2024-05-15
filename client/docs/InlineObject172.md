# InlineObject172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable alternate management interface | [optional] 
**VlanId** | Pointer to **int32** | Alternate management interface VLAN, must be between 1 and 4094 | [optional] 
**Protocols** | Pointer to **[]string** | Can be one or more of the following values: &#39;radius&#39;, &#39;snmp&#39;, &#39;syslog&#39; or &#39;ldap&#39; | [optional] 
**AccessPoints** | Pointer to [**[]NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints**](NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints.md) | Array of access point serial number and IP assignment. Note: accessPoints IP assignment is not applicable for template networks, in other words, do not put &#39;accessPoints&#39; in the body when updating template networks. Also, an empty &#39;accessPoints&#39; array will remove all previous static IP assignments | [optional] 

## Methods

### NewInlineObject172

`func NewInlineObject172() *InlineObject172`

NewInlineObject172 instantiates a new InlineObject172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject172WithDefaults

`func NewInlineObject172WithDefaults() *InlineObject172`

NewInlineObject172WithDefaults instantiates a new InlineObject172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject172) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject172) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject172) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject172) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineObject172) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineObject172) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineObject172) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineObject172) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetProtocols

`func (o *InlineObject172) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *InlineObject172) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *InlineObject172) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *InlineObject172) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetAccessPoints

`func (o *InlineObject172) GetAccessPoints() []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints`

GetAccessPoints returns the AccessPoints field if non-nil, zero value otherwise.

### GetAccessPointsOk

`func (o *InlineObject172) GetAccessPointsOk() (*[]NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints, bool)`

GetAccessPointsOk returns a tuple with the AccessPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoints

`func (o *InlineObject172) SetAccessPoints(v []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints)`

SetAccessPoints sets AccessPoints field to given value.

### HasAccessPoints

`func (o *InlineObject172) HasAccessPoints() bool`

HasAccessPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


