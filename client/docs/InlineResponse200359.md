# InlineResponse200359

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**Resolution** | Pointer to [**InlineResponse200359Resolution**](InlineResponse200359Resolution.md) |  | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200359Shipping**](InlineResponse200359Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200358Subscriptions**](InlineResponse200358Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200359

`func NewInlineResponse200359() *InlineResponse200359`

NewInlineResponse200359 instantiates a new InlineResponse200359 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200359WithDefaults

`func NewInlineResponse200359WithDefaults() *InlineResponse200359`

NewInlineResponse200359WithDefaults instantiates a new InlineResponse200359 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200359) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200359) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200359) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200359) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetResolution

`func (o *InlineResponse200359) GetResolution() InlineResponse200359Resolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InlineResponse200359) GetResolutionOk() (*InlineResponse200359Resolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InlineResponse200359) SetResolution(v InlineResponse200359Resolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InlineResponse200359) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200359) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200359) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200359) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200359) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200359) GetShipping() InlineResponse200359Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200359) GetShippingOk() (*InlineResponse200359Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200359) SetShipping(v InlineResponse200359Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200359) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200359) GetSubscriptions() []InlineResponse200358Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200359) GetSubscriptionsOk() (*[]InlineResponse200358Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200359) SetSubscriptions(v []InlineResponse200358Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200359) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


