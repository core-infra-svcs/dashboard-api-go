# InlineObject97

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeWindow** | Pointer to [**InlineResponse20096UpgradeWindow**](InlineResponse20096UpgradeWindow.md) |  | [optional] 
**Timezone** | Pointer to **string** | The timezone for the network | [optional] 
**Products** | Pointer to [**NetworksNetworkIdFirmwareUpgradesProducts**](NetworksNetworkIdFirmwareUpgradesProducts.md) |  | [optional] 

## Methods

### NewInlineObject97

`func NewInlineObject97() *InlineObject97`

NewInlineObject97 instantiates a new InlineObject97 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject97WithDefaults

`func NewInlineObject97WithDefaults() *InlineObject97`

NewInlineObject97WithDefaults instantiates a new InlineObject97 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeWindow

`func (o *InlineObject97) GetUpgradeWindow() InlineResponse20096UpgradeWindow`

GetUpgradeWindow returns the UpgradeWindow field if non-nil, zero value otherwise.

### GetUpgradeWindowOk

`func (o *InlineObject97) GetUpgradeWindowOk() (*InlineResponse20096UpgradeWindow, bool)`

GetUpgradeWindowOk returns a tuple with the UpgradeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeWindow

`func (o *InlineObject97) SetUpgradeWindow(v InlineResponse20096UpgradeWindow)`

SetUpgradeWindow sets UpgradeWindow field to given value.

### HasUpgradeWindow

`func (o *InlineObject97) HasUpgradeWindow() bool`

HasUpgradeWindow returns a boolean if a field has been set.

### GetTimezone

`func (o *InlineObject97) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InlineObject97) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InlineObject97) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InlineObject97) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetProducts

`func (o *InlineObject97) GetProducts() NetworksNetworkIdFirmwareUpgradesProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineObject97) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineObject97) SetProducts(v NetworksNetworkIdFirmwareUpgradesProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineObject97) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


