# InlineResponse200293

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainderLicenses** | Pointer to [**[]InlineResponse200292**](InlineResponse200292.md) | Remainder licenses created in the source organization as a result of moving a subset of the counts of a license | [optional] 
**MovedLicenses** | Pointer to [**[]InlineResponse200292**](InlineResponse200292.md) | Newly moved licenses created in the destination organization of the license move operation | [optional] 

## Methods

### NewInlineResponse200293

`func NewInlineResponse200293() *InlineResponse200293`

NewInlineResponse200293 instantiates a new InlineResponse200293 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200293WithDefaults

`func NewInlineResponse200293WithDefaults() *InlineResponse200293`

NewInlineResponse200293WithDefaults instantiates a new InlineResponse200293 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainderLicenses

`func (o *InlineResponse200293) GetRemainderLicenses() []InlineResponse200292`

GetRemainderLicenses returns the RemainderLicenses field if non-nil, zero value otherwise.

### GetRemainderLicensesOk

`func (o *InlineResponse200293) GetRemainderLicensesOk() (*[]InlineResponse200292, bool)`

GetRemainderLicensesOk returns a tuple with the RemainderLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderLicenses

`func (o *InlineResponse200293) SetRemainderLicenses(v []InlineResponse200292)`

SetRemainderLicenses sets RemainderLicenses field to given value.

### HasRemainderLicenses

`func (o *InlineResponse200293) HasRemainderLicenses() bool`

HasRemainderLicenses returns a boolean if a field has been set.

### GetMovedLicenses

`func (o *InlineResponse200293) GetMovedLicenses() []InlineResponse200292`

GetMovedLicenses returns the MovedLicenses field if non-nil, zero value otherwise.

### GetMovedLicensesOk

`func (o *InlineResponse200293) GetMovedLicensesOk() (*[]InlineResponse200292, bool)`

GetMovedLicensesOk returns a tuple with the MovedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedLicenses

`func (o *InlineResponse200293) SetMovedLicenses(v []InlineResponse200292)`

SetMovedLicenses sets MovedLicenses field to given value.

### HasMovedLicenses

`func (o *InlineResponse200293) HasMovedLicenses() bool`

HasMovedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


