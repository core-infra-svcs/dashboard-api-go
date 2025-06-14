# InlineObject74

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The VLAN ID of the new VLAN (must be between 1 and 4094) | 
**Name** | **string** | The name of the new VLAN | 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **NullableString** | The id of the desired group policy to apply to the VLAN | [optional] 
**TemplateVlanType** | Pointer to **string** | Type of subnetting of the VLAN. Applicable only for template network. | [optional] [default to "same"]
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all bound to the template networks. Applicable only for template network. | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdApplianceSingleLanIpv6**](NetworksNetworkIdApplianceSingleLanIpv6.md) |  | [optional] 
**DhcpHandling** | Pointer to **string** | The appliance&#39;s handling of DHCP requests on this VLAN. One of: &#39;Run a DHCP server&#39;, &#39;Relay DHCP to another server&#39; or &#39;Do not respond to DHCP requests&#39; | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: &#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39; | [optional] 
**MandatoryDhcp** | Pointer to [**NetworksNetworkIdApplianceVlansMandatoryDhcp**](NetworksNetworkIdApplianceVlansMandatoryDhcp.md) |  | [optional] 
**DhcpBootOptionsEnabled** | Pointer to **bool** | Use DHCP boot options specified in other properties | [optional] 
**DhcpOptions** | Pointer to [**[]NetworksNetworkIdApplianceVlansDhcpOptions**](NetworksNetworkIdApplianceVlansDhcpOptions.md) | The list of DHCP options that will be included in DHCP responses. Each object in the list should have \&quot;code\&quot;, \&quot;type\&quot;, and \&quot;value\&quot; properties. | [optional] 

## Methods

### NewInlineObject74

`func NewInlineObject74(id string, name string, ) *InlineObject74`

NewInlineObject74 instantiates a new InlineObject74 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject74WithDefaults

`func NewInlineObject74WithDefaults() *InlineObject74`

NewInlineObject74WithDefaults instantiates a new InlineObject74 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject74) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject74) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject74) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InlineObject74) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject74) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject74) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *InlineObject74) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject74) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject74) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject74) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineObject74) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineObject74) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineObject74) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineObject74) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineObject74) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject74) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject74) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineObject74) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### SetGroupPolicyIdNil

`func (o *InlineObject74) SetGroupPolicyIdNil(b bool)`

 SetGroupPolicyIdNil sets the value for GroupPolicyId to be an explicit nil

### UnsetGroupPolicyId
`func (o *InlineObject74) UnsetGroupPolicyId()`

UnsetGroupPolicyId ensures that no value is present for GroupPolicyId, not even an explicit nil
### GetTemplateVlanType

`func (o *InlineObject74) GetTemplateVlanType() string`

GetTemplateVlanType returns the TemplateVlanType field if non-nil, zero value otherwise.

### GetTemplateVlanTypeOk

`func (o *InlineObject74) GetTemplateVlanTypeOk() (*string, bool)`

GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVlanType

`func (o *InlineObject74) SetTemplateVlanType(v string)`

SetTemplateVlanType sets TemplateVlanType field to given value.

### HasTemplateVlanType

`func (o *InlineObject74) HasTemplateVlanType() bool`

HasTemplateVlanType returns a boolean if a field has been set.

### GetCidr

`func (o *InlineObject74) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineObject74) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineObject74) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineObject74) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *InlineObject74) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineObject74) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineObject74) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineObject74) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineObject74) GetIpv6() NetworksNetworkIdApplianceSingleLanIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineObject74) GetIpv6Ok() (*NetworksNetworkIdApplianceSingleLanIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineObject74) SetIpv6(v NetworksNetworkIdApplianceSingleLanIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineObject74) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetDhcpHandling

`func (o *InlineObject74) GetDhcpHandling() string`

GetDhcpHandling returns the DhcpHandling field if non-nil, zero value otherwise.

### GetDhcpHandlingOk

`func (o *InlineObject74) GetDhcpHandlingOk() (*string, bool)`

GetDhcpHandlingOk returns a tuple with the DhcpHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHandling

`func (o *InlineObject74) SetDhcpHandling(v string)`

SetDhcpHandling sets DhcpHandling field to given value.

### HasDhcpHandling

`func (o *InlineObject74) HasDhcpHandling() bool`

HasDhcpHandling returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *InlineObject74) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineObject74) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineObject74) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineObject74) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetMandatoryDhcp

`func (o *InlineObject74) GetMandatoryDhcp() NetworksNetworkIdApplianceVlansMandatoryDhcp`

GetMandatoryDhcp returns the MandatoryDhcp field if non-nil, zero value otherwise.

### GetMandatoryDhcpOk

`func (o *InlineObject74) GetMandatoryDhcpOk() (*NetworksNetworkIdApplianceVlansMandatoryDhcp, bool)`

GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcp

`func (o *InlineObject74) SetMandatoryDhcp(v NetworksNetworkIdApplianceVlansMandatoryDhcp)`

SetMandatoryDhcp sets MandatoryDhcp field to given value.

### HasMandatoryDhcp

`func (o *InlineObject74) HasMandatoryDhcp() bool`

HasMandatoryDhcp returns a boolean if a field has been set.

### GetDhcpBootOptionsEnabled

`func (o *InlineObject74) GetDhcpBootOptionsEnabled() bool`

GetDhcpBootOptionsEnabled returns the DhcpBootOptionsEnabled field if non-nil, zero value otherwise.

### GetDhcpBootOptionsEnabledOk

`func (o *InlineObject74) GetDhcpBootOptionsEnabledOk() (*bool, bool)`

GetDhcpBootOptionsEnabledOk returns a tuple with the DhcpBootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBootOptionsEnabled

`func (o *InlineObject74) SetDhcpBootOptionsEnabled(v bool)`

SetDhcpBootOptionsEnabled sets DhcpBootOptionsEnabled field to given value.

### HasDhcpBootOptionsEnabled

`func (o *InlineObject74) HasDhcpBootOptionsEnabled() bool`

HasDhcpBootOptionsEnabled returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *InlineObject74) GetDhcpOptions() []NetworksNetworkIdApplianceVlansDhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *InlineObject74) GetDhcpOptionsOk() (*[]NetworksNetworkIdApplianceVlansDhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *InlineObject74) SetDhcpOptions(v []NetworksNetworkIdApplianceVlansDhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *InlineObject74) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


