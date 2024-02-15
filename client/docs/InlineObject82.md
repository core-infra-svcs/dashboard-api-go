# InlineObject82

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | [**[]NetworksNetworkIdClientsProvisionClients**](NetworksNetworkIdClientsProvisionClients.md) | The array of clients to provision | 
**DevicePolicy** | **string** | The policy to apply to the specified client. Can be &#39;Group policy&#39;, &#39;Allowed&#39;, &#39;Blocked&#39;, &#39;Per connection&#39; or &#39;Normal&#39;. Required. | 
**GroupPolicyId** | Pointer to **string** | The ID of the desired group policy to apply to the client. Required if &#39;devicePolicy&#39; is set to \&quot;Group policy\&quot;. Otherwise this is ignored. | [optional] 
**PoliciesBySecurityAppliance** | Pointer to [**NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance**](NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance.md) |  | [optional] 
**PoliciesBySsid** | Pointer to [**NetworksNetworkIdClientsProvisionPoliciesBySsid**](NetworksNetworkIdClientsProvisionPoliciesBySsid.md) |  | [optional] 

## Methods

### NewInlineObject82

`func NewInlineObject82(clients []NetworksNetworkIdClientsProvisionClients, devicePolicy string, ) *InlineObject82`

NewInlineObject82 instantiates a new InlineObject82 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject82WithDefaults

`func NewInlineObject82WithDefaults() *InlineObject82`

NewInlineObject82WithDefaults instantiates a new InlineObject82 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *InlineObject82) GetClients() []NetworksNetworkIdClientsProvisionClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineObject82) GetClientsOk() (*[]NetworksNetworkIdClientsProvisionClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineObject82) SetClients(v []NetworksNetworkIdClientsProvisionClients)`

SetClients sets Clients field to given value.


### GetDevicePolicy

`func (o *InlineObject82) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineObject82) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineObject82) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.


### GetGroupPolicyId

`func (o *InlineObject82) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject82) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject82) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineObject82) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPoliciesBySecurityAppliance

`func (o *InlineObject82) GetPoliciesBySecurityAppliance() NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance`

GetPoliciesBySecurityAppliance returns the PoliciesBySecurityAppliance field if non-nil, zero value otherwise.

### GetPoliciesBySecurityApplianceOk

`func (o *InlineObject82) GetPoliciesBySecurityApplianceOk() (*NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance, bool)`

GetPoliciesBySecurityApplianceOk returns a tuple with the PoliciesBySecurityAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesBySecurityAppliance

`func (o *InlineObject82) SetPoliciesBySecurityAppliance(v NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance)`

SetPoliciesBySecurityAppliance sets PoliciesBySecurityAppliance field to given value.

### HasPoliciesBySecurityAppliance

`func (o *InlineObject82) HasPoliciesBySecurityAppliance() bool`

HasPoliciesBySecurityAppliance returns a boolean if a field has been set.

### GetPoliciesBySsid

`func (o *InlineObject82) GetPoliciesBySsid() NetworksNetworkIdClientsProvisionPoliciesBySsid`

GetPoliciesBySsid returns the PoliciesBySsid field if non-nil, zero value otherwise.

### GetPoliciesBySsidOk

`func (o *InlineObject82) GetPoliciesBySsidOk() (*NetworksNetworkIdClientsProvisionPoliciesBySsid, bool)`

GetPoliciesBySsidOk returns a tuple with the PoliciesBySsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesBySsid

`func (o *InlineObject82) SetPoliciesBySsid(v NetworksNetworkIdClientsProvisionPoliciesBySsid)`

SetPoliciesBySsid sets PoliciesBySsid field to given value.

### HasPoliciesBySsid

`func (o *InlineObject82) HasPoliciesBySsid() bool`

HasPoliciesBySsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


