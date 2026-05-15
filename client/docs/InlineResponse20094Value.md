# InlineResponse20094Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20094ValueSource**](InlineResponse20094ValueSource.md) |  | 
**Destination** | [**InlineResponse20094ValueDestination**](InlineResponse20094ValueDestination.md) |  | 

## Methods

### NewInlineResponse20094Value

`func NewInlineResponse20094Value(source InlineResponse20094ValueSource, destination InlineResponse20094ValueDestination, ) *InlineResponse20094Value`

NewInlineResponse20094Value instantiates a new InlineResponse20094Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20094ValueWithDefaults

`func NewInlineResponse20094ValueWithDefaults() *InlineResponse20094Value`

NewInlineResponse20094ValueWithDefaults instantiates a new InlineResponse20094Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20094Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20094Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20094Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20094Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20094Value) GetSource() InlineResponse20094ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20094Value) GetSourceOk() (*InlineResponse20094ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20094Value) SetSource(v InlineResponse20094ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20094Value) GetDestination() InlineResponse20094ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20094Value) GetDestinationOk() (*InlineResponse20094ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20094Value) SetDestination(v InlineResponse20094ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


