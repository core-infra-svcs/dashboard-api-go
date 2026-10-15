# InlineResponse200403Subscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription ID | [optional] 
**Name** | Pointer to **string** | Subscription name | [optional] 
**Description** | Pointer to **string** | Subscription description | [optional] 
**StartDate** | Pointer to **time.Time** | The date this subscription will begin | [optional] 
**EndDate** | Pointer to **time.Time** | The date this subscription will end | [optional] 
**IsClaimed** | Pointer to **bool** | Whether the subscription has been claimed | [optional] 
**Counts** | Pointer to [**InlineResponse200403Counts**](InlineResponse200403Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200403Subscriptions

`func NewInlineResponse200403Subscriptions() *InlineResponse200403Subscriptions`

NewInlineResponse200403Subscriptions instantiates a new InlineResponse200403Subscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200403SubscriptionsWithDefaults

`func NewInlineResponse200403SubscriptionsWithDefaults() *InlineResponse200403Subscriptions`

NewInlineResponse200403SubscriptionsWithDefaults instantiates a new InlineResponse200403Subscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *InlineResponse200403Subscriptions) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InlineResponse200403Subscriptions) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InlineResponse200403Subscriptions) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InlineResponse200403Subscriptions) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200403Subscriptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200403Subscriptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200403Subscriptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200403Subscriptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200403Subscriptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200403Subscriptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200403Subscriptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200403Subscriptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse200403Subscriptions) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse200403Subscriptions) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse200403Subscriptions) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse200403Subscriptions) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InlineResponse200403Subscriptions) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InlineResponse200403Subscriptions) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InlineResponse200403Subscriptions) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InlineResponse200403Subscriptions) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsClaimed

`func (o *InlineResponse200403Subscriptions) GetIsClaimed() bool`

GetIsClaimed returns the IsClaimed field if non-nil, zero value otherwise.

### GetIsClaimedOk

`func (o *InlineResponse200403Subscriptions) GetIsClaimedOk() (*bool, bool)`

GetIsClaimedOk returns a tuple with the IsClaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClaimed

`func (o *InlineResponse200403Subscriptions) SetIsClaimed(v bool)`

SetIsClaimed sets IsClaimed field to given value.

### HasIsClaimed

`func (o *InlineResponse200403Subscriptions) HasIsClaimed() bool`

HasIsClaimed returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200403Subscriptions) GetCounts() InlineResponse200403Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200403Subscriptions) GetCountsOk() (*InlineResponse200403Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200403Subscriptions) SetCounts(v InlineResponse200403Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200403Subscriptions) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


