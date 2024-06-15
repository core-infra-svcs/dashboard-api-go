# InlineResponse20064WanTrafficUplinkPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficFilters** | [**[]InlineResponse20064TrafficFilters**](InlineResponse20064TrafficFilters.md) | Traffic filters | 
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | 

## Methods

### NewInlineResponse20064WanTrafficUplinkPreferences

`func NewInlineResponse20064WanTrafficUplinkPreferences(trafficFilters []InlineResponse20064TrafficFilters, preferredUplink string, ) *InlineResponse20064WanTrafficUplinkPreferences`

NewInlineResponse20064WanTrafficUplinkPreferences instantiates a new InlineResponse20064WanTrafficUplinkPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20064WanTrafficUplinkPreferencesWithDefaults

`func NewInlineResponse20064WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20064WanTrafficUplinkPreferences`

NewInlineResponse20064WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20064WanTrafficUplinkPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficFilters

`func (o *InlineResponse20064WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20064TrafficFilters`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *InlineResponse20064WanTrafficUplinkPreferences) GetTrafficFiltersOk() (*[]InlineResponse20064TrafficFilters, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *InlineResponse20064WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20064TrafficFilters)`

SetTrafficFilters sets TrafficFilters field to given value.


### GetPreferredUplink

`func (o *InlineResponse20064WanTrafficUplinkPreferences) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *InlineResponse20064WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *InlineResponse20064WanTrafficUplinkPreferences) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


