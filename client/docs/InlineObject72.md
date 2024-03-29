# InlineObject72

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The site-to-site VPN mode. Can be one of &#39;none&#39;, &#39;spoke&#39; or &#39;hub&#39; | 
**Hubs** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs**](NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs.md) | The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required. | [optional] 
**Subnets** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets**](NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets.md) | The list of subnets and their VPN presence. | [optional] 

## Methods

### NewInlineObject72

`func NewInlineObject72(mode string, ) *InlineObject72`

NewInlineObject72 instantiates a new InlineObject72 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject72WithDefaults

`func NewInlineObject72WithDefaults() *InlineObject72`

NewInlineObject72WithDefaults instantiates a new InlineObject72 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject72) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject72) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject72) SetMode(v string)`

SetMode sets Mode field to given value.


### GetHubs

`func (o *InlineObject72) GetHubs() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineObject72) GetHubsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineObject72) SetHubs(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineObject72) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineObject72) GetSubnets() []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineObject72) GetSubnetsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineObject72) SetSubnets(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineObject72) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


