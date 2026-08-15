# InlineResponse200302Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | id of this alert profile | [optional] 
**Name** | Pointer to **string** | Name of this alert profile | [optional] 
**OrganizationId** | Pointer to **string** | Organization ID | [optional] 
**NetworkIds** | Pointer to **[]string** | Networks associated with this alert profile | [optional] 
**AlertTypes** | Pointer to **[]string** | Alert types associated with this alert profile | [optional] 
**AlertScheduleId** | Pointer to **NullableString** | ID of the alert schedule associated with this profile | [optional] 
**Configuration** | Pointer to [**InlineResponse200302Configuration**](InlineResponse200302Configuration.md) |  | [optional] 

## Methods

### NewInlineResponse200302Items

`func NewInlineResponse200302Items() *InlineResponse200302Items`

NewInlineResponse200302Items instantiates a new InlineResponse200302Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302ItemsWithDefaults

`func NewInlineResponse200302ItemsWithDefaults() *InlineResponse200302Items`

NewInlineResponse200302ItemsWithDefaults instantiates a new InlineResponse200302Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200302Items) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200302Items) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200302Items) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200302Items) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200302Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200302Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200302Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200302Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse200302Items) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse200302Items) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse200302Items) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse200302Items) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetNetworkIds

`func (o *InlineResponse200302Items) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *InlineResponse200302Items) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *InlineResponse200302Items) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *InlineResponse200302Items) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.

### GetAlertTypes

`func (o *InlineResponse200302Items) GetAlertTypes() []string`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *InlineResponse200302Items) GetAlertTypesOk() (*[]string, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *InlineResponse200302Items) SetAlertTypes(v []string)`

SetAlertTypes sets AlertTypes field to given value.

### HasAlertTypes

`func (o *InlineResponse200302Items) HasAlertTypes() bool`

HasAlertTypes returns a boolean if a field has been set.

### GetAlertScheduleId

`func (o *InlineResponse200302Items) GetAlertScheduleId() string`

GetAlertScheduleId returns the AlertScheduleId field if non-nil, zero value otherwise.

### GetAlertScheduleIdOk

`func (o *InlineResponse200302Items) GetAlertScheduleIdOk() (*string, bool)`

GetAlertScheduleIdOk returns a tuple with the AlertScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertScheduleId

`func (o *InlineResponse200302Items) SetAlertScheduleId(v string)`

SetAlertScheduleId sets AlertScheduleId field to given value.

### HasAlertScheduleId

`func (o *InlineResponse200302Items) HasAlertScheduleId() bool`

HasAlertScheduleId returns a boolean if a field has been set.

### SetAlertScheduleIdNil

`func (o *InlineResponse200302Items) SetAlertScheduleIdNil(b bool)`

 SetAlertScheduleIdNil sets the value for AlertScheduleId to be an explicit nil

### UnsetAlertScheduleId
`func (o *InlineResponse200302Items) UnsetAlertScheduleId()`

UnsetAlertScheduleId ensures that no value is present for AlertScheduleId, not even an explicit nil
### GetConfiguration

`func (o *InlineResponse200302Items) GetConfiguration() InlineResponse200302Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InlineResponse200302Items) GetConfigurationOk() (*InlineResponse200302Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InlineResponse200302Items) SetConfiguration(v InlineResponse200302Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *InlineResponse200302Items) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


