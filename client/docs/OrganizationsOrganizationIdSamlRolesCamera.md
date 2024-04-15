# OrganizationsOrganizationIdSamlRolesCamera

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgWide** | Pointer to **bool** | Whether or not SAML administrator has org-wide access | [optional] 
**Access** | Pointer to **string** | Camera access ability | [optional] 

## Methods

### NewOrganizationsOrganizationIdSamlRolesCamera

`func NewOrganizationsOrganizationIdSamlRolesCamera() *OrganizationsOrganizationIdSamlRolesCamera`

NewOrganizationsOrganizationIdSamlRolesCamera instantiates a new OrganizationsOrganizationIdSamlRolesCamera object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdSamlRolesCameraWithDefaults

`func NewOrganizationsOrganizationIdSamlRolesCameraWithDefaults() *OrganizationsOrganizationIdSamlRolesCamera`

NewOrganizationsOrganizationIdSamlRolesCameraWithDefaults instantiates a new OrganizationsOrganizationIdSamlRolesCamera object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgWide

`func (o *OrganizationsOrganizationIdSamlRolesCamera) GetOrgWide() bool`

GetOrgWide returns the OrgWide field if non-nil, zero value otherwise.

### GetOrgWideOk

`func (o *OrganizationsOrganizationIdSamlRolesCamera) GetOrgWideOk() (*bool, bool)`

GetOrgWideOk returns a tuple with the OrgWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgWide

`func (o *OrganizationsOrganizationIdSamlRolesCamera) SetOrgWide(v bool)`

SetOrgWide sets OrgWide field to given value.

### HasOrgWide

`func (o *OrganizationsOrganizationIdSamlRolesCamera) HasOrgWide() bool`

HasOrgWide returns a boolean if a field has been set.

### GetAccess

`func (o *OrganizationsOrganizationIdSamlRolesCamera) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *OrganizationsOrganizationIdSamlRolesCamera) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *OrganizationsOrganizationIdSamlRolesCamera) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *OrganizationsOrganizationIdSamlRolesCamera) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


