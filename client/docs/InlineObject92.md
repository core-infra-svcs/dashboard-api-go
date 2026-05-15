# InlineObject92

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new cluster | 
**Uplinks** | [**[]NetworksNetworkIdCampusGatewayClustersUplinks**](NetworksNetworkIdCampusGatewayClustersUplinks.md) | Uplink interface settings of the cluster | 
**Tunnels** | [**[]NetworksNetworkIdCampusGatewayClustersTunnels**](NetworksNetworkIdCampusGatewayClustersTunnels.md) | Tunnel interface settings of the cluster: Reuse uplink or specify tunnel interface | 
**Nameservers** | [**NetworksNetworkIdCampusGatewayClustersNameservers**](NetworksNetworkIdCampusGatewayClustersNameservers.md) |  | 
**PortChannels** | [**[]NetworksNetworkIdCampusGatewayClustersPortChannels**](NetworksNetworkIdCampusGatewayClustersPortChannels.md) | Port channel settings of the cluster | 
**Devices** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersDevices**](NetworksNetworkIdCampusGatewayClustersDevices.md) | Devices to be added to the cluster | [optional] 
**Notes** | Pointer to **string** | Notes about cluster with max size of 511 characters allowed | [optional] 

## Methods

### NewInlineObject92

`func NewInlineObject92(name string, uplinks []NetworksNetworkIdCampusGatewayClustersUplinks, tunnels []NetworksNetworkIdCampusGatewayClustersTunnels, nameservers NetworksNetworkIdCampusGatewayClustersNameservers, portChannels []NetworksNetworkIdCampusGatewayClustersPortChannels, ) *InlineObject92`

NewInlineObject92 instantiates a new InlineObject92 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject92WithDefaults

`func NewInlineObject92WithDefaults() *InlineObject92`

NewInlineObject92WithDefaults instantiates a new InlineObject92 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject92) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject92) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject92) SetName(v string)`

SetName sets Name field to given value.


### GetUplinks

`func (o *InlineObject92) GetUplinks() []NetworksNetworkIdCampusGatewayClustersUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineObject92) GetUplinksOk() (*[]NetworksNetworkIdCampusGatewayClustersUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineObject92) SetUplinks(v []NetworksNetworkIdCampusGatewayClustersUplinks)`

SetUplinks sets Uplinks field to given value.


### GetTunnels

`func (o *InlineObject92) GetTunnels() []NetworksNetworkIdCampusGatewayClustersTunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineObject92) GetTunnelsOk() (*[]NetworksNetworkIdCampusGatewayClustersTunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineObject92) SetTunnels(v []NetworksNetworkIdCampusGatewayClustersTunnels)`

SetTunnels sets Tunnels field to given value.


### GetNameservers

`func (o *InlineObject92) GetNameservers() NetworksNetworkIdCampusGatewayClustersNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineObject92) GetNameserversOk() (*NetworksNetworkIdCampusGatewayClustersNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineObject92) SetNameservers(v NetworksNetworkIdCampusGatewayClustersNameservers)`

SetNameservers sets Nameservers field to given value.


### GetPortChannels

`func (o *InlineObject92) GetPortChannels() []NetworksNetworkIdCampusGatewayClustersPortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineObject92) GetPortChannelsOk() (*[]NetworksNetworkIdCampusGatewayClustersPortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineObject92) SetPortChannels(v []NetworksNetworkIdCampusGatewayClustersPortChannels)`

SetPortChannels sets PortChannels field to given value.


### GetDevices

`func (o *InlineObject92) GetDevices() []NetworksNetworkIdCampusGatewayClustersDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject92) GetDevicesOk() (*[]NetworksNetworkIdCampusGatewayClustersDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject92) SetDevices(v []NetworksNetworkIdCampusGatewayClustersDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineObject92) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject92) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject92) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject92) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject92) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


