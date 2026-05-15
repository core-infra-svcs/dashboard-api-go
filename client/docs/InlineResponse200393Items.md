# InlineResponse200393Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReportedAt** | **time.Time** | Time, in ISO8601 format, when RSSI was last reported for the sensor and gateway pair. | 
**LastConnectedAt** | Pointer to **time.Time** | Time, in ISO8601 format, when the sensor and gateway pair last connected. | [optional] 
**Rssi** | **int32** | Last reported RSSI value. | 
**Network** | [**InlineResponse200393Network**](InlineResponse200393Network.md) |  | 
**Sensor** | [**InlineResponse200393Sensor**](InlineResponse200393Sensor.md) |  | 
**Gateway** | [**InlineResponse200393Gateway**](InlineResponse200393Gateway.md) |  | 

## Methods

### NewInlineResponse200393Items

`func NewInlineResponse200393Items(lastReportedAt time.Time, rssi int32, network InlineResponse200393Network, sensor InlineResponse200393Sensor, gateway InlineResponse200393Gateway, ) *InlineResponse200393Items`

NewInlineResponse200393Items instantiates a new InlineResponse200393Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200393ItemsWithDefaults

`func NewInlineResponse200393ItemsWithDefaults() *InlineResponse200393Items`

NewInlineResponse200393ItemsWithDefaults instantiates a new InlineResponse200393Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReportedAt

`func (o *InlineResponse200393Items) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200393Items) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200393Items) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.


### GetLastConnectedAt

`func (o *InlineResponse200393Items) GetLastConnectedAt() time.Time`

GetLastConnectedAt returns the LastConnectedAt field if non-nil, zero value otherwise.

### GetLastConnectedAtOk

`func (o *InlineResponse200393Items) GetLastConnectedAtOk() (*time.Time, bool)`

GetLastConnectedAtOk returns a tuple with the LastConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedAt

`func (o *InlineResponse200393Items) SetLastConnectedAt(v time.Time)`

SetLastConnectedAt sets LastConnectedAt field to given value.

### HasLastConnectedAt

`func (o *InlineResponse200393Items) HasLastConnectedAt() bool`

HasLastConnectedAt returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse200393Items) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse200393Items) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse200393Items) SetRssi(v int32)`

SetRssi sets Rssi field to given value.


### GetNetwork

`func (o *InlineResponse200393Items) GetNetwork() InlineResponse200393Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200393Items) GetNetworkOk() (*InlineResponse200393Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200393Items) SetNetwork(v InlineResponse200393Network)`

SetNetwork sets Network field to given value.


### GetSensor

`func (o *InlineResponse200393Items) GetSensor() InlineResponse200393Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *InlineResponse200393Items) GetSensorOk() (*InlineResponse200393Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *InlineResponse200393Items) SetSensor(v InlineResponse200393Sensor)`

SetSensor sets Sensor field to given value.


### GetGateway

`func (o *InlineResponse200393Items) GetGateway() InlineResponse200393Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200393Items) GetGatewayOk() (*InlineResponse200393Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200393Items) SetGateway(v InlineResponse200393Gateway)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


