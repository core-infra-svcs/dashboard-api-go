# InlineObject89

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the cluster | [optional] 
**Uplinks** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdUplinks**](NetworksNetworkIdCampusGatewayClustersClusterIdUplinks.md) | Uplink interface settings of the cluster | [optional] 
**Tunnels** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdTunnels**](NetworksNetworkIdCampusGatewayClustersClusterIdTunnels.md) | Tunnel interface settings of the cluster: Reuse uplink or specify tunnel interface | [optional] 
**Nameservers** | Pointer to [**InlineResponse20112Nameservers**](InlineResponse20112Nameservers.md) |  | [optional] 
**PortChannels** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels**](NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels.md) | Port channel settings of the cluster | [optional] 
**Devices** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdDevices**](NetworksNetworkIdCampusGatewayClustersClusterIdDevices.md) | Devices in the cluster | [optional] 
**Notes** | Pointer to **string** | Notes about cluster with max size of 511 characters allowed | [optional] 

## Methods

### NewInlineObject89

`func NewInlineObject89() *InlineObject89`

NewInlineObject89 instantiates a new InlineObject89 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject89WithDefaults

`func NewInlineObject89WithDefaults() *InlineObject89`

NewInlineObject89WithDefaults instantiates a new InlineObject89 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject89) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject89) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject89) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject89) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineObject89) GetUplinks() []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineObject89) GetUplinksOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineObject89) SetUplinks(v []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineObject89) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineObject89) GetTunnels() []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineObject89) GetTunnelsOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdTunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineObject89) SetTunnels(v []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineObject89) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineObject89) GetNameservers() InlineResponse20112Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineObject89) GetNameserversOk() (*InlineResponse20112Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineObject89) SetNameservers(v InlineResponse20112Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineObject89) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortChannels

`func (o *InlineObject89) GetPortChannels() []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineObject89) GetPortChannelsOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineObject89) SetPortChannels(v []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *InlineObject89) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### GetDevices

`func (o *InlineObject89) GetDevices() []NetworksNetworkIdCampusGatewayClustersClusterIdDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject89) GetDevicesOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject89) SetDevices(v []NetworksNetworkIdCampusGatewayClustersClusterIdDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineObject89) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject89) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject89) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject89) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject89) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


