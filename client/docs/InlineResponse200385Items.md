# InlineResponse200385Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveId** | Pointer to **string** | ID of the network move operation | [optional] 
**Initiator** | Pointer to [**InlineResponse200385Initiator**](InlineResponse200385Initiator.md) |  | [optional] 
**Organizations** | Pointer to [**InlineResponse200385Organizations**](InlineResponse200385Organizations.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200385Network**](InlineResponse200385Network.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the network move initiated | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp when the network move status last changed | [optional] 
**Result** | Pointer to [**InlineResponse200385Result**](InlineResponse200385Result.md) |  | [optional] 

## Methods

### NewInlineResponse200385Items

`func NewInlineResponse200385Items() *InlineResponse200385Items`

NewInlineResponse200385Items instantiates a new InlineResponse200385Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200385ItemsWithDefaults

`func NewInlineResponse200385ItemsWithDefaults() *InlineResponse200385Items`

NewInlineResponse200385ItemsWithDefaults instantiates a new InlineResponse200385Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveId

`func (o *InlineResponse200385Items) GetMoveId() string`

GetMoveId returns the MoveId field if non-nil, zero value otherwise.

### GetMoveIdOk

`func (o *InlineResponse200385Items) GetMoveIdOk() (*string, bool)`

GetMoveIdOk returns a tuple with the MoveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveId

`func (o *InlineResponse200385Items) SetMoveId(v string)`

SetMoveId sets MoveId field to given value.

### HasMoveId

`func (o *InlineResponse200385Items) HasMoveId() bool`

HasMoveId returns a boolean if a field has been set.

### GetInitiator

`func (o *InlineResponse200385Items) GetInitiator() InlineResponse200385Initiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InlineResponse200385Items) GetInitiatorOk() (*InlineResponse200385Initiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InlineResponse200385Items) SetInitiator(v InlineResponse200385Initiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InlineResponse200385Items) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetOrganizations

`func (o *InlineResponse200385Items) GetOrganizations() InlineResponse200385Organizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *InlineResponse200385Items) GetOrganizationsOk() (*InlineResponse200385Organizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *InlineResponse200385Items) SetOrganizations(v InlineResponse200385Organizations)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *InlineResponse200385Items) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200385Items) GetNetwork() InlineResponse200385Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200385Items) GetNetworkOk() (*InlineResponse200385Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200385Items) SetNetwork(v InlineResponse200385Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200385Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200385Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200385Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200385Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200385Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200385Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200385Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200385Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200385Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse200385Items) GetResult() InlineResponse200385Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse200385Items) GetResultOk() (*InlineResponse200385Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse200385Items) SetResult(v InlineResponse200385Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse200385Items) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


