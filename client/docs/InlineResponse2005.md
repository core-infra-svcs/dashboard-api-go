# InlineResponse2005

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription ID | [optional] 
**Networks** | Pointer to [**[]InlineResponse2005Networks**](InlineResponse2005Networks.md) | Unbound networks | [optional] 
**Errors** | Pointer to **[]string** | Array of errors if failed | [optional] 
**InsufficientEntitlements** | Pointer to [**[]AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements**](AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements.md) | A list of entitlements required to successfully bind the networks to the subscription | [optional] 

## Methods

### NewInlineResponse2005

`func NewInlineResponse2005() *InlineResponse2005`

NewInlineResponse2005 instantiates a new InlineResponse2005 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2005WithDefaults

`func NewInlineResponse2005WithDefaults() *InlineResponse2005`

NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *InlineResponse2005) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InlineResponse2005) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InlineResponse2005) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InlineResponse2005) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineResponse2005) GetNetworks() []InlineResponse2005Networks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse2005) GetNetworksOk() (*[]InlineResponse2005Networks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse2005) SetNetworks(v []InlineResponse2005Networks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineResponse2005) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse2005) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse2005) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse2005) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse2005) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInsufficientEntitlements

`func (o *InlineResponse2005) GetInsufficientEntitlements() []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements`

GetInsufficientEntitlements returns the InsufficientEntitlements field if non-nil, zero value otherwise.

### GetInsufficientEntitlementsOk

`func (o *InlineResponse2005) GetInsufficientEntitlementsOk() (*[]AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements, bool)`

GetInsufficientEntitlementsOk returns a tuple with the InsufficientEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientEntitlements

`func (o *InlineResponse2005) SetInsufficientEntitlements(v []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements)`

SetInsufficientEntitlements sets InsufficientEntitlements field to given value.

### HasInsufficientEntitlements

`func (o *InlineResponse2005) HasInsufficientEntitlements() bool`

HasInsufficientEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


