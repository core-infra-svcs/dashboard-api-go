# InlineResponse200358

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The order claim id | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Serials** | Pointer to **[]string** | All devices claimed in this order | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200358Subscriptions**](InlineResponse200358Subscriptions.md) | Details for subscriptions claimed in this order | [optional] 

## Methods

### NewInlineResponse200358

`func NewInlineResponse200358() *InlineResponse200358`

NewInlineResponse200358 instantiates a new InlineResponse200358 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200358WithDefaults

`func NewInlineResponse200358WithDefaults() *InlineResponse200358`

NewInlineResponse200358WithDefaults instantiates a new InlineResponse200358 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200358) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200358) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200358) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200358) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200358) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200358) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200358) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200358) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200358) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200358) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200358) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200358) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200358) GetSubscriptions() []InlineResponse200358Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200358) GetSubscriptionsOk() (*[]InlineResponse200358Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200358) SetSubscriptions(v []InlineResponse200358Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200358) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


