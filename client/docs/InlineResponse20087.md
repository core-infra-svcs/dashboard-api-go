# InlineResponse20087

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | Pointer to **[]string** | The serials of the devices | [optional] 
**Errors** | Pointer to [**[]InlineResponse20087Errors**](InlineResponse20087Errors.md) | Errors for devices that were not added | [optional] 

## Methods

### NewInlineResponse20087

`func NewInlineResponse20087() *InlineResponse20087`

NewInlineResponse20087 instantiates a new InlineResponse20087 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20087WithDefaults

`func NewInlineResponse20087WithDefaults() *InlineResponse20087`

NewInlineResponse20087WithDefaults instantiates a new InlineResponse20087 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineResponse20087) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse20087) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse20087) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse20087) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20087) GetErrors() []InlineResponse20087Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20087) GetErrorsOk() (*[]InlineResponse20087Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20087) SetErrors(v []InlineResponse20087Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20087) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


