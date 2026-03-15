# InlineResponse200323Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200323ShippingShipments**](InlineResponse200323ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200323ShippingPending**](InlineResponse200323ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200323Shipping

`func NewInlineResponse200323Shipping() *InlineResponse200323Shipping`

NewInlineResponse200323Shipping instantiates a new InlineResponse200323Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200323ShippingWithDefaults

`func NewInlineResponse200323ShippingWithDefaults() *InlineResponse200323Shipping`

NewInlineResponse200323ShippingWithDefaults instantiates a new InlineResponse200323Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200323Shipping) GetShipments() []InlineResponse200323ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200323Shipping) GetShipmentsOk() (*[]InlineResponse200323ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200323Shipping) SetShipments(v []InlineResponse200323ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200323Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200323Shipping) GetPending() InlineResponse200323ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200323Shipping) GetPendingOk() (*InlineResponse200323ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200323Shipping) SetPending(v InlineResponse200323ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200323Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


