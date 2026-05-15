# InlineResponse200246Failover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestIp** | Pointer to **string** | IP addressed reserved on DHCP server where SSID will terminate. | [optional] 
**HeartbeatInterval** | Pointer to **int32** | Idle timer interval in seconds. | [optional] 
**IdleTimeout** | Pointer to **int32** | Idle timer timeout in seconds. | [optional] 

## Methods

### NewInlineResponse200246Failover

`func NewInlineResponse200246Failover() *InlineResponse200246Failover`

NewInlineResponse200246Failover instantiates a new InlineResponse200246Failover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200246FailoverWithDefaults

`func NewInlineResponse200246FailoverWithDefaults() *InlineResponse200246Failover`

NewInlineResponse200246FailoverWithDefaults instantiates a new InlineResponse200246Failover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestIp

`func (o *InlineResponse200246Failover) GetRequestIp() string`

GetRequestIp returns the RequestIp field if non-nil, zero value otherwise.

### GetRequestIpOk

`func (o *InlineResponse200246Failover) GetRequestIpOk() (*string, bool)`

GetRequestIpOk returns a tuple with the RequestIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIp

`func (o *InlineResponse200246Failover) SetRequestIp(v string)`

SetRequestIp sets RequestIp field to given value.

### HasRequestIp

`func (o *InlineResponse200246Failover) HasRequestIp() bool`

HasRequestIp returns a boolean if a field has been set.

### GetHeartbeatInterval

`func (o *InlineResponse200246Failover) GetHeartbeatInterval() int32`

GetHeartbeatInterval returns the HeartbeatInterval field if non-nil, zero value otherwise.

### GetHeartbeatIntervalOk

`func (o *InlineResponse200246Failover) GetHeartbeatIntervalOk() (*int32, bool)`

GetHeartbeatIntervalOk returns a tuple with the HeartbeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatInterval

`func (o *InlineResponse200246Failover) SetHeartbeatInterval(v int32)`

SetHeartbeatInterval sets HeartbeatInterval field to given value.

### HasHeartbeatInterval

`func (o *InlineResponse200246Failover) HasHeartbeatInterval() bool`

HasHeartbeatInterval returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *InlineResponse200246Failover) GetIdleTimeout() int32`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *InlineResponse200246Failover) GetIdleTimeoutOk() (*int32, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *InlineResponse200246Failover) SetIdleTimeout(v int32)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *InlineResponse200246Failover) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


