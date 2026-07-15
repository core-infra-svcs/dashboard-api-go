# InlineResponse200389Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | ID of the firewall rule | [optional] 
**Name** | Pointer to **string** | Name of the firewall rule | [optional] 
**RulesetId** | Pointer to **string** | Firewall ruleset ID | [optional] 
**Policy** | Pointer to **string** | Rule policy (allow or deny) | [optional] 
**Enabled** | Pointer to **bool** | Whether the rule is enabled | [optional] 
**Priority** | Pointer to **int32** | Rule priority | [optional] 
**Description** | Pointer to **string** | Description of the firewall rule | [optional] 
**Sources** | Pointer to [**InlineResponse200389Sources**](InlineResponse200389Sources.md) |  | [optional] 
**Destinations** | Pointer to [**InlineResponse200389Destinations**](InlineResponse200389Destinations.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Rule creation time | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Rule last update time | [optional] 

## Methods

### NewInlineResponse200389Items

`func NewInlineResponse200389Items() *InlineResponse200389Items`

NewInlineResponse200389Items instantiates a new InlineResponse200389Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200389ItemsWithDefaults

`func NewInlineResponse200389ItemsWithDefaults() *InlineResponse200389Items`

NewInlineResponse200389ItemsWithDefaults instantiates a new InlineResponse200389Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200389Items) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200389Items) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200389Items) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200389Items) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200389Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200389Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200389Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200389Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRulesetId

`func (o *InlineResponse200389Items) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *InlineResponse200389Items) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *InlineResponse200389Items) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *InlineResponse200389Items) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200389Items) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200389Items) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200389Items) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200389Items) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200389Items) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200389Items) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200389Items) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200389Items) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200389Items) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200389Items) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200389Items) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200389Items) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200389Items) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200389Items) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200389Items) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200389Items) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *InlineResponse200389Items) GetSources() InlineResponse200389Sources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InlineResponse200389Items) GetSourcesOk() (*InlineResponse200389Sources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InlineResponse200389Items) SetSources(v InlineResponse200389Sources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *InlineResponse200389Items) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *InlineResponse200389Items) GetDestinations() InlineResponse200389Destinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *InlineResponse200389Items) GetDestinationsOk() (*InlineResponse200389Destinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *InlineResponse200389Items) SetDestinations(v InlineResponse200389Destinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *InlineResponse200389Items) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200389Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200389Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200389Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200389Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200389Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200389Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200389Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200389Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


