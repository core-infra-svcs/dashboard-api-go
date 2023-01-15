# InlineResponse20088

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The device MAC address. | [optional] 
**Name** | Pointer to **string** | The device name. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesNetwork**](OrganizationsOrganizationIdDevicesAvailabilitiesNetwork.md) |  | [optional] 
**ProductType** | Pointer to **string** | Device product type. | [optional] 
**Serial** | Pointer to **string** | The device serial number. | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device. | [optional] 
**Uplinks** | Pointer to [**[]OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks**](OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks.md) | List of device uplink addresses information. | [optional] 

## Methods

### NewInlineResponse20088

`func NewInlineResponse20088() *InlineResponse20088`

NewInlineResponse20088 instantiates a new InlineResponse20088 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20088WithDefaults

`func NewInlineResponse20088WithDefaults() *InlineResponse20088`

NewInlineResponse20088WithDefaults instantiates a new InlineResponse20088 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse20088) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20088) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20088) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20088) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20088) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20088) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20088) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20088) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20088) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20088) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20088) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20088) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse20088) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse20088) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse20088) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse20088) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse20088) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20088) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20088) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20088) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20088) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20088) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20088) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20088) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse20088) GetUplinks() []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse20088) GetUplinksOk() (*[]OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse20088) SetUplinks(v []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse20088) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


