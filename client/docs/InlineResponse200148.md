# InlineResponse200148

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | SNMP community string if access is &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]InlineResponse200148Users**](InlineResponse200148Users.md) | SNMP settings if access is &#39;users&#39;. | [optional] 

## Methods

### NewInlineResponse200148

`func NewInlineResponse200148() *InlineResponse200148`

NewInlineResponse200148 instantiates a new InlineResponse200148 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200148WithDefaults

`func NewInlineResponse200148WithDefaults() *InlineResponse200148`

NewInlineResponse200148WithDefaults instantiates a new InlineResponse200148 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *InlineResponse200148) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse200148) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse200148) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse200148) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *InlineResponse200148) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *InlineResponse200148) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *InlineResponse200148) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *InlineResponse200148) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *InlineResponse200148) GetUsers() []InlineResponse200148Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineResponse200148) GetUsersOk() (*[]InlineResponse200148Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineResponse200148) SetUsers(v []InlineResponse200148Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineResponse200148) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


