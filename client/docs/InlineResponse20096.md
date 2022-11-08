# InlineResponse20096

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID associated with the SAML role | [optional] 
**Role** | Pointer to **string** | The role of the SAML administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the SAML administrator on the organization | [optional] 
**Networks** | Pointer to [**[]InlineResponse20096Networks**](InlineResponse20096Networks.md) | The list of networks that the SAML administrator has privileges on | [optional] 
**Tags** | Pointer to [**[]InlineResponse20096Tags**](InlineResponse20096Tags.md) | The list of tags that the SAML administrator has privleges on | [optional] 

## Methods

### NewInlineResponse20096

`func NewInlineResponse20096() *InlineResponse20096`

NewInlineResponse20096 instantiates a new InlineResponse20096 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20096WithDefaults

`func NewInlineResponse20096WithDefaults() *InlineResponse20096`

NewInlineResponse20096WithDefaults instantiates a new InlineResponse20096 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20096) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20096) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20096) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20096) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *InlineResponse20096) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InlineResponse20096) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InlineResponse20096) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InlineResponse20096) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetOrgAccess

`func (o *InlineResponse20096) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineResponse20096) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineResponse20096) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *InlineResponse20096) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineResponse20096) GetNetworks() []InlineResponse20096Networks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse20096) GetNetworksOk() (*[]InlineResponse20096Networks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse20096) SetNetworks(v []InlineResponse20096Networks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineResponse20096) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20096) GetTags() []InlineResponse20096Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20096) GetTagsOk() (*[]InlineResponse20096Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20096) SetTags(v []InlineResponse20096Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20096) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


