# InlineResponse20047

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID | [optional] 
**OrganizationId** | Pointer to **string** | Organization ID | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**ProductTypes** | Pointer to **[]string** | List of the product types that the network supports | [optional] 
**TimeZone** | Pointer to **string** | Timezone of the network | [optional] 
**Tags** | Pointer to **[]string** | Network tags | [optional] 
**EnrollmentString** | Pointer to **string** | Enrollment string for the network | [optional] 
**Url** | Pointer to **string** | URL to the network Dashboard UI | [optional] 
**Notes** | Pointer to **string** | Notes for the network | [optional] 
**IsBoundToConfigTemplate** | Pointer to **bool** | If the network is bound to a config template | [optional] 

## Methods

### NewInlineResponse20047

`func NewInlineResponse20047() *InlineResponse20047`

NewInlineResponse20047 instantiates a new InlineResponse20047 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20047WithDefaults

`func NewInlineResponse20047WithDefaults() *InlineResponse20047`

NewInlineResponse20047WithDefaults instantiates a new InlineResponse20047 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20047) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20047) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20047) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20047) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse20047) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse20047) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse20047) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse20047) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20047) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20047) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20047) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20047) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse20047) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse20047) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse20047) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse20047) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetTimeZone

`func (o *InlineResponse20047) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InlineResponse20047) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InlineResponse20047) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InlineResponse20047) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20047) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20047) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20047) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20047) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnrollmentString

`func (o *InlineResponse20047) GetEnrollmentString() string`

GetEnrollmentString returns the EnrollmentString field if non-nil, zero value otherwise.

### GetEnrollmentStringOk

`func (o *InlineResponse20047) GetEnrollmentStringOk() (*string, bool)`

GetEnrollmentStringOk returns a tuple with the EnrollmentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentString

`func (o *InlineResponse20047) SetEnrollmentString(v string)`

SetEnrollmentString sets EnrollmentString field to given value.

### HasEnrollmentString

`func (o *InlineResponse20047) HasEnrollmentString() bool`

HasEnrollmentString returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20047) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20047) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20047) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20047) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20047) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20047) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20047) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20047) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsBoundToConfigTemplate

`func (o *InlineResponse20047) GetIsBoundToConfigTemplate() bool`

GetIsBoundToConfigTemplate returns the IsBoundToConfigTemplate field if non-nil, zero value otherwise.

### GetIsBoundToConfigTemplateOk

`func (o *InlineResponse20047) GetIsBoundToConfigTemplateOk() (*bool, bool)`

GetIsBoundToConfigTemplateOk returns a tuple with the IsBoundToConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBoundToConfigTemplate

`func (o *InlineResponse20047) SetIsBoundToConfigTemplate(v bool)`

SetIsBoundToConfigTemplate sets IsBoundToConfigTemplate field to given value.

### HasIsBoundToConfigTemplate

`func (o *InlineResponse20047) HasIsBoundToConfigTemplate() bool`

HasIsBoundToConfigTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


