# InlineResponse200213Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The stack ID | [optional] 
**Name** | Pointer to **string** | The stack name | [optional] 
**Members** | Pointer to **[]map[string]interface{}** | Stack member devices | [optional] 
**Clients** | Pointer to [**InlineResponse200213StackClients**](InlineResponse200213StackClients.md) |  | [optional] 

## Methods

### NewInlineResponse200213Stack

`func NewInlineResponse200213Stack() *InlineResponse200213Stack`

NewInlineResponse200213Stack instantiates a new InlineResponse200213Stack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200213StackWithDefaults

`func NewInlineResponse200213StackWithDefaults() *InlineResponse200213Stack`

NewInlineResponse200213StackWithDefaults instantiates a new InlineResponse200213Stack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200213Stack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200213Stack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200213Stack) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200213Stack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200213Stack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200213Stack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200213Stack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200213Stack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *InlineResponse200213Stack) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineResponse200213Stack) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineResponse200213Stack) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineResponse200213Stack) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200213Stack) GetClients() InlineResponse200213StackClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200213Stack) GetClientsOk() (*InlineResponse200213StackClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200213Stack) SetClients(v InlineResponse200213StackClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200213Stack) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


