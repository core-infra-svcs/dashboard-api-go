# InlineResponse200366Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200131Network**](InlineResponse200131Network.md) |  | 
**Available** | [**[]InlineResponse200366Available**](InlineResponse200366Available.md) | The list of roles that can be assigned to a syslog server for this network | 

## Methods

### NewInlineResponse200366Items

`func NewInlineResponse200366Items(network InlineResponse200131Network, available []InlineResponse200366Available, ) *InlineResponse200366Items`

NewInlineResponse200366Items instantiates a new InlineResponse200366Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200366ItemsWithDefaults

`func NewInlineResponse200366ItemsWithDefaults() *InlineResponse200366Items`

NewInlineResponse200366ItemsWithDefaults instantiates a new InlineResponse200366Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200366Items) GetNetwork() InlineResponse200131Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200366Items) GetNetworkOk() (*InlineResponse200131Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200366Items) SetNetwork(v InlineResponse200131Network)`

SetNetwork sets Network field to given value.


### GetAvailable

`func (o *InlineResponse200366Items) GetAvailable() []InlineResponse200366Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200366Items) GetAvailableOk() (*[]InlineResponse200366Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200366Items) SetAvailable(v []InlineResponse200366Available)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


