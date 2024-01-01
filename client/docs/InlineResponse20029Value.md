# InlineResponse20029Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20029ValueSource**](InlineResponse20029ValueSource.md) |  | 
**Destination** | [**InlineResponse20029ValueDestination**](InlineResponse20029ValueDestination.md) |  | 

## Methods

### NewInlineResponse20029Value

`func NewInlineResponse20029Value(source InlineResponse20029ValueSource, destination InlineResponse20029ValueDestination, ) *InlineResponse20029Value`

NewInlineResponse20029Value instantiates a new InlineResponse20029Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029ValueWithDefaults

`func NewInlineResponse20029ValueWithDefaults() *InlineResponse20029Value`

NewInlineResponse20029ValueWithDefaults instantiates a new InlineResponse20029Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20029Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20029Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20029Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20029Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20029Value) GetSource() InlineResponse20029ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20029Value) GetSourceOk() (*InlineResponse20029ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20029Value) SetSource(v InlineResponse20029ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20029Value) GetDestination() InlineResponse20029ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20029Value) GetDestinationOk() (*InlineResponse20029ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20029Value) SetDestination(v InlineResponse20029ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


