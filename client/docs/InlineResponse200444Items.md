# InlineResponse200444Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReportedAt** | **time.Time** | Time, in ISO8601 format, when RSSI was last reported for the sensor and gateway pair. | 
**LastConnectedAt** | Pointer to **time.Time** | Time, in ISO8601 format, when the sensor and gateway pair last connected. | [optional] 
**Rssi** | **int32** | Last reported RSSI value. | 
**Network** | [**InlineResponse200444Network**](InlineResponse200444Network.md) |  | 
**Sensor** | [**InlineResponse200444Sensor**](InlineResponse200444Sensor.md) |  | 
**Gateway** | [**InlineResponse200444Gateway**](InlineResponse200444Gateway.md) |  | 

## Methods

### NewInlineResponse200444Items

`func NewInlineResponse200444Items(lastReportedAt time.Time, rssi int32, network InlineResponse200444Network, sensor InlineResponse200444Sensor, gateway InlineResponse200444Gateway, ) *InlineResponse200444Items`

NewInlineResponse200444Items instantiates a new InlineResponse200444Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200444ItemsWithDefaults

`func NewInlineResponse200444ItemsWithDefaults() *InlineResponse200444Items`

NewInlineResponse200444ItemsWithDefaults instantiates a new InlineResponse200444Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReportedAt

`func (o *InlineResponse200444Items) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200444Items) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200444Items) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.


### GetLastConnectedAt

`func (o *InlineResponse200444Items) GetLastConnectedAt() time.Time`

GetLastConnectedAt returns the LastConnectedAt field if non-nil, zero value otherwise.

### GetLastConnectedAtOk

`func (o *InlineResponse200444Items) GetLastConnectedAtOk() (*time.Time, bool)`

GetLastConnectedAtOk returns a tuple with the LastConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedAt

`func (o *InlineResponse200444Items) SetLastConnectedAt(v time.Time)`

SetLastConnectedAt sets LastConnectedAt field to given value.

### HasLastConnectedAt

`func (o *InlineResponse200444Items) HasLastConnectedAt() bool`

HasLastConnectedAt returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse200444Items) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse200444Items) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse200444Items) SetRssi(v int32)`

SetRssi sets Rssi field to given value.


### GetNetwork

`func (o *InlineResponse200444Items) GetNetwork() InlineResponse200444Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200444Items) GetNetworkOk() (*InlineResponse200444Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200444Items) SetNetwork(v InlineResponse200444Network)`

SetNetwork sets Network field to given value.


### GetSensor

`func (o *InlineResponse200444Items) GetSensor() InlineResponse200444Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *InlineResponse200444Items) GetSensorOk() (*InlineResponse200444Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *InlineResponse200444Items) SetSensor(v InlineResponse200444Sensor)`

SetSensor sets Sensor field to given value.


### GetGateway

`func (o *InlineResponse200444Items) GetGateway() InlineResponse200444Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200444Items) GetGatewayOk() (*InlineResponse200444Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200444Items) SetGateway(v InlineResponse200444Gateway)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


