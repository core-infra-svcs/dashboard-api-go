# InlineResponse200390

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**Resolution** | Pointer to [**InlineResponse200390Resolution**](InlineResponse200390Resolution.md) |  | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200390Shipping**](InlineResponse200390Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200389Subscriptions**](InlineResponse200389Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200390

`func NewInlineResponse200390() *InlineResponse200390`

NewInlineResponse200390 instantiates a new InlineResponse200390 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200390WithDefaults

`func NewInlineResponse200390WithDefaults() *InlineResponse200390`

NewInlineResponse200390WithDefaults instantiates a new InlineResponse200390 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200390) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200390) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200390) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200390) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetResolution

`func (o *InlineResponse200390) GetResolution() InlineResponse200390Resolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InlineResponse200390) GetResolutionOk() (*InlineResponse200390Resolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InlineResponse200390) SetResolution(v InlineResponse200390Resolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InlineResponse200390) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200390) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200390) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200390) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200390) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200390) GetShipping() InlineResponse200390Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200390) GetShippingOk() (*InlineResponse200390Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200390) SetShipping(v InlineResponse200390Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200390) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200390) GetSubscriptions() []InlineResponse200389Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200390) GetSubscriptionsOk() (*[]InlineResponse200389Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200390) SetSubscriptions(v []InlineResponse200389Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200390) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


