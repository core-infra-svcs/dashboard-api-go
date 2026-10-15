# InlineResponse2027CountsByJobOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Pipeline job operation display name. Values are registered operation names and may grow additively. | [optional] 
**Total** | Pointer to **int32** | Total jobs for this operation | [optional] 
**ByStatus** | Pointer to [**InlineResponse200274CountsJobsByStatus**](InlineResponse200274CountsJobsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse2027CountsByJobOperation

`func NewInlineResponse2027CountsByJobOperation() *InlineResponse2027CountsByJobOperation`

NewInlineResponse2027CountsByJobOperation instantiates a new InlineResponse2027CountsByJobOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2027CountsByJobOperationWithDefaults

`func NewInlineResponse2027CountsByJobOperationWithDefaults() *InlineResponse2027CountsByJobOperation`

NewInlineResponse2027CountsByJobOperationWithDefaults instantiates a new InlineResponse2027CountsByJobOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2027CountsByJobOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2027CountsByJobOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2027CountsByJobOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2027CountsByJobOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse2027CountsByJobOperation) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse2027CountsByJobOperation) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse2027CountsByJobOperation) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse2027CountsByJobOperation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse2027CountsByJobOperation) GetByStatus() InlineResponse200274CountsJobsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse2027CountsByJobOperation) GetByStatusOk() (*InlineResponse200274CountsJobsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse2027CountsByJobOperation) SetByStatus(v InlineResponse200274CountsJobsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse2027CountsByJobOperation) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


