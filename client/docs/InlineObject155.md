# InlineObject155

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | The SNMP community string. Only relevant if &#39;access&#39; is set to &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]NetworksNetworkIdSnmpUsers**](NetworksNetworkIdSnmpUsers.md) | The list of SNMP users. Only relevant if &#39;access&#39; is set to &#39;users&#39;. | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdSnmpAuthentication**](NetworksNetworkIdSnmpAuthentication.md) |  | [optional] 
**Privacy** | Pointer to [**NetworksNetworkIdSnmpPrivacy**](NetworksNetworkIdSnmpPrivacy.md) |  | [optional] 

## Methods

### NewInlineObject155

`func NewInlineObject155() *InlineObject155`

NewInlineObject155 instantiates a new InlineObject155 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject155WithDefaults

`func NewInlineObject155WithDefaults() *InlineObject155`

NewInlineObject155WithDefaults instantiates a new InlineObject155 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineObject155) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineObject155) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineObject155) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineObject155) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineObject155) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineObject155) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineObject155) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineObject155) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineObject155) GetUsers() []NetworksNetworkIdSnmpUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineObject155) GetUsersOk() (*[]NetworksNetworkIdSnmpUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineObject155) SetUsers(v []NetworksNetworkIdSnmpUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineObject155) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject155) GetAuthentication() NetworksNetworkIdSnmpAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject155) GetAuthenticationOk() (*NetworksNetworkIdSnmpAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject155) SetAuthentication(v NetworksNetworkIdSnmpAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject155) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineObject155) GetPrivacy() NetworksNetworkIdSnmpPrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineObject155) GetPrivacyOk() (*NetworksNetworkIdSnmpPrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineObject155) SetPrivacy(v NetworksNetworkIdSnmpPrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineObject155) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


