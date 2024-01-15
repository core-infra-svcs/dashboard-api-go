# InlineResponse200161

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainderLicenses** | Pointer to [**[]InlineResponse200160**](InlineResponse200160.md) | Remainder licenses created in the source organization as a result of moving a subset of the counts of a license | [optional] 
**MovedLicenses** | Pointer to [**[]InlineResponse200160**](InlineResponse200160.md) | Newly moved licenses created in the destination organization of the license move operation | [optional] 

## Methods

### NewInlineResponse200161

`func NewInlineResponse200161() *InlineResponse200161`

NewInlineResponse200161 instantiates a new InlineResponse200161 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200161WithDefaults

`func NewInlineResponse200161WithDefaults() *InlineResponse200161`

NewInlineResponse200161WithDefaults instantiates a new InlineResponse200161 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainderLicenses

`func (o *InlineResponse200161) GetRemainderLicenses() []InlineResponse200160`

GetRemainderLicenses returns the RemainderLicenses field if non-nil, zero value otherwise.

### GetRemainderLicensesOk

`func (o *InlineResponse200161) GetRemainderLicensesOk() (*[]InlineResponse200160, bool)`

GetRemainderLicensesOk returns a tuple with the RemainderLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainderLicenses

`func (o *InlineResponse200161) SetRemainderLicenses(v []InlineResponse200160)`

SetRemainderLicenses sets RemainderLicenses field to given value.

### HasRemainderLicenses

`func (o *InlineResponse200161) HasRemainderLicenses() bool`

HasRemainderLicenses returns a boolean if a field has been set.

### GetMovedLicenses

`func (o *InlineResponse200161) GetMovedLicenses() []InlineResponse200160`

GetMovedLicenses returns the MovedLicenses field if non-nil, zero value otherwise.

### GetMovedLicensesOk

`func (o *InlineResponse200161) GetMovedLicensesOk() (*[]InlineResponse200160, bool)`

GetMovedLicensesOk returns a tuple with the MovedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedLicenses

`func (o *InlineResponse200161) SetMovedLicenses(v []InlineResponse200160)`

SetMovedLicenses sets MovedLicenses field to given value.

### HasMovedLicenses

`func (o *InlineResponse200161) HasMovedLicenses() bool`

HasMovedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


