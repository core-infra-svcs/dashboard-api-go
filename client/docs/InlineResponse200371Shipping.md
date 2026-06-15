# InlineResponse200371Shipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipments** | Pointer to [**[]InlineResponse200371ShippingShipments**](InlineResponse200371ShippingShipments.md) | Hardware shipments for this order | [optional] 
**Pending** | Pointer to [**InlineResponse200371ShippingPending**](InlineResponse200371ShippingPending.md) |  | [optional] 

## Methods

### NewInlineResponse200371Shipping

`func NewInlineResponse200371Shipping() *InlineResponse200371Shipping`

NewInlineResponse200371Shipping instantiates a new InlineResponse200371Shipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200371ShippingWithDefaults

`func NewInlineResponse200371ShippingWithDefaults() *InlineResponse200371Shipping`

NewInlineResponse200371ShippingWithDefaults instantiates a new InlineResponse200371Shipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipments

`func (o *InlineResponse200371Shipping) GetShipments() []InlineResponse200371ShippingShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *InlineResponse200371Shipping) GetShipmentsOk() (*[]InlineResponse200371ShippingShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *InlineResponse200371Shipping) SetShipments(v []InlineResponse200371ShippingShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *InlineResponse200371Shipping) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPending

`func (o *InlineResponse200371Shipping) GetPending() InlineResponse200371ShippingPending`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *InlineResponse200371Shipping) GetPendingOk() (*InlineResponse200371ShippingPending, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *InlineResponse200371Shipping) SetPending(v InlineResponse200371ShippingPending)`

SetPending sets Pending field to given value.

### HasPending

`func (o *InlineResponse200371Shipping) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


