# InlineResponse200349Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | The serial number of the Campus Gateway | 
**Network** | [**InlineResponse200349Network**](InlineResponse200349Network.md) |  | 
**ObservedAt** | **NullableTime** | When the node_stats value was last reported | 
**Load** | [**NullableInlineResponse200349Load**](InlineResponse200349Load.md) |  | 

## Methods

### NewInlineResponse200349Items

`func NewInlineResponse200349Items(serial string, network InlineResponse200349Network, observedAt NullableTime, load NullableInlineResponse200349Load, ) *InlineResponse200349Items`

NewInlineResponse200349Items instantiates a new InlineResponse200349Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200349ItemsWithDefaults

`func NewInlineResponse200349ItemsWithDefaults() *InlineResponse200349Items`

NewInlineResponse200349ItemsWithDefaults instantiates a new InlineResponse200349Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200349Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200349Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200349Items) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetNetwork

`func (o *InlineResponse200349Items) GetNetwork() InlineResponse200349Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200349Items) GetNetworkOk() (*InlineResponse200349Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200349Items) SetNetwork(v InlineResponse200349Network)`

SetNetwork sets Network field to given value.


### GetObservedAt

`func (o *InlineResponse200349Items) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *InlineResponse200349Items) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *InlineResponse200349Items) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### SetObservedAtNil

`func (o *InlineResponse200349Items) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *InlineResponse200349Items) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetLoad

`func (o *InlineResponse200349Items) GetLoad() InlineResponse200349Load`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *InlineResponse200349Items) GetLoadOk() (*InlineResponse200349Load, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *InlineResponse200349Items) SetLoad(v InlineResponse200349Load)`

SetLoad sets Load field to given value.


### SetLoadNil

`func (o *InlineResponse200349Items) SetLoadNil(b bool)`

 SetLoadNil sets the value for Load to be an explicit nil

### UnsetLoad
`func (o *InlineResponse200349Items) UnsetLoad()`

UnsetLoad ensures that no value is present for Load, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


