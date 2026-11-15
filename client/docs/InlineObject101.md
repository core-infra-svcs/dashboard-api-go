# InlineObject101

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **bool** | Whether to accept the confirmation warnings for these changes. Defaults to true, which accepts confirmation warnings and preserves the existing update behavior. Set to false to opt in to confirmation warnings. Review the warnings returned in confirmations, then retry with confirmed: true to accept them. Fatal errors must be corrected before retrying. | [optional] 
**Mode** | **string** | The site-to-site VPN mode. Can be one of &#39;none&#39;, &#39;spoke&#39; or &#39;hub&#39; | 
**NatTraversal** | Pointer to [**NetworksNetworkIdApplianceVpnSiteToSiteVpnNatTraversal**](NetworksNetworkIdApplianceVpnSiteToSiteVpnNatTraversal.md) |  | [optional] 
**Hubs** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs**](NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs.md) | The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required. | [optional] 
**Subnets** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets**](NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets.md) | The list of local VPN subnets and their VPN presence. VLAN entries include vlanId and subnetting where applicable. When the corresponding NFO is enabled, NAT settings are returned for every VLAN; non-VLAN entries omit vlanId, subnetting, and nat.legacy. | [optional] 
**Sgt** | Pointer to [**InlineResponse200114Sgt**](InlineResponse200114Sgt.md) |  | [optional] 
**Subnet** | Pointer to [**InlineResponse200114Subnet**](InlineResponse200114Subnet.md) |  | [optional] 
**HostTranslations** | Pointer to [**[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations**](NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations.md) | The list of VPN host translations. Host translations are supported starting from MX firmware version 26.1.2 | [optional] 

## Methods

### NewInlineObject101

`func NewInlineObject101(mode string, ) *InlineObject101`

NewInlineObject101 instantiates a new InlineObject101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject101WithDefaults

`func NewInlineObject101WithDefaults() *InlineObject101`

NewInlineObject101WithDefaults instantiates a new InlineObject101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *InlineObject101) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *InlineObject101) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *InlineObject101) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *InlineObject101) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetMode

`func (o *InlineObject101) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject101) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject101) SetMode(v string)`

SetMode sets Mode field to given value.


### GetNatTraversal

`func (o *InlineObject101) GetNatTraversal() NetworksNetworkIdApplianceVpnSiteToSiteVpnNatTraversal`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *InlineObject101) GetNatTraversalOk() (*NetworksNetworkIdApplianceVpnSiteToSiteVpnNatTraversal, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *InlineObject101) SetNatTraversal(v NetworksNetworkIdApplianceVpnSiteToSiteVpnNatTraversal)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *InlineObject101) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.

### GetHubs

`func (o *InlineObject101) GetHubs() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineObject101) GetHubsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineObject101) SetHubs(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineObject101) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineObject101) GetSubnets() []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineObject101) GetSubnetsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineObject101) SetSubnets(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineObject101) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSgt

`func (o *InlineObject101) GetSgt() InlineResponse200114Sgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineObject101) GetSgtOk() (*InlineResponse200114Sgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineObject101) SetSgt(v InlineResponse200114Sgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineObject101) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject101) GetSubnet() InlineResponse200114Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject101) GetSubnetOk() (*InlineResponse200114Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject101) SetSubnet(v InlineResponse200114Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject101) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetHostTranslations

`func (o *InlineObject101) GetHostTranslations() []NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations`

GetHostTranslations returns the HostTranslations field if non-nil, zero value otherwise.

### GetHostTranslationsOk

`func (o *InlineObject101) GetHostTranslationsOk() (*[]NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations, bool)`

GetHostTranslationsOk returns a tuple with the HostTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTranslations

`func (o *InlineObject101) SetHostTranslations(v []NetworksNetworkIdApplianceVpnSiteToSiteVpnHostTranslations)`

SetHostTranslations sets HostTranslations field to given value.

### HasHostTranslations

`func (o *InlineObject101) HasHostTranslations() bool`

HasHostTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


