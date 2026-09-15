# InlineResponse200402Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Network group identifier | [optional] 
**Name** | Pointer to **string** | Network group name | [optional] 
**Clients** | Pointer to [**InlineResponse200402Clients**](InlineResponse200402Clients.md) |  | [optional] 
**Statuses** | Pointer to [**InlineResponse200402Statuses**](InlineResponse200402Statuses.md) |  | [optional] 
**ProductTypes** | Pointer to **[]string** | Product types in network group | [optional] 

## Methods

### NewInlineResponse200402Items

`func NewInlineResponse200402Items() *InlineResponse200402Items`

NewInlineResponse200402Items instantiates a new InlineResponse200402Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200402ItemsWithDefaults

`func NewInlineResponse200402ItemsWithDefaults() *InlineResponse200402Items`

NewInlineResponse200402ItemsWithDefaults instantiates a new InlineResponse200402Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *InlineResponse200402Items) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InlineResponse200402Items) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InlineResponse200402Items) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InlineResponse200402Items) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200402Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200402Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200402Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200402Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200402Items) GetClients() InlineResponse200402Clients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200402Items) GetClientsOk() (*InlineResponse200402Clients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200402Items) SetClients(v InlineResponse200402Clients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200402Items) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStatuses

`func (o *InlineResponse200402Items) GetStatuses() InlineResponse200402Statuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *InlineResponse200402Items) GetStatusesOk() (*InlineResponse200402Statuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *InlineResponse200402Items) SetStatuses(v InlineResponse200402Statuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *InlineResponse200402Items) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200402Items) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200402Items) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200402Items) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200402Items) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


