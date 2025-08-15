# InlineResponse200254Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200254Network**](InlineResponse200254Network.md) |  | [optional] 
**ClusterId** | Pointer to **string** | ID of the cluster | [optional] 
**Name** | Pointer to **string** | Name of the cluster | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse20111Uplinks**](InlineResponse20111Uplinks.md) | Uplink settings of the cluster | [optional] 
**Tunnels** | Pointer to [**[]InlineResponse20111Tunnels**](InlineResponse20111Tunnels.md) | Tunnel settings of the cluster | [optional] 
**Nameservers** | Pointer to [**NetworksNetworkIdCampusGatewayClustersNameservers**](NetworksNetworkIdCampusGatewayClustersNameservers.md) |  | [optional] 
**PortChannels** | Pointer to [**[]InlineResponse20111PortChannels**](InlineResponse20111PortChannels.md) | Port channel settings of the cluster | [optional] 
**Devices** | Pointer to [**[]InlineResponse20111Devices**](InlineResponse20111Devices.md) | Devices in the cluster | [optional] 
**Notes** | Pointer to **string** | Notes about cluster | [optional] 
**Url** | Pointer to **string** | URL to display cluster details | [optional] 

## Methods

### NewInlineResponse200254Items

`func NewInlineResponse200254Items() *InlineResponse200254Items`

NewInlineResponse200254Items instantiates a new InlineResponse200254Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200254ItemsWithDefaults

`func NewInlineResponse200254ItemsWithDefaults() *InlineResponse200254Items`

NewInlineResponse200254ItemsWithDefaults instantiates a new InlineResponse200254Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200254Items) GetNetwork() InlineResponse200254Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200254Items) GetNetworkOk() (*InlineResponse200254Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200254Items) SetNetwork(v InlineResponse200254Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200254Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetClusterId

`func (o *InlineResponse200254Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200254Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200254Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200254Items) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200254Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200254Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200254Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200254Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200254Items) GetUplinks() []InlineResponse20111Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200254Items) GetUplinksOk() (*[]InlineResponse20111Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200254Items) SetUplinks(v []InlineResponse20111Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200254Items) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineResponse200254Items) GetTunnels() []InlineResponse20111Tunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineResponse200254Items) GetTunnelsOk() (*[]InlineResponse20111Tunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineResponse200254Items) SetTunnels(v []InlineResponse20111Tunnels)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineResponse200254Items) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse200254Items) GetNameservers() NetworksNetworkIdCampusGatewayClustersNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse200254Items) GetNameserversOk() (*NetworksNetworkIdCampusGatewayClustersNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse200254Items) SetNameservers(v NetworksNetworkIdCampusGatewayClustersNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse200254Items) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortChannels

`func (o *InlineResponse200254Items) GetPortChannels() []InlineResponse20111PortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineResponse200254Items) GetPortChannelsOk() (*[]InlineResponse20111PortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineResponse200254Items) SetPortChannels(v []InlineResponse20111PortChannels)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *InlineResponse200254Items) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200254Items) GetDevices() []InlineResponse20111Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200254Items) GetDevicesOk() (*[]InlineResponse20111Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200254Items) SetDevices(v []InlineResponse20111Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200254Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200254Items) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200254Items) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200254Items) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200254Items) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200254Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200254Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200254Items) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200254Items) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


