# InlineObject85

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**NetworksNetworkIdFirmwareUpgradesStagedEventsProducts**](NetworksNetworkIdFirmwareUpgradesStagedEventsProducts.md) |  | [optional] 
**Stages** | [**[]NetworksNetworkIdFirmwareUpgradesStagedEventsStages**](NetworksNetworkIdFirmwareUpgradesStagedEventsStages.md) | All firmware upgrade stages in the network with their start time. | 

## Methods

### NewInlineObject85

`func NewInlineObject85(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages, ) *InlineObject85`

NewInlineObject85 instantiates a new InlineObject85 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject85WithDefaults

`func NewInlineObject85WithDefaults() *InlineObject85`

NewInlineObject85WithDefaults instantiates a new InlineObject85 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineObject85) GetProducts() NetworksNetworkIdFirmwareUpgradesStagedEventsProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineObject85) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineObject85) SetProducts(v NetworksNetworkIdFirmwareUpgradesStagedEventsProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineObject85) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineObject85) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineObject85) GetStagesOk() (*[]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineObject85) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


