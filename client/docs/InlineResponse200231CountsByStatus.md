# InlineResponse200231CountsByStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | number of uplinks that are active and working | [optional] 
**Ready** | Pointer to **int32** | number of uplinks that are working but on standby | [optional] 
**Failed** | Pointer to **int32** | number of uplinks that were working but have failed | [optional] 
**Connecting** | Pointer to **int32** | number of uplinks currently connecting | [optional] 
**NotConnected** | Pointer to **int32** | number of uplinks currently where nothing is plugged in | [optional] 

## Methods

### NewInlineResponse200231CountsByStatus

`func NewInlineResponse200231CountsByStatus() *InlineResponse200231CountsByStatus`

NewInlineResponse200231CountsByStatus instantiates a new InlineResponse200231CountsByStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200231CountsByStatusWithDefaults

`func NewInlineResponse200231CountsByStatusWithDefaults() *InlineResponse200231CountsByStatus`

NewInlineResponse200231CountsByStatusWithDefaults instantiates a new InlineResponse200231CountsByStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *InlineResponse200231CountsByStatus) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200231CountsByStatus) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200231CountsByStatus) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200231CountsByStatus) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReady

`func (o *InlineResponse200231CountsByStatus) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *InlineResponse200231CountsByStatus) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *InlineResponse200231CountsByStatus) SetReady(v int32)`

SetReady sets Ready field to given value.

### HasReady

`func (o *InlineResponse200231CountsByStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetFailed

`func (o *InlineResponse200231CountsByStatus) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *InlineResponse200231CountsByStatus) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *InlineResponse200231CountsByStatus) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *InlineResponse200231CountsByStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetConnecting

`func (o *InlineResponse200231CountsByStatus) GetConnecting() int32`

GetConnecting returns the Connecting field if non-nil, zero value otherwise.

### GetConnectingOk

`func (o *InlineResponse200231CountsByStatus) GetConnectingOk() (*int32, bool)`

GetConnectingOk returns a tuple with the Connecting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnecting

`func (o *InlineResponse200231CountsByStatus) SetConnecting(v int32)`

SetConnecting sets Connecting field to given value.

### HasConnecting

`func (o *InlineResponse200231CountsByStatus) HasConnecting() bool`

HasConnecting returns a boolean if a field has been set.

### GetNotConnected

`func (o *InlineResponse200231CountsByStatus) GetNotConnected() int32`

GetNotConnected returns the NotConnected field if non-nil, zero value otherwise.

### GetNotConnectedOk

`func (o *InlineResponse200231CountsByStatus) GetNotConnectedOk() (*int32, bool)`

GetNotConnectedOk returns a tuple with the NotConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotConnected

`func (o *InlineResponse200231CountsByStatus) SetNotConnected(v int32)`

SetNotConnected sets NotConnected field to given value.

### HasNotConnected

`func (o *InlineResponse200231CountsByStatus) HasNotConnected() bool`

HasNotConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


