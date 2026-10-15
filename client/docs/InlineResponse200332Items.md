# InlineResponse200332Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200332Network**](InlineResponse200332Network.md) |  | [optional] 
**ClusterId** | Pointer to **string** | ID of the cluster | [optional] 
**Name** | Pointer to **string** | Name of the cluster | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse20118Uplinks**](InlineResponse20118Uplinks.md) | Uplink settings of the cluster | [optional] 
**Tunnels** | Pointer to [**[]InlineResponse20118Tunnels**](InlineResponse20118Tunnels.md) | Tunnel settings of the cluster | [optional] 
**Nameservers** | Pointer to [**InlineResponse20118Nameservers**](InlineResponse20118Nameservers.md) |  | [optional] 
**PortChannels** | Pointer to [**[]InlineResponse20118PortChannels**](InlineResponse20118PortChannels.md) | Port channel settings of the cluster | [optional] 
**Devices** | Pointer to [**[]InlineResponse20118Devices**](InlineResponse20118Devices.md) | Devices in the cluster | [optional] 
**Notes** | Pointer to **string** | Notes about cluster | [optional] 
**Url** | Pointer to **string** | URL to display cluster details | [optional] 

## Methods

### NewInlineResponse200332Items

`func NewInlineResponse200332Items() *InlineResponse200332Items`

NewInlineResponse200332Items instantiates a new InlineResponse200332Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200332ItemsWithDefaults

`func NewInlineResponse200332ItemsWithDefaults() *InlineResponse200332Items`

NewInlineResponse200332ItemsWithDefaults instantiates a new InlineResponse200332Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200332Items) GetNetwork() InlineResponse200332Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200332Items) GetNetworkOk() (*InlineResponse200332Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200332Items) SetNetwork(v InlineResponse200332Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200332Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetClusterId

`func (o *InlineResponse200332Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200332Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200332Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200332Items) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200332Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200332Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200332Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200332Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200332Items) GetUplinks() []InlineResponse20118Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200332Items) GetUplinksOk() (*[]InlineResponse20118Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200332Items) SetUplinks(v []InlineResponse20118Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200332Items) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineResponse200332Items) GetTunnels() []InlineResponse20118Tunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineResponse200332Items) GetTunnelsOk() (*[]InlineResponse20118Tunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineResponse200332Items) SetTunnels(v []InlineResponse20118Tunnels)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineResponse200332Items) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse200332Items) GetNameservers() InlineResponse20118Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse200332Items) GetNameserversOk() (*InlineResponse20118Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse200332Items) SetNameservers(v InlineResponse20118Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse200332Items) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortChannels

`func (o *InlineResponse200332Items) GetPortChannels() []InlineResponse20118PortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineResponse200332Items) GetPortChannelsOk() (*[]InlineResponse20118PortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineResponse200332Items) SetPortChannels(v []InlineResponse20118PortChannels)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *InlineResponse200332Items) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200332Items) GetDevices() []InlineResponse20118Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200332Items) GetDevicesOk() (*[]InlineResponse20118Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200332Items) SetDevices(v []InlineResponse20118Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200332Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200332Items) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200332Items) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200332Items) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200332Items) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200332Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200332Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200332Items) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200332Items) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


