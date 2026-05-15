# InlineObject230

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The alert type | 
**AlertCondition** | [**OrganizationsOrganizationIdAlertsProfilesAlertCondition1**](OrganizationsOrganizationIdAlertsProfilesAlertCondition1.md) |  | 
**Recipients** | [**OrganizationsOrganizationIdAlertsProfilesRecipients**](OrganizationsOrganizationIdAlertsProfilesRecipients.md) |  | 
**NetworkTags** | **[]string** | Networks with these tags will be monitored for the alert | 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewInlineObject230

`func NewInlineObject230(type_ string, alertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition1, recipients OrganizationsOrganizationIdAlertsProfilesRecipients, networkTags []string, ) *InlineObject230`

NewInlineObject230 instantiates a new InlineObject230 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject230WithDefaults

`func NewInlineObject230WithDefaults() *InlineObject230`

NewInlineObject230WithDefaults instantiates a new InlineObject230 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject230) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject230) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject230) SetType(v string)`

SetType sets Type field to given value.


### GetAlertCondition

`func (o *InlineObject230) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition1`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *InlineObject230) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition1, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *InlineObject230) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition1)`

SetAlertCondition sets AlertCondition field to given value.


### GetRecipients

`func (o *InlineObject230) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject230) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject230) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.


### GetNetworkTags

`func (o *InlineObject230) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineObject230) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineObject230) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.


### GetDescription

`func (o *InlineObject230) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject230) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject230) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject230) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


