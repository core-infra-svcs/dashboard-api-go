# InlineObject220

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The alert type | 
**AlertCondition** | [**OrganizationsOrganizationIdAlertsProfilesAlertCondition1**](OrganizationsOrganizationIdAlertsProfilesAlertCondition1.md) |  | 
**Recipients** | [**OrganizationsOrganizationIdAlertsProfilesRecipients**](OrganizationsOrganizationIdAlertsProfilesRecipients.md) |  | 
**NetworkTags** | **[]string** | Networks with these tags will be monitored for the alert | 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewInlineObject220

`func NewInlineObject220(type_ string, alertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition1, recipients OrganizationsOrganizationIdAlertsProfilesRecipients, networkTags []string, ) *InlineObject220`

NewInlineObject220 instantiates a new InlineObject220 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject220WithDefaults

`func NewInlineObject220WithDefaults() *InlineObject220`

NewInlineObject220WithDefaults instantiates a new InlineObject220 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject220) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject220) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject220) SetType(v string)`

SetType sets Type field to given value.


### GetAlertCondition

`func (o *InlineObject220) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition1`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *InlineObject220) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition1, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *InlineObject220) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition1)`

SetAlertCondition sets AlertCondition field to given value.


### GetRecipients

`func (o *InlineObject220) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject220) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject220) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.


### GetNetworkTags

`func (o *InlineObject220) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineObject220) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineObject220) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.


### GetDescription

`func (o *InlineObject220) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject220) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject220) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject220) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


