# InlineResponse200352

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the profile. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Time the profile was last updated in ISO8601Z. | [optional] 
**Name** | Pointer to **string** | Name of the profile (names should be unique for all profiles defined for the org). | [optional] 
**Description** | Pointer to **string** | The description of the profile. | [optional] 
**Rules** | Pointer to [**[]InlineResponse200349Rules**](InlineResponse200349Rules.md) | The rules associated with this profile. At least one rule and no more than two rules may be defined for a profile. | [optional] 

## Methods

### NewInlineResponse200352

`func NewInlineResponse200352() *InlineResponse200352`

NewInlineResponse200352 instantiates a new InlineResponse200352 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200352WithDefaults

`func NewInlineResponse200352WithDefaults() *InlineResponse200352`

NewInlineResponse200352WithDefaults instantiates a new InlineResponse200352 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200352) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200352) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200352) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200352) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200352) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200352) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200352) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200352) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200352) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200352) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200352) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200352) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200352) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200352) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200352) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200352) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200352) GetRules() []InlineResponse200349Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200352) GetRulesOk() (*[]InlineResponse200349Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200352) SetRules(v []InlineResponse200349Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200352) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


