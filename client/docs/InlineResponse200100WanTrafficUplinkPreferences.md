# InlineResponse200100WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse200100TrafficFilters**](InlineResponse200100TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39;, or any other valid uplink (wanX, cellular) if it applies to the network | 
**Vrf** | Pointer to [**InlineResponse200100Vrf**](InlineResponse200100Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse200100WanTrafficUplinkPreferences

`func NewInlineResponse200100WanTrafficUplinkPreferences(trafficFilters []InlineResponse200100TrafficFilters, preferredUplink string, ) *InlineResponse200100WanTrafficUplinkPreferences`

NewInlineResponse200100WanTrafficUplinkPreferences instantiates a new InlineResponse200100WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200100WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse200100WanTrafficUplinkPreferencesWithDefaults() *InlineResponse200100WanTrafficUplinkPreferences`

NewInlineResponse200100WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse200100WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse200100WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse200100TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse200100WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse200100TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse200100WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse200100TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse200100WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse200100WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse200100WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetVrf

`func (o *InlineResponse200100WanTrafficUplinkPreferences) GetVrf() InlineResponse200100Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200100WanTrafficUplinkPreferences) GetVrfOk() (*InlineResponse200100Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200100WanTrafficUplinkPreferences) SetVrf(v InlineResponse200100Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200100WanTrafficUplinkPreferences) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


