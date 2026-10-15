# InlineResponse200420

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the application category | [optional] 
**Name** | Pointer to **string** | Name of the application category | [optional] 
**Applications** | Pointer to [**[]OrganizationsOrganizationIdPoliciesGlobalFirewallApplicationCategoriesApplications**](OrganizationsOrganizationIdPoliciesGlobalFirewallApplicationCategoriesApplications.md) | List of applications in this category | [optional] 

## Methods

### NewInlineResponse200420

`func NewInlineResponse200420() *InlineResponse200420`

NewInlineResponse200420 instantiates a new InlineResponse200420 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200420WithDefaults

`func NewInlineResponse200420WithDefaults() *InlineResponse200420`

NewInlineResponse200420WithDefaults instantiates a new InlineResponse200420 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200420) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200420) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200420) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200420) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200420) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200420) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200420) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200420) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse200420) GetApplications() []OrganizationsOrganizationIdPoliciesGlobalFirewallApplicationCategoriesApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse200420) GetApplicationsOk() (*[]OrganizationsOrganizationIdPoliciesGlobalFirewallApplicationCategoriesApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse200420) SetApplications(v []OrganizationsOrganizationIdPoliciesGlobalFirewallApplicationCategoriesApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse200420) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


