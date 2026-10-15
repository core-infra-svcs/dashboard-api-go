# InlineResponse200337Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The unique identifier of the campus gateway cluster | 
**Name** | **string** | The name of the campus gateway cluster | 
**Url** | **string** | Dashboard URL for campus gateway cluster settings for this row&#39;s campus gateway network | 
**Network** | [**InlineResponse200337Network**](InlineResponse200337Network.md) |  | 
**Source** | [**InlineResponse200337Source**](InlineResponse200337Source.md) |  | 

## Methods

### NewInlineResponse200337Items

`func NewInlineResponse200337Items(clusterId string, name string, url string, network InlineResponse200337Network, source InlineResponse200337Source, ) *InlineResponse200337Items`

NewInlineResponse200337Items instantiates a new InlineResponse200337Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200337ItemsWithDefaults

`func NewInlineResponse200337ItemsWithDefaults() *InlineResponse200337Items`

NewInlineResponse200337ItemsWithDefaults instantiates a new InlineResponse200337Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200337Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200337Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200337Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetName

`func (o *InlineResponse200337Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200337Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200337Items) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *InlineResponse200337Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200337Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200337Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200337Items) GetNetwork() InlineResponse200337Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200337Items) GetNetworkOk() (*InlineResponse200337Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200337Items) SetNetwork(v InlineResponse200337Network)`

SetNetwork sets Network field to given value.


### GetSource

`func (o *InlineResponse200337Items) GetSource() InlineResponse200337Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse200337Items) GetSourceOk() (*InlineResponse200337Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse200337Items) SetSource(v InlineResponse200337Source)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


