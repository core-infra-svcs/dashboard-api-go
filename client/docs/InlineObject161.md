# InlineObject161

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | The SNMP community string. Only relevant if &#39;access&#39; is set to &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]NetworksNetworkIdSnmpUsers**](NetworksNetworkIdSnmpUsers.md) | The list of SNMP users. Only relevant if &#39;access&#39; is set to &#39;users&#39;. | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdSnmpAuthentication**](NetworksNetworkIdSnmpAuthentication.md) |  | [optional] 
**Privacy** | Pointer to [**NetworksNetworkIdSnmpPrivacy**](NetworksNetworkIdSnmpPrivacy.md) |  | [optional] 

## Methods

### NewInlineObject161

`func NewInlineObject161() *InlineObject161`

NewInlineObject161 instantiates a new InlineObject161 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject161WithDefaults

`func NewInlineObject161WithDefaults() *InlineObject161`

NewInlineObject161WithDefaults instantiates a new InlineObject161 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineObject161) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineObject161) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineObject161) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineObject161) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineObject161) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineObject161) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineObject161) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineObject161) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineObject161) GetUsers() []NetworksNetworkIdSnmpUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineObject161) GetUsersOk() (*[]NetworksNetworkIdSnmpUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineObject161) SetUsers(v []NetworksNetworkIdSnmpUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineObject161) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject161) GetAuthentication() NetworksNetworkIdSnmpAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject161) GetAuthenticationOk() (*NetworksNetworkIdSnmpAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject161) SetAuthentication(v NetworksNetworkIdSnmpAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject161) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineObject161) GetPrivacy() NetworksNetworkIdSnmpPrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineObject161) GetPrivacyOk() (*NetworksNetworkIdSnmpPrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineObject161) SetPrivacy(v NetworksNetworkIdSnmpPrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineObject161) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


