# InlineResponse200423Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReportedAt** | **time.Time** | Time, in ISO8601 format, when RSSI was last reported for the sensor and gateway pair. | 
**LastConnectedAt** | Pointer to **time.Time** | Time, in ISO8601 format, when the sensor and gateway pair last connected. | [optional] 
**Rssi** | **int32** | Last reported RSSI value. | 
**Network** | [**InlineResponse200423Network**](InlineResponse200423Network.md) |  | 
**Sensor** | [**InlineResponse200423Sensor**](InlineResponse200423Sensor.md) |  | 
**Gateway** | [**InlineResponse200423Gateway**](InlineResponse200423Gateway.md) |  | 

## Methods

### NewInlineResponse200423Items

`func NewInlineResponse200423Items(lastReportedAt time.Time, rssi int32, network InlineResponse200423Network, sensor InlineResponse200423Sensor, gateway InlineResponse200423Gateway, ) *InlineResponse200423Items`

NewInlineResponse200423Items instantiates a new InlineResponse200423Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200423ItemsWithDefaults

`func NewInlineResponse200423ItemsWithDefaults() *InlineResponse200423Items`

NewInlineResponse200423ItemsWithDefaults instantiates a new InlineResponse200423Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReportedAt

`func (o *InlineResponse200423Items) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200423Items) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200423Items) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.


### GetLastConnectedAt

`func (o *InlineResponse200423Items) GetLastConnectedAt() time.Time`

GetLastConnectedAt returns the LastConnectedAt field if non-nil, zero value otherwise.

### GetLastConnectedAtOk

`func (o *InlineResponse200423Items) GetLastConnectedAtOk() (*time.Time, bool)`

GetLastConnectedAtOk returns a tuple with the LastConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedAt

`func (o *InlineResponse200423Items) SetLastConnectedAt(v time.Time)`

SetLastConnectedAt sets LastConnectedAt field to given value.

### HasLastConnectedAt

`func (o *InlineResponse200423Items) HasLastConnectedAt() bool`

HasLastConnectedAt returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse200423Items) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse200423Items) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse200423Items) SetRssi(v int32)`

SetRssi sets Rssi field to given value.


### GetNetwork

`func (o *InlineResponse200423Items) GetNetwork() InlineResponse200423Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200423Items) GetNetworkOk() (*InlineResponse200423Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200423Items) SetNetwork(v InlineResponse200423Network)`

SetNetwork sets Network field to given value.


### GetSensor

`func (o *InlineResponse200423Items) GetSensor() InlineResponse200423Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *InlineResponse200423Items) GetSensorOk() (*InlineResponse200423Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *InlineResponse200423Items) SetSensor(v InlineResponse200423Sensor)`

SetSensor sets Sensor field to given value.


### GetGateway

`func (o *InlineResponse200423Items) GetGateway() InlineResponse200423Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200423Items) GetGatewayOk() (*InlineResponse200423Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200423Items) SetGateway(v InlineResponse200423Gateway)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


