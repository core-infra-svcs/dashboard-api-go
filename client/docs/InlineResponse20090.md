# InlineResponse20090

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeWindow** | Pointer to [**InlineResponse20090UpgradeWindow**](InlineResponse20090UpgradeWindow.md) |  | [optional] 
**Timezone** | Pointer to **string** | The timezone for the network | [optional] 
**Products** | Pointer to [**InlineResponse20090Products**](InlineResponse20090Products.md) |  | [optional] 

## Methods

### NewInlineResponse20090

`func NewInlineResponse20090() *InlineResponse20090`

NewInlineResponse20090 instantiates a new InlineResponse20090 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20090WithDefaults

`func NewInlineResponse20090WithDefaults() *InlineResponse20090`

NewInlineResponse20090WithDefaults instantiates a new InlineResponse20090 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeWindow

`func (o *InlineResponse20090) GetUpgradeWindow() InlineResponse20090UpgradeWindow`

GetUpgradeWindow returns the UpgradeWindow field if non-nil, zero value otherwise.

### GetUpgradeWindowOk

`func (o *InlineResponse20090) GetUpgradeWindowOk() (*InlineResponse20090UpgradeWindow, bool)`

GetUpgradeWindowOk returns a tuple with the UpgradeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeWindow

`func (o *InlineResponse20090) SetUpgradeWindow(v InlineResponse20090UpgradeWindow)`

SetUpgradeWindow sets UpgradeWindow field to given value.

### HasUpgradeWindow

`func (o *InlineResponse20090) HasUpgradeWindow() bool`

HasUpgradeWindow returns a boolean if a field has been set.

### GetTimezone

`func (o *InlineResponse20090) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InlineResponse20090) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InlineResponse20090) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InlineResponse20090) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetProducts

`func (o *InlineResponse20090) GetProducts() InlineResponse20090Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20090) GetProductsOk() (*InlineResponse20090Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20090) SetProducts(v InlineResponse20090Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20090) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


