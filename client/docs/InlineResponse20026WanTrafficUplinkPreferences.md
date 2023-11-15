# InlineResponse20026WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20026TrafficFilters**](InlineResponse20026TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewInlineResponse20026WanTrafficUplinkPreferences

`func NewInlineResponse20026WanTrafficUplinkPreferences(trafficFilters []InlineResponse20026TrafficFilters, preferredUplink string, ) *InlineResponse20026WanTrafficUplinkPreferences`

NewInlineResponse20026WanTrafficUplinkPreferences instantiates a new InlineResponse20026WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20026WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20026WanTrafficUplinkPreferences`

NewInlineResponse20026WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20026WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20026WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20026TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20026WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20026TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20026WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20026TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20026WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20026WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20026WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


