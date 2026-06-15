# InlineResponse200371

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The secure unique order claim number | [optional] 
**Resolution** | Pointer to [**InlineResponse200371Resolution**](InlineResponse200371Resolution.md) |  | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Shipping** | Pointer to [**InlineResponse200371Shipping**](InlineResponse200371Shipping.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200370Subscriptions**](InlineResponse200370Subscriptions.md) | All subscriptions contained in this order | [optional] 

## Methods

### NewInlineResponse200371

`func NewInlineResponse200371() *InlineResponse200371`

NewInlineResponse200371 instantiates a new InlineResponse200371 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200371WithDefaults

`func NewInlineResponse200371WithDefaults() *InlineResponse200371`

NewInlineResponse200371WithDefaults instantiates a new InlineResponse200371 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200371) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200371) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200371) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200371) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetResolution

`func (o *InlineResponse200371) GetResolution() InlineResponse200371Resolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *InlineResponse200371) GetResolutionOk() (*InlineResponse200371Resolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *InlineResponse200371) SetResolution(v InlineResponse200371Resolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *InlineResponse200371) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200371) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200371) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200371) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200371) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShipping

`func (o *InlineResponse200371) GetShipping() InlineResponse200371Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *InlineResponse200371) GetShippingOk() (*InlineResponse200371Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *InlineResponse200371) SetShipping(v InlineResponse200371Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *InlineResponse200371) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200371) GetSubscriptions() []InlineResponse200370Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200371) GetSubscriptionsOk() (*[]InlineResponse200370Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200371) SetSubscriptions(v []InlineResponse200370Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200371) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


