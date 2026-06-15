# InlineResponse200210Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The stack ID | [optional] 
**Name** | Pointer to **string** | The stack name | [optional] 
**Members** | Pointer to **[]map[string]interface{}** | Stack member devices | [optional] 
**Clients** | Pointer to [**InlineResponse200210StackClients**](InlineResponse200210StackClients.md) |  | [optional] 

## Methods

### NewInlineResponse200210Stack

`func NewInlineResponse200210Stack() *InlineResponse200210Stack`

NewInlineResponse200210Stack instantiates a new InlineResponse200210Stack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200210StackWithDefaults

`func NewInlineResponse200210StackWithDefaults() *InlineResponse200210Stack`

NewInlineResponse200210StackWithDefaults instantiates a new InlineResponse200210Stack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200210Stack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200210Stack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200210Stack) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200210Stack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200210Stack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200210Stack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200210Stack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200210Stack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *InlineResponse200210Stack) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineResponse200210Stack) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineResponse200210Stack) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineResponse200210Stack) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200210Stack) GetClients() InlineResponse200210StackClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200210Stack) GetClientsOk() (*InlineResponse200210StackClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200210Stack) SetClients(v InlineResponse200210StackClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200210Stack) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


