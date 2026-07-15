# InlineResponse200374Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200374ShippingShipments**](InlineResponse200374ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200374ShippingPending**](InlineResponse200374ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200374Shipping

`func NewInlineResponse200374Shipping() *InlineResponse200374Shipping`

NewInlineResponse200374Shipping instantiates a new InlineResponse200374Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200374ShippingWithDefaults

`func NewInlineResponse200374ShippingWithDefaults() *InlineResponse200374Shipping`

NewInlineResponse200374ShippingWithDefaults instantiates a new InlineResponse200374Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200374Shipping) GetShipments() []InlineResponse200374ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200374Shipping) GetShipmentsOk() (*[]InlineResponse200374ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200374Shipping) SetShipments(v []InlineResponse200374ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200374Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200374Shipping) GetPending() InlineResponse200374ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200374Shipping) GetPendingOk() (*InlineResponse200374ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200374Shipping) SetPending(v InlineResponse200374ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200374Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


