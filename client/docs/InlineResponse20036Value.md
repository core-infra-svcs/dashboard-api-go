# InlineResponse20036Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20036ValueSource**](InlineResponse20036ValueSource.md) |  | 
**Destination** | [**InlineResponse20036ValueDestination**](InlineResponse20036ValueDestination.md) |  | 

## Methods

### NewInlineResponse20036Value

`func NewInlineResponse20036Value(source InlineResponse20036ValueSource, destination InlineResponse20036ValueDestination, ) *InlineResponse20036Value`

NewInlineResponse20036Value instantiates a new InlineResponse20036Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20036ValueWithDefaults

`func NewInlineResponse20036ValueWithDefaults() *InlineResponse20036Value`

NewInlineResponse20036ValueWithDefaults instantiates a new InlineResponse20036Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20036Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20036Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20036Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20036Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20036Value) GetSource() InlineResponse20036ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20036Value) GetSourceOk() (*InlineResponse20036ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20036Value) SetSource(v InlineResponse20036ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20036Value) GetDestination() InlineResponse20036ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20036Value) GetDestinationOk() (*InlineResponse20036ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20036Value) SetDestination(v InlineResponse20036ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


