# InlineResponse200416

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**TrialId** | Pointer to **NullableString** | Trial identifier associated with the Unified Order | [optional] 
**Resolution** | Pointer to [**InlineResponse200416Resolution**](InlineResponse200416Resolution.md) |  | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200416Shipping**](InlineResponse200416Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200415Subscriptions**](InlineResponse200415Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200416

`func NewInlineResponse200416() *InlineResponse200416`

NewInlineResponse200416 instantiates a new InlineResponse200416 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200416WithDefaults

`func NewInlineResponse200416WithDefaults() *InlineResponse200416`

NewInlineResponse200416WithDefaults instantiates a new InlineResponse200416 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200416) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200416) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200416) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200416) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetTrialId

`func (o *InlineResponse200416) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *InlineResponse200416) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *InlineResponse200416) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *InlineResponse200416) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *InlineResponse200416) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *InlineResponse200416) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetResolution

`func (o *InlineResponse200416) GetResolution() InlineResponse200416Resolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InlineResponse200416) GetResolutionOk() (*InlineResponse200416Resolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InlineResponse200416) SetResolution(v InlineResponse200416Resolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InlineResponse200416) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200416) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200416) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200416) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200416) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200416) GetShipping() InlineResponse200416Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200416) GetShippingOk() (*InlineResponse200416Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200416) SetShipping(v InlineResponse200416Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200416) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200416) GetSubscriptions() []InlineResponse200415Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200416) GetSubscriptionsOk() (*[]InlineResponse200415Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200416) SetSubscriptions(v []InlineResponse200415Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200416) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


