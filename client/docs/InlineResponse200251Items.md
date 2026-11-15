# InlineResponse200251Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200251Network**](InlineResponse200251Network.md) |  | [optional] 
**Transitions** | Pointer to [**[]InlineResponse200251Transitions**](InlineResponse200251Transitions.md) | OWE transition pairs for the network | [optional] 

## Methods

### NewInlineResponse200251Items

`func NewInlineResponse200251Items() *InlineResponse200251Items`

NewInlineResponse200251Items instantiates a new InlineResponse200251Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200251ItemsWithDefaults

`func NewInlineResponse200251ItemsWithDefaults() *InlineResponse200251Items`

NewInlineResponse200251ItemsWithDefaults instantiates a new InlineResponse200251Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200251Items) GetNetwork() InlineResponse200251Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200251Items) GetNetworkOk() (*InlineResponse200251Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200251Items) SetNetwork(v InlineResponse200251Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200251Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTransitions

`func (o *InlineResponse200251Items) GetTransitions() []InlineResponse200251Transitions`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *InlineResponse200251Items) GetTransitionsOk() (*[]InlineResponse200251Transitions, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *InlineResponse200251Items) SetTransitions(v []InlineResponse200251Transitions)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *InlineResponse200251Items) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


