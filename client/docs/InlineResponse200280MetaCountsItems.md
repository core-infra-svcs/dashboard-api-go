# InlineResponse200280MetaCountsItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The non-negative total number of assignments in the filtered dataset. Maximum is bounded by the number of assignments accessible to the organization. | [optional] 
**Remaining** | Pointer to **int32** | The non-negative number of assignments available on subsequent pages; this value is less than or equal to total | [optional] 

## Methods

### NewInlineResponse200280MetaCountsItems

`func NewInlineResponse200280MetaCountsItems() *InlineResponse200280MetaCountsItems`

NewInlineResponse200280MetaCountsItems instantiates a new InlineResponse200280MetaCountsItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200280MetaCountsItemsWithDefaults

`func NewInlineResponse200280MetaCountsItemsWithDefaults() *InlineResponse200280MetaCountsItems`

NewInlineResponse200280MetaCountsItemsWithDefaults instantiates a new InlineResponse200280MetaCountsItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200280MetaCountsItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200280MetaCountsItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200280MetaCountsItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200280MetaCountsItems) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRemaining

`func (o *InlineResponse200280MetaCountsItems) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *InlineResponse200280MetaCountsItems) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *InlineResponse200280MetaCountsItems) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *InlineResponse200280MetaCountsItems) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


