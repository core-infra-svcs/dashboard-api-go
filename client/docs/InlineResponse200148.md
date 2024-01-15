# InlineResponse200148

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Serial** | Pointer to **string** | Serial of MX device | [optional] 
**Uplink** | Pointer to **string** | Uplink interface (wan1, wan2, or cellular) | [optional] 
**Ip** | Pointer to **string** | IP address of uplink | [optional] 
**TimeSeries** | Pointer to [**[]OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries**](OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries.md) | Loss and latency timeseries data | [optional] 

## Methods

### NewInlineResponse200148

`func NewInlineResponse200148() *InlineResponse200148`

NewInlineResponse200148 instantiates a new InlineResponse200148 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200148WithDefaults

`func NewInlineResponse200148WithDefaults() *InlineResponse200148`

NewInlineResponse200148WithDefaults instantiates a new InlineResponse200148 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200148) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200148) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200148) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200148) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200148) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200148) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200148) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200148) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200148) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200148) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200148) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200148) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse200148) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse200148) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse200148) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse200148) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTimeSeries

`func (o *InlineResponse200148) GetTimeSeries() []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries`

GetTimeSeries returns the TimeSeries field if non-nil, zero value otherwise.

### GetTimeSeriesOk

`func (o *InlineResponse200148) GetTimeSeriesOk() (*[]OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries, bool)`

GetTimeSeriesOk returns a tuple with the TimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeries

`func (o *InlineResponse200148) SetTimeSeries(v []OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries)`

SetTimeSeries sets TimeSeries field to given value.

### HasTimeSeries

`func (o *InlineResponse200148) HasTimeSeries() bool`

HasTimeSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


