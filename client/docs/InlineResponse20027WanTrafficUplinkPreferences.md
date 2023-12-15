# InlineResponse20027WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20027TrafficFilters**](InlineResponse20027TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewInlineResponse20027WanTrafficUplinkPreferences

`func NewInlineResponse20027WanTrafficUplinkPreferences(trafficFilters []InlineResponse20027TrafficFilters, preferredUplink string, ) *InlineResponse20027WanTrafficUplinkPreferences`

NewInlineResponse20027WanTrafficUplinkPreferences instantiates a new InlineResponse20027WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20027WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20027WanTrafficUplinkPreferences`

NewInlineResponse20027WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20027WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20027WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20027TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20027WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20027TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20027WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20027TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20027WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20027WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20027WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


