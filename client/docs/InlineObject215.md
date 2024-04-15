# InlineObject215

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new role. Must be unique. This parameter is required. | 
**AppliedOnDevices** | Pointer to [**[]OrganizationsOrganizationIdCameraRolesAppliedOnDevices**](OrganizationsOrganizationIdCameraRolesAppliedOnDevices.md) | Device tag on which this specified permission is applied. | [optional] 
**AppliedOnNetworks** | Pointer to [**[]OrganizationsOrganizationIdCameraRolesAppliedOnNetworks**](OrganizationsOrganizationIdCameraRolesAppliedOnNetworks.md) | Network tag on which this specified permission is applied. | [optional] 
**AppliedOrgWide** | Pointer to [**[]OrganizationsOrganizationIdCameraRolesAppliedOrgWide**](OrganizationsOrganizationIdCameraRolesAppliedOrgWide.md) | Permissions to be applied org wide. | [optional] 

## Methods

### NewInlineObject215

`func NewInlineObject215(name string, ) *InlineObject215`

NewInlineObject215 instantiates a new InlineObject215 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject215WithDefaults

`func NewInlineObject215WithDefaults() *InlineObject215`

NewInlineObject215WithDefaults instantiates a new InlineObject215 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject215) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject215) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject215) SetName(v string)`

SetName sets Name field to given value.


### GetAppliedOnDevices

`func (o *InlineObject215) GetAppliedOnDevices() []OrganizationsOrganizationIdCameraRolesAppliedOnDevices`

GetAppliedOnDevices returns the AppliedOnDevices field if non-nil, zero value otherwise.

### GetAppliedOnDevicesOk

`func (o *InlineObject215) GetAppliedOnDevicesOk() (*[]OrganizationsOrganizationIdCameraRolesAppliedOnDevices, bool)`

GetAppliedOnDevicesOk returns a tuple with the AppliedOnDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOnDevices

`func (o *InlineObject215) SetAppliedOnDevices(v []OrganizationsOrganizationIdCameraRolesAppliedOnDevices)`

SetAppliedOnDevices sets AppliedOnDevices field to given value.

### HasAppliedOnDevices

`func (o *InlineObject215) HasAppliedOnDevices() bool`

HasAppliedOnDevices returns a boolean if a field has been set.

### GetAppliedOnNetworks

`func (o *InlineObject215) GetAppliedOnNetworks() []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks`

GetAppliedOnNetworks returns the AppliedOnNetworks field if non-nil, zero value otherwise.

### GetAppliedOnNetworksOk

`func (o *InlineObject215) GetAppliedOnNetworksOk() (*[]OrganizationsOrganizationIdCameraRolesAppliedOnNetworks, bool)`

GetAppliedOnNetworksOk returns a tuple with the AppliedOnNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOnNetworks

`func (o *InlineObject215) SetAppliedOnNetworks(v []OrganizationsOrganizationIdCameraRolesAppliedOnNetworks)`

SetAppliedOnNetworks sets AppliedOnNetworks field to given value.

### HasAppliedOnNetworks

`func (o *InlineObject215) HasAppliedOnNetworks() bool`

HasAppliedOnNetworks returns a boolean if a field has been set.

### GetAppliedOrgWide

`func (o *InlineObject215) GetAppliedOrgWide() []OrganizationsOrganizationIdCameraRolesAppliedOrgWide`

GetAppliedOrgWide returns the AppliedOrgWide field if non-nil, zero value otherwise.

### GetAppliedOrgWideOk

`func (o *InlineObject215) GetAppliedOrgWideOk() (*[]OrganizationsOrganizationIdCameraRolesAppliedOrgWide, bool)`

GetAppliedOrgWideOk returns a tuple with the AppliedOrgWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOrgWide

`func (o *InlineObject215) SetAppliedOrgWide(v []OrganizationsOrganizationIdCameraRolesAppliedOrgWide)`

SetAppliedOrgWide sets AppliedOrgWide field to given value.

### HasAppliedOrgWide

`func (o *InlineObject215) HasAppliedOrgWide() bool`

HasAppliedOrgWide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


