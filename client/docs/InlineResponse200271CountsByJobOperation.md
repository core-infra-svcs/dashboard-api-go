# InlineResponse200271CountsByJobOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Pipeline job operation display name. Values are registered operation names and may grow additively. | [optional] 
**Total** | Pointer to **int32** | Total jobs for this operation | [optional] 
**ByStatus** | Pointer to [**InlineResponse200271CountsJobsByStatus**](InlineResponse200271CountsJobsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse200271CountsByJobOperation

`func NewInlineResponse200271CountsByJobOperation() *InlineResponse200271CountsByJobOperation`

NewInlineResponse200271CountsByJobOperation instantiates a new InlineResponse200271CountsByJobOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200271CountsByJobOperationWithDefaults

`func NewInlineResponse200271CountsByJobOperationWithDefaults() *InlineResponse200271CountsByJobOperation`

NewInlineResponse200271CountsByJobOperationWithDefaults instantiates a new InlineResponse200271CountsByJobOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200271CountsByJobOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200271CountsByJobOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200271CountsByJobOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200271CountsByJobOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200271CountsByJobOperation) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200271CountsByJobOperation) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200271CountsByJobOperation) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200271CountsByJobOperation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse200271CountsByJobOperation) GetByStatus() InlineResponse200271CountsJobsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200271CountsByJobOperation) GetByStatusOk() (*InlineResponse200271CountsJobsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200271CountsByJobOperation) SetByStatus(v InlineResponse200271CountsJobsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200271CountsByJobOperation) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


