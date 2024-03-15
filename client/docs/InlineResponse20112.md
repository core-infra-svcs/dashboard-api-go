# InlineResponse20112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message related to whether or not the device was found and can be imported. | [optional] 
**Udi** | Pointer to **string** | Device UDI certificate | [optional] 
**DeviceId** | Pointer to **string** | Import ID from the Import operation | [optional] 
**Status** | Pointer to **string** | The import status of the device | [optional] 
**ConfigParams** | Pointer to [**OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams**](OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams.md) |  | [optional] 

## Methods

### NewInlineResponse20112

`func NewInlineResponse20112() *InlineResponse20112`

NewInlineResponse20112 instantiates a new InlineResponse20112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20112WithDefaults

`func NewInlineResponse20112WithDefaults() *InlineResponse20112`

NewInlineResponse20112WithDefaults instantiates a new InlineResponse20112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InlineResponse20112) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse20112) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse20112) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse20112) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUdi

`func (o *InlineResponse20112) GetUdi() string`

GetUdi returns the Udi field if non-nil, zero value otherwise.

### GetUdiOk

`func (o *InlineResponse20112) GetUdiOk() (*string, bool)`

GetUdiOk returns a tuple with the Udi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdi

`func (o *InlineResponse20112) SetUdi(v string)`

SetUdi sets Udi field to given value.

### HasUdi

`func (o *InlineResponse20112) HasUdi() bool`

HasUdi returns a boolean if a field has been set.

### GetDeviceId

`func (o *InlineResponse20112) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *InlineResponse20112) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *InlineResponse20112) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *InlineResponse20112) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20112) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20112) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20112) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20112) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConfigParams

`func (o *InlineResponse20112) GetConfigParams() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *InlineResponse20112) GetConfigParamsOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *InlineResponse20112) SetConfigParams(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams)`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *InlineResponse20112) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


