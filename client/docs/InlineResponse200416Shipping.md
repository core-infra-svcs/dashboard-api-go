# InlineResponse200416Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200416ShippingShipments**](InlineResponse200416ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200416ShippingPending**](InlineResponse200416ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200416Shipping

`func NewInlineResponse200416Shipping() *InlineResponse200416Shipping`

NewInlineResponse200416Shipping instantiates a new InlineResponse200416Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200416ShippingWithDefaults

`func NewInlineResponse200416ShippingWithDefaults() *InlineResponse200416Shipping`

NewInlineResponse200416ShippingWithDefaults instantiates a new InlineResponse200416Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200416Shipping) GetShipments() []InlineResponse200416ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200416Shipping) GetShipmentsOk() (*[]InlineResponse200416ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200416Shipping) SetShipments(v []InlineResponse200416ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200416Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200416Shipping) GetPending() InlineResponse200416ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200416Shipping) GetPendingOk() (*InlineResponse200416ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200416Shipping) SetPending(v InlineResponse200416ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200416Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


