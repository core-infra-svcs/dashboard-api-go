# InlineResponse200185

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **time.Time** | Timestamp at which the event occurred | [optional] 
**Band** | Pointer to **string** | Wireless band the event occurred on | [optional] 
**SsidNumber** | Pointer to **int32** | Number of the SSID the event occurred in | [optional] 
**Type** | Pointer to **string** | Event type | [optional] 
**Subtype** | Pointer to **string** | Event subtype | [optional] 
**Severity** | Pointer to **string** | Event severity | [optional] 
**DurationMs** | Pointer to **int32** | Duration of the event in milliseconds | [optional] 
**Channel** | Pointer to **int32** | Wireless channel the event occurred over | [optional] 
**Rssi** | Pointer to **int32** | RSSI recorded at the time of the event | [optional] 
**EventData** | Pointer to **map[string]interface{}** | Additional information relevant to the given event. Properties vary based on event type. | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device the event occurred for | [optional] 
**CaptureId** | Pointer to **string** | Id of the packet capture triggered for the event, if any | [optional] 

## Methods

### NewInlineResponse200185

`func NewInlineResponse200185() *InlineResponse200185`

NewInlineResponse200185 instantiates a new InlineResponse200185 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200185WithDefaults

`func NewInlineResponse200185WithDefaults() *InlineResponse200185`

NewInlineResponse200185WithDefaults instantiates a new InlineResponse200185 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *InlineResponse200185) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *InlineResponse200185) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *InlineResponse200185) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *InlineResponse200185) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetBand

`func (o *InlineResponse200185) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *InlineResponse200185) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *InlineResponse200185) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *InlineResponse200185) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetSsidNumber

`func (o *InlineResponse200185) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200185) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200185) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200185) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200185) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200185) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200185) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200185) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *InlineResponse200185) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InlineResponse200185) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InlineResponse200185) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *InlineResponse200185) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSeverity

`func (o *InlineResponse200185) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InlineResponse200185) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InlineResponse200185) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InlineResponse200185) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDurationMs

`func (o *InlineResponse200185) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *InlineResponse200185) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *InlineResponse200185) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *InlineResponse200185) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse200185) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse200185) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse200185) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse200185) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse200185) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse200185) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse200185) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *InlineResponse200185) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetEventData

`func (o *InlineResponse200185) GetEventData() map[string]interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *InlineResponse200185) GetEventDataOk() (*map[string]interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *InlineResponse200185) SetEventData(v map[string]interface{})`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *InlineResponse200185) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200185) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200185) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200185) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200185) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetCaptureId

`func (o *InlineResponse200185) GetCaptureId() string`

GetCaptureId returns the CaptureId field if non-nil, zero value otherwise.

### GetCaptureIdOk

`func (o *InlineResponse200185) GetCaptureIdOk() (*string, bool)`

GetCaptureIdOk returns a tuple with the CaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureId

`func (o *InlineResponse200185) SetCaptureId(v string)`

SetCaptureId sets CaptureId field to given value.

### HasCaptureId

`func (o *InlineResponse200185) HasCaptureId() bool`

HasCaptureId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


