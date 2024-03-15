# OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices**](OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices.md) |  | 
**AfterAction** | **string** | What action to perform on devices.old after the device cloning is complete. &#39;remove from network&#39; will return the device to inventory, while &#39;release from organization inventory&#39; will free up the license attached to the device. | 

## Methods

### NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps

`func NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps(devices OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices, afterAction string, ) *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps`

NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps instantiates a new OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwapsWithDefaults

`func NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwapsWithDefaults() *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps`

NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwapsWithDefaults instantiates a new OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetDevices() OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetDevicesOk() (*OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) SetDevices(v OrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices)`

SetDevices sets Devices field to given value.


### GetAfterAction

`func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetAfterAction() string`

GetAfterAction returns the AfterAction field if non-nil, zero value otherwise.

### GetAfterActionOk

`func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) GetAfterActionOk() (*string, bool)`

GetAfterActionOk returns a tuple with the AfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterAction

`func (o *OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) SetAfterAction(v string)`

SetAfterAction sets AfterAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


