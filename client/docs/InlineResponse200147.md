# InlineResponse200147

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | SNMP community string if access is &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]InlineResponse200147Users**](InlineResponse200147Users.md) | SNMP settings if access is &#39;users&#39;. | [optional] 

## Methods

### NewInlineResponse200147

`func NewInlineResponse200147() *InlineResponse200147`

NewInlineResponse200147 instantiates a new InlineResponse200147 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200147WithDefaults

`func NewInlineResponse200147WithDefaults() *InlineResponse200147`

NewInlineResponse200147WithDefaults instantiates a new InlineResponse200147 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineResponse200147) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse200147) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse200147) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse200147) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineResponse200147) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineResponse200147) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineResponse200147) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineResponse200147) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineResponse200147) GetUsers() []InlineResponse200147Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineResponse200147) GetUsersOk() (*[]InlineResponse200147Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineResponse200147) SetUsers(v []InlineResponse200147Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineResponse200147) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


