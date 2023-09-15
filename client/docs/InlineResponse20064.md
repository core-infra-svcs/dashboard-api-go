# InlineResponse20064

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuPercentUsed** | Pointer to **float32** | The percentage of CPU used as a decimal format. | [optional] 
**MemFree** | Pointer to **int32** | Memory that is not yet in use by the system. | [optional] 
**MemWired** | Pointer to **int32** | Memory used for core OS functions on the device. | [optional] 
**MemActive** | Pointer to **int32** | The active RAM on the device. | [optional] 
**MemInactive** | Pointer to **int32** | The inactive RAM on the device. | [optional] 
**NetworkSent** | Pointer to **int32** | Network bandwith transmitted. | [optional] 
**NetworkReceived** | Pointer to **int32** | Network bandwith received. | [optional] 
**SwapUsed** | Pointer to **int32** | The amount of space being used on the startup disk to swap unused files to and from RAM. | [optional] 
**DiskUsage** | Pointer to [**NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage**](NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage.md) |  | [optional] 
**Ts** | Pointer to **string** | The time at which the performance was measured. | [optional] 

## Methods

### NewInlineResponse20064

`func NewInlineResponse20064() *InlineResponse20064`

NewInlineResponse20064 instantiates a new InlineResponse20064 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20064WithDefaults

`func NewInlineResponse20064WithDefaults() *InlineResponse20064`

NewInlineResponse20064WithDefaults instantiates a new InlineResponse20064 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuPercentUsed

`func (o *InlineResponse20064) GetCpuPercentUsed() float32`

GetCpuPercentUsed returns the CpuPercentUsed field if non-nil, zero value otherwise.

### GetCpuPercentUsedOk

`func (o *InlineResponse20064) GetCpuPercentUsedOk() (*float32, bool)`

GetCpuPercentUsedOk returns a tuple with the CpuPercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercentUsed

`func (o *InlineResponse20064) SetCpuPercentUsed(v float32)`

SetCpuPercentUsed sets CpuPercentUsed field to given value.

### HasCpuPercentUsed

`func (o *InlineResponse20064) HasCpuPercentUsed() bool`

HasCpuPercentUsed returns a boolean if a field has been set.

### GetMemFree

`func (o *InlineResponse20064) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *InlineResponse20064) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *InlineResponse20064) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.

### HasMemFree

`func (o *InlineResponse20064) HasMemFree() bool`

HasMemFree returns a boolean if a field has been set.

### GetMemWired

`func (o *InlineResponse20064) GetMemWired() int32`

GetMemWired returns the MemWired field if non-nil, zero value otherwise.

### GetMemWiredOk

`func (o *InlineResponse20064) GetMemWiredOk() (*int32, bool)`

GetMemWiredOk returns a tuple with the MemWired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemWired

`func (o *InlineResponse20064) SetMemWired(v int32)`

SetMemWired sets MemWired field to given value.

### HasMemWired

`func (o *InlineResponse20064) HasMemWired() bool`

HasMemWired returns a boolean if a field has been set.

### GetMemActive

`func (o *InlineResponse20064) GetMemActive() int32`

GetMemActive returns the MemActive field if non-nil, zero value otherwise.

### GetMemActiveOk

`func (o *InlineResponse20064) GetMemActiveOk() (*int32, bool)`

GetMemActiveOk returns a tuple with the MemActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemActive

`func (o *InlineResponse20064) SetMemActive(v int32)`

SetMemActive sets MemActive field to given value.

### HasMemActive

`func (o *InlineResponse20064) HasMemActive() bool`

HasMemActive returns a boolean if a field has been set.

### GetMemInactive

`func (o *InlineResponse20064) GetMemInactive() int32`

GetMemInactive returns the MemInactive field if non-nil, zero value otherwise.

### GetMemInactiveOk

`func (o *InlineResponse20064) GetMemInactiveOk() (*int32, bool)`

GetMemInactiveOk returns a tuple with the MemInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemInactive

`func (o *InlineResponse20064) SetMemInactive(v int32)`

SetMemInactive sets MemInactive field to given value.

### HasMemInactive

`func (o *InlineResponse20064) HasMemInactive() bool`

HasMemInactive returns a boolean if a field has been set.

### GetNetworkSent

`func (o *InlineResponse20064) GetNetworkSent() int32`

GetNetworkSent returns the NetworkSent field if non-nil, zero value otherwise.

### GetNetworkSentOk

`func (o *InlineResponse20064) GetNetworkSentOk() (*int32, bool)`

GetNetworkSentOk returns a tuple with the NetworkSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSent

`func (o *InlineResponse20064) SetNetworkSent(v int32)`

SetNetworkSent sets NetworkSent field to given value.

### HasNetworkSent

`func (o *InlineResponse20064) HasNetworkSent() bool`

HasNetworkSent returns a boolean if a field has been set.

### GetNetworkReceived

`func (o *InlineResponse20064) GetNetworkReceived() int32`

GetNetworkReceived returns the NetworkReceived field if non-nil, zero value otherwise.

### GetNetworkReceivedOk

`func (o *InlineResponse20064) GetNetworkReceivedOk() (*int32, bool)`

GetNetworkReceivedOk returns a tuple with the NetworkReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReceived

`func (o *InlineResponse20064) SetNetworkReceived(v int32)`

SetNetworkReceived sets NetworkReceived field to given value.

### HasNetworkReceived

`func (o *InlineResponse20064) HasNetworkReceived() bool`

HasNetworkReceived returns a boolean if a field has been set.

### GetSwapUsed

`func (o *InlineResponse20064) GetSwapUsed() int32`

GetSwapUsed returns the SwapUsed field if non-nil, zero value otherwise.

### GetSwapUsedOk

`func (o *InlineResponse20064) GetSwapUsedOk() (*int32, bool)`

GetSwapUsedOk returns a tuple with the SwapUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapUsed

`func (o *InlineResponse20064) SetSwapUsed(v int32)`

SetSwapUsed sets SwapUsed field to given value.

### HasSwapUsed

`func (o *InlineResponse20064) HasSwapUsed() bool`

HasSwapUsed returns a boolean if a field has been set.

### GetDiskUsage

`func (o *InlineResponse20064) GetDiskUsage() NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *InlineResponse20064) GetDiskUsageOk() (*NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *InlineResponse20064) SetDiskUsage(v NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *InlineResponse20064) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse20064) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse20064) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse20064) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse20064) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


