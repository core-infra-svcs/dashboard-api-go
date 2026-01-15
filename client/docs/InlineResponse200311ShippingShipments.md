# InlineResponse200311ShippingShipments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippedAt** | Pointer to **time.Time** | The date this shipment was shipped | [optional] 
**Number** | Pointer to **int32** | Which shipment the information is for | [optional] 
**Devices** | Pointer to [**[]InlineResponse200311ShippingDevices**](InlineResponse200311ShippingDevices.md) | All devices contained in this shipment | [optional] 

## Methods

### NewInlineResponse200311ShippingShipments

`func NewInlineResponse200311ShippingShipments() *InlineResponse200311ShippingShipments`

NewInlineResponse200311ShippingShipments instantiates a new InlineResponse200311ShippingShipments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311ShippingShipmentsWithDefaults

`func NewInlineResponse200311ShippingShipmentsWithDefaults() *InlineResponse200311ShippingShipments`

NewInlineResponse200311ShippingShipmentsWithDefaults instantiates a new InlineResponse200311ShippingShipments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippedAt

`func (o *InlineResponse200311ShippingShipments) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *InlineResponse200311ShippingShipments) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *InlineResponse200311ShippingShipments) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *InlineResponse200311ShippingShipments) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200311ShippingShipments) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200311ShippingShipments) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200311ShippingShipments) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200311ShippingShipments) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200311ShippingShipments) GetDevices() []InlineResponse200311ShippingDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200311ShippingShipments) GetDevicesOk() (*[]InlineResponse200311ShippingDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200311ShippingShipments) SetDevices(v []InlineResponse200311ShippingDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200311ShippingShipments) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


