# InlineObject262

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the alert profile | 
**NetworkIds** | **[]string** | Networks associated with this alert profile | 
**AlertTypes** | **[]string** | Alert types associated with this alert profile | 
**AlertScheduleId** | Pointer to **NullableString** | ID of the alert schedule associated with this profile | [optional] 
**Configuration** | [**OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration**](OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration.md) |  | 

## Methods

### NewInlineObject262

`func NewInlineObject262(name string, networkIds []string, alertTypes []string, configuration OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration, ) *InlineObject262`

NewInlineObject262 instantiates a new InlineObject262 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject262WithDefaults

`func NewInlineObject262WithDefaults() *InlineObject262`

NewInlineObject262WithDefaults instantiates a new InlineObject262 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject262) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject262) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject262) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkIds

`func (o *InlineObject262) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *InlineObject262) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *InlineObject262) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.


### GetAlertTypes

`func (o *InlineObject262) GetAlertTypes() []string`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *InlineObject262) GetAlertTypesOk() (*[]string, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *InlineObject262) SetAlertTypes(v []string)`

SetAlertTypes sets AlertTypes field to given value.


### GetAlertScheduleId

`func (o *InlineObject262) GetAlertScheduleId() string`

GetAlertScheduleId returns the AlertScheduleId field if non-nil, zero value otherwise.

### GetAlertScheduleIdOk

`func (o *InlineObject262) GetAlertScheduleIdOk() (*string, bool)`

GetAlertScheduleIdOk returns a tuple with the AlertScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertScheduleId

`func (o *InlineObject262) SetAlertScheduleId(v string)`

SetAlertScheduleId sets AlertScheduleId field to given value.

### HasAlertScheduleId

`func (o *InlineObject262) HasAlertScheduleId() bool`

HasAlertScheduleId returns a boolean if a field has been set.

### SetAlertScheduleIdNil

`func (o *InlineObject262) SetAlertScheduleIdNil(b bool)`

 SetAlertScheduleIdNil sets the value for AlertScheduleId to be an explicit nil

### UnsetAlertScheduleId
`func (o *InlineObject262) UnsetAlertScheduleId()`

UnsetAlertScheduleId ensures that no value is present for AlertScheduleId, not even an explicit nil
### GetConfiguration

`func (o *InlineObject262) GetConfiguration() OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InlineObject262) GetConfigurationOk() (*OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InlineObject262) SetConfiguration(v OrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


