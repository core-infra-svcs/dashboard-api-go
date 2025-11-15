# NetworksNetworkIdSwitchSettingsUplinkSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failback** | [**NetworksNetworkIdSwitchSettingsUplinkSelectionFailback**](NetworksNetworkIdSwitchSettingsUplinkSelectionFailback.md) |  | 
**Candidates** | **string** | &#39;all&#39; lets devices try any potential interface. &#39;designated&#39; restricts to specified candidates (configured via the Routing &amp; DHCP page). | 

## Methods

### NewNetworksNetworkIdSwitchSettingsUplinkSelection

`func NewNetworksNetworkIdSwitchSettingsUplinkSelection(failback NetworksNetworkIdSwitchSettingsUplinkSelectionFailback, candidates string, ) *NetworksNetworkIdSwitchSettingsUplinkSelection`

NewNetworksNetworkIdSwitchSettingsUplinkSelection instantiates a new NetworksNetworkIdSwitchSettingsUplinkSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchSettingsUplinkSelectionWithDefaults

`func NewNetworksNetworkIdSwitchSettingsUplinkSelectionWithDefaults() *NetworksNetworkIdSwitchSettingsUplinkSelection`

NewNetworksNetworkIdSwitchSettingsUplinkSelectionWithDefaults instantiates a new NetworksNetworkIdSwitchSettingsUplinkSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailback

`func (o *NetworksNetworkIdSwitchSettingsUplinkSelection) GetFailback() NetworksNetworkIdSwitchSettingsUplinkSelectionFailback`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *NetworksNetworkIdSwitchSettingsUplinkSelection) GetFailbackOk() (*NetworksNetworkIdSwitchSettingsUplinkSelectionFailback, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *NetworksNetworkIdSwitchSettingsUplinkSelection) SetFailback(v NetworksNetworkIdSwitchSettingsUplinkSelectionFailback)`

SetFailback sets Failback field to given value.


### GetCandidates

`func (o *NetworksNetworkIdSwitchSettingsUplinkSelection) GetCandidates() string`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *NetworksNetworkIdSwitchSettingsUplinkSelection) GetCandidatesOk() (*string, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *NetworksNetworkIdSwitchSettingsUplinkSelection) SetCandidates(v string)`

SetCandidates sets Candidates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


