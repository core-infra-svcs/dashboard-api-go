# InlineResponse200176

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Traffic application | [optional] 
**Destination** | Pointer to **string** | Traffic destination | [optional] 
**Protocol** | Pointer to **string** | Traffic protocol | [optional] 
**Port** | Pointer to **int32** | Traffic port | [optional] 
**Sent** | Pointer to **float32** | Traffic sent in kb | [optional] 
**Recv** | Pointer to **float32** | Traffic received in kb | [optional] 
**NumClients** | Pointer to **int32** | Number of clients with traffic | [optional] 
**ActiveTime** | Pointer to **int32** | Active time with traffic | [optional] 
**Flows** | Pointer to **int32** | Number of traffic flows | [optional] 

## Methods

### NewInlineResponse200176

`func NewInlineResponse200176() *InlineResponse200176`

NewInlineResponse200176 instantiates a new InlineResponse200176 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200176WithDefaults

`func NewInlineResponse200176WithDefaults() *InlineResponse200176`

NewInlineResponse200176WithDefaults instantiates a new InlineResponse200176 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *InlineResponse200176) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *InlineResponse200176) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *InlineResponse200176) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *InlineResponse200176) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse200176) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse200176) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse200176) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse200176) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse200176) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200176) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200176) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200176) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse200176) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse200176) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse200176) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse200176) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSent

`func (o *InlineResponse200176) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse200176) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse200176) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse200176) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse200176) GetRecv() float32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse200176) GetRecvOk() (*float32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse200176) SetRecv(v float32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse200176) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetNumClients

`func (o *InlineResponse200176) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *InlineResponse200176) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *InlineResponse200176) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *InlineResponse200176) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetActiveTime

`func (o *InlineResponse200176) GetActiveTime() int32`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *InlineResponse200176) GetActiveTimeOk() (*int32, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *InlineResponse200176) SetActiveTime(v int32)`

SetActiveTime sets ActiveTime field to given value.

### HasActiveTime

`func (o *InlineResponse200176) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.

### GetFlows

`func (o *InlineResponse200176) GetFlows() int32`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *InlineResponse200176) GetFlowsOk() (*int32, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *InlineResponse200176) SetFlows(v int32)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *InlineResponse200176) HasFlows() bool`

HasFlows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


