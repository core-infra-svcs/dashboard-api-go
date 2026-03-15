# InlineResponse20079WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20079TrafficFilters**](InlineResponse20079TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39;, or any other valid uplink(wanX) if it applies to the network | 

## Methods

### NewInlineResponse20079WanTrafficUplinkPreferences

`func NewInlineResponse20079WanTrafficUplinkPreferences(trafficFilters []InlineResponse20079TrafficFilters, preferredUplink string, ) *InlineResponse20079WanTrafficUplinkPreferences`

NewInlineResponse20079WanTrafficUplinkPreferences instantiates a new InlineResponse20079WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20079WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20079WanTrafficUplinkPreferences`

NewInlineResponse20079WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20079WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20079WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20079TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20079WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20079TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20079WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20079TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20079WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20079WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20079WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


