# InlineResponse2028CountsByJobOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Pipeline job operation display name. Values are registered operation names and may grow additively. | [optional] 
**Total** | Pointer to **int32** | Total jobs for this operation | [optional] 
**ByStatus** | Pointer to [**InlineResponse200277CountsJobsByStatus**](InlineResponse200277CountsJobsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse2028CountsByJobOperation

`func NewInlineResponse2028CountsByJobOperation() *InlineResponse2028CountsByJobOperation`

NewInlineResponse2028CountsByJobOperation instantiates a new InlineResponse2028CountsByJobOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2028CountsByJobOperationWithDefaults

`func NewInlineResponse2028CountsByJobOperationWithDefaults() *InlineResponse2028CountsByJobOperation`

NewInlineResponse2028CountsByJobOperationWithDefaults instantiates a new InlineResponse2028CountsByJobOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2028CountsByJobOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2028CountsByJobOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2028CountsByJobOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2028CountsByJobOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse2028CountsByJobOperation) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse2028CountsByJobOperation) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse2028CountsByJobOperation) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse2028CountsByJobOperation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse2028CountsByJobOperation) GetByStatus() InlineResponse200277CountsJobsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse2028CountsByJobOperation) GetByStatusOk() (*InlineResponse200277CountsJobsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse2028CountsByJobOperation) SetByStatus(v InlineResponse200277CountsJobsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse2028CountsByJobOperation) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


