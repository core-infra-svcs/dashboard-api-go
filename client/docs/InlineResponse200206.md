# InlineResponse200206

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the query range | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the query range | [optional] 
**TotalKbps** | Pointer to **int32** | Total usage in kilobytes-per-second | [optional] 
**SentKbps** | Pointer to **int32** | Sent kilobytes-per-second | [optional] 
**ReceivedKbps** | Pointer to **int32** | Received kilobytes-per-second | [optional] 

## Methods

### NewInlineResponse200206

`func NewInlineResponse200206() *InlineResponse200206`

NewInlineResponse200206 instantiates a new InlineResponse200206 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200206WithDefaults

`func NewInlineResponse200206WithDefaults() *InlineResponse200206`

NewInlineResponse200206WithDefaults instantiates a new InlineResponse200206 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200206) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200206) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200206) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200206) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200206) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200206) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200206) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200206) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetTotalKbps

`func (o *InlineResponse200206) GetTotalKbps() int32`

GetTotalKbps returns the TotalKbps field if non-nil, zero value otherwise.

### GetTotalKbpsOk

`func (o *InlineResponse200206) GetTotalKbpsOk() (*int32, bool)`

GetTotalKbpsOk returns a tuple with the TotalKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbps

`func (o *InlineResponse200206) SetTotalKbps(v int32)`

SetTotalKbps sets TotalKbps field to given value.

### HasTotalKbps

`func (o *InlineResponse200206) HasTotalKbps() bool`

HasTotalKbps returns a boolean if a field has been set.

### GetSentKbps

`func (o *InlineResponse200206) GetSentKbps() int32`

GetSentKbps returns the SentKbps field if non-nil, zero value otherwise.

### GetSentKbpsOk

`func (o *InlineResponse200206) GetSentKbpsOk() (*int32, bool)`

GetSentKbpsOk returns a tuple with the SentKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentKbps

`func (o *InlineResponse200206) SetSentKbps(v int32)`

SetSentKbps sets SentKbps field to given value.

### HasSentKbps

`func (o *InlineResponse200206) HasSentKbps() bool`

HasSentKbps returns a boolean if a field has been set.

### GetReceivedKbps

`func (o *InlineResponse200206) GetReceivedKbps() int32`

GetReceivedKbps returns the ReceivedKbps field if non-nil, zero value otherwise.

### GetReceivedKbpsOk

`func (o *InlineResponse200206) GetReceivedKbpsOk() (*int32, bool)`

GetReceivedKbpsOk returns a tuple with the ReceivedKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedKbps

`func (o *InlineResponse200206) SetReceivedKbps(v int32)`

SetReceivedKbps sets ReceivedKbps field to given value.

### HasReceivedKbps

`func (o *InlineResponse200206) HasReceivedKbps() bool`

HasReceivedKbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


