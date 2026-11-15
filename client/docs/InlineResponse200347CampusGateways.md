# InlineResponse200347CampusGateways

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MCG name | 
**Priority** | **int32** | CG priority for this AP (0 &#x3D; primary, 1 &#x3D; secondary) | 
**Serial** | **string** | MCG serial | 
**Mac** | **string** | MCG mac address | 
**Tunnel** | [**InlineResponse200347Tunnel**](InlineResponse200347Tunnel.md) |  | 
**Data** | [**InlineResponse200347Data**](InlineResponse200347Data.md) |  | 
**Url** | **string** | Link to campus gateway details page | 
**Cluster** | [**InlineResponse200347Cluster**](InlineResponse200347Cluster.md) |  | 

## Methods

### NewInlineResponse200347CampusGateways

`func NewInlineResponse200347CampusGateways(name string, priority int32, serial string, mac string, tunnel InlineResponse200347Tunnel, data InlineResponse200347Data, url string, cluster InlineResponse200347Cluster, ) *InlineResponse200347CampusGateways`

NewInlineResponse200347CampusGateways instantiates a new InlineResponse200347CampusGateways object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200347CampusGatewaysWithDefaults

`func NewInlineResponse200347CampusGatewaysWithDefaults() *InlineResponse200347CampusGateways`

NewInlineResponse200347CampusGatewaysWithDefaults instantiates a new InlineResponse200347CampusGateways object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200347CampusGateways) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200347CampusGateways) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200347CampusGateways) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *InlineResponse200347CampusGateways) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200347CampusGateways) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200347CampusGateways) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSerial

`func (o *InlineResponse200347CampusGateways) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200347CampusGateways) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200347CampusGateways) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200347CampusGateways) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200347CampusGateways) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200347CampusGateways) SetMac(v string)`

SetMac sets Mac field to given value.


### GetTunnel

`func (o *InlineResponse200347CampusGateways) GetTunnel() InlineResponse200347Tunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *InlineResponse200347CampusGateways) GetTunnelOk() (*InlineResponse200347Tunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *InlineResponse200347CampusGateways) SetTunnel(v InlineResponse200347Tunnel)`

SetTunnel sets Tunnel field to given value.


### GetData

`func (o *InlineResponse200347CampusGateways) GetData() InlineResponse200347Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200347CampusGateways) GetDataOk() (*InlineResponse200347Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200347CampusGateways) SetData(v InlineResponse200347Data)`

SetData sets Data field to given value.


### GetUrl

`func (o *InlineResponse200347CampusGateways) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200347CampusGateways) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200347CampusGateways) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCluster

`func (o *InlineResponse200347CampusGateways) GetCluster() InlineResponse200347Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200347CampusGateways) GetClusterOk() (*InlineResponse200347Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200347CampusGateways) SetCluster(v InlineResponse200347Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


