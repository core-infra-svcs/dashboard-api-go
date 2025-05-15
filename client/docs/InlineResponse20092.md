# InlineResponse20092

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | Pointer to **[]string** | The serials of the devices | [optional] 
**Errors** | Pointer to [**[]InlineResponse20092Errors**](InlineResponse20092Errors.md) | Errors for devices that were not added | [optional] 

## Methods

### NewInlineResponse20092

`func NewInlineResponse20092() *InlineResponse20092`

NewInlineResponse20092 instantiates a new InlineResponse20092 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20092WithDefaults

`func NewInlineResponse20092WithDefaults() *InlineResponse20092`

NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineResponse20092) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse20092) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse20092) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse20092) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20092) GetErrors() []InlineResponse20092Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20092) GetErrorsOk() (*[]InlineResponse20092Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20092) SetErrors(v []InlineResponse20092Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20092) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


