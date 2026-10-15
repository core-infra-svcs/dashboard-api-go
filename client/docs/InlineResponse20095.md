# InlineResponse20095

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Route ID | [optional] 
**IpVersion** | Pointer to **int32** | IP protocol version | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Enabled** | Pointer to **bool** | Whether the route is enabled or not | [optional] 
**Name** | Pointer to **string** | Name of the route | [optional] 
**Subnet** | Pointer to **string** | Subnet of the route | [optional] 
**GatewayIp** | Pointer to **string** | Gateway IP address (next hop) | [optional] 
**FixedIpAssignments** | Pointer to [**map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments**](NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments.md) | Fixed DHCP IP assignments on the route | [optional] 
**ReservedIpRanges** | Pointer to [**[]NetworksNetworkIdApplianceStaticRoutesReservedIpRanges**](NetworksNetworkIdApplianceStaticRoutesReservedIpRanges.md) | DHCP reserved IP ranges | [optional] 
**GatewayVlanId** | Pointer to **int32** | Gateway VLAN ID | [optional] 
**Vrf** | Pointer to [**NetworksNetworkIdApplianceStaticRoutesVrf**](NetworksNetworkIdApplianceStaticRoutesVrf.md) |  | [optional] 

## Methods

### NewInlineResponse20095

`func NewInlineResponse20095() *InlineResponse20095`

NewInlineResponse20095 instantiates a new InlineResponse20095 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095WithDefaults

`func NewInlineResponse20095WithDefaults() *InlineResponse20095`

NewInlineResponse20095WithDefaults instantiates a new InlineResponse20095 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20095) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20095) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20095) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20095) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse20095) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse20095) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse20095) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse20095) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse20095) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20095) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20095) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20095) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20095) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20095) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20095) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20095) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20095) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20095) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20095) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20095) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20095) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20095) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20095) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20095) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGatewayIp

`func (o *InlineResponse20095) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *InlineResponse20095) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *InlineResponse20095) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *InlineResponse20095) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineResponse20095) GetFixedIpAssignments() map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineResponse20095) GetFixedIpAssignmentsOk() (*map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineResponse20095) SetFixedIpAssignments(v map[string]NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineResponse20095) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineResponse20095) GetReservedIpRanges() []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineResponse20095) GetReservedIpRangesOk() (*[]NetworksNetworkIdApplianceStaticRoutesReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineResponse20095) SetReservedIpRanges(v []NetworksNetworkIdApplianceStaticRoutesReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineResponse20095) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetGatewayVlanId

`func (o *InlineResponse20095) GetGatewayVlanId() int32`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *InlineResponse20095) GetGatewayVlanIdOk() (*int32, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *InlineResponse20095) SetGatewayVlanId(v int32)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *InlineResponse20095) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20095) GetVrf() NetworksNetworkIdApplianceStaticRoutesVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20095) GetVrfOk() (*NetworksNetworkIdApplianceStaticRoutesVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20095) SetVrf(v NetworksNetworkIdApplianceStaticRoutesVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20095) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


