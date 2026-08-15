# InlineResponse200386Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200386ShippingShipments**](InlineResponse200386ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200386ShippingPending**](InlineResponse200386ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200386Shipping

`func NewInlineResponse200386Shipping() *InlineResponse200386Shipping`

NewInlineResponse200386Shipping instantiates a new InlineResponse200386Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200386ShippingWithDefaults

`func NewInlineResponse200386ShippingWithDefaults() *InlineResponse200386Shipping`

NewInlineResponse200386ShippingWithDefaults instantiates a new InlineResponse200386Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200386Shipping) GetShipments() []InlineResponse200386ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200386Shipping) GetShipmentsOk() (*[]InlineResponse200386ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200386Shipping) SetShipments(v []InlineResponse200386ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200386Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200386Shipping) GetPending() InlineResponse200386ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200386Shipping) GetPendingOk() (*InlineResponse200386ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200386Shipping) SetPending(v InlineResponse200386ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200386Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


