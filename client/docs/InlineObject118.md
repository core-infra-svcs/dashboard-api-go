# InlineObject118

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeWindow** | Pointer to [**NetworksNetworkIdFirmwareUpgradesUpgradeWindow**](NetworksNetworkIdFirmwareUpgradesUpgradeWindow.md) |  | [optional] 
**Timezone** | Pointer to **string** | The timezone for the network | [optional] 
**Products** | Pointer to [**NetworksNetworkIdFirmwareUpgradesProducts**](NetworksNetworkIdFirmwareUpgradesProducts.md) |  | [optional] 
**FeatureLossAcknowledgements** | Pointer to [**[]NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements**](NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements.md) | Acknowledges the target version has removed certain features that the network is currently using. | [optional] 

## Methods

### NewInlineObject118

`func NewInlineObject118() *InlineObject118`

NewInlineObject118 instantiates a new InlineObject118 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject118WithDefaults

`func NewInlineObject118WithDefaults() *InlineObject118`

NewInlineObject118WithDefaults instantiates a new InlineObject118 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeWindow

`func (o *InlineObject118) GetUpgradeWindow() NetworksNetworkIdFirmwareUpgradesUpgradeWindow`

GetUpgradeWindow returns the UpgradeWindow field if non-nil, zero value otherwise.

### GetUpgradeWindowOk

`func (o *InlineObject118) GetUpgradeWindowOk() (*NetworksNetworkIdFirmwareUpgradesUpgradeWindow, bool)`

GetUpgradeWindowOk returns a tuple with the UpgradeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeWindow

`func (o *InlineObject118) SetUpgradeWindow(v NetworksNetworkIdFirmwareUpgradesUpgradeWindow)`

SetUpgradeWindow sets UpgradeWindow field to given value.

### HasUpgradeWindow

`func (o *InlineObject118) HasUpgradeWindow() bool`

HasUpgradeWindow returns a boolean if a field has been set.

### GetTimezone

`func (o *InlineObject118) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InlineObject118) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InlineObject118) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InlineObject118) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetProducts

`func (o *InlineObject118) GetProducts() NetworksNetworkIdFirmwareUpgradesProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineObject118) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineObject118) SetProducts(v NetworksNetworkIdFirmwareUpgradesProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineObject118) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetFeatureLossAcknowledgements

`func (o *InlineObject118) GetFeatureLossAcknowledgements() []NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements`

GetFeatureLossAcknowledgements returns the FeatureLossAcknowledgements field if non-nil, zero value otherwise.

### GetFeatureLossAcknowledgementsOk

`func (o *InlineObject118) GetFeatureLossAcknowledgementsOk() (*[]NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements, bool)`

GetFeatureLossAcknowledgementsOk returns a tuple with the FeatureLossAcknowledgements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLossAcknowledgements

`func (o *InlineObject118) SetFeatureLossAcknowledgements(v []NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements)`

SetFeatureLossAcknowledgements sets FeatureLossAcknowledgements field to given value.

### HasFeatureLossAcknowledgements

`func (o *InlineObject118) HasFeatureLossAcknowledgements() bool`

HasFeatureLossAcknowledgements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


