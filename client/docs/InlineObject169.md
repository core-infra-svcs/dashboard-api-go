# InlineObject169

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the adaptive policy ACL | 
**Description** | Pointer to **string** | Description of the adaptive policy ACL | [optional] [default to ""]
**Rules** | [**[]OrganizationsOrganizationIdAdaptivePolicyAclsRules**](OrganizationsOrganizationIdAdaptivePolicyAclsRules.md) | An ordered array of the adaptive policy ACL rules. | 
**IpVersion** | **string** | IP version of adpative policy ACL. One of: &#39;any&#39;, &#39;ipv4&#39; or &#39;ipv6&#39; | 

## Methods

### NewInlineObject169

`func NewInlineObject169(name string, rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules, ipVersion string, ) *InlineObject169`

NewInlineObject169 instantiates a new InlineObject169 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject169WithDefaults

`func NewInlineObject169WithDefaults() *InlineObject169`

NewInlineObject169WithDefaults instantiates a new InlineObject169 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject169) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject169) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject169) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InlineObject169) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject169) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject169) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject169) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject169) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject169) GetRulesOk() (*[]OrganizationsOrganizationIdAdaptivePolicyAclsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject169) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules)`

SetRules sets Rules field to given value.


### GetIpVersion

`func (o *InlineObject169) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineObject169) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineObject169) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


