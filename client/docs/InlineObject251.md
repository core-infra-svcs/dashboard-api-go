# InlineObject251

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | The numbers of the orders that should be claimed | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices that should be claimed | [optional] 
**Licenses** | Pointer to [**[]OrganizationsOrganizationIdClaimLicenses**](OrganizationsOrganizationIdClaimLicenses.md) | The licenses that should be claimed | [optional] 

## Methods

### NewInlineObject251

`func NewInlineObject251() *InlineObject251`

NewInlineObject251 instantiates a new InlineObject251 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject251WithDefaults

`func NewInlineObject251WithDefaults() *InlineObject251`

NewInlineObject251WithDefaults instantiates a new InlineObject251 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *InlineObject251) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *InlineObject251) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *InlineObject251) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *InlineObject251) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject251) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject251) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject251) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject251) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetLicenses

`func (o *InlineObject251) GetLicenses() []OrganizationsOrganizationIdClaimLicenses`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *InlineObject251) GetLicensesOk() (*[]OrganizationsOrganizationIdClaimLicenses, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *InlineObject251) SetLicenses(v []OrganizationsOrganizationIdClaimLicenses)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *InlineObject251) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


