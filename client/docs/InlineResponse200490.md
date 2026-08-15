# InlineResponse200490

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**Group** | Pointer to [**InlineResponse200490Group**](InlineResponse200490Group.md) |  | [optional] 
**ClientsUrl** | Pointer to **string** | URL to the network&#39;s clients page in Dashboard | [optional] 
**Assignments** | Pointer to [**[]InlineResponse200490Assignments**](InlineResponse200490Assignments.md) | The SSID profile assignments for the network | [optional] 

## Methods

### NewInlineResponse200490

`func NewInlineResponse200490() *InlineResponse200490`

NewInlineResponse200490 instantiates a new InlineResponse200490 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200490WithDefaults

`func NewInlineResponse200490WithDefaults() *InlineResponse200490`

NewInlineResponse200490WithDefaults instantiates a new InlineResponse200490 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200490) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200490) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200490) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200490) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200490) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200490) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200490) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200490) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200490) GetGroup() InlineResponse200490Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200490) GetGroupOk() (*InlineResponse200490Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200490) SetGroup(v InlineResponse200490Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200490) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetClientsUrl

`func (o *InlineResponse200490) GetClientsUrl() string`

GetClientsUrl returns the ClientsUrl field if non-nil, zero value otherwise.

### GetClientsUrlOk

`func (o *InlineResponse200490) GetClientsUrlOk() (*string, bool)`

GetClientsUrlOk returns a tuple with the ClientsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsUrl

`func (o *InlineResponse200490) SetClientsUrl(v string)`

SetClientsUrl sets ClientsUrl field to given value.

### HasClientsUrl

`func (o *InlineResponse200490) HasClientsUrl() bool`

HasClientsUrl returns a boolean if a field has been set.

### GetAssignments

`func (o *InlineResponse200490) GetAssignments() []InlineResponse200490Assignments`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *InlineResponse200490) GetAssignmentsOk() (*[]InlineResponse200490Assignments, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *InlineResponse200490) SetAssignments(v []InlineResponse200490Assignments)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *InlineResponse200490) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


