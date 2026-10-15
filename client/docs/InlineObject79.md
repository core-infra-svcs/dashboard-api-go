# InlineObject79

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the route | 
**Subnet** | **string** | Subnet of the route | 
**GatewayIp** | **string** | Gateway IP address (next hop) | 
**GatewayVlanId** | Pointer to **int32** | Gateway VLAN ID | [optional] 
**Enabled** | Pointer to **bool** | Enable/disable the static route | [optional] 
**Vrf** | Pointer to [**NetworksNetworkIdApplianceStaticRoutesVrf1**](NetworksNetworkIdApplianceStaticRoutesVrf1.md) |  | [optional] 

## Methods

### NewInlineObject79

`func NewInlineObject79(name string, subnet string, gatewayIp string, ) *InlineObject79`

NewInlineObject79 instantiates a new InlineObject79 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject79WithDefaults

`func NewInlineObject79WithDefaults() *InlineObject79`

NewInlineObject79WithDefaults instantiates a new InlineObject79 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject79) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject79) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject79) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *InlineObject79) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject79) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject79) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetGatewayIp

`func (o *InlineObject79) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *InlineObject79) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *InlineObject79) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.


### GetGatewayVlanId

`func (o *InlineObject79) GetGatewayVlanId() int32`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *InlineObject79) GetGatewayVlanIdOk() (*int32, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *InlineObject79) SetGatewayVlanId(v int32)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *InlineObject79) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject79) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject79) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject79) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject79) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject79) GetVrf() NetworksNetworkIdApplianceStaticRoutesVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject79) GetVrfOk() (*NetworksNetworkIdApplianceStaticRoutesVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject79) SetVrf(v NetworksNetworkIdApplianceStaticRoutesVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject79) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


