# InlineResponse200318Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReportedAt** | **time.Time** | Time, in ISO8601 format, when RSSI was last reported for the sensor and gateway pair. | 
**LastConnectedAt** | Pointer to **time.Time** | Time, in ISO8601 format, when the sensor and gateway pair last connected. | [optional] 
**Rssi** | **int32** | Last reported RSSI value. | 
**Network** | [**InlineResponse200318Network**](InlineResponse200318Network.md) |  | 
**Sensor** | [**InlineResponse200318Sensor**](InlineResponse200318Sensor.md) |  | 
**Gateway** | [**InlineResponse200318Gateway**](InlineResponse200318Gateway.md) |  | 

## Methods

### NewInlineResponse200318Items

`func NewInlineResponse200318Items(lastReportedAt time.Time, rssi int32, network InlineResponse200318Network, sensor InlineResponse200318Sensor, gateway InlineResponse200318Gateway, ) *InlineResponse200318Items`

NewInlineResponse200318Items instantiates a new InlineResponse200318Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200318ItemsWithDefaults

`func NewInlineResponse200318ItemsWithDefaults() *InlineResponse200318Items`

NewInlineResponse200318ItemsWithDefaults instantiates a new InlineResponse200318Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReportedAt

`func (o *InlineResponse200318Items) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200318Items) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200318Items) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.


### GetLastConnectedAt

`func (o *InlineResponse200318Items) GetLastConnectedAt() time.Time`

GetLastConnectedAt returns the LastConnectedAt field if non-nil, zero value otherwise.

### GetLastConnectedAtOk

`func (o *InlineResponse200318Items) GetLastConnectedAtOk() (*time.Time, bool)`

GetLastConnectedAtOk returns a tuple with the LastConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedAt

`func (o *InlineResponse200318Items) SetLastConnectedAt(v time.Time)`

SetLastConnectedAt sets LastConnectedAt field to given value.

### HasLastConnectedAt

`func (o *InlineResponse200318Items) HasLastConnectedAt() bool`

HasLastConnectedAt returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse200318Items) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse200318Items) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse200318Items) SetRssi(v int32)`

SetRssi sets Rssi field to given value.


### GetNetwork

`func (o *InlineResponse200318Items) GetNetwork() InlineResponse200318Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200318Items) GetNetworkOk() (*InlineResponse200318Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200318Items) SetNetwork(v InlineResponse200318Network)`

SetNetwork sets Network field to given value.


### GetSensor

`func (o *InlineResponse200318Items) GetSensor() InlineResponse200318Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *InlineResponse200318Items) GetSensorOk() (*InlineResponse200318Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *InlineResponse200318Items) SetSensor(v InlineResponse200318Sensor)`

SetSensor sets Sensor field to given value.


### GetGateway

`func (o *InlineResponse200318Items) GetGateway() InlineResponse200318Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200318Items) GetGatewayOk() (*InlineResponse200318Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200318Items) SetGateway(v InlineResponse200318Gateway)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


