# InlineResponse200355Overall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total usage of all queried interfaces during the interval, unit is bit/sec | [optional] 
**Recv** | Pointer to **int32** | The received usage of all queried interfaces during the interval, unit is bit/sec | [optional] 
**Send** | Pointer to **int32** | The sent usage of all queried interfaces during the interval, unit is bit/sec | [optional] 

## Methods

### NewInlineResponse200355Overall

`func NewInlineResponse200355Overall() *InlineResponse200355Overall`

NewInlineResponse200355Overall instantiates a new InlineResponse200355Overall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200355OverallWithDefaults

`func NewInlineResponse200355OverallWithDefaults() *InlineResponse200355Overall`

NewInlineResponse200355OverallWithDefaults instantiates a new InlineResponse200355Overall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200355Overall) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200355Overall) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200355Overall) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200355Overall) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse200355Overall) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse200355Overall) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse200355Overall) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse200355Overall) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetSend

`func (o *InlineResponse200355Overall) GetSend() int32`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *InlineResponse200355Overall) GetSendOk() (*int32, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *InlineResponse200355Overall) SetSend(v int32)`

SetSend sets Send field to given value.

### HasSend

`func (o *InlineResponse200355Overall) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


