# InlineResponse20090WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39;, &#39;wan2&#39;, &#39;bestForVoIP&#39;, &#39;loadBalancing&#39;, &#39;defaultUplink&#39; or &#39;cellular&#39; | 
**FailOverCriterion** | Pointer to **string** | WAN failover and failback behavior | [optional] 
**PerformanceClass** | Pointer to [**NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass**](NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass.md) |  | [optional] 
**TrafficFilters** | [**[]NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters**](NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters.md) | Traffic filters | 
**Vrf** | Pointer to [**InlineResponse20090Vrf**](InlineResponse20090Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse20090WanTrafficUplinkPreferences

`func NewInlineResponse20090WanTrafficUplinkPreferences(preferredUplink string, trafficFilters []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters, ) *InlineResponse20090WanTrafficUplinkPreferences`

NewInlineResponse20090WanTrafficUplinkPreferences instantiates a new InlineResponse20090WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20090WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20090WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20090WanTrafficUplinkPreferences`

NewInlineResponse20090WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20090WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferredUplink

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20090WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetFailOverCriterion

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetFailOverCriterion() string`

GetFailOverCriterion returns the FailOverCriterion field if non-nil, zero value otherwise.

### GetFailOverCriterionOk

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool)`

GetFailOverCriterionOk returns a tuple with the FailOverCriterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOverCriterion

`func (o *InlineResponse20090WanTrafficUplinkPreferences) SetFailOverCriterion(v string)`

SetFailOverCriterion sets FailOverCriterion field to given value.

### HasFailOverCriterion

`func (o *InlineResponse20090WanTrafficUplinkPreferences) HasFailOverCriterion() bool`

HasFailOverCriterion returns a boolean if a field has been set.

### GetPerformanceClass

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetPerformanceClass() NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass`

GetPerformanceClass returns the PerformanceClass field if non-nil, zero value otherwise.

### GetPerformanceClassOk

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetPerformanceClassOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass, bool)`

GetPerformanceClassOk returns a tuple with the PerformanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceClass

`func (o *InlineResponse20090WanTrafficUplinkPreferences) SetPerformanceClass(v NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass)`

SetPerformanceClass sets PerformanceClass field to given value.

### HasPerformanceClass

`func (o *InlineResponse20090WanTrafficUplinkPreferences) HasPerformanceClass() bool`

HasPerformanceClass returns a boolean if a field has been set.

### GetTrafficFilters

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetTrafficFilters() []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20090WanTrafficUplinkPreferences) SetTrafficFilters(v []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetVrf

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetVrf() InlineResponse20090Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20090WanTrafficUplinkPreferences) GetVrfOk() (*InlineResponse20090Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20090WanTrafficUplinkPreferences) SetVrf(v InlineResponse20090Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20090WanTrafficUplinkPreferences) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


