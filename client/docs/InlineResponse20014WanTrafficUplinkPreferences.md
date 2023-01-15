# InlineResponse20014WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20014TrafficFilters**](InlineResponse20014TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewInlineResponse20014WanTrafficUplinkPreferences

`func NewInlineResponse20014WanTrafficUplinkPreferences(trafficFilters []InlineResponse20014TrafficFilters, preferredUplink string, ) *InlineResponse20014WanTrafficUplinkPreferences`

NewInlineResponse20014WanTrafficUplinkPreferences instantiates a new InlineResponse20014WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20014WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20014WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20014WanTrafficUplinkPreferences`

NewInlineResponse20014WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20014WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20014WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20014TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20014WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20014TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20014WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20014TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20014WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20014WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20014WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


