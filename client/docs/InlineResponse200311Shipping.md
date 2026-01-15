# InlineResponse200311Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200311ShippingShipments**](InlineResponse200311ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200311ShippingPending**](InlineResponse200311ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200311Shipping

`func NewInlineResponse200311Shipping() *InlineResponse200311Shipping`

NewInlineResponse200311Shipping instantiates a new InlineResponse200311Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311ShippingWithDefaults

`func NewInlineResponse200311ShippingWithDefaults() *InlineResponse200311Shipping`

NewInlineResponse200311ShippingWithDefaults instantiates a new InlineResponse200311Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200311Shipping) GetShipments() []InlineResponse200311ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200311Shipping) GetShipmentsOk() (*[]InlineResponse200311ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200311Shipping) SetShipments(v []InlineResponse200311ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200311Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200311Shipping) GetPending() InlineResponse200311ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200311Shipping) GetPendingOk() (*InlineResponse200311ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200311Shipping) SetPending(v InlineResponse200311ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200311Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


