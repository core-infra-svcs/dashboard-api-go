# InlineResponse200418Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveId** | Pointer to **string** | ID of the network move operation | [optional] 
**Initiator** | Pointer to [**InlineResponse200418Initiator**](InlineResponse200418Initiator.md) |  | [optional] 
**Organizations** | Pointer to [**InlineResponse200418Organizations**](InlineResponse200418Organizations.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200418Network**](InlineResponse200418Network.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the network move initiated | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp when the network move status last changed | [optional] 
**Result** | Pointer to [**InlineResponse200418Result**](InlineResponse200418Result.md) |  | [optional] 

## Methods

### NewInlineResponse200418Items

`func NewInlineResponse200418Items() *InlineResponse200418Items`

NewInlineResponse200418Items instantiates a new InlineResponse200418Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200418ItemsWithDefaults

`func NewInlineResponse200418ItemsWithDefaults() *InlineResponse200418Items`

NewInlineResponse200418ItemsWithDefaults instantiates a new InlineResponse200418Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveId

`func (o *InlineResponse200418Items) GetMoveId() string`

GetMoveId returns the MoveId field if non-nil, zero value otherwise.

### GetMoveIdOk

`func (o *InlineResponse200418Items) GetMoveIdOk() (*string, bool)`

GetMoveIdOk returns a tuple with the MoveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveId

`func (o *InlineResponse200418Items) SetMoveId(v string)`

SetMoveId sets MoveId field to given value.

### HasMoveId

`func (o *InlineResponse200418Items) HasMoveId() bool`

HasMoveId returns a boolean if a field has been set.

### GetInitiator

`func (o *InlineResponse200418Items) GetInitiator() InlineResponse200418Initiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InlineResponse200418Items) GetInitiatorOk() (*InlineResponse200418Initiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InlineResponse200418Items) SetInitiator(v InlineResponse200418Initiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InlineResponse200418Items) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetOrganizations

`func (o *InlineResponse200418Items) GetOrganizations() InlineResponse200418Organizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *InlineResponse200418Items) GetOrganizationsOk() (*InlineResponse200418Organizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *InlineResponse200418Items) SetOrganizations(v InlineResponse200418Organizations)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *InlineResponse200418Items) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200418Items) GetNetwork() InlineResponse200418Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200418Items) GetNetworkOk() (*InlineResponse200418Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200418Items) SetNetwork(v InlineResponse200418Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200418Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200418Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200418Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200418Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200418Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200418Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200418Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200418Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200418Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse200418Items) GetResult() InlineResponse200418Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse200418Items) GetResultOk() (*InlineResponse200418Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse200418Items) SetResult(v InlineResponse200418Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse200418Items) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


