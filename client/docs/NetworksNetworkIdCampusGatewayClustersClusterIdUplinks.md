# NetworksNetworkIdCampusGatewayClustersClusterIdUplinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | **string** | Interface identifier, should be set to man1 | 
**Vlan** | Pointer to **int32** | VLAN ID of the interface | [optional] 
**Addresses** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdAddresses**](NetworksNetworkIdCampusGatewayClustersClusterIdAddresses.md) | Addresses of the interface | [optional] 

## Methods

### NewNetworksNetworkIdCampusGatewayClustersClusterIdUplinks

`func NewNetworksNetworkIdCampusGatewayClustersClusterIdUplinks(interface_ string, ) *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks`

NewNetworksNetworkIdCampusGatewayClustersClusterIdUplinks instantiates a new NetworksNetworkIdCampusGatewayClustersClusterIdUplinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdCampusGatewayClustersClusterIdUplinksWithDefaults

`func NewNetworksNetworkIdCampusGatewayClustersClusterIdUplinksWithDefaults() *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks`

NewNetworksNetworkIdCampusGatewayClustersClusterIdUplinksWithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersClusterIdUplinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) GetAddresses() []NetworksNetworkIdCampusGatewayClustersClusterIdAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) GetAddressesOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) SetAddresses(v []NetworksNetworkIdCampusGatewayClustersClusterIdAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdUplinks) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


