# InlineObject147

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency code of this node group&#39;s billing plans | [optional] 
**Plans** | Pointer to [**[]NetworksNetworkIdWirelessBillingPlans**](NetworksNetworkIdWirelessBillingPlans.md) | Array of billing plans in the node group. (Can configure a maximum of 5) | [optional] 

## Methods

### NewInlineObject147

`func NewInlineObject147() *InlineObject147`

NewInlineObject147 instantiates a new InlineObject147 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject147WithDefaults

`func NewInlineObject147WithDefaults() *InlineObject147`

NewInlineObject147WithDefaults instantiates a new InlineObject147 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InlineObject147) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineObject147) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineObject147) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineObject147) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *InlineObject147) GetPlans() []NetworksNetworkIdWirelessBillingPlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InlineObject147) GetPlansOk() (*[]NetworksNetworkIdWirelessBillingPlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InlineObject147) SetPlans(v []NetworksNetworkIdWirelessBillingPlans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InlineObject147) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


