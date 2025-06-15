# InlineResponse20110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ID of the cluster | [optional] 
**Name** | Pointer to **string** | Name of the cluster | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse20110Uplinks**](InlineResponse20110Uplinks.md) | Uplink settings of the cluster | [optional] 
**Tunnels** | Pointer to [**[]InlineResponse20110Tunnels**](InlineResponse20110Tunnels.md) | Tunnel settings of the cluster | [optional] 
**Nameservers** | Pointer to [**NetworksNetworkIdCampusGatewayClustersNameservers**](NetworksNetworkIdCampusGatewayClustersNameservers.md) |  | [optional] 
**PortChannels** | Pointer to [**[]InlineResponse20110PortChannels**](InlineResponse20110PortChannels.md) | Port channel settings of the cluster | [optional] 
**Devices** | Pointer to [**[]InlineResponse20110Devices**](InlineResponse20110Devices.md) | Devices in the cluster | [optional] 
**Notes** | Pointer to **string** | Notes about cluster | [optional] 
**Url** | Pointer to **string** | URL to display cluster details | [optional] 

## Methods

### NewInlineResponse20110

`func NewInlineResponse20110() *InlineResponse20110`

NewInlineResponse20110 instantiates a new InlineResponse20110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20110WithDefaults

`func NewInlineResponse20110WithDefaults() *InlineResponse20110`

NewInlineResponse20110WithDefaults instantiates a new InlineResponse20110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse20110) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse20110) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse20110) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse20110) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20110) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20110) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20110) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20110) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse20110) GetUplinks() []InlineResponse20110Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse20110) GetUplinksOk() (*[]InlineResponse20110Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse20110) SetUplinks(v []InlineResponse20110Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse20110) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineResponse20110) GetTunnels() []InlineResponse20110Tunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineResponse20110) GetTunnelsOk() (*[]InlineResponse20110Tunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineResponse20110) SetTunnels(v []InlineResponse20110Tunnels)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineResponse20110) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse20110) GetNameservers() NetworksNetworkIdCampusGatewayClustersNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse20110) GetNameserversOk() (*NetworksNetworkIdCampusGatewayClustersNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse20110) SetNameservers(v NetworksNetworkIdCampusGatewayClustersNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse20110) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortChannels

`func (o *InlineResponse20110) GetPortChannels() []InlineResponse20110PortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineResponse20110) GetPortChannelsOk() (*[]InlineResponse20110PortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineResponse20110) SetPortChannels(v []InlineResponse20110PortChannels)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *InlineResponse20110) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse20110) GetDevices() []InlineResponse20110Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse20110) GetDevicesOk() (*[]InlineResponse20110Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse20110) SetDevices(v []InlineResponse20110Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse20110) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20110) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20110) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20110) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20110) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20110) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20110) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20110) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20110) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


