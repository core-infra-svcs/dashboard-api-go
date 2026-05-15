# InlineResponse200359Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200359ShippingShipments**](InlineResponse200359ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200359ShippingPending**](InlineResponse200359ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200359Shipping

`func NewInlineResponse200359Shipping() *InlineResponse200359Shipping`

NewInlineResponse200359Shipping instantiates a new InlineResponse200359Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200359ShippingWithDefaults

`func NewInlineResponse200359ShippingWithDefaults() *InlineResponse200359Shipping`

NewInlineResponse200359ShippingWithDefaults instantiates a new InlineResponse200359Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200359Shipping) GetShipments() []InlineResponse200359ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200359Shipping) GetShipmentsOk() (*[]InlineResponse200359ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200359Shipping) SetShipments(v []InlineResponse200359ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200359Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200359Shipping) GetPending() InlineResponse200359ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200359Shipping) GetPendingOk() (*InlineResponse200359ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200359Shipping) SetPending(v InlineResponse200359ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200359Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


