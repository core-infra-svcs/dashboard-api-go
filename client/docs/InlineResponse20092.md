# InlineResponse20092

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to **float32** | Usage received by the client on a given day | [optional] 
**Sent** | Pointer to **float32** | Usage sent by the client on a given day | [optional] 
**Ts** | Pointer to **time.Time** | The day&#39;s timestamp | [optional] 

## Methods

### NewInlineResponse20092

`func NewInlineResponse20092() *InlineResponse20092`

NewInlineResponse20092 instantiates a new InlineResponse20092 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20092WithDefaults

`func NewInlineResponse20092WithDefaults() *InlineResponse20092`

NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *InlineResponse20092) GetReceived() float32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InlineResponse20092) GetReceivedOk() (*float32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InlineResponse20092) SetReceived(v float32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InlineResponse20092) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSent

`func (o *InlineResponse20092) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse20092) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse20092) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse20092) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse20092) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse20092) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse20092) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse20092) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


