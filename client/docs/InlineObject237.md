# InlineObject237

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the group | [optional] 
**Sgt** | Pointer to **int32** | SGT value of the group | [optional] 
**Description** | Pointer to **string** | Description of the group | [optional] 
**PolicyObjects** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects**](OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects.md) | The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group&#39;s SGT value if no other tagging scheme is being used (each requires one unique attribute) | [optional] 

## Methods

### NewInlineObject237

`func NewInlineObject237() *InlineObject237`

NewInlineObject237 instantiates a new InlineObject237 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject237WithDefaults

`func NewInlineObject237WithDefaults() *InlineObject237`

NewInlineObject237WithDefaults instantiates a new InlineObject237 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject237) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject237) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject237) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject237) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSgt

`func (o *InlineObject237) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineObject237) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineObject237) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineObject237) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject237) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject237) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject237) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject237) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineObject237) GetPolicyObjects() []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineObject237) GetPolicyObjectsOk() (*[]OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineObject237) SetPolicyObjects(v []OrganizationsOrganizationIdAdaptivePolicyGroupsPolicyObjects)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineObject237) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


