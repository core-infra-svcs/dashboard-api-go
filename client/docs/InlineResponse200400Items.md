# InlineResponse200400Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveId** | Pointer to **string** | ID of the network move operation | [optional] 
**Initiator** | Pointer to [**InlineResponse200400Initiator**](InlineResponse200400Initiator.md) |  | [optional] 
**Organizations** | Pointer to [**InlineResponse200400Organizations**](InlineResponse200400Organizations.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200400Network**](InlineResponse200400Network.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the network move initiated | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp when the network move status last changed | [optional] 
**Result** | Pointer to [**InlineResponse200400Result**](InlineResponse200400Result.md) |  | [optional] 

## Methods

### NewInlineResponse200400Items

`func NewInlineResponse200400Items() *InlineResponse200400Items`

NewInlineResponse200400Items instantiates a new InlineResponse200400Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200400ItemsWithDefaults

`func NewInlineResponse200400ItemsWithDefaults() *InlineResponse200400Items`

NewInlineResponse200400ItemsWithDefaults instantiates a new InlineResponse200400Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveId

`func (o *InlineResponse200400Items) GetMoveId() string`

GetMoveId returns the MoveId field if non-nil, zero value otherwise.

### GetMoveIdOk

`func (o *InlineResponse200400Items) GetMoveIdOk() (*string, bool)`

GetMoveIdOk returns a tuple with the MoveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveId

`func (o *InlineResponse200400Items) SetMoveId(v string)`

SetMoveId sets MoveId field to given value.

### HasMoveId

`func (o *InlineResponse200400Items) HasMoveId() bool`

HasMoveId returns a boolean if a field has been set.

### GetInitiator

`func (o *InlineResponse200400Items) GetInitiator() InlineResponse200400Initiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InlineResponse200400Items) GetInitiatorOk() (*InlineResponse200400Initiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InlineResponse200400Items) SetInitiator(v InlineResponse200400Initiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InlineResponse200400Items) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetOrganizations

`func (o *InlineResponse200400Items) GetOrganizations() InlineResponse200400Organizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *InlineResponse200400Items) GetOrganizationsOk() (*InlineResponse200400Organizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *InlineResponse200400Items) SetOrganizations(v InlineResponse200400Organizations)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *InlineResponse200400Items) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200400Items) GetNetwork() InlineResponse200400Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200400Items) GetNetworkOk() (*InlineResponse200400Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200400Items) SetNetwork(v InlineResponse200400Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200400Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200400Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200400Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200400Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200400Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200400Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200400Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200400Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200400Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse200400Items) GetResult() InlineResponse200400Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse200400Items) GetResultOk() (*InlineResponse200400Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse200400Items) SetResult(v InlineResponse200400Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse200400Items) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


