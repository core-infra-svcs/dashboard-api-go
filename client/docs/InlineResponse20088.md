# InlineResponse20088

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | The start time from which daily traffic data was collected | [optional] 
**Application** | Pointer to **string** | The name of the application the client is connected to | [optional] 
**Destination** | Pointer to **string** | The IP or web address the client is connected to | [optional] 
**Protocol** | Pointer to **string** | The client protocol | [optional] 
**Port** | Pointer to **int32** | The port the client is connected to | [optional] 
**Recv** | Pointer to **float32** | Usage received by the client | [optional] 
**Sent** | Pointer to **float32** | Usage sent by the client | [optional] 
**NumFlows** | Pointer to **int32** | The number of flows the client has | [optional] 
**ActiveSeconds** | Pointer to **int32** | The amount of seconds the client was active | [optional] 

## Methods

### NewInlineResponse20088

`func NewInlineResponse20088() *InlineResponse20088`

NewInlineResponse20088 instantiates a new InlineResponse20088 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20088WithDefaults

`func NewInlineResponse20088WithDefaults() *InlineResponse20088`

NewInlineResponse20088WithDefaults instantiates a new InlineResponse20088 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse20088) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse20088) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse20088) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse20088) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetApplication

`func (o *InlineResponse20088) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *InlineResponse20088) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *InlineResponse20088) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *InlineResponse20088) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse20088) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20088) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20088) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse20088) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20088) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20088) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20088) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20088) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse20088) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20088) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20088) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20088) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse20088) GetRecv() float32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse20088) GetRecvOk() (*float32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse20088) SetRecv(v float32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse20088) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetSent

`func (o *InlineResponse20088) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse20088) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse20088) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse20088) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetNumFlows

`func (o *InlineResponse20088) GetNumFlows() int32`

GetNumFlows returns the NumFlows field if non-nil, zero value otherwise.

### GetNumFlowsOk

`func (o *InlineResponse20088) GetNumFlowsOk() (*int32, bool)`

GetNumFlowsOk returns a tuple with the NumFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFlows

`func (o *InlineResponse20088) SetNumFlows(v int32)`

SetNumFlows sets NumFlows field to given value.

### HasNumFlows

`func (o *InlineResponse20088) HasNumFlows() bool`

HasNumFlows returns a boolean if a field has been set.

### GetActiveSeconds

`func (o *InlineResponse20088) GetActiveSeconds() int32`

GetActiveSeconds returns the ActiveSeconds field if non-nil, zero value otherwise.

### GetActiveSecondsOk

`func (o *InlineResponse20088) GetActiveSecondsOk() (*int32, bool)`

GetActiveSecondsOk returns a tuple with the ActiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSeconds

`func (o *InlineResponse20088) SetActiveSeconds(v int32)`

SetActiveSeconds sets ActiveSeconds field to given value.

### HasActiveSeconds

`func (o *InlineResponse20088) HasActiveSeconds() bool`

HasActiveSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


