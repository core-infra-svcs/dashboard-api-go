# NetworksNetworkIdApplianceVlansIpv61PrefixAssignments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonomous** | Pointer to **bool** | Auto assign a /64 prefix from the origin to the VLAN | [optional] 
**Disabled** | Pointer to **bool** | Disable the prefix assignment. Required when IPv6 is enabled on the VLAN. | [optional] 
**StaticPrefix** | Pointer to **NullableString** | Manual configuration of a /64 prefix on the VLAN | [optional] 
**StaticApplianceIp6** | Pointer to **NullableString** | Manual configuration of the IPv6 Appliance IP | [optional] 
**Origin** | Pointer to [**NetworksNetworkIdApplianceSingleLanIpv6Origin**](NetworksNetworkIdApplianceSingleLanIpv6Origin.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdApplianceVlansIpv61PrefixAssignments

`func NewNetworksNetworkIdApplianceVlansIpv61PrefixAssignments() *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments`

NewNetworksNetworkIdApplianceVlansIpv61PrefixAssignments instantiates a new NetworksNetworkIdApplianceVlansIpv61PrefixAssignments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceVlansIpv61PrefixAssignmentsWithDefaults

`func NewNetworksNetworkIdApplianceVlansIpv61PrefixAssignmentsWithDefaults() *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments`

NewNetworksNetworkIdApplianceVlansIpv61PrefixAssignmentsWithDefaults instantiates a new NetworksNetworkIdApplianceVlansIpv61PrefixAssignments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomous

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetAutonomous() bool`

GetAutonomous returns the Autonomous field if non-nil, zero value otherwise.

### GetAutonomousOk

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetAutonomousOk() (*bool, bool)`

GetAutonomousOk returns a tuple with the Autonomous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomous

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetAutonomous(v bool)`

SetAutonomous sets Autonomous field to given value.

### HasAutonomous

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) HasAutonomous() bool`

HasAutonomous returns a boolean if a field has been set.

### GetDisabled

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetStaticPrefix

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetStaticPrefix() string`

GetStaticPrefix returns the StaticPrefix field if non-nil, zero value otherwise.

### GetStaticPrefixOk

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetStaticPrefixOk() (*string, bool)`

GetStaticPrefixOk returns a tuple with the StaticPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticPrefix

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetStaticPrefix(v string)`

SetStaticPrefix sets StaticPrefix field to given value.

### HasStaticPrefix

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) HasStaticPrefix() bool`

HasStaticPrefix returns a boolean if a field has been set.

### SetStaticPrefixNil

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetStaticPrefixNil(b bool)`

 SetStaticPrefixNil sets the value for StaticPrefix to be an explicit nil

### UnsetStaticPrefix
`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) UnsetStaticPrefix()`

UnsetStaticPrefix ensures that no value is present for StaticPrefix, not even an explicit nil
### GetStaticApplianceIp6

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetStaticApplianceIp6() string`

GetStaticApplianceIp6 returns the StaticApplianceIp6 field if non-nil, zero value otherwise.

### GetStaticApplianceIp6Ok

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetStaticApplianceIp6Ok() (*string, bool)`

GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticApplianceIp6

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetStaticApplianceIp6(v string)`

SetStaticApplianceIp6 sets StaticApplianceIp6 field to given value.

### HasStaticApplianceIp6

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) HasStaticApplianceIp6() bool`

HasStaticApplianceIp6 returns a boolean if a field has been set.

### SetStaticApplianceIp6Nil

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetStaticApplianceIp6Nil(b bool)`

 SetStaticApplianceIp6Nil sets the value for StaticApplianceIp6 to be an explicit nil

### UnsetStaticApplianceIp6
`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) UnsetStaticApplianceIp6()`

UnsetStaticApplianceIp6 ensures that no value is present for StaticApplianceIp6, not even an explicit nil
### GetOrigin

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetOrigin() NetworksNetworkIdApplianceSingleLanIpv6Origin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) GetOriginOk() (*NetworksNetworkIdApplianceSingleLanIpv6Origin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) SetOrigin(v NetworksNetworkIdApplianceSingleLanIpv6Origin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *NetworksNetworkIdApplianceVlansIpv61PrefixAssignments) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


