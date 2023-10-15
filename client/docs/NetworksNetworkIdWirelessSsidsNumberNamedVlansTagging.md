# NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not traffic should be directed to use specific VLAN names. | [optional] 
**DefaultVlanName** | Pointer to **string** | The default VLAN name used to tag traffic in the absence of a matching AP tag. | [optional] 
**ByApTags** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags**](NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags.md) | The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag. | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging

`func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging`

NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingWithDefaults() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging`

NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultVlanName

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetDefaultVlanName() string`

GetDefaultVlanName returns the DefaultVlanName field if non-nil, zero value otherwise.

### GetDefaultVlanNameOk

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetDefaultVlanNameOk() (*string, bool)`

GetDefaultVlanNameOk returns a tuple with the DefaultVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanName

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) SetDefaultVlanName(v string)`

SetDefaultVlanName sets DefaultVlanName field to given value.

### HasDefaultVlanName

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) HasDefaultVlanName() bool`

HasDefaultVlanName returns a boolean if a field has been set.

### GetByApTags

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetByApTags() []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags`

GetByApTags returns the ByApTags field if non-nil, zero value otherwise.

### GetByApTagsOk

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetByApTagsOk() (*[]NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags, bool)`

GetByApTagsOk returns a tuple with the ByApTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByApTags

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) SetByApTags(v []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags)`

SetByApTags sets ByApTags field to given value.

### HasByApTags

`func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) HasByApTags() bool`

HasByApTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


