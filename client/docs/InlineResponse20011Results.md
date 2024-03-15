# InlineResponse20011Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**InlineResponse20011ResultsLoss**](InlineResponse20011ResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**InlineResponse20011ResultsLatencies**](InlineResponse20011ResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]InlineResponse20011ResultsReplies**](InlineResponse20011ResultsReplies.md) | Received packets | [optional] 

## Methods

### NewInlineResponse20011Results

`func NewInlineResponse20011Results() *InlineResponse20011Results`

NewInlineResponse20011Results instantiates a new InlineResponse20011Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011ResultsWithDefaults

`func NewInlineResponse20011ResultsWithDefaults() *InlineResponse20011Results`

NewInlineResponse20011ResultsWithDefaults instantiates a new InlineResponse20011Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *InlineResponse20011Results) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse20011Results) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse20011Results) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse20011Results) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *InlineResponse20011Results) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InlineResponse20011Results) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InlineResponse20011Results) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InlineResponse20011Results) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *InlineResponse20011Results) GetLoss() InlineResponse20011ResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *InlineResponse20011Results) GetLossOk() (*InlineResponse20011ResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *InlineResponse20011Results) SetLoss(v InlineResponse20011ResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *InlineResponse20011Results) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *InlineResponse20011Results) GetLatencies() InlineResponse20011ResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *InlineResponse20011Results) GetLatenciesOk() (*InlineResponse20011ResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *InlineResponse20011Results) SetLatencies(v InlineResponse20011ResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *InlineResponse20011Results) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *InlineResponse20011Results) GetReplies() []InlineResponse20011ResultsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *InlineResponse20011Results) GetRepliesOk() (*[]InlineResponse20011ResultsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *InlineResponse20011Results) SetReplies(v []InlineResponse20011ResultsReplies)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *InlineResponse20011Results) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


