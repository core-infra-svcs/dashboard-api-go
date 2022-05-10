# NetworksNetworkIdGroupPoliciesVlanTagging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How VLAN tagging is applied. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**VlanId** | Pointer to **string** | The ID of the vlan you want to tag. This only applies if &#39;settings&#39; is set to &#39;custom&#39;. | [optional] 

## Methods

### NewNetworksNetworkIdGroupPoliciesVlanTagging

`func NewNetworksNetworkIdGroupPoliciesVlanTagging() *NetworksNetworkIdGroupPoliciesVlanTagging`

NewNetworksNetworkIdGroupPoliciesVlanTagging instantiates a new NetworksNetworkIdGroupPoliciesVlanTagging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdGroupPoliciesVlanTaggingWithDefaults

`func NewNetworksNetworkIdGroupPoliciesVlanTaggingWithDefaults() *NetworksNetworkIdGroupPoliciesVlanTagging`

NewNetworksNetworkIdGroupPoliciesVlanTaggingWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesVlanTagging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetVlanId

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NetworksNetworkIdGroupPoliciesVlanTagging) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


