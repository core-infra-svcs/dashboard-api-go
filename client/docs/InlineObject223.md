# InlineObject223

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Identity PSK | [optional] 
**Passphrase** | Pointer to **string** | The passphrase for client authentication | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy to be applied to clients | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Timestamp for when the Identity PSK expires, or &#39;null&#39; to never expire | [optional] 

## Methods

### NewInlineObject223

`func NewInlineObject223() *InlineObject223`

NewInlineObject223 instantiates a new InlineObject223 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject223WithDefaults

`func NewInlineObject223WithDefaults() *InlineObject223`

NewInlineObject223WithDefaults instantiates a new InlineObject223 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject223) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject223) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject223) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject223) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassphrase

`func (o *InlineObject223) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *InlineObject223) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *InlineObject223) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *InlineObject223) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineObject223) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject223) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject223) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineObject223) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InlineObject223) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineObject223) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineObject223) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineObject223) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *InlineObject223) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *InlineObject223) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


