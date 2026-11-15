# InlineObject268

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Controls how the configuration payload in the request body is applied to the resource. This parameter dictates the declarative mode:  * **&#x60;complete&#x60;**: The request body represents the entire desired configuration for this resource. Any existing configurations that are not included in the payload will be removed. * **&#x60;partial&#x60; (default)**: The request body contains only the configurations to be created or modified. Existing configurations that are not specified in the payload will be preserved.  | [optional] 
**Recursive** | Pointer to **bool** | Controls how the configuration payload in the request body applies to the rule group hierarchy. When true, the API applies each declared override to the rule group itself and its descendants unless the payload explicitly sets a descendant override. When false (default), the API applies overrides only to the rule groups listed in the payload.  | [optional] 
**Items** | [**[]OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems**](OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems.md) | Desired overrides state | 

## Methods

### NewInlineObject268

`func NewInlineObject268(items []OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems, ) *InlineObject268`

NewInlineObject268 instantiates a new InlineObject268 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject268WithDefaults

`func NewInlineObject268WithDefaults() *InlineObject268`

NewInlineObject268WithDefaults instantiates a new InlineObject268 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject268) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject268) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject268) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject268) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRecursive

`func (o *InlineObject268) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *InlineObject268) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *InlineObject268) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *InlineObject268) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetItems

`func (o *InlineObject268) GetItems() []OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineObject268) GetItemsOk() (*[]OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineObject268) SetItems(v []OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


