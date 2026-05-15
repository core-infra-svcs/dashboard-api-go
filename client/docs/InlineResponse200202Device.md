# InlineResponse200202Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number of the device | [optional] 
**Name** | Pointer to **string** | The name of the device | [optional] 
**Model** | Pointer to **string** | The model of the device | [optional] 
**ProductType** | Pointer to **string** | The product type of the device | [optional] 
**Status** | Pointer to **string** | The status of the device | [optional] 
**LastReportedAt** | Pointer to **string** | The last time the device reported | [optional] 
**Clients** | Pointer to [**InlineResponse200202DeviceClients**](InlineResponse200202DeviceClients.md) |  | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200202DeviceUplinks**](InlineResponse200202DeviceUplinks.md) | Uplink information | [optional] 

## Methods

### NewInlineResponse200202Device

`func NewInlineResponse200202Device() *InlineResponse200202Device`

NewInlineResponse200202Device instantiates a new InlineResponse200202Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200202DeviceWithDefaults

`func NewInlineResponse200202DeviceWithDefaults() *InlineResponse200202Device`

NewInlineResponse200202DeviceWithDefaults instantiates a new InlineResponse200202Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200202Device) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200202Device) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200202Device) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200202Device) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200202Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200202Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200202Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200202Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200202Device) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200202Device) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200202Device) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200202Device) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200202Device) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200202Device) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200202Device) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200202Device) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200202Device) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200202Device) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200202Device) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200202Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *InlineResponse200202Device) GetLastReportedAt() string`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200202Device) GetLastReportedAtOk() (*string, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200202Device) SetLastReportedAt(v string)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *InlineResponse200202Device) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200202Device) GetClients() InlineResponse200202DeviceClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200202Device) GetClientsOk() (*InlineResponse200202DeviceClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200202Device) SetClients(v InlineResponse200202DeviceClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200202Device) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200202Device) GetUplinks() []InlineResponse200202DeviceUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200202Device) GetUplinksOk() (*[]InlineResponse200202DeviceUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200202Device) SetUplinks(v []InlineResponse200202DeviceUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200202Device) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


