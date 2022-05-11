# NetworksNetworkIdClientsProvisionPoliciesBySsid0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePolicy** | **string** | The policy to apply to the specified client. Can be &#39;Allowed&#39;, &#39;Blocked&#39;, &#39;Normal&#39; or &#39;Group policy&#39;. Required. | 
**GroupPolicyId** | Pointer to **string** | The ID of the desired group policy to apply to the client. Required if &#39;devicePolicy&#39; is set to \&quot;Group policy\&quot;. Otherwise this is ignored. | [optional] 

## Methods

### NewNetworksNetworkIdClientsProvisionPoliciesBySsid0

`func NewNetworksNetworkIdClientsProvisionPoliciesBySsid0(devicePolicy string, ) *NetworksNetworkIdClientsProvisionPoliciesBySsid0`

NewNetworksNetworkIdClientsProvisionPoliciesBySsid0 instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySsid0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdClientsProvisionPoliciesBySsid0WithDefaults

`func NewNetworksNetworkIdClientsProvisionPoliciesBySsid0WithDefaults() *NetworksNetworkIdClientsProvisionPoliciesBySsid0`

NewNetworksNetworkIdClientsProvisionPoliciesBySsid0WithDefaults instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySsid0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePolicy

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.


### GetGroupPolicyId

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


