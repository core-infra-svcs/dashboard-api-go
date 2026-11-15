# InlineObject325

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | **string** | The unique order claim id | 
**Subscriptions** | Pointer to [**[]OrganizationsOrganizationIdInventoryOrdersClaimSubscriptions**](OrganizationsOrganizationIdInventoryOrdersClaimSubscriptions.md) | The individual subscriptions to claim | [optional] 

## Methods

### NewInlineObject325

`func NewInlineObject325(claimId string, ) *InlineObject325`

NewInlineObject325 instantiates a new InlineObject325 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject325WithDefaults

`func NewInlineObject325WithDefaults() *InlineObject325`

NewInlineObject325WithDefaults instantiates a new InlineObject325 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *InlineObject325) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *InlineObject325) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *InlineObject325) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.


### GetSubscriptions

`func (o *InlineObject325) GetSubscriptions() []OrganizationsOrganizationIdInventoryOrdersClaimSubscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *InlineObject325) GetSubscriptionsOk() (*[]OrganizationsOrganizationIdInventoryOrdersClaimSubscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *InlineObject325) SetSubscriptions(v []OrganizationsOrganizationIdInventoryOrdersClaimSubscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *InlineObject325) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


