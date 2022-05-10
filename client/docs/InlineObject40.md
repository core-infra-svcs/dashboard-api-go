# InlineObject40

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the static route | [optional] 
**Subnet** | Pointer to **string** | The subnet of the static route | [optional] 
**GatewayIp** | Pointer to **string** | The gateway IP (next hop) of the static route | [optional] 
**GatewayVlanId** | Pointer to **string** | The gateway IP (next hop) VLAN ID of the static route | [optional] 
**Enabled** | Pointer to **bool** | The enabled state of the static route | [optional] 
**FixedIpAssignments** | Pointer to **map[string]interface{}** | The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \&quot;ip\&quot; and \&quot;name\&quot; string fields. See the sample request/response for more details. | [optional] 
**ReservedIpRanges** | Pointer to [**[]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges**](NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges.md) | The DHCP reserved IP ranges on the static route | [optional] 

## Methods

### NewInlineObject40

`func NewInlineObject40() *InlineObject40`

NewInlineObject40 instantiates a new InlineObject40 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject40WithDefaults

`func NewInlineObject40WithDefaults() *InlineObject40`

NewInlineObject40WithDefaults instantiates a new InlineObject40 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject40) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject40) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject40) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject40) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject40) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject40) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject40) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject40) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGatewayIp

`func (o *InlineObject40) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *InlineObject40) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *InlineObject40) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *InlineObject40) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayVlanId

`func (o *InlineObject40) GetGatewayVlanId() string`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *InlineObject40) GetGatewayVlanIdOk() (*string, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *InlineObject40) SetGatewayVlanId(v string)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *InlineObject40) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject40) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject40) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject40) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject40) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineObject40) GetFixedIpAssignments() map[string]interface{}`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineObject40) GetFixedIpAssignmentsOk() (*map[string]interface{}, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineObject40) SetFixedIpAssignments(v map[string]interface{})`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineObject40) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineObject40) GetReservedIpRanges() []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineObject40) GetReservedIpRangesOk() (*[]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineObject40) SetReservedIpRanges(v []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineObject40) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


