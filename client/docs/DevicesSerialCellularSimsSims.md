# DevicesSerialCellularSimsSims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. Use &#39;sim3&#39; for eSIM. | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using &#39;simOrdering&#39;. | [optional] [default to false]
**Apns** | Pointer to [**[]InlineResponse20015Apns**](InlineResponse20015Apns.md) | APN configurations. If empty, the default APN will be used. | [optional] 
**SimOrder** | Pointer to **int32** | Priority of SIM slot being configured. Use a value between 1 and total number of SIMs available. The value must be unique for each SIM. | [optional] 

## Methods

### NewDevicesSerialCellularSimsSims

`func NewDevicesSerialCellularSimsSims() *DevicesSerialCellularSimsSims`

NewDevicesSerialCellularSimsSims instantiates a new DevicesSerialCellularSimsSims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialCellularSimsSimsWithDefaults

`func NewDevicesSerialCellularSimsSimsWithDefaults() *DevicesSerialCellularSimsSims`

NewDevicesSerialCellularSimsSimsWithDefaults instantiates a new DevicesSerialCellularSimsSims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *DevicesSerialCellularSimsSims) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *DevicesSerialCellularSimsSims) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *DevicesSerialCellularSimsSims) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *DevicesSerialCellularSimsSims) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsPrimary

`func (o *DevicesSerialCellularSimsSims) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *DevicesSerialCellularSimsSims) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *DevicesSerialCellularSimsSims) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *DevicesSerialCellularSimsSims) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetApns

`func (o *DevicesSerialCellularSimsSims) GetApns() []InlineResponse20015Apns`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *DevicesSerialCellularSimsSims) GetApnsOk() (*[]InlineResponse20015Apns, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *DevicesSerialCellularSimsSims) SetApns(v []InlineResponse20015Apns)`

SetApns sets Apns field to given value.

### HasApns

`func (o *DevicesSerialCellularSimsSims) HasApns() bool`

HasApns returns a boolean if a field has been set.

### GetSimOrder

`func (o *DevicesSerialCellularSimsSims) GetSimOrder() int32`

GetSimOrder returns the SimOrder field if non-nil, zero value otherwise.

### GetSimOrderOk

`func (o *DevicesSerialCellularSimsSims) GetSimOrderOk() (*int32, bool)`

GetSimOrderOk returns a tuple with the SimOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrder

`func (o *DevicesSerialCellularSimsSims) SetSimOrder(v int32)`

SetSimOrder sets SimOrder field to given value.

### HasSimOrder

`func (o *DevicesSerialCellularSimsSims) HasSimOrder() bool`

HasSimOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


