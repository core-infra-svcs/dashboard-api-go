# InlineResponse200101

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainderLicenses** | Pointer to [**[]InlineResponse200100**](InlineResponse200100.md) | Remainder licenses created in the source organization as a result of moving a subset of the counts of a license | [optional] 
**MovedLicenses** | Pointer to [**[]InlineResponse200100**](InlineResponse200100.md) | Newly moved licenses created in the destination organization of the license move operation | [optional] 

## Methods

### NewInlineResponse200101

`func NewInlineResponse200101() *InlineResponse200101`

NewInlineResponse200101 instantiates a new InlineResponse200101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101WithDefaults

`func NewInlineResponse200101WithDefaults() *InlineResponse200101`

NewInlineResponse200101WithDefaults instantiates a new InlineResponse200101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainderLicenses

`func (o *InlineResponse200101) GetRemainderLicenses() []InlineResponse200100`

GetRemainderLicenses returns the RemainderLicenses field if non-nil, zero value otherwise.

### GetRemainderLicensesOk

`func (o *InlineResponse200101) GetRemainderLicensesOk() (*[]InlineResponse200100, bool)`

GetRemainderLicensesOk returns a tuple with the RemainderLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderLicenses

`func (o *InlineResponse200101) SetRemainderLicenses(v []InlineResponse200100)`

SetRemainderLicenses sets RemainderLicenses field to given value.

### HasRemainderLicenses

`func (o *InlineResponse200101) HasRemainderLicenses() bool`

HasRemainderLicenses returns a boolean if a field has been set.

### GetMovedLicenses

`func (o *InlineResponse200101) GetMovedLicenses() []InlineResponse200100`

GetMovedLicenses returns the MovedLicenses field if non-nil, zero value otherwise.

### GetMovedLicensesOk

`func (o *InlineResponse200101) GetMovedLicensesOk() (*[]InlineResponse200100, bool)`

GetMovedLicensesOk returns a tuple with the MovedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedLicenses

`func (o *InlineResponse200101) SetMovedLicenses(v []InlineResponse200100)`

SetMovedLicenses sets MovedLicenses field to given value.

### HasMovedLicenses

`func (o *InlineResponse200101) HasMovedLicenses() bool`

HasMovedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


