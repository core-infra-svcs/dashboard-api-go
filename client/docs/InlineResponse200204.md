# InlineResponse200204

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

### NewInlineResponse200204

`func NewInlineResponse200204() *InlineResponse200204`

NewInlineResponse200204 instantiates a new InlineResponse200204 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200204WithDefaults

`func NewInlineResponse200204WithDefaults() *InlineResponse200204`

NewInlineResponse200204WithDefaults instantiates a new InlineResponse200204 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200204) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200204) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200204) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200204) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200204) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200204) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200204) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200204) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse200204) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse200204) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse200204) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse200204) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetPassphrase

`func (o *InlineResponse200204) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *InlineResponse200204) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *InlineResponse200204) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *InlineResponse200204) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetWifiPersonalNetworkId

`func (o *InlineResponse200204) GetWifiPersonalNetworkId() string`

GetWifiPersonalNetworkId returns the WifiPersonalNetworkId field if non-nil, zero value otherwise.

### GetWifiPersonalNetworkIdOk

`func (o *InlineResponse200204) GetWifiPersonalNetworkIdOk() (*string, bool)`

GetWifiPersonalNetworkIdOk returns a tuple with the WifiPersonalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiPersonalNetworkId

`func (o *InlineResponse200204) SetWifiPersonalNetworkId(v string)`

SetWifiPersonalNetworkId sets WifiPersonalNetworkId field to given value.

### HasWifiPersonalNetworkId

`func (o *InlineResponse200204) HasWifiPersonalNetworkId() bool`

HasWifiPersonalNetworkId returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200204) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200204) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200204) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200204) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InlineResponse200204) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineResponse200204) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineResponse200204) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineResponse200204) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


