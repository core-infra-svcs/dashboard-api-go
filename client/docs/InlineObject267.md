# InlineObject267

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Controls how the configuration payload in the request body is applied to the resource. This parameter dictates the declarative mode:  * **&#x60;complete&#x60;**: The request body represents the entire desired configuration for this resource. Any existing configurations that are not included in the payload will be removed. This effectively performs a full replacement or overwrite of the resource&#39;s configuration. * **&#x60;partial&#x60; (default)**: The request body contains only the configurations to be created or modified. Existing configurations that are not specified in the payload will be preserved. This performs a merge or partial update, applying only the changes specified.  | [optional] 
**Items** | [**[]OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems**](OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems.md) | Desired overrides state | 

## Methods

### NewInlineObject267

`func NewInlineObject267(items []OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems, ) *InlineObject267`

NewInlineObject267 instantiates a new InlineObject267 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject267WithDefaults

`func NewInlineObject267WithDefaults() *InlineObject267`

NewInlineObject267WithDefaults instantiates a new InlineObject267 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject267) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject267) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject267) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject267) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetItems

`func (o *InlineObject267) GetItems() []OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineObject267) GetItemsOk() (*[]OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineObject267) SetItems(v []OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


