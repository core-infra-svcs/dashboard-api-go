# OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether or not the sms channel is enabled for this alert type | 
**Recipients** | Pointer to **[]string** | List of sms recipient phone numbers to receive the notifications | [optional] 

## Methods

### NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms

`func NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms(enabled bool, ) *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms`

NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms instantiates a new OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSmsWithDefaults

`func NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSmsWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms`

NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSmsWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRecipients

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *OrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsSms) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


