# InlineObject109

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the cluster | [optional] 
**Uplinks** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdUplinks**](NetworksNetworkIdCampusGatewayClustersClusterIdUplinks.md) | Uplink interface settings of the cluster | [optional] 
**Tunnels** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdTunnels**](NetworksNetworkIdCampusGatewayClustersClusterIdTunnels.md) | Tunnel interface settings of the cluster: Reuse uplink or specify tunnel interface | [optional] 
**Nameservers** | Pointer to [**InlineResponse20119Nameservers**](InlineResponse20119Nameservers.md) |  | [optional] 
**PortChannels** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels**](NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels.md) | Port channel settings of the cluster | [optional] 
**Devices** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersClusterIdDevices**](NetworksNetworkIdCampusGatewayClustersClusterIdDevices.md) | Devices in the cluster | [optional] 
**Notes** | Pointer to **string** | Notes about cluster with max size of 511 characters allowed | [optional] 

## Methods

### NewInlineObject109

`func NewInlineObject109() *InlineObject109`

NewInlineObject109 instantiates a new InlineObject109 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject109WithDefaults

`func NewInlineObject109WithDefaults() *InlineObject109`

NewInlineObject109WithDefaults instantiates a new InlineObject109 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject109) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject109) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject109) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject109) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineObject109) GetUplinks() []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineObject109) GetUplinksOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineObject109) SetUplinks(v []NetworksNetworkIdCampusGatewayClustersClusterIdUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineObject109) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineObject109) GetTunnels() []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineObject109) GetTunnelsOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdTunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineObject109) SetTunnels(v []NetworksNetworkIdCampusGatewayClustersClusterIdTunnels)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineObject109) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineObject109) GetNameservers() InlineResponse20119Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineObject109) GetNameserversOk() (*InlineResponse20119Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineObject109) SetNameservers(v InlineResponse20119Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineObject109) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortChannels

`func (o *InlineObject109) GetPortChannels() []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineObject109) GetPortChannelsOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineObject109) SetPortChannels(v []NetworksNetworkIdCampusGatewayClustersClusterIdPortChannels)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *InlineObject109) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### GetDevices

`func (o *InlineObject109) GetDevices() []NetworksNetworkIdCampusGatewayClustersClusterIdDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject109) GetDevicesOk() (*[]NetworksNetworkIdCampusGatewayClustersClusterIdDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject109) SetDevices(v []NetworksNetworkIdCampusGatewayClustersClusterIdDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineObject109) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject109) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject109) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject109) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject109) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


