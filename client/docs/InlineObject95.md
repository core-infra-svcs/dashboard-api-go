# InlineObject95

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeWindow** | Pointer to [**InlineResponse20096UpgradeWindow**](InlineResponse20096UpgradeWindow.md) |  | [optional] 
**Timezone** | Pointer to **string** | The timezone for the network | [optional] 
**Products** | Pointer to [**NetworksNetworkIdFirmwareUpgradesProducts**](NetworksNetworkIdFirmwareUpgradesProducts.md) |  | [optional] 

## Methods

### NewInlineObject95

`func NewInlineObject95() *InlineObject95`

NewInlineObject95 instantiates a new InlineObject95 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject95WithDefaults

`func NewInlineObject95WithDefaults() *InlineObject95`

NewInlineObject95WithDefaults instantiates a new InlineObject95 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeWindow

`func (o *InlineObject95) GetUpgradeWindow() InlineResponse20096UpgradeWindow`

GetUpgradeWindow returns the UpgradeWindow field if non-nil, zero value otherwise.

### GetUpgradeWindowOk

`func (o *InlineObject95) GetUpgradeWindowOk() (*InlineResponse20096UpgradeWindow, bool)`

GetUpgradeWindowOk returns a tuple with the UpgradeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeWindow

`func (o *InlineObject95) SetUpgradeWindow(v InlineResponse20096UpgradeWindow)`

SetUpgradeWindow sets UpgradeWindow field to given value.

### HasUpgradeWindow

`func (o *InlineObject95) HasUpgradeWindow() bool`

HasUpgradeWindow returns a boolean if a field has been set.

### GetTimezone

`func (o *InlineObject95) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InlineObject95) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InlineObject95) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InlineObject95) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetProducts

`func (o *InlineObject95) GetProducts() NetworksNetworkIdFirmwareUpgradesProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineObject95) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineObject95) SetProducts(v NetworksNetworkIdFirmwareUpgradesProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineObject95) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


