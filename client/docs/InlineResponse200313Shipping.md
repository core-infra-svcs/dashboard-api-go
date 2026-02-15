# InlineResponse200313Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200313ShippingShipments**](InlineResponse200313ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200313ShippingPending**](InlineResponse200313ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200313Shipping

`func NewInlineResponse200313Shipping() *InlineResponse200313Shipping`

NewInlineResponse200313Shipping instantiates a new InlineResponse200313Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200313ShippingWithDefaults

`func NewInlineResponse200313ShippingWithDefaults() *InlineResponse200313Shipping`

NewInlineResponse200313ShippingWithDefaults instantiates a new InlineResponse200313Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200313Shipping) GetShipments() []InlineResponse200313ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200313Shipping) GetShipmentsOk() (*[]InlineResponse200313ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200313Shipping) SetShipments(v []InlineResponse200313ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200313Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200313Shipping) GetPending() InlineResponse200313ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200313Shipping) GetPendingOk() (*InlineResponse200313ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200313Shipping) SetPending(v InlineResponse200313ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200313Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


