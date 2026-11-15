# InlineResponse2025

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the background job | [optional] 
**Action** | Pointer to **string** | The operation being performed | [optional] 
**Status** | Pointer to **string** | Current status of the job | [optional] 
**Message** | Pointer to **string** | Job message or error details | [optional] 
**Result** | Pointer to [**InlineResponse2025Result**](InlineResponse2025Result.md) |  | [optional] 

## Methods

### NewInlineResponse2025

`func NewInlineResponse2025() *InlineResponse2025`

NewInlineResponse2025 instantiates a new InlineResponse2025 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2025WithDefaults

`func NewInlineResponse2025WithDefaults() *InlineResponse2025`

NewInlineResponse2025WithDefaults instantiates a new InlineResponse2025 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2025) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2025) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2025) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2025) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *InlineResponse2025) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InlineResponse2025) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InlineResponse2025) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InlineResponse2025) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2025) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2025) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2025) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2025) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *InlineResponse2025) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse2025) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse2025) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse2025) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse2025) GetResult() InlineResponse2025Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse2025) GetResultOk() (*InlineResponse2025Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse2025) SetResult(v InlineResponse2025Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse2025) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


