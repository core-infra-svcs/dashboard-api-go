# InlineResponse200484FailoverLast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Wireless LAN controller last redundancy switchover time | [optional] 
**Reason** | Pointer to **string** | Wireless LAN controller last redundancy switchover reason | [optional] 

## Methods

### NewInlineResponse200484FailoverLast

`func NewInlineResponse200484FailoverLast() *InlineResponse200484FailoverLast`

NewInlineResponse200484FailoverLast instantiates a new InlineResponse200484FailoverLast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200484FailoverLastWithDefaults

`func NewInlineResponse200484FailoverLastWithDefaults() *InlineResponse200484FailoverLast`

NewInlineResponse200484FailoverLastWithDefaults instantiates a new InlineResponse200484FailoverLast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200484FailoverLast) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200484FailoverLast) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200484FailoverLast) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200484FailoverLast) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetReason

`func (o *InlineResponse200484FailoverLast) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200484FailoverLast) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200484FailoverLast) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200484FailoverLast) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


