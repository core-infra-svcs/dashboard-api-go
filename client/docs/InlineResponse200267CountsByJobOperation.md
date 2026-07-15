# InlineResponse200267CountsByJobOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Pipeline job operation display name. Values are registered operation names and may grow additively. | [optional] 
**Total** | Pointer to **int32** | Total jobs for this operation | [optional] 
**ByStatus** | Pointer to [**InlineResponse200267CountsJobsByStatus**](InlineResponse200267CountsJobsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse200267CountsByJobOperation

`func NewInlineResponse200267CountsByJobOperation() *InlineResponse200267CountsByJobOperation`

NewInlineResponse200267CountsByJobOperation instantiates a new InlineResponse200267CountsByJobOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200267CountsByJobOperationWithDefaults

`func NewInlineResponse200267CountsByJobOperationWithDefaults() *InlineResponse200267CountsByJobOperation`

NewInlineResponse200267CountsByJobOperationWithDefaults instantiates a new InlineResponse200267CountsByJobOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200267CountsByJobOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200267CountsByJobOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200267CountsByJobOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200267CountsByJobOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200267CountsByJobOperation) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200267CountsByJobOperation) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200267CountsByJobOperation) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200267CountsByJobOperation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse200267CountsByJobOperation) GetByStatus() InlineResponse200267CountsJobsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200267CountsByJobOperation) GetByStatusOk() (*InlineResponse200267CountsJobsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200267CountsByJobOperation) SetByStatus(v InlineResponse200267CountsJobsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200267CountsByJobOperation) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


