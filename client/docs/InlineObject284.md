# InlineObject284

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | ID of the cluster to be provisioned | 
**Network** | [**OrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork**](OrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork.md) |  | 
**Name** | **string** | Name of the new cluster | 
**Uplinks** | [**[]NetworksNetworkIdCampusGatewayClustersUplinks**](NetworksNetworkIdCampusGatewayClustersUplinks.md) | Uplink interface settings of the cluster | 
**Tunnels** | [**[]NetworksNetworkIdCampusGatewayClustersTunnels**](NetworksNetworkIdCampusGatewayClustersTunnels.md) | Tunnel interface settings of the cluster: Reuse uplink or specify tunnel interface | 
**Nameservers** | [**NetworksNetworkIdCampusGatewayClustersNameservers**](NetworksNetworkIdCampusGatewayClustersNameservers.md) |  | 
**PortChannels** | [**[]OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels**](OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels.md) | Port channel settings of the cluster | 
**Devices** | Pointer to [**[]NetworksNetworkIdCampusGatewayClustersDevices**](NetworksNetworkIdCampusGatewayClustersDevices.md) | Devices to be added to the cluster | [optional] 
**Failover** | Pointer to [**OrganizationsOrganizationIdCampusGatewayClustersProvisionFailover**](OrganizationsOrganizationIdCampusGatewayClustersProvisionFailover.md) |  | [optional] 
**Notes** | Pointer to **string** | Notes about cluster with max size of 511 characters allowed | [optional] 

## Methods

### NewInlineObject284

`func NewInlineObject284(clusterId string, network OrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork, name string, uplinks []NetworksNetworkIdCampusGatewayClustersUplinks, tunnels []NetworksNetworkIdCampusGatewayClustersTunnels, nameservers NetworksNetworkIdCampusGatewayClustersNameservers, portChannels []OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels, ) *InlineObject284`

NewInlineObject284 instantiates a new InlineObject284 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject284WithDefaults

`func NewInlineObject284WithDefaults() *InlineObject284`

NewInlineObject284WithDefaults instantiates a new InlineObject284 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineObject284) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineObject284) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineObject284) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetNetwork

`func (o *InlineObject284) GetNetwork() OrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject284) GetNetworkOk() (*OrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject284) SetNetwork(v OrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork)`

SetNetwork sets Network field to given value.


### GetName

`func (o *InlineObject284) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject284) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject284) SetName(v string)`

SetName sets Name field to given value.


### GetUplinks

`func (o *InlineObject284) GetUplinks() []NetworksNetworkIdCampusGatewayClustersUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineObject284) GetUplinksOk() (*[]NetworksNetworkIdCampusGatewayClustersUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineObject284) SetUplinks(v []NetworksNetworkIdCampusGatewayClustersUplinks)`

SetUplinks sets Uplinks field to given value.


### GetTunnels

`func (o *InlineObject284) GetTunnels() []NetworksNetworkIdCampusGatewayClustersTunnels`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineObject284) GetTunnelsOk() (*[]NetworksNetworkIdCampusGatewayClustersTunnels, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineObject284) SetTunnels(v []NetworksNetworkIdCampusGatewayClustersTunnels)`

SetTunnels sets Tunnels field to given value.


### GetNameservers

`func (o *InlineObject284) GetNameservers() NetworksNetworkIdCampusGatewayClustersNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineObject284) GetNameserversOk() (*NetworksNetworkIdCampusGatewayClustersNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineObject284) SetNameservers(v NetworksNetworkIdCampusGatewayClustersNameservers)`

SetNameservers sets Nameservers field to given value.


### GetPortChannels

`func (o *InlineObject284) GetPortChannels() []OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *InlineObject284) GetPortChannelsOk() (*[]OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *InlineObject284) SetPortChannels(v []OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels)`

SetPortChannels sets PortChannels field to given value.


### GetDevices

`func (o *InlineObject284) GetDevices() []NetworksNetworkIdCampusGatewayClustersDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject284) GetDevicesOk() (*[]NetworksNetworkIdCampusGatewayClustersDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject284) SetDevices(v []NetworksNetworkIdCampusGatewayClustersDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineObject284) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetFailover

`func (o *InlineObject284) GetFailover() OrganizationsOrganizationIdCampusGatewayClustersProvisionFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineObject284) GetFailoverOk() (*OrganizationsOrganizationIdCampusGatewayClustersProvisionFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineObject284) SetFailover(v OrganizationsOrganizationIdCampusGatewayClustersProvisionFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineObject284) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject284) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject284) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject284) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject284) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


