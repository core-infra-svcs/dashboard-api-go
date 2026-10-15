# InlineResponse200366

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the profile. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Time the profile was last updated in ISO8601Z. | [optional] 
**Name** | Pointer to **string** | Name of the profile (names should be unique for all profiles defined for the org). | [optional] 
**Description** | Pointer to **string** | The description of the profile. | [optional] 
**Rules** | Pointer to [**[]InlineResponse200363Rules**](InlineResponse200363Rules.md) | The rules associated with this profile. At least one rule and no more than two rules may be defined for a profile. | [optional] 

## Methods

### NewInlineResponse200366

`func NewInlineResponse200366() *InlineResponse200366`

NewInlineResponse200366 instantiates a new InlineResponse200366 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200366WithDefaults

`func NewInlineResponse200366WithDefaults() *InlineResponse200366`

NewInlineResponse200366WithDefaults instantiates a new InlineResponse200366 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200366) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200366) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200366) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200366) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200366) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200366) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200366) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200366) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200366) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200366) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200366) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200366) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200366) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200366) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200366) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200366) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200366) GetRules() []InlineResponse200363Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200366) GetRulesOk() (*[]InlineResponse200363Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200366) SetRules(v []InlineResponse200363Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200366) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


