# InlineResponse20068WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20068TrafficFilters**](InlineResponse20068TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewInlineResponse20068WanTrafficUplinkPreferences

`func NewInlineResponse20068WanTrafficUplinkPreferences(trafficFilters []InlineResponse20068TrafficFilters, preferredUplink string, ) *InlineResponse20068WanTrafficUplinkPreferences`

NewInlineResponse20068WanTrafficUplinkPreferences instantiates a new InlineResponse20068WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20068WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20068WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20068WanTrafficUplinkPreferences`

NewInlineResponse20068WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20068WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20068WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20068TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20068WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20068TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20068WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20068TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20068WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20068WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20068WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


