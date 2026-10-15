# InlineObject121

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**NetworksNetworkIdFirmwareUpgradesStagedEventsProducts**](NetworksNetworkIdFirmwareUpgradesStagedEventsProducts.md) |  | [optional] 
**Stages** | [**[]NetworksNetworkIdFirmwareUpgradesStagedEventsStages**](NetworksNetworkIdFirmwareUpgradesStagedEventsStages.md) | All firmware upgrade stages in the network with their start time. | 

## Methods

### NewInlineObject121

`func NewInlineObject121(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages, ) *InlineObject121`

NewInlineObject121 instantiates a new InlineObject121 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject121WithDefaults

`func NewInlineObject121WithDefaults() *InlineObject121`

NewInlineObject121WithDefaults instantiates a new InlineObject121 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineObject121) GetProducts() NetworksNetworkIdFirmwareUpgradesStagedEventsProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineObject121) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineObject121) SetProducts(v NetworksNetworkIdFirmwareUpgradesStagedEventsProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineObject121) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineObject121) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineObject121) GetStagesOk() (*[]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineObject121) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


