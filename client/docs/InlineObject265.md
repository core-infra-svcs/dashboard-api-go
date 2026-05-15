# InlineObject265

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the profile to be added. This must be unique. | 
**Description** | **string** | Description of the profile to be added. | 
**Rules** | [**[]OrganizationsOrganizationIdDevicesCellularDataProfilesRules**](OrganizationsOrganizationIdDevicesCellularDataProfilesRules.md) | The rules associated with this profile. At least one rule and no more than two rules may be defined for a profile. | 

## Methods

### NewInlineObject265

`func NewInlineObject265(name string, description string, rules []OrganizationsOrganizationIdDevicesCellularDataProfilesRules, ) *InlineObject265`

NewInlineObject265 instantiates a new InlineObject265 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject265WithDefaults

`func NewInlineObject265WithDefaults() *InlineObject265`

NewInlineObject265WithDefaults instantiates a new InlineObject265 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject265) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject265) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject265) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InlineObject265) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject265) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject265) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *InlineObject265) GetRules() []OrganizationsOrganizationIdDevicesCellularDataProfilesRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject265) GetRulesOk() (*[]OrganizationsOrganizationIdDevicesCellularDataProfilesRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject265) SetRules(v []OrganizationsOrganizationIdDevicesCellularDataProfilesRules)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


