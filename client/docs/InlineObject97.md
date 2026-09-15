# InlineObject97

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The site-to-site VPN mode. Can be one of &#39;none&#39;, &#39;spoke&#39; or &#39;hub&#39; | 
**Hubs** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs**](NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs.md) | The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required. | [optional] 
**Subnets** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets**](NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets.md) | The list of subnets and their VPN presence. | [optional] 
**Sgt** | Pointer to [**InlineResponse200111Sgt**](InlineResponse200111Sgt.md) |  | [optional] 
**Subnet** | Pointer to [**InlineResponse200111Subnet**](InlineResponse200111Subnet.md) |  | [optional] 
**HostTranslations** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations**](NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations.md) | The list of VPN host translations. Host translations are supported starting from MX firmware version 26.1.2 | [optional] 

## Methods

### NewInlineObject97

`func NewInlineObject97(mode string, ) *InlineObject97`

NewInlineObject97 instantiates a new InlineObject97 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject97WithDefaults

`func NewInlineObject97WithDefaults() *InlineObject97`

NewInlineObject97WithDefaults instantiates a new InlineObject97 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject97) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject97) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject97) SetMode(v string)`

SetMode sets Mode field to given value.


### GetHubs

`func (o *InlineObject97) GetHubs() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineObject97) GetHubsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineObject97) SetHubs(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineObject97) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineObject97) GetSubnets() []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineObject97) GetSubnetsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineObject97) SetSubnets(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineObject97) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSgt

`func (o *InlineObject97) GetSgt() InlineResponse200111Sgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineObject97) GetSgtOk() (*InlineResponse200111Sgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineObject97) SetSgt(v InlineResponse200111Sgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineObject97) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject97) GetSubnet() InlineResponse200111Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject97) GetSubnetOk() (*InlineResponse200111Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject97) SetSubnet(v InlineResponse200111Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject97) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetHostTranslations

`func (o *InlineObject97) GetHostTranslations() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations`

GetHostTranslations returns the HostTranslations field if non-nil, zero value otherwise.

### GetHostTranslationsOk

`func (o *InlineObject97) GetHostTranslationsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations, bool)`

GetHostTranslationsOk returns a tuple with the HostTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTranslations

`func (o *InlineObject97) SetHostTranslations(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations)`

SetHostTranslations sets HostTranslations field to given value.

### HasHostTranslations

`func (o *InlineObject97) HasHostTranslations() bool`

HasHostTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


