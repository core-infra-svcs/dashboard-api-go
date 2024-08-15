# InlineResponse20046Alerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of alert | 
**Enabled** | Pointer to **bool** | A boolean depicting if the alert is turned on or off | [optional] 
**AlertDestinations** | Pointer to [**InlineResponse20046AlertDestinations**](InlineResponse20046AlertDestinations.md) |  | [optional] 
**Filters** | Pointer to [**InlineResponse20046Filters**](InlineResponse20046Filters.md) |  | [optional] 

## Methods

### NewInlineResponse20046Alerts

`func NewInlineResponse20046Alerts(type_ string, ) *InlineResponse20046Alerts`

NewInlineResponse20046Alerts instantiates a new InlineResponse20046Alerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20046AlertsWithDefaults

`func NewInlineResponse20046AlertsWithDefaults() *InlineResponse20046Alerts`

NewInlineResponse20046AlertsWithDefaults instantiates a new InlineResponse20046Alerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20046Alerts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20046Alerts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20046Alerts) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *InlineResponse20046Alerts) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20046Alerts) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20046Alerts) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20046Alerts) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *InlineResponse20046Alerts) GetAlertDestinations() InlineResponse20046AlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *InlineResponse20046Alerts) GetAlertDestinationsOk() (*InlineResponse20046AlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *InlineResponse20046Alerts) SetAlertDestinations(v InlineResponse20046AlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.

### HasAlertDestinations

`func (o *InlineResponse20046Alerts) HasAlertDestinations() bool`

HasAlertDestinations returns a boolean if a field has been set.

### GetFilters

`func (o *InlineResponse20046Alerts) GetFilters() InlineResponse20046Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InlineResponse20046Alerts) GetFiltersOk() (*InlineResponse20046Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InlineResponse20046Alerts) SetFilters(v InlineResponse20046Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InlineResponse20046Alerts) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


