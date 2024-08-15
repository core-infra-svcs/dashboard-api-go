# InlineResponse200137

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | SNMP community string if access is &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]InlineResponse200137Users**](InlineResponse200137Users.md) | SNMP settings if access is &#39;users&#39;. | [optional] 

## Methods

### NewInlineResponse200137

`func NewInlineResponse200137() *InlineResponse200137`

NewInlineResponse200137 instantiates a new InlineResponse200137 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200137WithDefaults

`func NewInlineResponse200137WithDefaults() *InlineResponse200137`

NewInlineResponse200137WithDefaults instantiates a new InlineResponse200137 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineResponse200137) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse200137) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse200137) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse200137) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineResponse200137) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineResponse200137) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineResponse200137) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineResponse200137) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineResponse200137) GetUsers() []InlineResponse200137Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineResponse200137) GetUsersOk() (*[]InlineResponse200137Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineResponse200137) SetUsers(v []InlineResponse200137Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineResponse200137) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


