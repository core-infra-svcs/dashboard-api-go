# InlineResponse200519

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**Group** | Pointer to [**InlineResponse200519Group**](InlineResponse200519Group.md) |  | [optional] 
**ClientsUrl** | Pointer to **string** | URL to the network&#39;s clients page in Dashboard | [optional] 
**Assignments** | Pointer to [**[]InlineResponse200519Assignments**](InlineResponse200519Assignments.md) | The SSID profile assignments for the network | [optional] 

## Methods

### NewInlineResponse200519

`func NewInlineResponse200519() *InlineResponse200519`

NewInlineResponse200519 instantiates a new InlineResponse200519 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200519WithDefaults

`func NewInlineResponse200519WithDefaults() *InlineResponse200519`

NewInlineResponse200519WithDefaults instantiates a new InlineResponse200519 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200519) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200519) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200519) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200519) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200519) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200519) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200519) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200519) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200519) GetGroup() InlineResponse200519Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200519) GetGroupOk() (*InlineResponse200519Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200519) SetGroup(v InlineResponse200519Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200519) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetClientsUrl

`func (o *InlineResponse200519) GetClientsUrl() string`

GetClientsUrl returns the ClientsUrl field if non-nil, zero value otherwise.

### GetClientsUrlOk

`func (o *InlineResponse200519) GetClientsUrlOk() (*string, bool)`

GetClientsUrlOk returns a tuple with the ClientsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsUrl

`func (o *InlineResponse200519) SetClientsUrl(v string)`

SetClientsUrl sets ClientsUrl field to given value.

### HasClientsUrl

`func (o *InlineResponse200519) HasClientsUrl() bool`

HasClientsUrl returns a boolean if a field has been set.

### GetAssignments

`func (o *InlineResponse200519) GetAssignments() []InlineResponse200519Assignments`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *InlineResponse200519) GetAssignmentsOk() (*[]InlineResponse200519Assignments, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *InlineResponse200519) SetAssignments(v []InlineResponse200519Assignments)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *InlineResponse200519) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


