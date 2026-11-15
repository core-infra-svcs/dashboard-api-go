# InlineResponse200103WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse200103TrafficFilters**](InlineResponse200103TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39;, or any other valid uplink (wanX, cellular) if it applies to the network | 
**Vrf** | Pointer to [**InlineResponse200103Vrf**](InlineResponse200103Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse200103WanTrafficUplinkPreferences

`func NewInlineResponse200103WanTrafficUplinkPreferences(trafficFilters []InlineResponse200103TrafficFilters, preferredUplink string, ) *InlineResponse200103WanTrafficUplinkPreferences`

NewInlineResponse200103WanTrafficUplinkPreferences instantiates a new InlineResponse200103WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse200103WanTrafficUplinkPreferencesWithDefaults() *InlineResponse200103WanTrafficUplinkPreferences`

NewInlineResponse200103WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse200103WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse200103WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse200103TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse200103WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse200103TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse200103WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse200103TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse200103WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse200103WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse200103WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetVrf

`func (o *InlineResponse200103WanTrafficUplinkPreferences) GetVrf() InlineResponse200103Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200103WanTrafficUplinkPreferences) GetVrfOk() (*InlineResponse200103Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200103WanTrafficUplinkPreferences) SetVrf(v InlineResponse200103Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200103WanTrafficUplinkPreferences) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


