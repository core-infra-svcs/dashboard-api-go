# InlineResponse200385

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | Pointer to **string** | The order claim id | [optional] 
**Number** | Pointer to **string** | Order Number | [optional] 
**Serials** | Pointer to **[]string** | All devices claimed in this order | [optional] 
**Subscriptions** | Pointer to [**[]InlineResponse200385Subscriptions**](InlineResponse200385Subscriptions.md) | Details for subscriptions claimed in this order | [optional] 

## Methods

### NewInlineResponse200385

`func NewInlineResponse200385() *InlineResponse200385`

NewInlineResponse200385 instantiates a new InlineResponse200385 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200385WithDefaults

`func NewInlineResponse200385WithDefaults() *InlineResponse200385`

NewInlineResponse200385WithDefaults instantiates a new InlineResponse200385 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineResponse200385) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineResponse200385) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineResponse200385) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.

### HasClaimId

`func (o *InlineResponse200385) HasClaimId() bool`

HasClaimId returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200385) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200385) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200385) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200385) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200385) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200385) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200385) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200385) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSubscriptions

`func (o *InlineResponse200385) GetSubscriptions() []InlineResponse200385Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineResponse200385) GetSubscriptionsOk() (*[]InlineResponse200385Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineResponse200385) SetSubscriptions(v []InlineResponse200385Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineResponse200385) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


