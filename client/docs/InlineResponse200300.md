# InlineResponse200300

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert type | 
**Title** | **string** | Alert type title | 
**CategoryType** | **string** | Alert category type | 
**Severities** | [**[]OrganizationsOrganizationIdAssuranceAlertsTaxonomyTypesSeverities**](OrganizationsOrganizationIdAssuranceAlertsTaxonomyTypesSeverities.md) | List of possible severities for the alert type | 
**DeviceTypes** | **[]string** | List of possible device types for the alert type | 

## Methods

### NewInlineResponse200300

`func NewInlineResponse200300(type_ string, title string, categoryType string, severities []OrganizationsOrganizationIdAssuranceAlertsTaxonomyTypesSeverities, deviceTypes []string, ) *InlineResponse200300`

NewInlineResponse200300 instantiates a new InlineResponse200300 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200300WithDefaults

`func NewInlineResponse200300WithDefaults() *InlineResponse200300`

NewInlineResponse200300WithDefaults instantiates a new InlineResponse200300 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200300) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200300) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200300) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *InlineResponse200300) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse200300) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse200300) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCategoryType

`func (o *InlineResponse200300) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *InlineResponse200300) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *InlineResponse200300) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.


### GetSeverities

`func (o *InlineResponse200300) GetSeverities() []OrganizationsOrganizationIdAssuranceAlertsTaxonomyTypesSeverities`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *InlineResponse200300) GetSeveritiesOk() (*[]OrganizationsOrganizationIdAssuranceAlertsTaxonomyTypesSeverities, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *InlineResponse200300) SetSeverities(v []OrganizationsOrganizationIdAssuranceAlertsTaxonomyTypesSeverities)`

SetSeverities sets Severities field to given value.


### GetDeviceTypes

`func (o *InlineResponse200300) GetDeviceTypes() []string`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *InlineResponse200300) GetDeviceTypesOk() (*[]string, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *InlineResponse200300) SetDeviceTypes(v []string)`

SetDeviceTypes sets DeviceTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


