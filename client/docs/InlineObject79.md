# InlineObject79

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The site-to-site VPN mode. Can be one of &#39;none&#39;, &#39;spoke&#39; or &#39;hub&#39; | 
**Hubs** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs**](NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs.md) | The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required. | [optional] 
**Subnets** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets**](NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets.md) | The list of subnets and their VPN presence. | [optional] 
**Subnet** | Pointer to [**InlineResponse20080Subnet**](InlineResponse20080Subnet.md) |  | [optional] 

## Methods

### NewInlineObject79

`func NewInlineObject79(mode string, ) *InlineObject79`

NewInlineObject79 instantiates a new InlineObject79 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject79WithDefaults

`func NewInlineObject79WithDefaults() *InlineObject79`

NewInlineObject79WithDefaults instantiates a new InlineObject79 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject79) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject79) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject79) SetMode(v string)`

SetMode sets Mode field to given value.


### GetHubs

`func (o *InlineObject79) GetHubs() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineObject79) GetHubsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineObject79) SetHubs(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineObject79) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineObject79) GetSubnets() []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineObject79) GetSubnetsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineObject79) SetSubnets(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineObject79) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject79) GetSubnet() InlineResponse20080Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject79) GetSubnetOk() (*InlineResponse20080Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject79) SetSubnet(v InlineResponse20080Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject79) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


