# InlineResponse20065DoorLocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DoorLockId** | Pointer to **string** | The Door Lock ID | [optional] 
**Name** | Pointer to **string** | The Door lock name | [optional] 
**ShortId** | Pointer to **string** | The Door lock short ID | [optional] 
**Lqi** | Pointer to **string** | The Door lock L quality index | [optional] 
**Rssi** | Pointer to **string** | The Door lock received signal strength indicator | [optional] 
**Status** | Pointer to **string** | The Door lock status, either online or offline. | [optional] 
**Eui64** | Pointer to **string** | The Door lock unique identifier per network | [optional] 
**EnrolledAt** | Pointer to **time.Time** | Date time the door lock was enrolled | [optional] 
**LastSeenAt** | Pointer to **time.Time** | Date time the door lock last send any messages | [optional] 
**Network** | Pointer to [**InlineResponse20065Network**](InlineResponse20065Network.md) |  | [optional] 
**Gateway** | Pointer to [**InlineResponse20065Gateway**](InlineResponse20065Gateway.md) |  | [optional] 

## Methods

### NewInlineResponse20065DoorLocks

`func NewInlineResponse20065DoorLocks() *InlineResponse20065DoorLocks`

NewInlineResponse20065DoorLocks instantiates a new InlineResponse20065DoorLocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20065DoorLocksWithDefaults

`func NewInlineResponse20065DoorLocksWithDefaults() *InlineResponse20065DoorLocks`

NewInlineResponse20065DoorLocksWithDefaults instantiates a new InlineResponse20065DoorLocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoorLockId

`func (o *InlineResponse20065DoorLocks) GetDoorLockId() string`

GetDoorLockId returns the DoorLockId field if non-nil, zero value otherwise.

### GetDoorLockIdOk

`func (o *InlineResponse20065DoorLocks) GetDoorLockIdOk() (*string, bool)`

GetDoorLockIdOk returns a tuple with the DoorLockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLockId

`func (o *InlineResponse20065DoorLocks) SetDoorLockId(v string)`

SetDoorLockId sets DoorLockId field to given value.

### HasDoorLockId

`func (o *InlineResponse20065DoorLocks) HasDoorLockId() bool`

HasDoorLockId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20065DoorLocks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20065DoorLocks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20065DoorLocks) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20065DoorLocks) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortId

`func (o *InlineResponse20065DoorLocks) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *InlineResponse20065DoorLocks) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *InlineResponse20065DoorLocks) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *InlineResponse20065DoorLocks) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetLqi

`func (o *InlineResponse20065DoorLocks) GetLqi() string`

GetLqi returns the Lqi field if non-nil, zero value otherwise.

### GetLqiOk

`func (o *InlineResponse20065DoorLocks) GetLqiOk() (*string, bool)`

GetLqiOk returns a tuple with the Lqi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLqi

`func (o *InlineResponse20065DoorLocks) SetLqi(v string)`

SetLqi sets Lqi field to given value.

### HasLqi

`func (o *InlineResponse20065DoorLocks) HasLqi() bool`

HasLqi returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse20065DoorLocks) GetRssi() string`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse20065DoorLocks) GetRssiOk() (*string, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse20065DoorLocks) SetRssi(v string)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *InlineResponse20065DoorLocks) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20065DoorLocks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20065DoorLocks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20065DoorLocks) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20065DoorLocks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEui64

`func (o *InlineResponse20065DoorLocks) GetEui64() string`

GetEui64 returns the Eui64 field if non-nil, zero value otherwise.

### GetEui64Ok

`func (o *InlineResponse20065DoorLocks) GetEui64Ok() (*string, bool)`

GetEui64Ok returns a tuple with the Eui64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEui64

`func (o *InlineResponse20065DoorLocks) SetEui64(v string)`

SetEui64 sets Eui64 field to given value.

### HasEui64

`func (o *InlineResponse20065DoorLocks) HasEui64() bool`

HasEui64 returns a boolean if a field has been set.

### GetEnrolledAt

`func (o *InlineResponse20065DoorLocks) GetEnrolledAt() time.Time`

GetEnrolledAt returns the EnrolledAt field if non-nil, zero value otherwise.

### GetEnrolledAtOk

`func (o *InlineResponse20065DoorLocks) GetEnrolledAtOk() (*time.Time, bool)`

GetEnrolledAtOk returns a tuple with the EnrolledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledAt

`func (o *InlineResponse20065DoorLocks) SetEnrolledAt(v time.Time)`

SetEnrolledAt sets EnrolledAt field to given value.

### HasEnrolledAt

`func (o *InlineResponse20065DoorLocks) HasEnrolledAt() bool`

HasEnrolledAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *InlineResponse20065DoorLocks) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *InlineResponse20065DoorLocks) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *InlineResponse20065DoorLocks) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *InlineResponse20065DoorLocks) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20065DoorLocks) GetNetwork() InlineResponse20065Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20065DoorLocks) GetNetworkOk() (*InlineResponse20065Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20065DoorLocks) SetNetwork(v InlineResponse20065Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20065DoorLocks) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse20065DoorLocks) GetGateway() InlineResponse20065Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse20065DoorLocks) GetGatewayOk() (*InlineResponse20065Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse20065DoorLocks) SetGateway(v InlineResponse20065Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse20065DoorLocks) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


