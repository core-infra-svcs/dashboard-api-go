# InlineResponse200307

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainderLicenses** | Pointer to [**[]InlineResponse200306**](InlineResponse200306.md) | Remainder licenses created in the source organization as a result of moving a subset of the counts of a license | [optional] 
**MovedLicenses** | Pointer to [**[]InlineResponse200306**](InlineResponse200306.md) | Newly moved licenses created in the destination organization of the license move operation | [optional] 

## Methods

### NewInlineResponse200307

`func NewInlineResponse200307() *InlineResponse200307`

NewInlineResponse200307 instantiates a new InlineResponse200307 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200307WithDefaults

`func NewInlineResponse200307WithDefaults() *InlineResponse200307`

NewInlineResponse200307WithDefaults instantiates a new InlineResponse200307 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainderLicenses

`func (o *InlineResponse200307) GetRemainderLicenses() []InlineResponse200306`

GetRemainderLicenses returns the RemainderLicenses field if non-nil, zero value otherwise.

### GetRemainderLicensesOk

`func (o *InlineResponse200307) GetRemainderLicensesOk() (*[]InlineResponse200306, bool)`

GetRemainderLicensesOk returns a tuple with the RemainderLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderLicenses

`func (o *InlineResponse200307) SetRemainderLicenses(v []InlineResponse200306)`

SetRemainderLicenses sets RemainderLicenses field to given value.

### HasRemainderLicenses

`func (o *InlineResponse200307) HasRemainderLicenses() bool`

HasRemainderLicenses returns a boolean if a field has been set.

### GetMovedLicenses

`func (o *InlineResponse200307) GetMovedLicenses() []InlineResponse200306`

GetMovedLicenses returns the MovedLicenses field if non-nil, zero value otherwise.

### GetMovedLicensesOk

`func (o *InlineResponse200307) GetMovedLicensesOk() (*[]InlineResponse200306, bool)`

GetMovedLicensesOk returns a tuple with the MovedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedLicenses

`func (o *InlineResponse200307) SetMovedLicenses(v []InlineResponse200306)`

SetMovedLicenses sets MovedLicenses field to given value.

### HasMovedLicenses

`func (o *InlineResponse200307) HasMovedLicenses() bool`

HasMovedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


