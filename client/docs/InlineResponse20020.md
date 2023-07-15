# InlineResponse20020

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | The subnet of the single LAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the single LAN | [optional] 
**MandatoryDhcp** | Pointer to [**InlineResponse20020MandatoryDhcp**](InlineResponse20020MandatoryDhcp.md) |  | [optional] 
**Ipv6** | Pointer to [**InlineResponse20020Ipv6**](InlineResponse20020Ipv6.md) |  | [optional] 

## Methods

### NewInlineResponse20020

`func NewInlineResponse20020() *InlineResponse20020`

NewInlineResponse20020 instantiates a new InlineResponse20020 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20020WithDefaults

`func NewInlineResponse20020WithDefaults() *InlineResponse20020`

NewInlineResponse20020WithDefaults instantiates a new InlineResponse20020 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *InlineResponse20020) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20020) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20020) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20020) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineResponse20020) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineResponse20020) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineResponse20020) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineResponse20020) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *InlineResponse20020) GetMandatoryDhcp() InlineResponse20020MandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *InlineResponse20020) GetMandatoryDhcpOk() (*InlineResponse20020MandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *InlineResponse20020) SetMandatoryDhcp(v InlineResponse20020MandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *InlineResponse20020) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20020) GetIpv6() InlineResponse20020Ipv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20020) GetIpv6Ok() (*InlineResponse20020Ipv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20020) SetIpv6(v InlineResponse20020Ipv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20020) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


