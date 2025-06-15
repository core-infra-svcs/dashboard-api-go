# NetworksNetworkIdCampusGatewayClustersClusterIdTunnels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uplink** | Pointer to [**NetworksNetworkIdCampusGatewayClustersUplink**](NetworksNetworkIdCampusGatewayClustersUplink.md) |  | [optional] 
**Interface** | Pointer to **string** | Interface identifier, should be set to tun1. Specify this interface only if uplink isn&#39;t reused as tunnel | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the interface | [optional] 
**Addresses** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdAddresses1**](NetworksNetworkIdCampusGatewayClustersClusterIdAddresses1.md) | Addresses of the interface | [optional] 

## Methods

### NewNetworksNetworkIdCampusGatewayClustersClusterIdTunnels

`func NewNetworksNetworkIdCampusGatewayClustersClusterIdTunnels() *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels`

NewNetworksNetworkIdCampusGatewayClustersClusterIdTunnels instantiates a new NetworksNetworkIdCampusGatewayClustersClusterIdTunnels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdCampusGatewayClustersClusterIdTunnelsWithDefaults

`func NewNetworksNetworkIdCampusGatewayClustersClusterIdTunnelsWithDefaults() *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels`

NewNetworksNetworkIdCampusGatewayClustersClusterIdTunnelsWithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersClusterIdTunnels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplink

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetUplink() NetworksNetworkIdCampusGatewayClustersUplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetUplinkOk() (*NetworksNetworkIdCampusGatewayClustersUplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) SetUplink(v NetworksNetworkIdCampusGatewayClustersUplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetAddresses() []NetworksNetworkIdCampusGatewayClustersClusterIdAddresses1`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) GetAddressesOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdAddresses1, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) SetAddresses(v []NetworksNetworkIdCampusGatewayClustersClusterIdAddresses1)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersClusterIdTunnels) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


