# InlineResponse200335Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200335ShippingShipments**](InlineResponse200335ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200335ShippingPending**](InlineResponse200335ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200335Shipping

`func NewInlineResponse200335Shipping() *InlineResponse200335Shipping`

NewInlineResponse200335Shipping instantiates a new InlineResponse200335Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200335ShippingWithDefaults

`func NewInlineResponse200335ShippingWithDefaults() *InlineResponse200335Shipping`

NewInlineResponse200335ShippingWithDefaults instantiates a new InlineResponse200335Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200335Shipping) GetShipments() []InlineResponse200335ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200335Shipping) GetShipmentsOk() (*[]InlineResponse200335ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200335Shipping) SetShipments(v []InlineResponse200335ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200335Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200335Shipping) GetPending() InlineResponse200335ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200335Shipping) GetPendingOk() (*InlineResponse200335ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200335Shipping) SetPending(v InlineResponse200335ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200335Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


