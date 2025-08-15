# InlineResponse200240

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

### NewInlineResponse200240

`func NewInlineResponse200240(id string, categoryType string, network OrganizationsOrganizationIdAssuranceAlertsNetwork, startedAt time.Time, type_ string, title string, severity string, ) *InlineResponse200240`

NewInlineResponse200240 instantiates a new InlineResponse200240 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200240WithDefaults

`func NewInlineResponse200240WithDefaults() *InlineResponse200240`

NewInlineResponse200240WithDefaults instantiates a new InlineResponse200240 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200240) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200240) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200240) SetId(v string)`

SetId sets Id field to given value.


### GetCategoryType

`func (o *InlineResponse200240) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *InlineResponse200240) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *InlineResponse200240) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.


### GetNetwork

`func (o *InlineResponse200240) GetNetwork() OrganizationsOrganizationIdAssuranceAlertsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200240) GetNetworkOk() (*OrganizationsOrganizationIdAssuranceAlertsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200240) SetNetwork(v OrganizationsOrganizationIdAssuranceAlertsNetwork)`

SetNetwork sets Network field to given value.


### GetStartedAt

`func (o *InlineResponse200240) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InlineResponse200240) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InlineResponse200240) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetResolvedAt

`func (o *InlineResponse200240) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *InlineResponse200240) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *InlineResponse200240) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *InlineResponse200240) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetDismissedAt

`func (o *InlineResponse200240) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *InlineResponse200240) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *InlineResponse200240) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.

### HasDismissedAt

`func (o *InlineResponse200240) HasDismissedAt() bool`

HasDismissedAt returns a boolean if a field has been set.

### GetDeviceType

`func (o *InlineResponse200240) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InlineResponse200240) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InlineResponse200240) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InlineResponse200240) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200240) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200240) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200240) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *InlineResponse200240) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse200240) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse200240) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InlineResponse200240) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200240) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200240) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200240) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *InlineResponse200240) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InlineResponse200240) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InlineResponse200240) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetScope

`func (o *InlineResponse200240) GetScope() OrganizationsOrganizationIdAssuranceAlertsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse200240) GetScopeOk() (*OrganizationsOrganizationIdAssuranceAlertsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse200240) SetScope(v OrganizationsOrganizationIdAssuranceAlertsScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse200240) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


