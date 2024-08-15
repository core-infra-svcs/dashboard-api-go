# InlineResponse20086

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | Pointer to **[]string** | The serials of the devices | [optional] 
**Errors** | Pointer to [**[]InlineResponse20086Errors**](InlineResponse20086Errors.md) | Errors for devices that were not added | [optional] 

## Methods

### NewInlineResponse20086

`func NewInlineResponse20086() *InlineResponse20086`

NewInlineResponse20086 instantiates a new InlineResponse20086 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20086WithDefaults

`func NewInlineResponse20086WithDefaults() *InlineResponse20086`

NewInlineResponse20086WithDefaults instantiates a new InlineResponse20086 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineResponse20086) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse20086) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse20086) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse20086) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20086) GetErrors() []InlineResponse20086Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20086) GetErrorsOk() (*[]InlineResponse20086Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20086) SetErrors(v []InlineResponse20086Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20086) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


