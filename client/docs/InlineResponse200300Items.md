# InlineResponse200300Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | Identifier for the rule | [optional] 
**Gid** | Pointer to **int32** | Snort generator identifier | [optional] 
**Sid** | Pointer to **int32** | Snort signature identifier | [optional] 
**Name** | Pointer to **string** | Display name for the rule | [optional] 
**Description** | Pointer to **string** | Description of the rule | [optional] 
**Base** | Pointer to [**InlineResponse200300Base**](InlineResponse200300Base.md) |  | [optional] 
**Override** | Pointer to [**NullableInlineResponse200300Override**](InlineResponse200300Override.md) |  | [optional] 
**Policy** | Pointer to [**InlineResponse200300Policy**](InlineResponse200300Policy.md) |  | [optional] 

## Methods

### NewInlineResponse200300Items

`func NewInlineResponse200300Items() *InlineResponse200300Items`

NewInlineResponse200300Items instantiates a new InlineResponse200300Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200300ItemsWithDefaults

`func NewInlineResponse200300ItemsWithDefaults() *InlineResponse200300Items`

NewInlineResponse200300ItemsWithDefaults instantiates a new InlineResponse200300Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200300Items) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200300Items) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200300Items) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200300Items) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetGid

`func (o *InlineResponse200300Items) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *InlineResponse200300Items) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *InlineResponse200300Items) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *InlineResponse200300Items) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetSid

`func (o *InlineResponse200300Items) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *InlineResponse200300Items) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *InlineResponse200300Items) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *InlineResponse200300Items) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200300Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200300Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200300Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200300Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200300Items) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200300Items) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200300Items) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200300Items) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBase

`func (o *InlineResponse200300Items) GetBase() InlineResponse200300Base`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *InlineResponse200300Items) GetBaseOk() (*InlineResponse200300Base, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *InlineResponse200300Items) SetBase(v InlineResponse200300Base)`

SetBase sets Base field to given value.

### HasBase

`func (o *InlineResponse200300Items) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetOverride

`func (o *InlineResponse200300Items) GetOverride() InlineResponse200300Override`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *InlineResponse200300Items) GetOverrideOk() (*InlineResponse200300Override, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *InlineResponse200300Items) SetOverride(v InlineResponse200300Override)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *InlineResponse200300Items) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### SetOverrideNil

`func (o *InlineResponse200300Items) SetOverrideNil(b bool)`

 SetOverrideNil sets the value for Override to be an explicit nil

### UnsetOverride
`func (o *InlineResponse200300Items) UnsetOverride()`

UnsetOverride ensures that no value is present for Override, not even an explicit nil
### GetPolicy

`func (o *InlineResponse200300Items) GetPolicy() InlineResponse200300Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200300Items) GetPolicyOk() (*InlineResponse200300Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200300Items) SetPolicy(v InlineResponse200300Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200300Items) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


