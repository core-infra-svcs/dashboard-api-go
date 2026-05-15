# InlineResponse200403

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | The success or failure status of the API | [optional] 
**Message** | Pointer to **string** | The error message of the API | [optional] 

## Methods

### NewInlineResponse200403

`func NewInlineResponse200403() *InlineResponse200403`

NewInlineResponse200403 instantiates a new InlineResponse200403 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200403WithDefaults

`func NewInlineResponse200403WithDefaults() *InlineResponse200403`

NewInlineResponse200403WithDefaults instantiates a new InlineResponse200403 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200403) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200403) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200403) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200403) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *InlineResponse200403) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse200403) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse200403) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse200403) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


