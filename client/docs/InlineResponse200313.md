# InlineResponse200313

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200313Shipping**](InlineResponse200313Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200312Subscriptions**](InlineResponse200312Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200313

`func NewInlineResponse200313() *InlineResponse200313`

NewInlineResponse200313 instantiates a new InlineResponse200313 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200313WithDefaults

`func NewInlineResponse200313WithDefaults() *InlineResponse200313`

NewInlineResponse200313WithDefaults instantiates a new InlineResponse200313 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200313) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200313) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200313) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200313) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200313) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200313) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200313) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200313) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200313) GetShipping() InlineResponse200313Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200313) GetShippingOk() (*InlineResponse200313Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200313) SetShipping(v InlineResponse200313Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200313) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200313) GetSubscriptions() []InlineResponse200312Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200313) GetSubscriptionsOk() (*[]InlineResponse200312Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200313) SetSubscriptions(v []InlineResponse200312Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200313) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


