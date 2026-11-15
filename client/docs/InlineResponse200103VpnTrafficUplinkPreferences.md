# InlineResponse200103VpnTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse200103TrafficFilters1**](InlineResponse200103TrafficFilters1.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39;, &#39;wan2&#39;, &#39;bestForVoIP&#39;, &#39;loadBalancing&#39;, &#39;defaultUplink&#39; or &#39;cellular&#39;, or any other valid uplink (wanX, cellular) if it applies to the network | 
**Vrf** | Pointer to [**InlineResponse200103Vrf1**](InlineResponse200103Vrf1.md) |  | [optional] 
**FailOverCriterion** | Pointer to **string** | Fail over criterion for uplink preference rule. Must be one of: &#39;poorPerformance&#39; or &#39;uplinkDown&#39; | [optional] 
**PerformanceClass** | Pointer to [**InlineResponse200103PerformanceClass**](InlineResponse200103PerformanceClass.md) |  | [optional] 

## Methods

### NewInlineResponse200103VpnTrafficUplinkPreferences

`func NewInlineResponse200103VpnTrafficUplinkPreferences(trafficFilters []InlineResponse200103TrafficFilters1, preferredUplink string, ) *InlineResponse200103VpnTrafficUplinkPreferences`

NewInlineResponse200103VpnTrafficUplinkPreferences instantiates a new InlineResponse200103VpnTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103VpnTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse200103VpnTrafficUplinkPreferencesWithDefaults() *InlineResponse200103VpnTrafficUplinkPreferences`

NewInlineResponse200103VpnTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse200103VpnTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse200103TrafficFilters1`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse200103TrafficFilters1, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse200103TrafficFilters1)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetVrf

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetVrf() InlineResponse200103Vrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetVrfOk() (*InlineResponse200103Vrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) SetVrf(v InlineResponse200103Vrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetFailOverCriterion

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetFailOverCriterion() string`

GetFailOverCriterion returns the FailOverCriterion field if non-nil, zero value otherwise.

### GetFailOverCriterionOk

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool)`

GetFailOverCriterionOk returns a tuple with the FailOverCriterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOverCriterion

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) SetFailOverCriterion(v string)`

SetFailOverCriterion sets FailOverCriterion field to given value.

### HasFailOverCriterion

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) HasFailOverCriterion() bool`

HasFailOverCriterion returns a boolean if a field has been set.

### GetPerformanceClass

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetPerformanceClass() InlineResponse200103PerformanceClass`

GetPerformanceClass returns the PerformanceClass field if non-nil, zero value otherwise.

### GetPerformanceClassOk

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) GetPerformanceClassOk() (*InlineResponse200103PerformanceClass, bool)`

GetPerformanceClassOk returns a tuple with the PerformanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceClass

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) SetPerformanceClass(v InlineResponse200103PerformanceClass)`

SetPerformanceClass sets PerformanceClass field to given value.

### HasPerformanceClass

`func (o *InlineResponse200103VpnTrafficUplinkPreferences) HasPerformanceClass() bool`

HasPerformanceClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


