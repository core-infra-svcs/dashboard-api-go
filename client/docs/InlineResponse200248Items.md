# InlineResponse200248Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200248Network**](InlineResponse200248Network.md) |  | [optional] 
**Transitions** | Pointer to [**[]InlineResponse200248Transitions**](InlineResponse200248Transitions.md) | OWE transition pairs for the network | [optional] 

## Methods

### NewInlineResponse200248Items

`func NewInlineResponse200248Items() *InlineResponse200248Items`

NewInlineResponse200248Items instantiates a new InlineResponse200248Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200248ItemsWithDefaults

`func NewInlineResponse200248ItemsWithDefaults() *InlineResponse200248Items`

NewInlineResponse200248ItemsWithDefaults instantiates a new InlineResponse200248Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200248Items) GetNetwork() InlineResponse200248Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200248Items) GetNetworkOk() (*InlineResponse200248Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200248Items) SetNetwork(v InlineResponse200248Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200248Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTransitions

`func (o *InlineResponse200248Items) GetTransitions() []InlineResponse200248Transitions`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *InlineResponse200248Items) GetTransitionsOk() (*[]InlineResponse200248Transitions, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *InlineResponse200248Items) SetTransitions(v []InlineResponse200248Transitions)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *InlineResponse200248Items) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


