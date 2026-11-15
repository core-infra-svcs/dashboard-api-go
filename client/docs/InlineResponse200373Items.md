# InlineResponse200373Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the profile. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Time the profile was last updated in ISO8601Z. | [optional] 
**Name** | Pointer to **string** | Name of the profile (names should be unique for all profiles defined for the org). | [optional] 
**Description** | Pointer to **string** | The description of the profile that was set when it was created. | [optional] 
**Counts** | Pointer to [**InlineResponse200373Counts**](InlineResponse200373Counts.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200373Rules**](InlineResponse200373Rules.md) | The data management rules defined for the profile. At least one rule and no more than two rules may be defined for a profile. | [optional] 

## Methods

### NewInlineResponse200373Items

`func NewInlineResponse200373Items() *InlineResponse200373Items`

NewInlineResponse200373Items instantiates a new InlineResponse200373Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200373ItemsWithDefaults

`func NewInlineResponse200373ItemsWithDefaults() *InlineResponse200373Items`

NewInlineResponse200373ItemsWithDefaults instantiates a new InlineResponse200373Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200373Items) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200373Items) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200373Items) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200373Items) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200373Items) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200373Items) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200373Items) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200373Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200373Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200373Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200373Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200373Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200373Items) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200373Items) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200373Items) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200373Items) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200373Items) GetCounts() InlineResponse200373Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200373Items) GetCountsOk() (*InlineResponse200373Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200373Items) SetCounts(v InlineResponse200373Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200373Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200373Items) GetRules() []InlineResponse200373Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200373Items) GetRulesOk() (*[]InlineResponse200373Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200373Items) SetRules(v []InlineResponse200373Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200373Items) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


