# InlineResponse200311

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200311Shipping**](InlineResponse200311Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200310Subscriptions**](InlineResponse200310Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200311

`func NewInlineResponse200311() *InlineResponse200311`

NewInlineResponse200311 instantiates a new InlineResponse200311 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311WithDefaults

`func NewInlineResponse200311WithDefaults() *InlineResponse200311`

NewInlineResponse200311WithDefaults instantiates a new InlineResponse200311 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200311) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200311) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200311) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200311) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200311) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200311) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200311) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200311) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200311) GetShipping() InlineResponse200311Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200311) GetShippingOk() (*InlineResponse200311Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200311) SetShipping(v InlineResponse200311Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200311) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200311) GetSubscriptions() []InlineResponse200310Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200311) GetSubscriptionsOk() (*[]InlineResponse200310Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200311) SetSubscriptions(v []InlineResponse200310Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200311) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


