# InlineResponse200404Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current status of the network move operation. Possible values are pending, in progress, failed, completed, and invalidated | [optional] 
**Reason** | Pointer to **string** | More information about the status of the network move operation | [optional] 

## Methods

### NewInlineResponse200404Result

`func NewInlineResponse200404Result() *InlineResponse200404Result`

NewInlineResponse200404Result instantiates a new InlineResponse200404Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200404ResultWithDefaults

`func NewInlineResponse200404ResultWithDefaults() *InlineResponse200404Result`

NewInlineResponse200404ResultWithDefaults instantiates a new InlineResponse200404Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200404Result) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200404Result) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200404Result) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200404Result) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *InlineResponse200404Result) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200404Result) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200404Result) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200404Result) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


