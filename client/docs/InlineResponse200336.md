# InlineResponse200336

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the profile. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Time the profile was last updated in ISO8601Z. | [optional] 
**Name** | Pointer to **string** | Name of the profile (names should be unique for all profiles defined for the org). | [optional] 
**Description** | Pointer to **string** | The description of the profile. | [optional] 
**Rules** | Pointer to [**[]InlineResponse200333Rules**](InlineResponse200333Rules.md) | The rules associated with this profile. At least one rule and no more than two rules may be defined for a profile. | [optional] 

## Methods

### NewInlineResponse200336

`func NewInlineResponse200336() *InlineResponse200336`

NewInlineResponse200336 instantiates a new InlineResponse200336 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200336WithDefaults

`func NewInlineResponse200336WithDefaults() *InlineResponse200336`

NewInlineResponse200336WithDefaults instantiates a new InlineResponse200336 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200336) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200336) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200336) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200336) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200336) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200336) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200336) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200336) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200336) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200336) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200336) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200336) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200336) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200336) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200336) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200336) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200336) GetRules() []InlineResponse200333Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200336) GetRulesOk() (*[]InlineResponse200333Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200336) SetRules(v []InlineResponse200333Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200336) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


