# InlineResponse20014Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20014ValueSource**](InlineResponse20014ValueSource.md) |  | 
**Destination** | [**InlineResponse20014ValueDestination**](InlineResponse20014ValueDestination.md) |  | 

## Methods

### NewInlineResponse20014Value

`func NewInlineResponse20014Value(source InlineResponse20014ValueSource, destination InlineResponse20014ValueDestination, ) *InlineResponse20014Value`

NewInlineResponse20014Value instantiates a new InlineResponse20014Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20014ValueWithDefaults

`func NewInlineResponse20014ValueWithDefaults() *InlineResponse20014Value`

NewInlineResponse20014ValueWithDefaults instantiates a new InlineResponse20014Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20014Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20014Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20014Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20014Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20014Value) GetSource() InlineResponse20014ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20014Value) GetSourceOk() (*InlineResponse20014ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20014Value) SetSource(v InlineResponse20014ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20014Value) GetDestination() InlineResponse20014ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20014Value) GetDestinationOk() (*InlineResponse20014ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20014Value) SetDestination(v InlineResponse20014ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


