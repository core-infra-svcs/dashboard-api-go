# InlineObject66

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the route | 
**Subnet** | **string** | Subnet of the route | 
**GatewayIp** | **string** | Gateway IP address (next hop) | 
**GatewayVlanId** | Pointer to **string** | Gateway VLAN ID | [optional] 

## Methods

### NewInlineObject66

`func NewInlineObject66(name string, subnet string, gatewayIp string, ) *InlineObject66`

NewInlineObject66 instantiates a new InlineObject66 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject66WithDefaults

`func NewInlineObject66WithDefaults() *InlineObject66`

NewInlineObject66WithDefaults instantiates a new InlineObject66 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject66) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject66) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject66) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *InlineObject66) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject66) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject66) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetGatewayIp

`func (o *InlineObject66) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *InlineObject66) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *InlineObject66) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.


### GetGatewayVlanId

`func (o *InlineObject66) GetGatewayVlanId() string`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *InlineObject66) GetGatewayVlanIdOk() (*string, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *InlineObject66) SetGatewayVlanId(v string)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *InlineObject66) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


