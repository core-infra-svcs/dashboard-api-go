# NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Traffic filter type. Must be &#39;custom&#39;, &#39;major_application&#39;, &#39;application (NBAR)&#39;, if type is &#39;application&#39;, you can pass either an NBAR App Category or Application | 
**Value** | [**NetworksNetworkIdApplianceSdwanInternetPoliciesValue**](NetworksNetworkIdApplianceSdwanInternetPoliciesValue.md) |  | 

## Methods

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters(type_ string, value NetworksNetworkIdApplianceSdwanInternetPoliciesValue, ) *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFiltersWithDefaults

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFiltersWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFiltersWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) GetValue() NetworksNetworkIdApplianceSdwanInternetPoliciesValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) GetValueOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) SetValue(v NetworksNetworkIdApplianceSdwanInternetPoliciesValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


