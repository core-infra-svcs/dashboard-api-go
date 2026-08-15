# InlineResponse200400Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current status of the network move operation. Possible values are pending, in progress, failed, completed, and invalidated | [optional] 
**Reason** | Pointer to **string** | More information about the status of the network move operation | [optional] 

## Methods

### NewInlineResponse200400Result

`func NewInlineResponse200400Result() *InlineResponse200400Result`

NewInlineResponse200400Result instantiates a new InlineResponse200400Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200400ResultWithDefaults

`func NewInlineResponse200400ResultWithDefaults() *InlineResponse200400Result`

NewInlineResponse200400ResultWithDefaults instantiates a new InlineResponse200400Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200400Result) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200400Result) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200400Result) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200400Result) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *InlineResponse200400Result) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200400Result) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200400Result) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200400Result) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


