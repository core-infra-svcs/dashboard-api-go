# InlineResponse200320Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of impacted metric | [optional] 
**ByFailureType** | Pointer to [**[]InlineResponse200320CountsByFailureType**](InlineResponse200320CountsByFailureType.md) | List of failure types | [optional] 

## Methods

### NewInlineResponse200320Counts

`func NewInlineResponse200320Counts() *InlineResponse200320Counts`

NewInlineResponse200320Counts instantiates a new InlineResponse200320Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200320CountsWithDefaults

`func NewInlineResponse200320CountsWithDefaults() *InlineResponse200320Counts`

NewInlineResponse200320CountsWithDefaults instantiates a new InlineResponse200320Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200320Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200320Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200320Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200320Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByFailureType

`func (o *InlineResponse200320Counts) GetByFailureType() []InlineResponse200320CountsByFailureType`

GetByFailureType returns the ByFailureType field if non-nil, zero value otherwise.

### GetByFailureTypeOk

`func (o *InlineResponse200320Counts) GetByFailureTypeOk() (*[]InlineResponse200320CountsByFailureType, bool)`

GetByFailureTypeOk returns a tuple with the ByFailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByFailureType

`func (o *InlineResponse200320Counts) SetByFailureType(v []InlineResponse200320CountsByFailureType)`

SetByFailureType sets ByFailureType field to given value.

### HasByFailureType

`func (o *InlineResponse200320Counts) HasByFailureType() bool`

HasByFailureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


