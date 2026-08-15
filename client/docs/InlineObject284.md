# InlineObject284

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the profile. | [optional] 
**Description** | Pointer to **string** | New description of the profile. | [optional] 
**Rules** | [**[]OrganizationsOrganizationIdDevicesCellularDataProfilesRules**](OrganizationsOrganizationIdDevicesCellularDataProfilesRules.md) | The rules associated with this profile. At least one rule and no more than two rules may be defined for a profile. | 

## Methods

### NewInlineObject284

`func NewInlineObject284(rules []OrganizationsOrganizationIdDevicesCellularDataProfilesRules, ) *InlineObject284`

NewInlineObject284 instantiates a new InlineObject284 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject284WithDefaults

`func NewInlineObject284WithDefaults() *InlineObject284`

NewInlineObject284WithDefaults instantiates a new InlineObject284 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineObject284) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineObject284) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineObject284) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineObject284) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject284) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject284) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject284) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject284) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject284) GetRules() []OrganizationsOrganizationIdDevicesCellularDataProfilesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject284) GetRulesOk() (*[]OrganizationsOrganizationIdDevicesCellularDataProfilesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject284) SetRules(v []OrganizationsOrganizationIdDevicesCellularDataProfilesRules)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


