# InlineResponse200327CampusGateways

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MCG name | 
**Priority** | **int32** | CG priority for this AP (0 &#x3D; primary, 1 &#x3D; secondary) | 
**Serial** | **string** | MCG serial | 
**Mac** | **string** | MCG mac address | 
**Tunnel** | [**InlineResponse200327Tunnel**](InlineResponse200327Tunnel.md) |  | 
**Data** | [**InlineResponse200327Data**](InlineResponse200327Data.md) |  | 
**Url** | **string** | Link to campus gateway details page | 
**Cluster** | [**InlineResponse200327Cluster**](InlineResponse200327Cluster.md) |  | 

## Methods

### NewInlineResponse200327CampusGateways

`func NewInlineResponse200327CampusGateways(name string, priority int32, serial string, mac string, tunnel InlineResponse200327Tunnel, data InlineResponse200327Data, url string, cluster InlineResponse200327Cluster, ) *InlineResponse200327CampusGateways`

NewInlineResponse200327CampusGateways instantiates a new InlineResponse200327CampusGateways object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327CampusGatewaysWithDefaults

`func NewInlineResponse200327CampusGatewaysWithDefaults() *InlineResponse200327CampusGateways`

NewInlineResponse200327CampusGatewaysWithDefaults instantiates a new InlineResponse200327CampusGateways object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200327CampusGateways) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200327CampusGateways) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200327CampusGateways) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *InlineResponse200327CampusGateways) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200327CampusGateways) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200327CampusGateways) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSerial

`func (o *InlineResponse200327CampusGateways) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200327CampusGateways) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200327CampusGateways) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200327CampusGateways) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200327CampusGateways) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200327CampusGateways) SetMac(v string)`

SetMac sets Mac field to given value.


### GetTunnel

`func (o *InlineResponse200327CampusGateways) GetTunnel() InlineResponse200327Tunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *InlineResponse200327CampusGateways) GetTunnelOk() (*InlineResponse200327Tunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *InlineResponse200327CampusGateways) SetTunnel(v InlineResponse200327Tunnel)`

SetTunnel sets Tunnel field to given value.


### GetData

`func (o *InlineResponse200327CampusGateways) GetData() InlineResponse200327Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200327CampusGateways) GetDataOk() (*InlineResponse200327Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200327CampusGateways) SetData(v InlineResponse200327Data)`

SetData sets Data field to given value.


### GetUrl

`func (o *InlineResponse200327CampusGateways) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200327CampusGateways) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200327CampusGateways) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCluster

`func (o *InlineResponse200327CampusGateways) GetCluster() InlineResponse200327Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200327CampusGateways) GetClusterOk() (*InlineResponse200327Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200327CampusGateways) SetCluster(v InlineResponse200327Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


