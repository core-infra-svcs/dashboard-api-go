# InlineResponse200325Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of ports | [optional] 
**ByStatus** | Pointer to [**InlineResponse200325CountsByStatus**](InlineResponse200325CountsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse200325Counts

`func NewInlineResponse200325Counts() *InlineResponse200325Counts`

NewInlineResponse200325Counts instantiates a new InlineResponse200325Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200325CountsWithDefaults

`func NewInlineResponse200325CountsWithDefaults() *InlineResponse200325Counts`

NewInlineResponse200325CountsWithDefaults instantiates a new InlineResponse200325Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200325Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200325Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200325Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200325Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse200325Counts) GetByStatus() InlineResponse200325CountsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200325Counts) GetByStatusOk() (*InlineResponse200325CountsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200325Counts) SetByStatus(v InlineResponse200325CountsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200325Counts) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


