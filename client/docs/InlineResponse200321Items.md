# InlineResponse200321Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | SSID number | 
**SsidId** | **string** | SSID identifier | 
**Name** | **string** | SSID name | 
**Url** | **string** | SSID management URL | 
**Network** | [**InlineResponse200321Network**](InlineResponse200321Network.md) |  | 
**Cluster** | [**InlineResponse200320Cluster**](InlineResponse200320Cluster.md) |  | 

## Methods

### NewInlineResponse200321Items

`func NewInlineResponse200321Items(number int32, ssidId string, name string, url string, network InlineResponse200321Network, cluster InlineResponse200320Cluster, ) *InlineResponse200321Items`

NewInlineResponse200321Items instantiates a new InlineResponse200321Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200321ItemsWithDefaults

`func NewInlineResponse200321ItemsWithDefaults() *InlineResponse200321Items`

NewInlineResponse200321ItemsWithDefaults instantiates a new InlineResponse200321Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse200321Items) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200321Items) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200321Items) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetSsidId

`func (o *InlineResponse200321Items) GetSsidId() string`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *InlineResponse200321Items) GetSsidIdOk() (*string, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *InlineResponse200321Items) SetSsidId(v string)`

SetSsidId sets SsidId field to given value.


### GetName

`func (o *InlineResponse200321Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200321Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200321Items) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *InlineResponse200321Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200321Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200321Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200321Items) GetNetwork() InlineResponse200321Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200321Items) GetNetworkOk() (*InlineResponse200321Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200321Items) SetNetwork(v InlineResponse200321Network)`

SetNetwork sets Network field to given value.


### GetCluster

`func (o *InlineResponse200321Items) GetCluster() InlineResponse200320Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200321Items) GetClusterOk() (*InlineResponse200320Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200321Items) SetCluster(v InlineResponse200320Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


