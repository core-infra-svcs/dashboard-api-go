# InlineResponse200359ShippingShipments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippedAt** | Pointer to **time.Time** | The date this shipment was shipped | [optional] 
**Number** | Pointer to **int32** | Which shipment the information is for | [optional] 
**Devices** | Pointer to [**[]InlineResponse200359ShippingDevices**](InlineResponse200359ShippingDevices.md) | All devices contained in this shipment | [optional] 

## Methods

### NewInlineResponse200359ShippingShipments

`func NewInlineResponse200359ShippingShipments() *InlineResponse200359ShippingShipments`

NewInlineResponse200359ShippingShipments instantiates a new InlineResponse200359ShippingShipments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200359ShippingShipmentsWithDefaults

`func NewInlineResponse200359ShippingShipmentsWithDefaults() *InlineResponse200359ShippingShipments`

NewInlineResponse200359ShippingShipmentsWithDefaults instantiates a new InlineResponse200359ShippingShipments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippedAt

`func (o *InlineResponse200359ShippingShipments) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *InlineResponse200359ShippingShipments) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *InlineResponse200359ShippingShipments) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *InlineResponse200359ShippingShipments) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200359ShippingShipments) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200359ShippingShipments) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200359ShippingShipments) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200359ShippingShipments) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200359ShippingShipments) GetDevices() []InlineResponse200359ShippingDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200359ShippingShipments) GetDevicesOk() (*[]InlineResponse200359ShippingDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200359ShippingShipments) SetDevices(v []InlineResponse200359ShippingDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200359ShippingShipments) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


