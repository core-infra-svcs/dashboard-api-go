# InlineResponse200435Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of ports | [optional] 
**ByStatus** | Pointer to [**InlineResponse200435CountsByStatus**](InlineResponse200435CountsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse200435Counts

`func NewInlineResponse200435Counts() *InlineResponse200435Counts`

NewInlineResponse200435Counts instantiates a new InlineResponse200435Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200435CountsWithDefaults

`func NewInlineResponse200435CountsWithDefaults() *InlineResponse200435Counts`

NewInlineResponse200435CountsWithDefaults instantiates a new InlineResponse200435Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200435Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200435Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200435Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200435Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse200435Counts) GetByStatus() InlineResponse200435CountsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200435Counts) GetByStatusOk() (*InlineResponse200435CountsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200435Counts) SetByStatus(v InlineResponse200435CountsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200435Counts) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


