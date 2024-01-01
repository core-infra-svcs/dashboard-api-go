# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription ID | [optional] 
**Networks** | Pointer to [**[]InlineResponse2002Networks**](InlineResponse2002Networks.md) | Unbound networks | [optional] 
**Errors** | Pointer to **[]string** | Array of errors if failed | [optional] 
**InsufficientEntitlements** | Pointer to [**[]InlineResponse2002InsufficientEntitlements**](InlineResponse2002InsufficientEntitlements.md) | A list of entitlements required to successfully bind the networks to the subscription | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *InlineResponse2002) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InlineResponse2002) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InlineResponse2002) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InlineResponse2002) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineResponse2002) GetNetworks() []InlineResponse2002Networks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse2002) GetNetworksOk() (*[]InlineResponse2002Networks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse2002) SetNetworks(v []InlineResponse2002Networks)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineResponse2002) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse2002) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse2002) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse2002) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse2002) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInsufficientEntitlements

`func (o *InlineResponse2002) GetInsufficientEntitlements() []InlineResponse2002InsufficientEntitlements`

GetInsufficientEntitlements returns the InsufficientEntitlements field if non-nil, zero value otherwise.

### GetInsufficientEntitlementsOk

`func (o *InlineResponse2002) GetInsufficientEntitlementsOk() (*[]InlineResponse2002InsufficientEntitlements, bool)`

GetInsufficientEntitlementsOk returns a tuple with the InsufficientEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientEntitlements

`func (o *InlineResponse2002) SetInsufficientEntitlements(v []InlineResponse2002InsufficientEntitlements)`

SetInsufficientEntitlements sets InsufficientEntitlements field to given value.

### HasInsufficientEntitlements

`func (o *InlineResponse2002) HasInsufficientEntitlements() bool`

HasInsufficientEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


