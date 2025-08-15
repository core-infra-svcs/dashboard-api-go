# InlineResponse200263

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices claimed | [optional] 
**Licenses** | Pointer to [**[]InlineResponse200263Licenses**](InlineResponse200263Licenses.md) | The licenses claimed | [optional] 

## Methods

### NewInlineResponse200263

`func NewInlineResponse200263() *InlineResponse200263`

NewInlineResponse200263 instantiates a new InlineResponse200263 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200263WithDefaults

`func NewInlineResponse200263WithDefaults() *InlineResponse200263`

NewInlineResponse200263WithDefaults instantiates a new InlineResponse200263 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *InlineResponse200263) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *InlineResponse200263) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *InlineResponse200263) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *InlineResponse200263) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200263) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200263) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200263) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200263) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *InlineResponse200263) GetLicenses() []InlineResponse200263Licenses`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *InlineResponse200263) GetLicensesOk() (*[]InlineResponse200263Licenses, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *InlineResponse200263) SetLicenses(v []InlineResponse200263Licenses)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *InlineResponse200263) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


