# InlineResponse200209

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclId** | Pointer to **string** | ID of the adaptive policy ACL | [optional] 
**Name** | Pointer to **string** | Name of the adaptive policy ACL | [optional] 
**Description** | Pointer to **string** | Description of the adaptive policy ACL | [optional] 
**IpVersion** | Pointer to **string** | IP version of adpative policy ACL | [optional] 
**Rules** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyAclsRules**](OrganizationsOrganizationIdAdaptivePolicyAclsRules.md) | An ordered array of the adaptive policy ACL rules | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the adaptive policy ACL was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the adaptive policy ACL was last updated | [optional] 

## Methods

### NewInlineResponse200209

`func NewInlineResponse200209() *InlineResponse200209`

NewInlineResponse200209 instantiates a new InlineResponse200209 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200209WithDefaults

`func NewInlineResponse200209WithDefaults() *InlineResponse200209`

NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclId

`func (o *InlineResponse200209) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *InlineResponse200209) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *InlineResponse200209) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *InlineResponse200209) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200209) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200209) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200209) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200209) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200209) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200209) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200209) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200209) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse200209) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse200209) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse200209) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse200209) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200209) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200209) GetRulesOk() (*[]OrganizationsOrganizationIdAdaptivePolicyAclsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200209) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200209) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200209) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200209) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200209) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200209) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200209) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200209) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200209) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200209) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


