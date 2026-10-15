# InlineObject248

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The alert type | 
**AlertCondition** | [**OrganizationsOrganizationIdAlertsProfilesAlertCondition1**](OrganizationsOrganizationIdAlertsProfilesAlertCondition1.md) |  | 
**Recipients** | [**OrganizationsOrganizationIdAlertsProfilesRecipients**](OrganizationsOrganizationIdAlertsProfilesRecipients.md) |  | 
**NetworkTags** | **[]string** | Networks with these tags will be monitored for the alert | 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewInlineObject248

`func NewInlineObject248(type_ string, alertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition1, recipients OrganizationsOrganizationIdAlertsProfilesRecipients, networkTags []string, ) *InlineObject248`

NewInlineObject248 instantiates a new InlineObject248 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject248WithDefaults

`func NewInlineObject248WithDefaults() *InlineObject248`

NewInlineObject248WithDefaults instantiates a new InlineObject248 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject248) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject248) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject248) SetType(v string)`

SetType sets Type field to given value.


### GetAlertCondition

`func (o *InlineObject248) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition1`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *InlineObject248) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition1, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *InlineObject248) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition1)`

SetAlertCondition sets AlertCondition field to given value.


### GetRecipients

`func (o *InlineObject248) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject248) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject248) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.


### GetNetworkTags

`func (o *InlineObject248) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineObject248) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineObject248) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.


### GetDescription

`func (o *InlineObject248) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject248) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject248) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject248) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


