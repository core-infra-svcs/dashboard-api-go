# InlineObject217

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the alert config enabled | [optional] 
**Type** | Pointer to **string** | The alert type | [optional] 
**AlertCondition** | Pointer to [**OrganizationsOrganizationIdAlertsProfilesAlertCondition1**](OrganizationsOrganizationIdAlertsProfilesAlertCondition1.md) |  | [optional] 
**Recipients** | Pointer to [**OrganizationsOrganizationIdAlertsProfilesRecipients**](OrganizationsOrganizationIdAlertsProfilesRecipients.md) |  | [optional] 
**NetworkTags** | Pointer to **[]string** | Networks with these tags will be monitored for the alert | [optional] 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewInlineObject217

`func NewInlineObject217() *InlineObject217`

NewInlineObject217 instantiates a new InlineObject217 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject217WithDefaults

`func NewInlineObject217WithDefaults() *InlineObject217`

NewInlineObject217WithDefaults instantiates a new InlineObject217 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject217) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject217) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject217) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject217) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineObject217) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject217) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject217) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject217) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlertCondition

`func (o *InlineObject217) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition1`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *InlineObject217) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition1, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *InlineObject217) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition1)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *InlineObject217) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineObject217) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject217) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject217) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineObject217) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineObject217) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineObject217) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineObject217) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineObject217) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject217) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject217) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject217) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject217) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


