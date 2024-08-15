# InlineResponse200205

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The alert config ID | [optional] 
**Type** | Pointer to **string** | The alert type | [optional] 
**Enabled** | Pointer to **bool** | Is the alert config enabled | [optional] 
**AlertCondition** | Pointer to [**OrganizationsOrganizationIdAlertsProfilesAlertCondition**](OrganizationsOrganizationIdAlertsProfilesAlertCondition.md) |  | [optional] 
**Recipients** | Pointer to [**OrganizationsOrganizationIdAlertsProfilesRecipients**](OrganizationsOrganizationIdAlertsProfilesRecipients.md) |  | [optional] 
**NetworkTags** | Pointer to **[]string** | Networks with these tags will be monitored for the alert | [optional] 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewInlineResponse200205

`func NewInlineResponse200205() *InlineResponse200205`

NewInlineResponse200205 instantiates a new InlineResponse200205 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200205WithDefaults

`func NewInlineResponse200205WithDefaults() *InlineResponse200205`

NewInlineResponse200205WithDefaults instantiates a new InlineResponse200205 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200205) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200205) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200205) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200205) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200205) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200205) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200205) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200205) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200205) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200205) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200205) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200205) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlertCondition

`func (o *InlineResponse200205) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *InlineResponse200205) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *InlineResponse200205) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *InlineResponse200205) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineResponse200205) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200205) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200205) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineResponse200205) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineResponse200205) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineResponse200205) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineResponse200205) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineResponse200205) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200205) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200205) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200205) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200205) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


