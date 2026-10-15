# InlineResponse200362Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the profile. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Time the profile was last updated in ISO8601Z. | [optional] 
**Name** | Pointer to **string** | Name of the profile (names should be unique for all profiles defined for the org). | [optional] 
**Description** | Pointer to **string** | The description of the profile that was set when it was created. | [optional] 
**Counts** | Pointer to [**InlineResponse200362Counts**](InlineResponse200362Counts.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200362Rules**](InlineResponse200362Rules.md) | The data management rules defined for the profile. At least one rule and no more than two rules may be defined for a profile. | [optional] 

## Methods

### NewInlineResponse200362Items

`func NewInlineResponse200362Items() *InlineResponse200362Items`

NewInlineResponse200362Items instantiates a new InlineResponse200362Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200362ItemsWithDefaults

`func NewInlineResponse200362ItemsWithDefaults() *InlineResponse200362Items`

NewInlineResponse200362ItemsWithDefaults instantiates a new InlineResponse200362Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200362Items) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200362Items) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200362Items) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200362Items) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200362Items) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200362Items) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200362Items) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200362Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200362Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200362Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200362Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200362Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200362Items) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200362Items) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200362Items) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200362Items) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200362Items) GetCounts() InlineResponse200362Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200362Items) GetCountsOk() (*InlineResponse200362Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200362Items) SetCounts(v InlineResponse200362Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200362Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200362Items) GetRules() []InlineResponse200362Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200362Items) GetRulesOk() (*[]InlineResponse200362Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200362Items) SetRules(v []InlineResponse200362Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200362Items) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


