# InlineResponse200404Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200404ShippingShipments**](InlineResponse200404ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200404ShippingPending**](InlineResponse200404ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200404Shipping

`func NewInlineResponse200404Shipping() *InlineResponse200404Shipping`

NewInlineResponse200404Shipping instantiates a new InlineResponse200404Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200404ShippingWithDefaults

`func NewInlineResponse200404ShippingWithDefaults() *InlineResponse200404Shipping`

NewInlineResponse200404ShippingWithDefaults instantiates a new InlineResponse200404Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200404Shipping) GetShipments() []InlineResponse200404ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200404Shipping) GetShipmentsOk() (*[]InlineResponse200404ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200404Shipping) SetShipments(v []InlineResponse200404ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200404Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200404Shipping) GetPending() InlineResponse200404ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200404Shipping) GetPendingOk() (*InlineResponse200404ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200404Shipping) SetPending(v InlineResponse200404ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200404Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


