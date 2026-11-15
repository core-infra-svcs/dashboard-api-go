# NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L3** | Pointer to **[]string** | Layer 3 allow list of source IP addresses excluded from inspection. Maximum 600 entries. | [optional] 
**L7** | Pointer to **[]string** | Layer 7 allow list of domains excluded from inspection. Maximum 600 entries. | [optional] 
**Applications** | Pointer to [**[]NetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsApplications**](NetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsApplications.md) | Specific applications to be excluded from inspection. Maximum 600 entries. | [optional] 

## Methods

### NewNetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists

`func NewNetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists() *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists`

NewNetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists instantiates a new NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsWithDefaults

`func NewNetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsWithDefaults() *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists`

NewNetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsWithDefaults instantiates a new NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL3

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) GetL3() []string`

GetL3 returns the L3 field if non-nil, zero value otherwise.

### GetL3Ok

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) GetL3Ok() (*[]string, bool)`

GetL3Ok returns a tuple with the L3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) SetL3(v []string)`

SetL3 sets L3 field to given value.

### HasL3

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) HasL3() bool`

HasL3 returns a boolean if a field has been set.

### GetL7

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) GetL7() []string`

GetL7 returns the L7 field if non-nil, zero value otherwise.

### GetL7Ok

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) GetL7Ok() (*[]string, bool)`

GetL7Ok returns a tuple with the L7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) SetL7(v []string)`

SetL7 sets L7 field to given value.

### HasL7

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) HasL7() bool`

HasL7 returns a boolean if a field has been set.

### GetApplications

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) GetApplications() []NetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) GetApplicationsOk() (*[]NetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) SetApplications(v []NetworksNetworkIdApplianceSecurityHttpsInspectionAllowListsApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *NetworksNetworkIdApplianceSecurityHttpsInspectionAllowLists) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


