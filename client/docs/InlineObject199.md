# InlineObject199

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Identity PSK | 
**Passphrase** | Pointer to **string** | The passphrase for client authentication. If left blank, one will be auto-generated. | [optional] 
**GroupPolicyId** | **string** | The group policy to be applied to clients | 
**ExpiresAt** | Pointer to **time.Time** | Timestamp for when the Identity PSK expires. Will not expire if left blank. | [optional] 

## Methods

### NewInlineObject199

`func NewInlineObject199(name string, groupPolicyId string, ) *InlineObject199`

NewInlineObject199 instantiates a new InlineObject199 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject199WithDefaults

`func NewInlineObject199WithDefaults() *InlineObject199`

NewInlineObject199WithDefaults instantiates a new InlineObject199 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject199) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject199) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject199) SetName(v string)`

SetName sets Name field to given value.


### GetPassphrase

`func (o *InlineObject199) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *InlineObject199) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *InlineObject199) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *InlineObject199) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineObject199) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject199) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject199) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.


### GetExpiresAt

`func (o *InlineObject199) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineObject199) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineObject199) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineObject199) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


