# InlineResponse20063Alerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of alert | 
**Enabled** | Pointer to **bool** | A boolean depicting if the alert is turned on or off | [optional] 
**AlertDestinations** | Pointer to [**InlineResponse20063AlertDestinations**](InlineResponse20063AlertDestinations.md) |  | [optional] 
**Filters** | Pointer to [**InlineResponse20063Filters**](InlineResponse20063Filters.md) |  | [optional] 

## Methods

### NewInlineResponse20063Alerts

`func NewInlineResponse20063Alerts(type_ string, ) *InlineResponse20063Alerts`

NewInlineResponse20063Alerts instantiates a new InlineResponse20063Alerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20063AlertsWithDefaults

`func NewInlineResponse20063AlertsWithDefaults() *InlineResponse20063Alerts`

NewInlineResponse20063AlertsWithDefaults instantiates a new InlineResponse20063Alerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20063Alerts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20063Alerts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20063Alerts) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *InlineResponse20063Alerts) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20063Alerts) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20063Alerts) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20063Alerts) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *InlineResponse20063Alerts) GetAlertDestinations() InlineResponse20063AlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *InlineResponse20063Alerts) GetAlertDestinationsOk() (*InlineResponse20063AlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *InlineResponse20063Alerts) SetAlertDestinations(v InlineResponse20063AlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.

### HasAlertDestinations

`func (o *InlineResponse20063Alerts) HasAlertDestinations() bool`

HasAlertDestinations returns a boolean if a field has been set.

### GetFilters

`func (o *InlineResponse20063Alerts) GetFilters() InlineResponse20063Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InlineResponse20063Alerts) GetFiltersOk() (*InlineResponse20063Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InlineResponse20063Alerts) SetFilters(v InlineResponse20063Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InlineResponse20063Alerts) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


