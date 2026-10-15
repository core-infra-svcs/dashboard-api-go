# OrganizationsOrganizationIdAdminsNetworks1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The network ID | 
**Access** | **string** | The privilege of the dashboard administrator on the network. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;guest-ambassador&#39;, &#39;monitor-only&#39;, &#39;switchport&#39; or &#39;switch-port-manager&#39; | 
**PacketCaptureAllowed** | Pointer to **bool** | True if admin has access to the packet capture utility. Only allowed if switch-port-manager or switchport is the access value | [optional] 
**SwitchPortTags** | Pointer to **[]string** | Switch port tags the admin will have access to. Only allowed if switch-port-manager or switchport is the access value | [optional] 

## Methods

### NewOrganizationsOrganizationIdAdminsNetworks1

`func NewOrganizationsOrganizationIdAdminsNetworks1(id string, access string, ) *OrganizationsOrganizationIdAdminsNetworks1`

NewOrganizationsOrganizationIdAdminsNetworks1 instantiates a new OrganizationsOrganizationIdAdminsNetworks1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAdminsNetworks1WithDefaults

`func NewOrganizationsOrganizationIdAdminsNetworks1WithDefaults() *OrganizationsOrganizationIdAdminsNetworks1`

NewOrganizationsOrganizationIdAdminsNetworks1WithDefaults instantiates a new OrganizationsOrganizationIdAdminsNetworks1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationsOrganizationIdAdminsNetworks1) SetId(v string)`

SetId sets Id field to given value.


### GetAccess

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *OrganizationsOrganizationIdAdminsNetworks1) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetPacketCaptureAllowed

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetPacketCaptureAllowed() bool`

GetPacketCaptureAllowed returns the PacketCaptureAllowed field if non-nil, zero value otherwise.

### GetPacketCaptureAllowedOk

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetPacketCaptureAllowedOk() (*bool, bool)`

GetPacketCaptureAllowedOk returns a tuple with the PacketCaptureAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketCaptureAllowed

`func (o *OrganizationsOrganizationIdAdminsNetworks1) SetPacketCaptureAllowed(v bool)`

SetPacketCaptureAllowed sets PacketCaptureAllowed field to given value.

### HasPacketCaptureAllowed

`func (o *OrganizationsOrganizationIdAdminsNetworks1) HasPacketCaptureAllowed() bool`

HasPacketCaptureAllowed returns a boolean if a field has been set.

### GetSwitchPortTags

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetSwitchPortTags() []string`

GetSwitchPortTags returns the SwitchPortTags field if non-nil, zero value otherwise.

### GetSwitchPortTagsOk

`func (o *OrganizationsOrganizationIdAdminsNetworks1) GetSwitchPortTagsOk() (*[]string, bool)`

GetSwitchPortTagsOk returns a tuple with the SwitchPortTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortTags

`func (o *OrganizationsOrganizationIdAdminsNetworks1) SetSwitchPortTags(v []string)`

SetSwitchPortTags sets SwitchPortTags field to given value.

### HasSwitchPortTags

`func (o *OrganizationsOrganizationIdAdminsNetworks1) HasSwitchPortTags() bool`

HasSwitchPortTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


