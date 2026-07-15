# InlineResponse200374

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**Resolution** | Pointer to [**InlineResponse200374Resolution**](InlineResponse200374Resolution.md) |  | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200374Shipping**](InlineResponse200374Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200373Subscriptions**](InlineResponse200373Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200374

`func NewInlineResponse200374() *InlineResponse200374`

NewInlineResponse200374 instantiates a new InlineResponse200374 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200374WithDefaults

`func NewInlineResponse200374WithDefaults() *InlineResponse200374`

NewInlineResponse200374WithDefaults instantiates a new InlineResponse200374 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200374) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200374) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200374) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200374) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetResolution

`func (o *InlineResponse200374) GetResolution() InlineResponse200374Resolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InlineResponse200374) GetResolutionOk() (*InlineResponse200374Resolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InlineResponse200374) SetResolution(v InlineResponse200374Resolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InlineResponse200374) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200374) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200374) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200374) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200374) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200374) GetShipping() InlineResponse200374Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200374) GetShippingOk() (*InlineResponse200374Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200374) SetShipping(v InlineResponse200374Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200374) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200374) GetSubscriptions() []InlineResponse200373Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200374) GetSubscriptionsOk() (*[]InlineResponse200373Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200374) SetSubscriptions(v []InlineResponse200373Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200374) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


