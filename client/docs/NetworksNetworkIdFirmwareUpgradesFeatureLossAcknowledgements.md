# NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | The network ID to acknowledge the failing check for. Must refer to a node group. | 
**ProductType** | Pointer to **string** | The product type within the given network to acknowledge feature loss for (if the network is a combined network). Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | [optional] 
**RemovedFeatureCheckNames** | **[]string** | The names of the removed feature checks to acknowledge. | 

## Methods

### NewNetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements

`func NewNetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements(networkId string, removedFeatureCheckNames []string, ) *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements`

NewNetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements instantiates a new NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgementsWithDefaults

`func NewNetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgementsWithDefaults() *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements`

NewNetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgementsWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetProductType

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRemovedFeatureCheckNames

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) GetRemovedFeatureCheckNames() []string`

GetRemovedFeatureCheckNames returns the RemovedFeatureCheckNames field if non-nil, zero value otherwise.

### GetRemovedFeatureCheckNamesOk

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) GetRemovedFeatureCheckNamesOk() (*[]string, bool)`

GetRemovedFeatureCheckNamesOk returns a tuple with the RemovedFeatureCheckNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedFeatureCheckNames

`func (o *NetworksNetworkIdFirmwareUpgradesFeatureLossAcknowledgements) SetRemovedFeatureCheckNames(v []string)`

SetRemovedFeatureCheckNames sets RemovedFeatureCheckNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


