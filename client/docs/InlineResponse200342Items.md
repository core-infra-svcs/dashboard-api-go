# InlineResponse200342Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | The serial number of the Campus Gateway | 
**Network** | [**InlineResponse200342Network**](InlineResponse200342Network.md) |  | 
**ObservedAt** | **NullableTime** | When the node_stats value was last reported | 
**Load** | [**NullableInlineResponse200342Load**](InlineResponse200342Load.md) |  | 

## Methods

### NewInlineResponse200342Items

`func NewInlineResponse200342Items(serial string, network InlineResponse200342Network, observedAt NullableTime, load NullableInlineResponse200342Load, ) *InlineResponse200342Items`

NewInlineResponse200342Items instantiates a new InlineResponse200342Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200342ItemsWithDefaults

`func NewInlineResponse200342ItemsWithDefaults() *InlineResponse200342Items`

NewInlineResponse200342ItemsWithDefaults instantiates a new InlineResponse200342Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200342Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200342Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200342Items) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetNetwork

`func (o *InlineResponse200342Items) GetNetwork() InlineResponse200342Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200342Items) GetNetworkOk() (*InlineResponse200342Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200342Items) SetNetwork(v InlineResponse200342Network)`

SetNetwork sets Network field to given value.


### GetObservedAt

`func (o *InlineResponse200342Items) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *InlineResponse200342Items) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *InlineResponse200342Items) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### SetObservedAtNil

`func (o *InlineResponse200342Items) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *InlineResponse200342Items) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetLoad

`func (o *InlineResponse200342Items) GetLoad() InlineResponse200342Load`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *InlineResponse200342Items) GetLoadOk() (*InlineResponse200342Load, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *InlineResponse200342Items) SetLoad(v InlineResponse200342Load)`

SetLoad sets Load field to given value.


### SetLoadNil

`func (o *InlineResponse200342Items) SetLoadNil(b bool)`

 SetLoadNil sets the value for Load to be an explicit nil

### UnsetLoad
`func (o *InlineResponse200342Items) UnsetLoad()`

UnsetLoad ensures that no value is present for Load, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


