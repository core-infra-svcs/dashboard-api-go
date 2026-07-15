# InlineResponse200185

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | SNMP community string if access is &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]InlineResponse200185Users**](InlineResponse200185Users.md) | SNMP settings if access is &#39;users&#39;. | [optional] 
**Authentication** | Pointer to [**InlineResponse200185Authentication**](InlineResponse200185Authentication.md) |  | [optional] 
**Privacy** | Pointer to [**InlineResponse200185Privacy**](InlineResponse200185Privacy.md) |  | [optional] 

## Methods

### NewInlineResponse200185

`func NewInlineResponse200185() *InlineResponse200185`

NewInlineResponse200185 instantiates a new InlineResponse200185 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200185WithDefaults

`func NewInlineResponse200185WithDefaults() *InlineResponse200185`

NewInlineResponse200185WithDefaults instantiates a new InlineResponse200185 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineResponse200185) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse200185) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse200185) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse200185) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineResponse200185) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineResponse200185) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineResponse200185) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineResponse200185) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineResponse200185) GetUsers() []InlineResponse200185Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineResponse200185) GetUsersOk() (*[]InlineResponse200185Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineResponse200185) SetUsers(v []InlineResponse200185Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineResponse200185) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200185) GetAuthentication() InlineResponse200185Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200185) GetAuthenticationOk() (*InlineResponse200185Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200185) SetAuthentication(v InlineResponse200185Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200185) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineResponse200185) GetPrivacy() InlineResponse200185Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineResponse200185) GetPrivacyOk() (*InlineResponse200185Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineResponse200185) SetPrivacy(v InlineResponse200185Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineResponse200185) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


