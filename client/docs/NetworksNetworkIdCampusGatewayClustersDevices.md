# NetworksNetworkIdCampusGatewayClustersDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | Serial of the device | 
**Uplinks** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersUplinks1**](NetworksNetworkIdCampusGatewayClustersUplinks1.md) | Uplink settings of the device when uplink assignment mode of cluster is static | [optional] 
**Tunnels** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersTunnels1**](NetworksNetworkIdCampusGatewayClustersTunnels1.md) | Tunnel settings of the device when tunnel interface of cluster is specified | [optional] 

## Methods

### NewNetworksNetworkIdCampusGatewayClustersDevices

`func NewNetworksNetworkIdCampusGatewayClustersDevices(serial string, ) *NetworksNetworkIdCampusGatewayClustersDevices`

NewNetworksNetworkIdCampusGatewayClustersDevices instantiates a new NetworksNetworkIdCampusGatewayClustersDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdCampusGatewayClustersDevicesWithDefaults

`func NewNetworksNetworkIdCampusGatewayClustersDevicesWithDefaults() *NetworksNetworkIdCampusGatewayClustersDevices`

NewNetworksNetworkIdCampusGatewayClustersDevicesWithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetUplinks

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) GetUplinks() []NetworksNetworkIdCampusGatewayClustersUplinks1`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) GetUplinksOk() (*[]NetworksNetworkIdCampusGatewayClustersUplinks1, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) SetUplinks(v []NetworksNetworkIdCampusGatewayClustersUplinks1)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) GetTunnels() []NetworksNetworkIdCampusGatewayClustersTunnels1`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) GetTunnelsOk() (*[]NetworksNetworkIdCampusGatewayClustersTunnels1, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) SetTunnels(v []NetworksNetworkIdCampusGatewayClustersTunnels1)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *NetworksNetworkIdCampusGatewayClustersDevices) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


