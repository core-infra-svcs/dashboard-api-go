# InlineResponse2004Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**InlineResponse2004ResultsLoss**](InlineResponse2004ResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**InlineResponse2004ResultsLatencies**](InlineResponse2004ResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]InlineResponse2004ResultsReplies**](InlineResponse2004ResultsReplies.md) | Received packets | [optional] 

## Methods

### NewInlineResponse2004Results

`func NewInlineResponse2004Results() *InlineResponse2004Results`

NewInlineResponse2004Results instantiates a new InlineResponse2004Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004ResultsWithDefaults

`func NewInlineResponse2004ResultsWithDefaults() *InlineResponse2004Results`

NewInlineResponse2004ResultsWithDefaults instantiates a new InlineResponse2004Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *InlineResponse2004Results) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse2004Results) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse2004Results) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse2004Results) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *InlineResponse2004Results) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InlineResponse2004Results) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InlineResponse2004Results) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InlineResponse2004Results) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *InlineResponse2004Results) GetLoss() InlineResponse2004ResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *InlineResponse2004Results) GetLossOk() (*InlineResponse2004ResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *InlineResponse2004Results) SetLoss(v InlineResponse2004ResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *InlineResponse2004Results) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *InlineResponse2004Results) GetLatencies() InlineResponse2004ResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *InlineResponse2004Results) GetLatenciesOk() (*InlineResponse2004ResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *InlineResponse2004Results) SetLatencies(v InlineResponse2004ResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *InlineResponse2004Results) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *InlineResponse2004Results) GetReplies() []InlineResponse2004ResultsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *InlineResponse2004Results) GetRepliesOk() (*[]InlineResponse2004ResultsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *InlineResponse2004Results) SetReplies(v []InlineResponse2004ResultsReplies)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *InlineResponse2004Results) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


