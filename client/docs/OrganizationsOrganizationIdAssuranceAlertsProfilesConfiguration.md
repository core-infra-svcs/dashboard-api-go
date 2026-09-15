# OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the alerts are enabled | 
**Maintenance** | Pointer to [**InlineResponse200304ConfigurationMaintenance**](InlineResponse200304ConfigurationMaintenance.md) |  | [optional] 
**AlertDestinations** | [**OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations**](OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations.md) |  | 

## Methods

### NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration

`func NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration(enabled bool, alertDestinations OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations, ) *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration`

NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration instantiates a new OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationWithDefaults

`func NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration`

NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaintenance

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) GetMaintenance() InlineResponse200304ConfigurationMaintenance`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) GetMaintenanceOk() (*InlineResponse200304ConfigurationMaintenance, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) SetMaintenance(v InlineResponse200304ConfigurationMaintenance)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) GetAlertDestinations() OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) GetAlertDestinationsOk() (*OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration) SetAlertDestinations(v OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


