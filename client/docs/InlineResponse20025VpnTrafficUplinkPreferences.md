# InlineResponse20025VpnTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20025TrafficFilters1**](InlineResponse20025TrafficFilters1.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39;, &#39;wan2&#39;, &#39;bestForVoIP&#39;, &#39;loadBalancing&#39; or &#39;defaultUplink&#39; | 
**FailOverCriterion** | Pointer to **string** | Fail over criterion for uplink preference rule. Must be one of: &#39;poorPerformance&#39; or &#39;uplinkDown&#39; | [optional] 
**PerformanceClass** | Pointer to [**InlineResponse20025PerformanceClass**](InlineResponse20025PerformanceClass.md) |  | [optional] 

## Methods

### NewInlineResponse20025VpnTrafficUplinkPreferences

`func NewInlineResponse20025VpnTrafficUplinkPreferences(trafficFilters []InlineResponse20025TrafficFilters1, preferredUplink string, ) *InlineResponse20025VpnTrafficUplinkPreferences`

NewInlineResponse20025VpnTrafficUplinkPreferences instantiates a new InlineResponse20025VpnTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20025VpnTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20025VpnTrafficUplinkPreferencesWithDefaults() *InlineResponse20025VpnTrafficUplinkPreferences`

NewInlineResponse20025VpnTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20025VpnTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20025TrafficFilters1`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20025TrafficFilters1, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20025TrafficFilters1)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetFailOverCriterion

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetFailOverCriterion() string`

GetFailOverCriterion returns the FailOverCriterion field if non-nil, zero value otherwise.

### GetFailOverCriterionOk

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool)`

GetFailOverCriterionOk returns a tuple with the FailOverCriterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOverCriterion

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) SetFailOverCriterion(v string)`

SetFailOverCriterion sets FailOverCriterion field to given value.

### HasFailOverCriterion

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) HasFailOverCriterion() bool`

HasFailOverCriterion returns a boolean if a field has been set.

### GetPerformanceClass

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetPerformanceClass() InlineResponse20025PerformanceClass`

GetPerformanceClass returns the PerformanceClass field if non-nil, zero value otherwise.

### GetPerformanceClassOk

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) GetPerformanceClassOk() (*InlineResponse20025PerformanceClass, bool)`

GetPerformanceClassOk returns a tuple with the PerformanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceClass

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) SetPerformanceClass(v InlineResponse20025PerformanceClass)`

SetPerformanceClass sets PerformanceClass field to given value.

### HasPerformanceClass

`func (o *InlineResponse20025VpnTrafficUplinkPreferences) HasPerformanceClass() bool`

HasPerformanceClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


