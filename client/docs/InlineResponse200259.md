# InlineResponse200259

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Organization ID | [optional] 
**Name** | Pointer to **string** | Organization name | [optional] 
**Url** | Pointer to **string** | Organization URL | [optional] 
**Api** | Pointer to [**OrganizationsApi**](OrganizationsApi.md) |  | [optional] 
**Licensing** | Pointer to [**OrganizationsLicensing**](OrganizationsLicensing.md) |  | [optional] 
**Cloud** | Pointer to [**OrganizationsCloud**](OrganizationsCloud.md) |  | [optional] 
**Management** | Pointer to [**OrganizationsManagement**](OrganizationsManagement.md) |  | [optional] 
**Privacy** | Pointer to **map[string]interface{}** | Privacy-related settings for the organization. | [optional] 

## Methods

### NewInlineResponse200259

`func NewInlineResponse200259() *InlineResponse200259`

NewInlineResponse200259 instantiates a new InlineResponse200259 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200259WithDefaults

`func NewInlineResponse200259WithDefaults() *InlineResponse200259`

NewInlineResponse200259WithDefaults instantiates a new InlineResponse200259 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200259) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200259) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200259) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200259) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200259) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200259) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200259) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200259) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200259) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200259) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200259) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200259) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApi

`func (o *InlineResponse200259) GetApi() OrganizationsApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *InlineResponse200259) GetApiOk() (*OrganizationsApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *InlineResponse200259) SetApi(v OrganizationsApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *InlineResponse200259) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetLicensing

`func (o *InlineResponse200259) GetLicensing() OrganizationsLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *InlineResponse200259) GetLicensingOk() (*OrganizationsLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *InlineResponse200259) SetLicensing(v OrganizationsLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *InlineResponse200259) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetCloud

`func (o *InlineResponse200259) GetCloud() OrganizationsCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InlineResponse200259) GetCloudOk() (*OrganizationsCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InlineResponse200259) SetCloud(v OrganizationsCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InlineResponse200259) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetManagement

`func (o *InlineResponse200259) GetManagement() OrganizationsManagement`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *InlineResponse200259) GetManagementOk() (*OrganizationsManagement, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *InlineResponse200259) SetManagement(v OrganizationsManagement)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *InlineResponse200259) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetPrivacy

`func (o *InlineResponse200259) GetPrivacy() map[string]interface{}`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *InlineResponse200259) GetPrivacyOk() (*map[string]interface{}, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *InlineResponse200259) SetPrivacy(v map[string]interface{})`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *InlineResponse200259) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


