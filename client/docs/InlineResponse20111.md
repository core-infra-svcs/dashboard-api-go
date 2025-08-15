# InlineResponse20111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewInlineResponse20111

`func NewInlineResponse20111() *InlineResponse20111`

NewInlineResponse20111 instantiates a new InlineResponse20111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20111WithDefaults

`func NewInlineResponse20111WithDefaults() *InlineResponse20111`

NewInlineResponse20111WithDefaults instantiates a new InlineResponse20111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse20111) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse20111) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse20111) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse20111) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20111) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20111) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20111) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20111) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse20111) GetUplinks() []InlineResponse20111Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse20111) GetUplinksOk() (*[]InlineResponse20111Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse20111) SetUplinks(v []InlineResponse20111Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse20111) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineResponse20111) GetTunnels() []InlineResponse20111Tunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineResponse20111) GetTunnelsOk() (*[]InlineResponse20111Tunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineResponse20111) SetTunnels(v []InlineResponse20111Tunnels)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineResponse20111) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse20111) GetNameservers() NetworksNetworkIdCampusGatewayClustersNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse20111) GetNameserversOk() (*NetworksNetworkIdCampusGatewayClustersNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse20111) SetNameservers(v NetworksNetworkIdCampusGatewayClustersNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse20111) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortChannels

`func (o *InlineResponse20111) GetPortChannels() []InlineResponse20111PortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineResponse20111) GetPortChannelsOk() (*[]InlineResponse20111PortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineResponse20111) SetPortChannels(v []InlineResponse20111PortChannels)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *InlineResponse20111) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse20111) GetDevices() []InlineResponse20111Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse20111) GetDevicesOk() (*[]InlineResponse20111Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse20111) SetDevices(v []InlineResponse20111Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse20111) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20111) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20111) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20111) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20111) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20111) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20111) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20111) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20111) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


