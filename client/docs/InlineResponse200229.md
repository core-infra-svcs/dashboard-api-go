# InlineResponse200229

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | Profile ID | [optional] 
**Name** | Pointer to **string** | Name of profile | [optional] 
**Hostnames** | Pointer to **[]string** | The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB. | [optional] 
**Nameservers** | Pointer to [**OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers**](OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers.md) |  | [optional] 

## Methods

### NewInlineResponse200229

`func NewInlineResponse200229() *InlineResponse200229`

NewInlineResponse200229 instantiates a new InlineResponse200229 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200229WithDefaults

`func NewInlineResponse200229WithDefaults() *InlineResponse200229`

NewInlineResponse200229WithDefaults instantiates a new InlineResponse200229 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200229) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200229) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200229) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200229) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200229) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200229) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200229) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200229) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostnames

`func (o *InlineResponse200229) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *InlineResponse200229) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *InlineResponse200229) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *InlineResponse200229) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse200229) GetNameservers() OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse200229) GetNameserversOk() (*OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse200229) SetNameservers(v OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse200229) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


