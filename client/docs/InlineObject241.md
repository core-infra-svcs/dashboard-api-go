# InlineObject241

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of profile | [optional] 
**Hostnames** | Pointer to **[]string** | The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB. | [optional] 
**Nameservers** | Pointer to [**OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers**](OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers.md) |  | [optional] 

## Methods

### NewInlineObject241

`func NewInlineObject241() *InlineObject241`

NewInlineObject241 instantiates a new InlineObject241 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject241WithDefaults

`func NewInlineObject241WithDefaults() *InlineObject241`

NewInlineObject241WithDefaults instantiates a new InlineObject241 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject241) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject241) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject241) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject241) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostnames

`func (o *InlineObject241) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *InlineObject241) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *InlineObject241) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *InlineObject241) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineObject241) GetNameservers() OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineObject241) GetNameserversOk() (*OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineObject241) SetNameservers(v OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineObject241) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


