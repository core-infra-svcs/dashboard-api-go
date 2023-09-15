# InlineResponse200107

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Identity PSK | [optional] 
**Id** | Pointer to **string** | The unique identifier of the Identity PSK | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy to be applied to clients | [optional] 
**Passphrase** | Pointer to **string** | The passphrase for client authentication | [optional] 
**WifiPersonalNetworkId** | Pointer to **string** | The WiFi Personal Network unique identifier | [optional] 
**Email** | Pointer to **string** | The email associated with the System&#39;s Manager User | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp for when the Identity PSK expires, or &#39;null&#39; to never expire | [optional] 

## Methods

### NewInlineResponse200107

`func NewInlineResponse200107() *InlineResponse200107`

NewInlineResponse200107 instantiates a new InlineResponse200107 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200107WithDefaults

`func NewInlineResponse200107WithDefaults() *InlineResponse200107`

NewInlineResponse200107WithDefaults instantiates a new InlineResponse200107 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200107) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200107) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200107) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200107) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200107) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200107) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200107) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200107) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200107) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200107) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200107) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200107) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPassphrase

`func (o *InlineResponse200107) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *InlineResponse200107) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *InlineResponse200107) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *InlineResponse200107) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetWifiPersonalNetworkId

`func (o *InlineResponse200107) GetWifiPersonalNetworkId() string`

GetWifiPersonalNetworkId returns the WifiPersonalNetworkId field if non-nil, zero value otherwise.

### GetWifiPersonalNetworkIdOk

`func (o *InlineResponse200107) GetWifiPersonalNetworkIdOk() (*string, bool)`

GetWifiPersonalNetworkIdOk returns a tuple with the WifiPersonalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiPersonalNetworkId

`func (o *InlineResponse200107) SetWifiPersonalNetworkId(v string)`

SetWifiPersonalNetworkId sets WifiPersonalNetworkId field to given value.

### HasWifiPersonalNetworkId

`func (o *InlineResponse200107) HasWifiPersonalNetworkId() bool`

HasWifiPersonalNetworkId returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200107) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200107) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200107) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200107) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InlineResponse200107) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineResponse200107) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineResponse200107) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineResponse200107) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


