# InlineResponse200228Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The network identifier | [optional] 
**OrganizationId** | Pointer to **string** | The organization identifier | [optional] 
**Name** | Pointer to **string** | The network name | [optional] 
**ProductTypes** | Pointer to **[]string** |  | [optional] 
**TimeZone** | Pointer to **string** | The network&#39;s timezone | [optional] 
**Tags** | Pointer to **[]string** | The network tags | [optional] 
**EnrollmentString** | Pointer to **string** | The network enrollment string | [optional] 
**Url** | Pointer to **string** | The network URL | [optional] 
**Notes** | Pointer to **string** | The notes for the network | [optional] 
**IsBoundToConfigTemplate** | Pointer to **bool** | If the network is bound to a config template | [optional] 

## Methods

### NewInlineResponse200228Network

`func NewInlineResponse200228Network() *InlineResponse200228Network`

NewInlineResponse200228Network instantiates a new InlineResponse200228Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200228NetworkWithDefaults

`func NewInlineResponse200228NetworkWithDefaults() *InlineResponse200228Network`

NewInlineResponse200228NetworkWithDefaults instantiates a new InlineResponse200228Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200228Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200228Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200228Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200228Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse200228Network) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse200228Network) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse200228Network) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse200228Network) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200228Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200228Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200228Network) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200228Network) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200228Network) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200228Network) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200228Network) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200228Network) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetTimeZone

`func (o *InlineResponse200228Network) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InlineResponse200228Network) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InlineResponse200228Network) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InlineResponse200228Network) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200228Network) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200228Network) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200228Network) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200228Network) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnrollmentString

`func (o *InlineResponse200228Network) GetEnrollmentString() string`

GetEnrollmentString returns the EnrollmentString field if non-nil, zero value otherwise.

### GetEnrollmentStringOk

`func (o *InlineResponse200228Network) GetEnrollmentStringOk() (*string, bool)`

GetEnrollmentStringOk returns a tuple with the EnrollmentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentString

`func (o *InlineResponse200228Network) SetEnrollmentString(v string)`

SetEnrollmentString sets EnrollmentString field to given value.

### HasEnrollmentString

`func (o *InlineResponse200228Network) HasEnrollmentString() bool`

HasEnrollmentString returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200228Network) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200228Network) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200228Network) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200228Network) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200228Network) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200228Network) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200228Network) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200228Network) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsBoundToConfigTemplate

`func (o *InlineResponse200228Network) GetIsBoundToConfigTemplate() bool`

GetIsBoundToConfigTemplate returns the IsBoundToConfigTemplate field if non-nil, zero value otherwise.

### GetIsBoundToConfigTemplateOk

`func (o *InlineResponse200228Network) GetIsBoundToConfigTemplateOk() (*bool, bool)`

GetIsBoundToConfigTemplateOk returns a tuple with the IsBoundToConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBoundToConfigTemplate

`func (o *InlineResponse200228Network) SetIsBoundToConfigTemplate(v bool)`

SetIsBoundToConfigTemplate sets IsBoundToConfigTemplate field to given value.

### HasIsBoundToConfigTemplate

`func (o *InlineResponse200228Network) HasIsBoundToConfigTemplate() bool`

HasIsBoundToConfigTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


