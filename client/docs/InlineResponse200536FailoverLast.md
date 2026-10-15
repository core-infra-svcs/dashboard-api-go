# InlineResponse200536FailoverLast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Wireless LAN controller last redundancy switchover time | [optional] 
**Reason** | Pointer to **string** | Wireless LAN controller last redundancy switchover reason | [optional] 

## Methods

### NewInlineResponse200536FailoverLast

`func NewInlineResponse200536FailoverLast() *InlineResponse200536FailoverLast`

NewInlineResponse200536FailoverLast instantiates a new InlineResponse200536FailoverLast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200536FailoverLastWithDefaults

`func NewInlineResponse200536FailoverLastWithDefaults() *InlineResponse200536FailoverLast`

NewInlineResponse200536FailoverLastWithDefaults instantiates a new InlineResponse200536FailoverLast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200536FailoverLast) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200536FailoverLast) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200536FailoverLast) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200536FailoverLast) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetReason

`func (o *InlineResponse200536FailoverLast) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200536FailoverLast) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200536FailoverLast) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200536FailoverLast) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


