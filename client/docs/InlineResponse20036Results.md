# InlineResponse20036Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** | Number of packets sent | [optional] 
**Received** | Pointer to **int32** | Number of packets received | [optional] 
**Loss** | Pointer to [**InlineResponse20036ResultsLoss**](InlineResponse20036ResultsLoss.md) |  | [optional] 
**Latencies** | Pointer to [**InlineResponse20036ResultsLatencies**](InlineResponse20036ResultsLatencies.md) |  | [optional] 
**Replies** | Pointer to [**[]InlineResponse20036ResultsReplies**](InlineResponse20036ResultsReplies.md) | Received packets | [optional] 

## Methods

### NewInlineResponse20036Results

`func NewInlineResponse20036Results() *InlineResponse20036Results`

NewInlineResponse20036Results instantiates a new InlineResponse20036Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20036ResultsWithDefaults

`func NewInlineResponse20036ResultsWithDefaults() *InlineResponse20036Results`

NewInlineResponse20036ResultsWithDefaults instantiates a new InlineResponse20036Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *InlineResponse20036Results) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse20036Results) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse20036Results) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse20036Results) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetReceived

`func (o *InlineResponse20036Results) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *InlineResponse20036Results) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *InlineResponse20036Results) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *InlineResponse20036Results) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetLoss

`func (o *InlineResponse20036Results) GetLoss() InlineResponse20036ResultsLoss`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *InlineResponse20036Results) GetLossOk() (*InlineResponse20036ResultsLoss, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *InlineResponse20036Results) SetLoss(v InlineResponse20036ResultsLoss)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *InlineResponse20036Results) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetLatencies

`func (o *InlineResponse20036Results) GetLatencies() InlineResponse20036ResultsLatencies`

GetLatencies returns the Latencies field if non-nil, zero value otherwise.

### GetLatenciesOk

`func (o *InlineResponse20036Results) GetLatenciesOk() (*InlineResponse20036ResultsLatencies, bool)`

GetLatenciesOk returns a tuple with the Latencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencies

`func (o *InlineResponse20036Results) SetLatencies(v InlineResponse20036ResultsLatencies)`

SetLatencies sets Latencies field to given value.

### HasLatencies

`func (o *InlineResponse20036Results) HasLatencies() bool`

HasLatencies returns a boolean if a field has been set.

### GetReplies

`func (o *InlineResponse20036Results) GetReplies() []InlineResponse20036ResultsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *InlineResponse20036Results) GetRepliesOk() (*[]InlineResponse20036ResultsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *InlineResponse20036Results) SetReplies(v []InlineResponse20036ResultsReplies)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *InlineResponse20036Results) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


