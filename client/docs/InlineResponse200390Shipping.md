# InlineResponse200390Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200390ShippingShipments**](InlineResponse200390ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200390ShippingPending**](InlineResponse200390ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200390Shipping

`func NewInlineResponse200390Shipping() *InlineResponse200390Shipping`

NewInlineResponse200390Shipping instantiates a new InlineResponse200390Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200390ShippingWithDefaults

`func NewInlineResponse200390ShippingWithDefaults() *InlineResponse200390Shipping`

NewInlineResponse200390ShippingWithDefaults instantiates a new InlineResponse200390Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200390Shipping) GetShipments() []InlineResponse200390ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200390Shipping) GetShipmentsOk() (*[]InlineResponse200390ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200390Shipping) SetShipments(v []InlineResponse200390ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200390Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200390Shipping) GetPending() InlineResponse200390ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200390Shipping) GetPendingOk() (*InlineResponse200390ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200390Shipping) SetPending(v InlineResponse200390ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200390Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


