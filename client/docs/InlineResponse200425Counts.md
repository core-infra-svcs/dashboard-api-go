# InlineResponse200425Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByStatus** | Pointer to [**InlineResponse200425CountsByStatus**](InlineResponse200425CountsByStatus.md) |  | [optional] 
**Total** | Pointer to **int32** | Total number of enrolled SSE sites | [optional] 

## Methods

### NewInlineResponse200425Counts

`func NewInlineResponse200425Counts() *InlineResponse200425Counts`

NewInlineResponse200425Counts instantiates a new InlineResponse200425Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200425CountsWithDefaults

`func NewInlineResponse200425CountsWithDefaults() *InlineResponse200425Counts`

NewInlineResponse200425CountsWithDefaults instantiates a new InlineResponse200425Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByStatus

`func (o *InlineResponse200425Counts) GetByStatus() InlineResponse200425CountsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200425Counts) GetByStatusOk() (*InlineResponse200425CountsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200425Counts) SetByStatus(v InlineResponse200425CountsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200425Counts) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200425Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200425Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200425Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200425Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


