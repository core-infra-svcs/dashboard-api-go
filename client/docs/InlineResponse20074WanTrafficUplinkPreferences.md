# InlineResponse20074WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20074TrafficFilters**](InlineResponse20074TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39;, or any other valid uplink(wanX) if it applies to the network | 

## Methods

### NewInlineResponse20074WanTrafficUplinkPreferences

`func NewInlineResponse20074WanTrafficUplinkPreferences(trafficFilters []InlineResponse20074TrafficFilters, preferredUplink string, ) *InlineResponse20074WanTrafficUplinkPreferences`

NewInlineResponse20074WanTrafficUplinkPreferences instantiates a new InlineResponse20074WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20074WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20074WanTrafficUplinkPreferences`

NewInlineResponse20074WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20074WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20074WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20074TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20074WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20074TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20074WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20074TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20074WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20074WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20074WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


