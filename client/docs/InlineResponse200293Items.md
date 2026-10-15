# InlineResponse200293Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **NullableString** | Identifier for the policy (present for default base policies). | [optional] 
**Name** | Pointer to **string** | Display name of the policy. | [optional] 
**Description** | Pointer to **NullableString** | Details about the policy. | [optional] 
**BasePolicy** | Pointer to [**NullableInlineResponse200293BasePolicy**](InlineResponse200293BasePolicy.md) |  | [optional] 
**IsBasePolicy** | Pointer to **bool** | Indicates whether the policy is one of the default base policies. | [optional] 

## Methods

### NewInlineResponse200293Items

`func NewInlineResponse200293Items() *InlineResponse200293Items`

NewInlineResponse200293Items instantiates a new InlineResponse200293Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200293ItemsWithDefaults

`func NewInlineResponse200293ItemsWithDefaults() *InlineResponse200293Items`

NewInlineResponse200293ItemsWithDefaults instantiates a new InlineResponse200293Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *InlineResponse200293Items) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *InlineResponse200293Items) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *InlineResponse200293Items) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *InlineResponse200293Items) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *InlineResponse200293Items) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *InlineResponse200293Items) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetName

`func (o *InlineResponse200293Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200293Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200293Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200293Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200293Items) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200293Items) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200293Items) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200293Items) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse200293Items) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse200293Items) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBasePolicy

`func (o *InlineResponse200293Items) GetBasePolicy() InlineResponse200293BasePolicy`

GetBasePolicy returns the BasePolicy field if non-nil, zero value otherwise.

### GetBasePolicyOk

`func (o *InlineResponse200293Items) GetBasePolicyOk() (*InlineResponse200293BasePolicy, bool)`

GetBasePolicyOk returns a tuple with the BasePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePolicy

`func (o *InlineResponse200293Items) SetBasePolicy(v InlineResponse200293BasePolicy)`

SetBasePolicy sets BasePolicy field to given value.

### HasBasePolicy

`func (o *InlineResponse200293Items) HasBasePolicy() bool`

HasBasePolicy returns a boolean if a field has been set.

### SetBasePolicyNil

`func (o *InlineResponse200293Items) SetBasePolicyNil(b bool)`

 SetBasePolicyNil sets the value for BasePolicy to be an explicit nil

### UnsetBasePolicy
`func (o *InlineResponse200293Items) UnsetBasePolicy()`

UnsetBasePolicy ensures that no value is present for BasePolicy, not even an explicit nil
### GetIsBasePolicy

`func (o *InlineResponse200293Items) GetIsBasePolicy() bool`

GetIsBasePolicy returns the IsBasePolicy field if non-nil, zero value otherwise.

### GetIsBasePolicyOk

`func (o *InlineResponse200293Items) GetIsBasePolicyOk() (*bool, bool)`

GetIsBasePolicyOk returns a tuple with the IsBasePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBasePolicy

`func (o *InlineResponse200293Items) SetIsBasePolicy(v bool)`

SetIsBasePolicy sets IsBasePolicy field to given value.

### HasIsBasePolicy

`func (o *InlineResponse200293Items) HasIsBasePolicy() bool`

HasIsBasePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


