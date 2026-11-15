# InlineResponse200428Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Network group identifier | [optional] 
**Name** | Pointer to **string** | Network group name | [optional] 
**Clients** | Pointer to [**InlineResponse200428Clients**](InlineResponse200428Clients.md) |  | [optional] 
**Statuses** | Pointer to [**InlineResponse200428Statuses**](InlineResponse200428Statuses.md) |  | [optional] 
**ProductTypes** | Pointer to **[]string** | Product types in network group | [optional] 

## Methods

### NewInlineResponse200428Items

`func NewInlineResponse200428Items() *InlineResponse200428Items`

NewInlineResponse200428Items instantiates a new InlineResponse200428Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200428ItemsWithDefaults

`func NewInlineResponse200428ItemsWithDefaults() *InlineResponse200428Items`

NewInlineResponse200428ItemsWithDefaults instantiates a new InlineResponse200428Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *InlineResponse200428Items) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InlineResponse200428Items) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InlineResponse200428Items) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InlineResponse200428Items) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200428Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200428Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200428Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200428Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200428Items) GetClients() InlineResponse200428Clients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200428Items) GetClientsOk() (*InlineResponse200428Clients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200428Items) SetClients(v InlineResponse200428Clients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200428Items) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStatuses

`func (o *InlineResponse200428Items) GetStatuses() InlineResponse200428Statuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *InlineResponse200428Items) GetStatusesOk() (*InlineResponse200428Statuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *InlineResponse200428Items) SetStatuses(v InlineResponse200428Statuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *InlineResponse200428Items) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200428Items) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200428Items) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200428Items) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200428Items) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


