# InlineResponse200370Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveId** | Pointer to **string** | ID of the network move operation | [optional] 
**Initiator** | Pointer to [**InlineResponse200370Initiator**](InlineResponse200370Initiator.md) |  | [optional] 
**Organizations** | Pointer to [**InlineResponse200370Organizations**](InlineResponse200370Organizations.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200370Network**](InlineResponse200370Network.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the network move initiated | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp when the network move status last changed | [optional] 
**Result** | Pointer to [**InlineResponse200370Result**](InlineResponse200370Result.md) |  | [optional] 

## Methods

### NewInlineResponse200370Items

`func NewInlineResponse200370Items() *InlineResponse200370Items`

NewInlineResponse200370Items instantiates a new InlineResponse200370Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200370ItemsWithDefaults

`func NewInlineResponse200370ItemsWithDefaults() *InlineResponse200370Items`

NewInlineResponse200370ItemsWithDefaults instantiates a new InlineResponse200370Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveId

`func (o *InlineResponse200370Items) GetMoveId() string`

GetMoveId returns the MoveId field if non-nil, zero value otherwise.

### GetMoveIdOk

`func (o *InlineResponse200370Items) GetMoveIdOk() (*string, bool)`

GetMoveIdOk returns a tuple with the MoveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveId

`func (o *InlineResponse200370Items) SetMoveId(v string)`

SetMoveId sets MoveId field to given value.

### HasMoveId

`func (o *InlineResponse200370Items) HasMoveId() bool`

HasMoveId returns a boolean if a field has been set.

### GetInitiator

`func (o *InlineResponse200370Items) GetInitiator() InlineResponse200370Initiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InlineResponse200370Items) GetInitiatorOk() (*InlineResponse200370Initiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InlineResponse200370Items) SetInitiator(v InlineResponse200370Initiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InlineResponse200370Items) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetOrganizations

`func (o *InlineResponse200370Items) GetOrganizations() InlineResponse200370Organizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *InlineResponse200370Items) GetOrganizationsOk() (*InlineResponse200370Organizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *InlineResponse200370Items) SetOrganizations(v InlineResponse200370Organizations)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *InlineResponse200370Items) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200370Items) GetNetwork() InlineResponse200370Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200370Items) GetNetworkOk() (*InlineResponse200370Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200370Items) SetNetwork(v InlineResponse200370Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200370Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200370Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200370Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200370Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200370Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200370Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200370Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200370Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200370Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse200370Items) GetResult() InlineResponse200370Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse200370Items) GetResultOk() (*InlineResponse200370Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse200370Items) SetResult(v InlineResponse200370Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse200370Items) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


