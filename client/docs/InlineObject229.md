# InlineObject229

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the organization | [optional] 
**Management** | Pointer to [**OrganizationsManagement1**](OrganizationsManagement1.md) |  | [optional] 
**Api** | Pointer to [**OrganizationsOrganizationIdApi**](OrganizationsOrganizationIdApi.md) |  | [optional] 
**Privacy** | Pointer to **map[string]interface{}** | Privacy-related settings for the organization. | [optional] 

## Methods

### NewInlineObject229

`func NewInlineObject229() *InlineObject229`

NewInlineObject229 instantiates a new InlineObject229 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject229WithDefaults

`func NewInlineObject229WithDefaults() *InlineObject229`

NewInlineObject229WithDefaults instantiates a new InlineObject229 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject229) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject229) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject229) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject229) HasName() bool`

HasName returns a boolean if a field has been set.

### GetManagement

`func (o *InlineObject229) GetManagement() OrganizationsManagement1`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *InlineObject229) GetManagementOk() (*OrganizationsManagement1, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *InlineObject229) SetManagement(v OrganizationsManagement1)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *InlineObject229) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetApi

`func (o *InlineObject229) GetApi() OrganizationsOrganizationIdApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *InlineObject229) GetApiOk() (*OrganizationsOrganizationIdApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *InlineObject229) SetApi(v OrganizationsOrganizationIdApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *InlineObject229) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineObject229) GetPrivacy() map[string]interface{}`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineObject229) GetPrivacyOk() (*map[string]interface{}, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineObject229) SetPrivacy(v map[string]interface{})`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineObject229) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


