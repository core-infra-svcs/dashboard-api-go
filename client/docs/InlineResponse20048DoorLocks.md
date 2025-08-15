# InlineResponse20048DoorLocks

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
**Network** | Pointer to [**InlineResponse20048Network**](InlineResponse20048Network.md) |  | [optional] 
**Gateway** | Pointer to [**InlineResponse20048Gateway**](InlineResponse20048Gateway.md) |  | [optional] 

## Methods

### NewInlineResponse20048DoorLocks

`func NewInlineResponse20048DoorLocks() *InlineResponse20048DoorLocks`

NewInlineResponse20048DoorLocks instantiates a new InlineResponse20048DoorLocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20048DoorLocksWithDefaults

`func NewInlineResponse20048DoorLocksWithDefaults() *InlineResponse20048DoorLocks`

NewInlineResponse20048DoorLocksWithDefaults instantiates a new InlineResponse20048DoorLocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoorLockId

`func (o *InlineResponse20048DoorLocks) GetDoorLockId() string`

GetDoorLockId returns the DoorLockId field if non-nil, zero value otherwise.

### GetDoorLockIdOk

`func (o *InlineResponse20048DoorLocks) GetDoorLockIdOk() (*string, bool)`

GetDoorLockIdOk returns a tuple with the DoorLockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLockId

`func (o *InlineResponse20048DoorLocks) SetDoorLockId(v string)`

SetDoorLockId sets DoorLockId field to given value.

### HasDoorLockId

`func (o *InlineResponse20048DoorLocks) HasDoorLockId() bool`

HasDoorLockId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20048DoorLocks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20048DoorLocks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20048DoorLocks) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20048DoorLocks) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortId

`func (o *InlineResponse20048DoorLocks) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *InlineResponse20048DoorLocks) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *InlineResponse20048DoorLocks) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *InlineResponse20048DoorLocks) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetLqi

`func (o *InlineResponse20048DoorLocks) GetLqi() string`

GetLqi returns the Lqi field if non-nil, zero value otherwise.

### GetLqiOk

`func (o *InlineResponse20048DoorLocks) GetLqiOk() (*string, bool)`

GetLqiOk returns a tuple with the Lqi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLqi

`func (o *InlineResponse20048DoorLocks) SetLqi(v string)`

SetLqi sets Lqi field to given value.

### HasLqi

`func (o *InlineResponse20048DoorLocks) HasLqi() bool`

HasLqi returns a boolean if a field has been set.

### GetRssi

`func (o *InlineResponse20048DoorLocks) GetRssi() string`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *InlineResponse20048DoorLocks) GetRssiOk() (*string, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *InlineResponse20048DoorLocks) SetRssi(v string)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *InlineResponse20048DoorLocks) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20048DoorLocks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20048DoorLocks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20048DoorLocks) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20048DoorLocks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEui64

`func (o *InlineResponse20048DoorLocks) GetEui64() string`

GetEui64 returns the Eui64 field if non-nil, zero value otherwise.

### GetEui64Ok

`func (o *InlineResponse20048DoorLocks) GetEui64Ok() (*string, bool)`

GetEui64Ok returns a tuple with the Eui64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEui64

`func (o *InlineResponse20048DoorLocks) SetEui64(v string)`

SetEui64 sets Eui64 field to given value.

### HasEui64

`func (o *InlineResponse20048DoorLocks) HasEui64() bool`

HasEui64 returns a boolean if a field has been set.

### GetEnrolledAt

`func (o *InlineResponse20048DoorLocks) GetEnrolledAt() time.Time`

GetEnrolledAt returns the EnrolledAt field if non-nil, zero value otherwise.

### GetEnrolledAtOk

`func (o *InlineResponse20048DoorLocks) GetEnrolledAtOk() (*time.Time, bool)`

GetEnrolledAtOk returns a tuple with the EnrolledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledAt

`func (o *InlineResponse20048DoorLocks) SetEnrolledAt(v time.Time)`

SetEnrolledAt sets EnrolledAt field to given value.

### HasEnrolledAt

`func (o *InlineResponse20048DoorLocks) HasEnrolledAt() bool`

HasEnrolledAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *InlineResponse20048DoorLocks) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *InlineResponse20048DoorLocks) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *InlineResponse20048DoorLocks) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *InlineResponse20048DoorLocks) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20048DoorLocks) GetNetwork() InlineResponse20048Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20048DoorLocks) GetNetworkOk() (*InlineResponse20048Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20048DoorLocks) SetNetwork(v InlineResponse20048Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20048DoorLocks) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse20048DoorLocks) GetGateway() InlineResponse20048Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse20048DoorLocks) GetGatewayOk() (*InlineResponse20048Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse20048DoorLocks) SetGateway(v InlineResponse20048Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse20048DoorLocks) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


