# InlineObject178

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The alert type | 
**AlertCondition** | [**OrganizationsOrganizationIdAlertsProfilesAlertCondition**](OrganizationsOrganizationIdAlertsProfilesAlertCondition.md) |  | 
**Recipients** | [**OrganizationsOrganizationIdAlertsProfilesRecipients**](OrganizationsOrganizationIdAlertsProfilesRecipients.md) |  | 
**NetworkTags** | **[]string** | Networks with these tags will be monitored for the alert | 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewInlineObject178

`func NewInlineObject178(type_ string, alertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition, recipients OrganizationsOrganizationIdAlertsProfilesRecipients, networkTags []string, ) *InlineObject178`

NewInlineObject178 instantiates a new InlineObject178 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject178WithDefaults

`func NewInlineObject178WithDefaults() *InlineObject178`

NewInlineObject178WithDefaults instantiates a new InlineObject178 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineObject178) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject178) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject178) SetType(v string)`

SetType sets Type field to given value.


### GetAlertCondition

`func (o *InlineObject178) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *InlineObject178) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *InlineObject178) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition)`

SetAlertCondition sets AlertCondition field to given value.


### GetRecipients

`func (o *InlineObject178) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject178) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject178) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.


### GetNetworkTags

`func (o *InlineObject178) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineObject178) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineObject178) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.


### GetDescription

`func (o *InlineObject178) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject178) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject178) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject178) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


