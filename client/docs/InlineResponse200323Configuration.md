# InlineResponse200323Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the alerts are enabled | [optional] 
**Maintenance** | Pointer to [**InlineResponse200323ConfigurationMaintenance**](InlineResponse200323ConfigurationMaintenance.md) |  | [optional] 
**AlertDestinations** | Pointer to [**InlineResponse200323ConfigurationAlertDestinations**](InlineResponse200323ConfigurationAlertDestinations.md) |  | [optional] 

## Methods

### NewInlineResponse200323Configuration

`func NewInlineResponse200323Configuration() *InlineResponse200323Configuration`

NewInlineResponse200323Configuration instantiates a new InlineResponse200323Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200323ConfigurationWithDefaults

`func NewInlineResponse200323ConfigurationWithDefaults() *InlineResponse200323Configuration`

NewInlineResponse200323ConfigurationWithDefaults instantiates a new InlineResponse200323Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200323Configuration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200323Configuration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200323Configuration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200323Configuration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaintenance

`func (o *InlineResponse200323Configuration) GetMaintenance() InlineResponse200323ConfigurationMaintenance`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *InlineResponse200323Configuration) GetMaintenanceOk() (*InlineResponse200323ConfigurationMaintenance, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *InlineResponse200323Configuration) SetMaintenance(v InlineResponse200323ConfigurationMaintenance)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *InlineResponse200323Configuration) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *InlineResponse200323Configuration) GetAlertDestinations() InlineResponse200323ConfigurationAlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *InlineResponse200323Configuration) GetAlertDestinationsOk() (*InlineResponse200323ConfigurationAlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *InlineResponse200323Configuration) SetAlertDestinations(v InlineResponse200323ConfigurationAlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.

### HasAlertDestinations

`func (o *InlineResponse200323Configuration) HasAlertDestinations() bool`

HasAlertDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


