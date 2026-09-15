# InlineResponse200100Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse200100ValueSource**](InlineResponse200100ValueSource.md) |  | 
**Destination** | [**InlineResponse200100ValueDestination**](InlineResponse200100ValueDestination.md) |  | 

## Methods

### NewInlineResponse200100Value

`func NewInlineResponse200100Value(source InlineResponse200100ValueSource, destination InlineResponse200100ValueDestination, ) *InlineResponse200100Value`

NewInlineResponse200100Value instantiates a new InlineResponse200100Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200100ValueWithDefaults

`func NewInlineResponse200100ValueWithDefaults() *InlineResponse200100Value`

NewInlineResponse200100ValueWithDefaults instantiates a new InlineResponse200100Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse200100Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200100Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200100Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200100Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse200100Value) GetSource() InlineResponse200100ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse200100Value) GetSourceOk() (*InlineResponse200100ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse200100Value) SetSource(v InlineResponse200100ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse200100Value) GetDestination() InlineResponse200100ValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse200100Value) GetDestinationOk() (*InlineResponse200100ValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse200100Value) SetDestination(v InlineResponse200100ValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


