# InlineResponse200242

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the health alert | 
**CategoryType** | **string** | Category type that the health alert belongs to | 
**Network** | [**OrganizationsOrganizationIdAssuranceAlertsNetwork**](OrganizationsOrganizationIdAssuranceAlertsNetwork.md) |  | 
**StartedAt** | **time.Time** | Time when the alert started | 
**ResolvedAt** | Pointer to **time.Time** | Time when the alert was resolved | [optional] 
**DismissedAt** | Pointer to **time.Time** | Time when the alert was dismissed | [optional] 
**DeviceType** | Pointer to **string** | Device Type that the alert occurred on | [optional] 
**Type** | **string** | Alert Type | 
**Title** | **string** | Human Readable Title for Alert type | 
**Description** | Pointer to **string** | Description of the alert | [optional] 
**Severity** | **string** | Alert severity | 
**Scope** | Pointer to [**OrganizationsOrganizationIdAssuranceAlertsScope**](OrganizationsOrganizationIdAssuranceAlertsScope.md) |  | [optional] 

## Methods

### NewInlineResponse200242

`func NewInlineResponse200242(id string, categoryType string, network OrganizationsOrganizationIdAssuranceAlertsNetwork, startedAt time.Time, type_ string, title string, severity string, ) *InlineResponse200242`

NewInlineResponse200242 instantiates a new InlineResponse200242 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200242WithDefaults

`func NewInlineResponse200242WithDefaults() *InlineResponse200242`

NewInlineResponse200242WithDefaults instantiates a new InlineResponse200242 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200242) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200242) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200242) SetId(v string)`

SetId sets Id field to given value.


### GetCategoryType

`func (o *InlineResponse200242) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *InlineResponse200242) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *InlineResponse200242) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.


### GetNetwork

`func (o *InlineResponse200242) GetNetwork() OrganizationsOrganizationIdAssuranceAlertsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200242) GetNetworkOk() (*OrganizationsOrganizationIdAssuranceAlertsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200242) SetNetwork(v OrganizationsOrganizationIdAssuranceAlertsNetwork)`

SetNetwork sets Network field to given value.


### GetStartedAt

`func (o *InlineResponse200242) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InlineResponse200242) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InlineResponse200242) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetResolvedAt

`func (o *InlineResponse200242) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *InlineResponse200242) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *InlineResponse200242) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *InlineResponse200242) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetDismissedAt

`func (o *InlineResponse200242) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *InlineResponse200242) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *InlineResponse200242) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.

### HasDismissedAt

`func (o *InlineResponse200242) HasDismissedAt() bool`

HasDismissedAt returns a boolean if a field has been set.

### GetDeviceType

`func (o *InlineResponse200242) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InlineResponse200242) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InlineResponse200242) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InlineResponse200242) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200242) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200242) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200242) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *InlineResponse200242) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse200242) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse200242) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InlineResponse200242) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200242) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200242) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200242) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InlineResponse200242) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InlineResponse200242) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InlineResponse200242) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetScope

`func (o *InlineResponse200242) GetScope() OrganizationsOrganizationIdAssuranceAlertsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse200242) GetScopeOk() (*OrganizationsOrganizationIdAssuranceAlertsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse200242) SetScope(v OrganizationsOrganizationIdAssuranceAlertsScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse200242) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


