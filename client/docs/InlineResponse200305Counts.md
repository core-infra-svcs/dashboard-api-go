# InlineResponse200305Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of impacted metric | [optional] 
**ByFailureType** | Pointer to [**[]InlineResponse200305CountsByFailureType**](InlineResponse200305CountsByFailureType.md) | List of failure types | [optional] 

## Methods

### NewInlineResponse200305Counts

`func NewInlineResponse200305Counts() *InlineResponse200305Counts`

NewInlineResponse200305Counts instantiates a new InlineResponse200305Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200305CountsWithDefaults

`func NewInlineResponse200305CountsWithDefaults() *InlineResponse200305Counts`

NewInlineResponse200305CountsWithDefaults instantiates a new InlineResponse200305Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200305Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200305Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200305Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200305Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByFailureType

`func (o *InlineResponse200305Counts) GetByFailureType() []InlineResponse200305CountsByFailureType`

GetByFailureType returns the ByFailureType field if non-nil, zero value otherwise.

### GetByFailureTypeOk

`func (o *InlineResponse200305Counts) GetByFailureTypeOk() (*[]InlineResponse200305CountsByFailureType, bool)`

GetByFailureTypeOk returns a tuple with the ByFailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByFailureType

`func (o *InlineResponse200305Counts) SetByFailureType(v []InlineResponse200305CountsByFailureType)`

SetByFailureType sets ByFailureType field to given value.

### HasByFailureType

`func (o *InlineResponse200305Counts) HasByFailureType() bool`

HasByFailureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


