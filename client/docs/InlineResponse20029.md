# InlineResponse20029

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Name of the MG. | [optional] 
**DeviceLanIp** | Pointer to **string** | Lan IP of the MG | [optional] 
**DeviceSubnet** | Pointer to **string** | Subnet configuration of the MG. | [optional] 
**FixedIpAssignments** | Pointer to [**[]InlineResponse20029FixedIpAssignments**](InlineResponse20029FixedIpAssignments.md) | list of all fixed IP assignments for a single MG | [optional] 
**ReservedIpRanges** | Pointer to [**[]InlineResponse20029ReservedIpRanges**](InlineResponse20029ReservedIpRanges.md) | list of all reserved IP ranges for a single MG | [optional] 

## Methods

### NewInlineResponse20029

`func NewInlineResponse20029() *InlineResponse20029`

NewInlineResponse20029 instantiates a new InlineResponse20029 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029WithDefaults

`func NewInlineResponse20029WithDefaults() *InlineResponse20029`

NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *InlineResponse20029) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *InlineResponse20029) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *InlineResponse20029) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *InlineResponse20029) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceLanIp

`func (o *InlineResponse20029) GetDeviceLanIp() string`

GetDeviceLanIp returns the DeviceLanIp field if non-nil, zero value otherwise.

### GetDeviceLanIpOk

`func (o *InlineResponse20029) GetDeviceLanIpOk() (*string, bool)`

GetDeviceLanIpOk returns a tuple with the DeviceLanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLanIp

`func (o *InlineResponse20029) SetDeviceLanIp(v string)`

SetDeviceLanIp sets DeviceLanIp field to given value.

### HasDeviceLanIp

`func (o *InlineResponse20029) HasDeviceLanIp() bool`

HasDeviceLanIp returns a boolean if a field has been set.

### GetDeviceSubnet

`func (o *InlineResponse20029) GetDeviceSubnet() string`

GetDeviceSubnet returns the DeviceSubnet field if non-nil, zero value otherwise.

### GetDeviceSubnetOk

`func (o *InlineResponse20029) GetDeviceSubnetOk() (*string, bool)`

GetDeviceSubnetOk returns a tuple with the DeviceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSubnet

`func (o *InlineResponse20029) SetDeviceSubnet(v string)`

SetDeviceSubnet sets DeviceSubnet field to given value.

### HasDeviceSubnet

`func (o *InlineResponse20029) HasDeviceSubnet() bool`

HasDeviceSubnet returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *InlineResponse20029) GetFixedIpAssignments() []InlineResponse20029FixedIpAssignments`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *InlineResponse20029) GetFixedIpAssignmentsOk() (*[]InlineResponse20029FixedIpAssignments, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *InlineResponse20029) SetFixedIpAssignments(v []InlineResponse20029FixedIpAssignments)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *InlineResponse20029) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *InlineResponse20029) GetReservedIpRanges() []InlineResponse20029ReservedIpRanges`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *InlineResponse20029) GetReservedIpRangesOk() (*[]InlineResponse20029ReservedIpRanges, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *InlineResponse20029) SetReservedIpRanges(v []InlineResponse20029ReservedIpRanges)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *InlineResponse20029) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


