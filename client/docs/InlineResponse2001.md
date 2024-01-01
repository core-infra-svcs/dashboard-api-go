# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription&#39;s ID | [optional] 
**Name** | Pointer to **string** | Subscription name | [optional] 
**Description** | Pointer to **string** | Subscription description | [optional] 
**Status** | Pointer to **string** | Subscription status | [optional] 
**StartDate** | Pointer to **time.Time** | Subscription start date | [optional] 
**EndDate** | Pointer to **time.Time** | Subscription expiration date | [optional] 
**WebOrderId** | Pointer to **string** | Web order id | [optional] 
**ProductTypes** | Pointer to **[]string** | Products the subscription has entitlements for | [optional] 
**Entitlements** | Pointer to [**[]AdministeredLicensingSubscriptionSubscriptionsEntitlements**](AdministeredLicensingSubscriptionSubscriptionsEntitlements.md) | Entitlement info | [optional] 
**Counts** | Pointer to [**AdministeredLicensingSubscriptionSubscriptionsCounts**](AdministeredLicensingSubscriptionSubscriptionsCounts.md) |  | [optional] 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001() *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *InlineResponse2001) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InlineResponse2001) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InlineResponse2001) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InlineResponse2001) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2001) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2001) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2001) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2001) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse2001) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse2001) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse2001) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse2001) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2001) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2001) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2001) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2001) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse2001) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse2001) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse2001) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse2001) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InlineResponse2001) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InlineResponse2001) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InlineResponse2001) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InlineResponse2001) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetWebOrderId

`func (o *InlineResponse2001) GetWebOrderId() string`

GetWebOrderId returns the WebOrderId field if non-nil, zero value otherwise.

### GetWebOrderIdOk

`func (o *InlineResponse2001) GetWebOrderIdOk() (*string, bool)`

GetWebOrderIdOk returns a tuple with the WebOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebOrderId

`func (o *InlineResponse2001) SetWebOrderId(v string)`

SetWebOrderId sets WebOrderId field to given value.

### HasWebOrderId

`func (o *InlineResponse2001) HasWebOrderId() bool`

HasWebOrderId returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse2001) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse2001) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse2001) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse2001) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetEntitlements

`func (o *InlineResponse2001) GetEntitlements() []AdministeredLicensingSubscriptionSubscriptionsEntitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *InlineResponse2001) GetEntitlementsOk() (*[]AdministeredLicensingSubscriptionSubscriptionsEntitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *InlineResponse2001) SetEntitlements(v []AdministeredLicensingSubscriptionSubscriptionsEntitlements)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *InlineResponse2001) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse2001) GetCounts() AdministeredLicensingSubscriptionSubscriptionsCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse2001) GetCountsOk() (*AdministeredLicensingSubscriptionSubscriptionsCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse2001) SetCounts(v AdministeredLicensingSubscriptionSubscriptionsCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse2001) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


