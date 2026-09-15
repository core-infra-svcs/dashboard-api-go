# InlineResponse200324Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The unique identifier of the campus gateway cluster | 
**Name** | **string** | The name of the campus gateway cluster | 
**Url** | **string** | Dashboard URL for campus gateway cluster settings for this row&#39;s campus gateway network | 
**Network** | [**InlineResponse200324Network**](InlineResponse200324Network.md) |  | 
**Source** | [**InlineResponse200324Source**](InlineResponse200324Source.md) |  | 

## Methods

### NewInlineResponse200324Items

`func NewInlineResponse200324Items(clusterId string, name string, url string, network InlineResponse200324Network, source InlineResponse200324Source, ) *InlineResponse200324Items`

NewInlineResponse200324Items instantiates a new InlineResponse200324Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200324ItemsWithDefaults

`func NewInlineResponse200324ItemsWithDefaults() *InlineResponse200324Items`

NewInlineResponse200324ItemsWithDefaults instantiates a new InlineResponse200324Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200324Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200324Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200324Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *InlineResponse200324Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200324Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200324Items) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *InlineResponse200324Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200324Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200324Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200324Items) GetNetwork() InlineResponse200324Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200324Items) GetNetworkOk() (*InlineResponse200324Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200324Items) SetNetwork(v InlineResponse200324Network)`

SetNetwork sets Network field to given value.


### GetSource

`func (o *InlineResponse200324Items) GetSource() InlineResponse200324Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse200324Items) GetSourceOk() (*InlineResponse200324Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse200324Items) SetSource(v InlineResponse200324Source)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


