# InlineResponse200430Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveId** | Pointer to **string** | ID of the network move operation | [optional] 
**Initiator** | Pointer to [**InlineResponse200430Initiator**](InlineResponse200430Initiator.md) |  | [optional] 
**Organizations** | Pointer to [**InlineResponse200430Organizations**](InlineResponse200430Organizations.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200430Network**](InlineResponse200430Network.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the network move initiated | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp when the network move status last changed | [optional] 
**Result** | Pointer to [**InlineResponse200430Result**](InlineResponse200430Result.md) |  | [optional] 

## Methods

### NewInlineResponse200430Items

`func NewInlineResponse200430Items() *InlineResponse200430Items`

NewInlineResponse200430Items instantiates a new InlineResponse200430Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200430ItemsWithDefaults

`func NewInlineResponse200430ItemsWithDefaults() *InlineResponse200430Items`

NewInlineResponse200430ItemsWithDefaults instantiates a new InlineResponse200430Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveId

`func (o *InlineResponse200430Items) GetMoveId() string`

GetMoveId returns the MoveId field if non-nil, zero value otherwise.

### GetMoveIdOk

`func (o *InlineResponse200430Items) GetMoveIdOk() (*string, bool)`

GetMoveIdOk returns a tuple with the MoveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveId

`func (o *InlineResponse200430Items) SetMoveId(v string)`

SetMoveId sets MoveId field to given value.

### HasMoveId

`func (o *InlineResponse200430Items) HasMoveId() bool`

HasMoveId returns a boolean if a field has been set.

### GetInitiator

`func (o *InlineResponse200430Items) GetInitiator() InlineResponse200430Initiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InlineResponse200430Items) GetInitiatorOk() (*InlineResponse200430Initiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InlineResponse200430Items) SetInitiator(v InlineResponse200430Initiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InlineResponse200430Items) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetOrganizations

`func (o *InlineResponse200430Items) GetOrganizations() InlineResponse200430Organizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *InlineResponse200430Items) GetOrganizationsOk() (*InlineResponse200430Organizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *InlineResponse200430Items) SetOrganizations(v InlineResponse200430Organizations)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *InlineResponse200430Items) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200430Items) GetNetwork() InlineResponse200430Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200430Items) GetNetworkOk() (*InlineResponse200430Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200430Items) SetNetwork(v InlineResponse200430Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200430Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200430Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200430Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200430Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200430Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200430Items) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200430Items) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200430Items) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200430Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetResult

`func (o *InlineResponse200430Items) GetResult() InlineResponse200430Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InlineResponse200430Items) GetResultOk() (*InlineResponse200430Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InlineResponse200430Items) SetResult(v InlineResponse200430Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *InlineResponse200430Items) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


