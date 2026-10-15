# InlineResponse200416Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Network group identifier | [optional] 
**Name** | Pointer to **string** | Network group name | [optional] 
**Clients** | Pointer to [**InlineResponse200416Clients**](InlineResponse200416Clients.md) |  | [optional] 
**Statuses** | Pointer to [**InlineResponse200416Statuses**](InlineResponse200416Statuses.md) |  | [optional] 
**ProductTypes** | Pointer to **[]string** | Product types in network group | [optional] 

## Methods

### NewInlineResponse200416Items

`func NewInlineResponse200416Items() *InlineResponse200416Items`

NewInlineResponse200416Items instantiates a new InlineResponse200416Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200416ItemsWithDefaults

`func NewInlineResponse200416ItemsWithDefaults() *InlineResponse200416Items`

NewInlineResponse200416ItemsWithDefaults instantiates a new InlineResponse200416Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *InlineResponse200416Items) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InlineResponse200416Items) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InlineResponse200416Items) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InlineResponse200416Items) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200416Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200416Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200416Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200416Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200416Items) GetClients() InlineResponse200416Clients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200416Items) GetClientsOk() (*InlineResponse200416Clients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200416Items) SetClients(v InlineResponse200416Clients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200416Items) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStatuses

`func (o *InlineResponse200416Items) GetStatuses() InlineResponse200416Statuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *InlineResponse200416Items) GetStatusesOk() (*InlineResponse200416Statuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *InlineResponse200416Items) SetStatuses(v InlineResponse200416Statuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *InlineResponse200416Items) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200416Items) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200416Items) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200416Items) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200416Items) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


