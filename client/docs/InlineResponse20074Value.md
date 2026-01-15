# InlineResponse20074Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20074ValueSource**](InlineResponse20074ValueSource.md) |  | 
**Destination** | [**InlineResponse20074ValueDestination**](InlineResponse20074ValueDestination.md) |  | 

## Methods

### NewInlineResponse20074Value

`func NewInlineResponse20074Value(source InlineResponse20074ValueSource, destination InlineResponse20074ValueDestination, ) *InlineResponse20074Value`

NewInlineResponse20074Value instantiates a new InlineResponse20074Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074ValueWithDefaults

`func NewInlineResponse20074ValueWithDefaults() *InlineResponse20074Value`

NewInlineResponse20074ValueWithDefaults instantiates a new InlineResponse20074Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20074Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20074Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20074Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20074Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20074Value) GetSource() InlineResponse20074ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20074Value) GetSourceOk() (*InlineResponse20074ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20074Value) SetSource(v InlineResponse20074ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20074Value) GetDestination() InlineResponse20074ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20074Value) GetDestinationOk() (*InlineResponse20074ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20074Value) SetDestination(v InlineResponse20074ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


