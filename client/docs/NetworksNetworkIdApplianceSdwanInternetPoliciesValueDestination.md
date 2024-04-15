# NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address (e.g.\&quot;192.168.10.1\&quot;, which is the same as \&quot;192.168.10.1/32\&quot;), or \&quot;any\&quot; | [optional] 
**Applications** | Pointer to [**[]NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications**](NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications.md) | list of application objects (either majorApplication or nbar) | [optional] 

## Methods

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination() *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationWithDefaults

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetApplications

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) GetApplications() []NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) GetApplicationsOk() (*[]NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) SetApplications(v []NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


