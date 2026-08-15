# InlineResponse200323CampusGateways

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MCG name | 
**Priority** | **int32** | CG priority for this AP (0 &#x3D; primary, 1 &#x3D; secondary) | 
**Serial** | **string** | MCG serial | 
**Mac** | **string** | MCG mac address | 
**Tunnel** | [**InlineResponse200323Tunnel**](InlineResponse200323Tunnel.md) |  | 
**Data** | [**InlineResponse200323Data**](InlineResponse200323Data.md) |  | 
**Url** | **string** | Link to campus gateway details page | 
**Cluster** | [**InlineResponse200323Cluster**](InlineResponse200323Cluster.md) |  | 

## Methods

### NewInlineResponse200323CampusGateways

`func NewInlineResponse200323CampusGateways(name string, priority int32, serial string, mac string, tunnel InlineResponse200323Tunnel, data InlineResponse200323Data, url string, cluster InlineResponse200323Cluster, ) *InlineResponse200323CampusGateways`

NewInlineResponse200323CampusGateways instantiates a new InlineResponse200323CampusGateways object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200323CampusGatewaysWithDefaults

`func NewInlineResponse200323CampusGatewaysWithDefaults() *InlineResponse200323CampusGateways`

NewInlineResponse200323CampusGatewaysWithDefaults instantiates a new InlineResponse200323CampusGateways object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200323CampusGateways) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200323CampusGateways) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200323CampusGateways) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *InlineResponse200323CampusGateways) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200323CampusGateways) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200323CampusGateways) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSerial

`func (o *InlineResponse200323CampusGateways) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200323CampusGateways) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200323CampusGateways) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200323CampusGateways) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200323CampusGateways) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200323CampusGateways) SetMac(v string)`

SetMac sets Mac field to given value.


### GetTunnel

`func (o *InlineResponse200323CampusGateways) GetTunnel() InlineResponse200323Tunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *InlineResponse200323CampusGateways) GetTunnelOk() (*InlineResponse200323Tunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *InlineResponse200323CampusGateways) SetTunnel(v InlineResponse200323Tunnel)`

SetTunnel sets Tunnel field to given value.


### GetData

`func (o *InlineResponse200323CampusGateways) GetData() InlineResponse200323Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200323CampusGateways) GetDataOk() (*InlineResponse200323Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200323CampusGateways) SetData(v InlineResponse200323Data)`

SetData sets Data field to given value.


### GetUrl

`func (o *InlineResponse200323CampusGateways) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200323CampusGateways) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200323CampusGateways) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCluster

`func (o *InlineResponse200323CampusGateways) GetCluster() InlineResponse200323Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200323CampusGateways) GetClusterOk() (*InlineResponse200323Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200323CampusGateways) SetCluster(v InlineResponse200323Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


