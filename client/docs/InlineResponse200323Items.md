# InlineResponse200323Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | SSID number | 
**SsidId** | **string** | SSID identifier | 
**Name** | **string** | SSID name | 
**Url** | **string** | SSID management URL | 
**Network** | [**InlineResponse200323Network**](InlineResponse200323Network.md) |  | 
**Cluster** | [**InlineResponse200322Cluster**](InlineResponse200322Cluster.md) |  | 

## Methods

### NewInlineResponse200323Items

`func NewInlineResponse200323Items(number int32, ssidId string, name string, url string, network InlineResponse200323Network, cluster InlineResponse200322Cluster, ) *InlineResponse200323Items`

NewInlineResponse200323Items instantiates a new InlineResponse200323Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200323ItemsWithDefaults

`func NewInlineResponse200323ItemsWithDefaults() *InlineResponse200323Items`

NewInlineResponse200323ItemsWithDefaults instantiates a new InlineResponse200323Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse200323Items) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200323Items) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200323Items) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetSsidId

`func (o *InlineResponse200323Items) GetSsidId() string`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *InlineResponse200323Items) GetSsidIdOk() (*string, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *InlineResponse200323Items) SetSsidId(v string)`

SetSsidId sets SsidId field to given value.


### GetName

`func (o *InlineResponse200323Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200323Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200323Items) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *InlineResponse200323Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200323Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200323Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200323Items) GetNetwork() InlineResponse200323Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200323Items) GetNetworkOk() (*InlineResponse200323Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200323Items) SetNetwork(v InlineResponse200323Network)`

SetNetwork sets Network field to given value.


### GetCluster

`func (o *InlineResponse200323Items) GetCluster() InlineResponse200322Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200323Items) GetClusterOk() (*InlineResponse200322Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200323Items) SetCluster(v InlineResponse200322Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


