# InlineResponse200404Items

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
**Sources** | Pointer to [**InlineResponse200404Sources**](InlineResponse200404Sources.md) |  | [optional] 
**Destinations** | Pointer to [**InlineResponse200404Destinations**](InlineResponse200404Destinations.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Rule creation time | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Rule last update time | [optional] 

## Methods

### NewInlineResponse200404Items

`func NewInlineResponse200404Items() *InlineResponse200404Items`

NewInlineResponse200404Items instantiates a new InlineResponse200404Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200404ItemsWithDefaults

`func NewInlineResponse200404ItemsWithDefaults() *InlineResponse200404Items`

NewInlineResponse200404ItemsWithDefaults instantiates a new InlineResponse200404Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200404Items) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200404Items) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200404Items) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200404Items) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200404Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200404Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200404Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200404Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRulesetId

`func (o *InlineResponse200404Items) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *InlineResponse200404Items) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *InlineResponse200404Items) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *InlineResponse200404Items) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200404Items) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200404Items) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200404Items) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200404Items) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200404Items) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200404Items) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200404Items) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200404Items) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse200404Items) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200404Items) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200404Items) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200404Items) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200404Items) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200404Items) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200404Items) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200404Items) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *InlineResponse200404Items) GetSources() InlineResponse200404Sources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InlineResponse200404Items) GetSourcesOk() (*InlineResponse200404Sources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InlineResponse200404Items) SetSources(v InlineResponse200404Sources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *InlineResponse200404Items) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *InlineResponse200404Items) GetDestinations() InlineResponse200404Destinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *InlineResponse200404Items) GetDestinationsOk() (*InlineResponse200404Destinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *InlineResponse200404Items) SetDestinations(v InlineResponse200404Destinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *InlineResponse200404Items) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200404Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200404Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200404Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200404Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200404Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200404Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200404Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200404Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


