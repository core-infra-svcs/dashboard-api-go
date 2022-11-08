# InlineResponse2016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message related to whether or not the device was found and can be imported. | [optional] 
**Udi** | Pointer to **string** | Device UDI certificate | [optional] 
**DeviceId** | Pointer to **string** | Import ID from the Import operation | [optional] 
**Status** | Pointer to **string** | The import status of the device | [optional] 
**ConfigParams** | Pointer to [**OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams**](OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams.md) |  | [optional] 

## Methods

### NewInlineResponse2016

`func NewInlineResponse2016() *InlineResponse2016`

NewInlineResponse2016 instantiates a new InlineResponse2016 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2016WithDefaults

`func NewInlineResponse2016WithDefaults() *InlineResponse2016`

NewInlineResponse2016WithDefaults instantiates a new InlineResponse2016 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InlineResponse2016) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse2016) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse2016) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse2016) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUdi

`func (o *InlineResponse2016) GetUdi() string`

GetUdi returns the Udi field if non-nil, zero value otherwise.

### GetUdiOk

`func (o *InlineResponse2016) GetUdiOk() (*string, bool)`

GetUdiOk returns a tuple with the Udi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdi

`func (o *InlineResponse2016) SetUdi(v string)`

SetUdi sets Udi field to given value.

### HasUdi

`func (o *InlineResponse2016) HasUdi() bool`

HasUdi returns a boolean if a field has been set.

### GetDeviceId

`func (o *InlineResponse2016) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *InlineResponse2016) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *InlineResponse2016) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *InlineResponse2016) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2016) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2016) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2016) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2016) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConfigParams

`func (o *InlineResponse2016) GetConfigParams() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *InlineResponse2016) GetConfigParamsOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *InlineResponse2016) SetConfigParams(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams)`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *InlineResponse2016) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


