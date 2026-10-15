# InlineObject80

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the route | [optional] 
**Subnet** | Pointer to **string** | Subnet of the route | [optional] 
**GatewayIp** | Pointer to **string** | Gateway IP address (next hop) | [optional] 
**GatewayVlanId** | Pointer to **NullableInt32** | Gateway VLAN ID | [optional] 
**Enabled** | Pointer to **bool** | Whether the route should be enabled or not | [optional] 
**FixedIpAssignments** | Pointer to [**map[string]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdFixedIpAssignments**](NetworksNetworkIdApplianceStaticRoutesStaticRouteIdFixedIpAssignments.md) | Fixed DHCP IP assignments on the route | [optional] 
**ReservedIpRanges** | Pointer to [**[]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges**](NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges.md) | DHCP reserved IP ranges | [optional] 
**Vrf** | Pointer to [**NetworksNetworkIdApplianceStaticRoutesVrf1**](NetworksNetworkIdApplianceStaticRoutesVrf1.md) |  | [optional] 

## Methods

### NewInlineObject80

`func NewInlineObject80() *InlineObject80`

NewInlineObject80 instantiates a new InlineObject80 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject80WithDefaults

`func NewInlineObject80WithDefaults() *InlineObject80`

NewInlineObject80WithDefaults instantiates a new InlineObject80 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject80) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject80) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject80) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject80) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject80) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject80) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject80) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject80) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGatewayIp

`func (o *InlineObject80) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *InlineObject80) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *InlineObject80) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *InlineObject80) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayVlanId

`func (o *InlineObject80) GetGatewayVlanId() int32`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *InlineObject80) GetGatewayVlanIdOk() (*int32, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *InlineObject80) SetGatewayVlanId(v int32)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *InlineObject80) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.

### SetGatewayVlanIdNil

`func (o *InlineObject80) SetGatewayVlanIdNil(b bool)`

 SetGatewayVlanIdNil sets the value for GatewayVlanId to be an explicit nil

### UnsetGatewayVlanId
`func (o *InlineObject80) UnsetGatewayVlanId()`

UnsetGatewayVlanId ensures that no value is present for GatewayVlanId, not even an explicit nil
### GetEnabled

`func (o *InlineObject80) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject80) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject80) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject80) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineObject80) GetFixedIpAssignments() map[string]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdFixedIpAssignments`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineObject80) GetFixedIpAssignmentsOk() (*map[string]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdFixedIpAssignments, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineObject80) SetFixedIpAssignments(v map[string]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdFixedIpAssignments)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineObject80) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineObject80) GetReservedIpRanges() []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineObject80) GetReservedIpRangesOk() (*[]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineObject80) SetReservedIpRanges(v []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineObject80) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject80) GetVrf() NetworksNetworkIdApplianceStaticRoutesVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject80) GetVrfOk() (*NetworksNetworkIdApplianceStaticRoutesVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject80) SetVrf(v NetworksNetworkIdApplianceStaticRoutesVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject80) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


