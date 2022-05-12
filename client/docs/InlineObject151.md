# InlineObject151

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the adaptive policy ACL | [optional] 
**Description** | Pointer to **string** | Description of the adaptive policy ACL | [optional] 
**Rules** | Pointer to [**[]OrganizationsOrganizationIdAdaptivePolicyAclsRules**](OrganizationsOrganizationIdAdaptivePolicyAclsRules.md) | An ordered array of the adaptive policy ACL rules. An empty array will clear the rules. | [optional] 
**IpVersion** | Pointer to **string** | IP version of adpative policy ACL. One of: &#39;any&#39;, &#39;ipv4&#39; or &#39;ipv6&#39; | [optional] 

## Methods

### NewInlineObject151

`func NewInlineObject151() *InlineObject151`

NewInlineObject151 instantiates a new InlineObject151 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject151WithDefaults

`func NewInlineObject151WithDefaults() *InlineObject151`

NewInlineObject151WithDefaults instantiates a new InlineObject151 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject151) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject151) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject151) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject151) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject151) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject151) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject151) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject151) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject151) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject151) GetRulesOk() (*[]OrganizationsOrganizationIdAdaptivePolicyAclsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject151) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject151) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineObject151) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineObject151) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineObject151) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineObject151) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

