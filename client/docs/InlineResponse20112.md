# InlineResponse20112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The VLAN ID of the VLAN | [optional] 
**InterfaceId** | Pointer to **string** | The interface ID of the VLAN | [optional] 
**Name** | Pointer to **string** | The name of the VLAN | [optional] 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **string** | The id of the desired group policy to apply to the VLAN | [optional] 
**TemplateVlanType** | Pointer to **string** | Type of subnetting of the VLAN. Applicable only for template network. | [optional] [default to "same"]
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all bound to the template networks. Applicable only for template network. | [optional] 
**MandatoryDhcp** | Pointer to [**NetworksNetworkIdApplianceVlansMandatoryDhcp**](NetworksNetworkIdApplianceVlansMandatoryDhcp.md) |  | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdApplianceVlansIpv6**](NetworksNetworkIdApplianceVlansIpv6.md) |  | [optional] 
**Vrf** | Pointer to [**NetworksNetworkIdApplianceVlansVrf**](NetworksNetworkIdApplianceVlansVrf.md) |  | [optional] 

## Methods

### NewInlineResponse20112

`func NewInlineResponse20112() *InlineResponse20112`

NewInlineResponse20112 instantiates a new InlineResponse20112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20112WithDefaults

`func NewInlineResponse20112WithDefaults() *InlineResponse20112`

NewInlineResponse20112WithDefaults instantiates a new InlineResponse20112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20112) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20112) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20112) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20112) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *InlineResponse20112) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *InlineResponse20112) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *InlineResponse20112) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *InlineResponse20112) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20112) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20112) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20112) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20112) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20112) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20112) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20112) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20112) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineResponse20112) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineResponse20112) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineResponse20112) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineResponse20112) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse20112) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse20112) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse20112) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse20112) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetTemplateVlanType

`func (o *InlineResponse20112) GetTemplateVlanType() string`

GetTemplateVlanType returns the TemplateVlanType field if non-nil, zero value otherwise.

### GetTemplateVlanTypeOk

`func (o *InlineResponse20112) GetTemplateVlanTypeOk() (*string, bool)`

GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVlanType

`func (o *InlineResponse20112) SetTemplateVlanType(v string)`

SetTemplateVlanType sets TemplateVlanType field to given value.

### HasTemplateVlanType

`func (o *InlineResponse20112) HasTemplateVlanType() bool`

HasTemplateVlanType returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20112) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20112) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20112) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20112) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *InlineResponse20112) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineResponse20112) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineResponse20112) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineResponse20112) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *InlineResponse20112) GetMandatoryDhcp() NetworksNetworkIdApplianceVlansMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *InlineResponse20112) GetMandatoryDhcpOk() (*NetworksNetworkIdApplianceVlansMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *InlineResponse20112) SetMandatoryDhcp(v NetworksNetworkIdApplianceVlansMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *InlineResponse20112) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20112) GetIpv6() NetworksNetworkIdApplianceVlansIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20112) GetIpv6Ok() (*NetworksNetworkIdApplianceVlansIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20112) SetIpv6(v NetworksNetworkIdApplianceVlansIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20112) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20112) GetVrf() NetworksNetworkIdApplianceVlansVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20112) GetVrfOk() (*NetworksNetworkIdApplianceVlansVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20112) SetVrf(v NetworksNetworkIdApplianceVlansVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20112) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


