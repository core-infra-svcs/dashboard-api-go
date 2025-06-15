# NetworksNetworkIdCampusGatewayClustersUplinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | **string** | Interface identifier, should be set to man1 | 
**Vlan** | **int32** | VLAN ID of the interface | 
**Addresses** | [**[]NetworksNetworkIdCampusGatewayClustersAddresses**](NetworksNetworkIdCampusGatewayClustersAddresses.md) | Addresses of the interface | 

## Methods

### NewNetworksNetworkIdCampusGatewayClustersUplinks

`func NewNetworksNetworkIdCampusGatewayClustersUplinks(interface_ string, vlan int32, addresses []NetworksNetworkIdCampusGatewayClustersAddresses, ) *NetworksNetworkIdCampusGatewayClustersUplinks`

NewNetworksNetworkIdCampusGatewayClustersUplinks instantiates a new NetworksNetworkIdCampusGatewayClustersUplinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdCampusGatewayClustersUplinksWithDefaults

`func NewNetworksNetworkIdCampusGatewayClustersUplinksWithDefaults() *NetworksNetworkIdCampusGatewayClustersUplinks`

NewNetworksNetworkIdCampusGatewayClustersUplinksWithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersUplinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) SetVlan(v int32)`

SetVlan sets Vlan field to given value.


### GetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) GetAddresses() []NetworksNetworkIdCampusGatewayClustersAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) GetAddressesOk() (*[]NetworksNetworkIdCampusGatewayClustersAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersUplinks) SetAddresses(v []NetworksNetworkIdCampusGatewayClustersAddresses)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


