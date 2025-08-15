# InlineObject229

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of profile | 
**Hostnames** | **[]string** | The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB. | 
**Nameservers** | [**OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers**](OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers.md) |  | 

## Methods

### NewInlineObject229

`func NewInlineObject229(name string, hostnames []string, nameservers OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers, ) *InlineObject229`

NewInlineObject229 instantiates a new InlineObject229 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject229WithDefaults

`func NewInlineObject229WithDefaults() *InlineObject229`

NewInlineObject229WithDefaults instantiates a new InlineObject229 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject229) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject229) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject229) SetName(v string)`

SetName sets Name field to given value.


### GetHostnames

`func (o *InlineObject229) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *InlineObject229) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *InlineObject229) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetNameservers

`func (o *InlineObject229) GetNameservers() OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineObject229) GetNameserversOk() (*OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineObject229) SetNameservers(v OrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers)`

SetNameservers sets Nameservers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


