# InlineResponse200261

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainderLicenses** | Pointer to [**[]InlineResponse200260**](InlineResponse200260.md) | Remainder licenses created in the source organization as a result of moving a subset of the counts of a license | [optional] 
**MovedLicenses** | Pointer to [**[]InlineResponse200260**](InlineResponse200260.md) | Newly moved licenses created in the destination organization of the license move operation | [optional] 

## Methods

### NewInlineResponse200261

`func NewInlineResponse200261() *InlineResponse200261`

NewInlineResponse200261 instantiates a new InlineResponse200261 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200261WithDefaults

`func NewInlineResponse200261WithDefaults() *InlineResponse200261`

NewInlineResponse200261WithDefaults instantiates a new InlineResponse200261 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainderLicenses

`func (o *InlineResponse200261) GetRemainderLicenses() []InlineResponse200260`

GetRemainderLicenses returns the RemainderLicenses field if non-nil, zero value otherwise.

### GetRemainderLicensesOk

`func (o *InlineResponse200261) GetRemainderLicensesOk() (*[]InlineResponse200260, bool)`

GetRemainderLicensesOk returns a tuple with the RemainderLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderLicenses

`func (o *InlineResponse200261) SetRemainderLicenses(v []InlineResponse200260)`

SetRemainderLicenses sets RemainderLicenses field to given value.

### HasRemainderLicenses

`func (o *InlineResponse200261) HasRemainderLicenses() bool`

HasRemainderLicenses returns a boolean if a field has been set.

### GetMovedLicenses

`func (o *InlineResponse200261) GetMovedLicenses() []InlineResponse200260`

GetMovedLicenses returns the MovedLicenses field if non-nil, zero value otherwise.

### GetMovedLicensesOk

`func (o *InlineResponse200261) GetMovedLicensesOk() (*[]InlineResponse200260, bool)`

GetMovedLicensesOk returns a tuple with the MovedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedLicenses

`func (o *InlineResponse200261) SetMovedLicenses(v []InlineResponse200260)`

SetMovedLicenses sets MovedLicenses field to given value.

### HasMovedLicenses

`func (o *InlineResponse200261) HasMovedLicenses() bool`

HasMovedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


