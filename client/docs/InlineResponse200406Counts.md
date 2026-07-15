# InlineResponse200406Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByStatus** | Pointer to [**InlineResponse200406CountsByStatus**](InlineResponse200406CountsByStatus.md) |  | [optional] 
**Total** | Pointer to **int32** | Total number of enrolled SSE sites | [optional] 

## Methods

### NewInlineResponse200406Counts

`func NewInlineResponse200406Counts() *InlineResponse200406Counts`

NewInlineResponse200406Counts instantiates a new InlineResponse200406Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200406CountsWithDefaults

`func NewInlineResponse200406CountsWithDefaults() *InlineResponse200406Counts`

NewInlineResponse200406CountsWithDefaults instantiates a new InlineResponse200406Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByStatus

`func (o *InlineResponse200406Counts) GetByStatus() InlineResponse200406CountsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200406Counts) GetByStatusOk() (*InlineResponse200406CountsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200406Counts) SetByStatus(v InlineResponse200406CountsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200406Counts) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200406Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200406Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200406Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200406Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


