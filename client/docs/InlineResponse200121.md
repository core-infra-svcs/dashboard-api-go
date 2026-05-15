# InlineResponse200121

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | Pointer to **[]string** | The serials of the devices | [optional] 
**Errors** | Pointer to [**[]InlineResponse200121Errors**](InlineResponse200121Errors.md) | Errors for devices that were not added | [optional] 

## Methods

### NewInlineResponse200121

`func NewInlineResponse200121() *InlineResponse200121`

NewInlineResponse200121 instantiates a new InlineResponse200121 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200121WithDefaults

`func NewInlineResponse200121WithDefaults() *InlineResponse200121`

NewInlineResponse200121WithDefaults instantiates a new InlineResponse200121 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineResponse200121) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200121) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200121) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200121) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200121) GetErrors() []InlineResponse200121Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200121) GetErrorsOk() (*[]InlineResponse200121Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200121) SetErrors(v []InlineResponse200121Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200121) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


