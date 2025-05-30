# InlineResponse200213Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **int32** | Number of adaptive policy groups currently in the organization. | [optional] 
**CustomGroups** | Pointer to **int32** | Number of user-created adaptive policy groups currently in the organization. | [optional] 
**CustomAcls** | Pointer to **int32** | Number of user-created adaptive policy ACLs currently in the organization. | [optional] 
**Policies** | Pointer to **int32** | Number of adaptive policies currently in the organization. | [optional] 
**DenyPolicies** | Pointer to **int32** | Number of adaptive policies currently in the organization that deny all traffic. | [optional] 
**AllowPolicies** | Pointer to **int32** | Number of adaptive policies currently in the organization that allow all traffic. | [optional] 
**PolicyObjects** | Pointer to **int32** | Number of policy objects (with the adaptive policy type) currently in the organization. | [optional] 

## Methods

### NewInlineResponse200213Counts

`func NewInlineResponse200213Counts() *InlineResponse200213Counts`

NewInlineResponse200213Counts instantiates a new InlineResponse200213Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200213CountsWithDefaults

`func NewInlineResponse200213CountsWithDefaults() *InlineResponse200213Counts`

NewInlineResponse200213CountsWithDefaults instantiates a new InlineResponse200213Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *InlineResponse200213Counts) GetGroups() int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *InlineResponse200213Counts) GetGroupsOk() (*int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *InlineResponse200213Counts) SetGroups(v int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *InlineResponse200213Counts) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetCustomGroups

`func (o *InlineResponse200213Counts) GetCustomGroups() int32`

GetCustomGroups returns the CustomGroups field if non-nil, zero value otherwise.

### GetCustomGroupsOk

`func (o *InlineResponse200213Counts) GetCustomGroupsOk() (*int32, bool)`

GetCustomGroupsOk returns a tuple with the CustomGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGroups

`func (o *InlineResponse200213Counts) SetCustomGroups(v int32)`

SetCustomGroups sets CustomGroups field to given value.

### HasCustomGroups

`func (o *InlineResponse200213Counts) HasCustomGroups() bool`

HasCustomGroups returns a boolean if a field has been set.

### GetCustomAcls

`func (o *InlineResponse200213Counts) GetCustomAcls() int32`

GetCustomAcls returns the CustomAcls field if non-nil, zero value otherwise.

### GetCustomAclsOk

`func (o *InlineResponse200213Counts) GetCustomAclsOk() (*int32, bool)`

GetCustomAclsOk returns a tuple with the CustomAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAcls

`func (o *InlineResponse200213Counts) SetCustomAcls(v int32)`

SetCustomAcls sets CustomAcls field to given value.

### HasCustomAcls

`func (o *InlineResponse200213Counts) HasCustomAcls() bool`

HasCustomAcls returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200213Counts) GetPolicies() int32`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200213Counts) GetPoliciesOk() (*int32, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200213Counts) SetPolicies(v int32)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200213Counts) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetDenyPolicies

`func (o *InlineResponse200213Counts) GetDenyPolicies() int32`

GetDenyPolicies returns the DenyPolicies field if non-nil, zero value otherwise.

### GetDenyPoliciesOk

`func (o *InlineResponse200213Counts) GetDenyPoliciesOk() (*int32, bool)`

GetDenyPoliciesOk returns a tuple with the DenyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPolicies

`func (o *InlineResponse200213Counts) SetDenyPolicies(v int32)`

SetDenyPolicies sets DenyPolicies field to given value.

### HasDenyPolicies

`func (o *InlineResponse200213Counts) HasDenyPolicies() bool`

HasDenyPolicies returns a boolean if a field has been set.

### GetAllowPolicies

`func (o *InlineResponse200213Counts) GetAllowPolicies() int32`

GetAllowPolicies returns the AllowPolicies field if non-nil, zero value otherwise.

### GetAllowPoliciesOk

`func (o *InlineResponse200213Counts) GetAllowPoliciesOk() (*int32, bool)`

GetAllowPoliciesOk returns a tuple with the AllowPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPolicies

`func (o *InlineResponse200213Counts) SetAllowPolicies(v int32)`

SetAllowPolicies sets AllowPolicies field to given value.

### HasAllowPolicies

`func (o *InlineResponse200213Counts) HasAllowPolicies() bool`

HasAllowPolicies returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *InlineResponse200213Counts) GetPolicyObjects() int32`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *InlineResponse200213Counts) GetPolicyObjectsOk() (*int32, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *InlineResponse200213Counts) SetPolicyObjects(v int32)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *InlineResponse200213Counts) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


