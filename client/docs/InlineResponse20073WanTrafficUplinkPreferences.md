# InlineResponse20073WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20073TrafficFilters**](InlineResponse20073TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39;, or any other valid uplink(wanX) if it applies to the network | 

## Methods

### NewInlineResponse20073WanTrafficUplinkPreferences

`func NewInlineResponse20073WanTrafficUplinkPreferences(trafficFilters []InlineResponse20073TrafficFilters, preferredUplink string, ) *InlineResponse20073WanTrafficUplinkPreferences`

NewInlineResponse20073WanTrafficUplinkPreferences instantiates a new InlineResponse20073WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20073WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20073WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20073WanTrafficUplinkPreferences`

NewInlineResponse20073WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20073WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20073WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20073TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20073WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20073TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20073WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20073TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20073WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20073WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20073WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


