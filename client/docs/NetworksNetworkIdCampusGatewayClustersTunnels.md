# NetworksNetworkIdCampusGatewayClustersTunnels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uplink** | Pointer to [**NetworksNetworkIdCampusGatewayClustersUplink**](NetworksNetworkIdCampusGatewayClustersUplink.md) |  | [optional] 
**Interface** | Pointer to **string** | Interface identifier, should be set to tun1. Specify this interface only if uplink isn&#39;t reused as tunnel | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the interface | [optional] 
**Addresses** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersAddresses1**](NetworksNetworkIdCampusGatewayClustersAddresses1.md) | Addresses of the interface | [optional] 

## Methods

### NewNetworksNetworkIdCampusGatewayClustersTunnels

`func NewNetworksNetworkIdCampusGatewayClustersTunnels() *NetworksNetworkIdCampusGatewayClustersTunnels`

NewNetworksNetworkIdCampusGatewayClustersTunnels instantiates a new NetworksNetworkIdCampusGatewayClustersTunnels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdCampusGatewayClustersTunnelsWithDefaults

`func NewNetworksNetworkIdCampusGatewayClustersTunnelsWithDefaults() *NetworksNetworkIdCampusGatewayClustersTunnels`

NewNetworksNetworkIdCampusGatewayClustersTunnelsWithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersTunnels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplink

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetUplink() NetworksNetworkIdCampusGatewayClustersUplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetUplinkOk() (*NetworksNetworkIdCampusGatewayClustersUplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) SetUplink(v NetworksNetworkIdCampusGatewayClustersUplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetAddresses() []NetworksNetworkIdCampusGatewayClustersAddresses1`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) GetAddressesOk() (*[]NetworksNetworkIdCampusGatewayClustersAddresses1, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) SetAddresses(v []NetworksNetworkIdCampusGatewayClustersAddresses1)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworksNetworkIdCampusGatewayClustersTunnels) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


