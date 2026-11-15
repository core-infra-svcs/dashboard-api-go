# InlineResponse200103Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse200103ValueSource**](InlineResponse200103ValueSource.md) |  | 
**Destination** | [**InlineResponse200103ValueDestination**](InlineResponse200103ValueDestination.md) |  | 

## Methods

### NewInlineResponse200103Value

`func NewInlineResponse200103Value(source InlineResponse200103ValueSource, destination InlineResponse200103ValueDestination, ) *InlineResponse200103Value`

NewInlineResponse200103Value instantiates a new InlineResponse200103Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103ValueWithDefaults

`func NewInlineResponse200103ValueWithDefaults() *InlineResponse200103Value`

NewInlineResponse200103ValueWithDefaults instantiates a new InlineResponse200103Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse200103Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200103Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200103Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200103Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse200103Value) GetSource() InlineResponse200103ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse200103Value) GetSourceOk() (*InlineResponse200103ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse200103Value) SetSource(v InlineResponse200103ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse200103Value) GetDestination() InlineResponse200103ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse200103Value) GetDestinationOk() (*InlineResponse200103ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse200103Value) SetDestination(v InlineResponse200103ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


