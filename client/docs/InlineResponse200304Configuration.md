# InlineResponse200304Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the alerts are enabled | [optional] 
**Maintenance** | Pointer to [**InlineResponse200304ConfigurationMaintenance**](InlineResponse200304ConfigurationMaintenance.md) |  | [optional] 
**AlertDestinations** | Pointer to [**InlineResponse200304ConfigurationAlertDestinations**](InlineResponse200304ConfigurationAlertDestinations.md) |  | [optional] 

## Methods

### NewInlineResponse200304Configuration

`func NewInlineResponse200304Configuration() *InlineResponse200304Configuration`

NewInlineResponse200304Configuration instantiates a new InlineResponse200304Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200304ConfigurationWithDefaults

`func NewInlineResponse200304ConfigurationWithDefaults() *InlineResponse200304Configuration`

NewInlineResponse200304ConfigurationWithDefaults instantiates a new InlineResponse200304Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200304Configuration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200304Configuration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200304Configuration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200304Configuration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaintenance

`func (o *InlineResponse200304Configuration) GetMaintenance() InlineResponse200304ConfigurationMaintenance`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *InlineResponse200304Configuration) GetMaintenanceOk() (*InlineResponse200304ConfigurationMaintenance, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *InlineResponse200304Configuration) SetMaintenance(v InlineResponse200304ConfigurationMaintenance)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *InlineResponse200304Configuration) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *InlineResponse200304Configuration) GetAlertDestinations() InlineResponse200304ConfigurationAlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *InlineResponse200304Configuration) GetAlertDestinationsOk() (*InlineResponse200304ConfigurationAlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *InlineResponse200304Configuration) SetAlertDestinations(v InlineResponse200304ConfigurationAlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.

### HasAlertDestinations

`func (o *InlineResponse200304Configuration) HasAlertDestinations() bool`

HasAlertDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


