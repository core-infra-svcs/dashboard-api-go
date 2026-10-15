# InlineResponse200351

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices claimed | [optional] 
**Licenses** | Pointer to [**[]InlineResponse200351Licenses**](InlineResponse200351Licenses.md) | The licenses claimed | [optional] 

## Methods

### NewInlineResponse200351

`func NewInlineResponse200351() *InlineResponse200351`

NewInlineResponse200351 instantiates a new InlineResponse200351 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200351WithDefaults

`func NewInlineResponse200351WithDefaults() *InlineResponse200351`

NewInlineResponse200351WithDefaults instantiates a new InlineResponse200351 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *InlineResponse200351) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *InlineResponse200351) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *InlineResponse200351) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *InlineResponse200351) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200351) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200351) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200351) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200351) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *InlineResponse200351) GetLicenses() []InlineResponse200351Licenses`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *InlineResponse200351) GetLicensesOk() (*[]InlineResponse200351Licenses, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *InlineResponse200351) SetLicenses(v []InlineResponse200351Licenses)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *InlineResponse200351) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


