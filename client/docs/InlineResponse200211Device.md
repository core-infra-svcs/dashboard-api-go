# InlineResponse200211Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number of the device | [optional] 
**Name** | Pointer to **string** | The name of the device | [optional] 
**Model** | Pointer to **string** | The model of the device | [optional] 
**ProductType** | Pointer to **string** | The product type of the device | [optional] 
**Status** | Pointer to **string** | The status of the device | [optional] 
**LastReportedAt** | Pointer to **string** | The last time the device reported | [optional] 
**Clients** | Pointer to [**InlineResponse200211DeviceClients**](InlineResponse200211DeviceClients.md) |  | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200211DeviceUplinks**](InlineResponse200211DeviceUplinks.md) | Uplink information | [optional] 

## Methods

### NewInlineResponse200211Device

`func NewInlineResponse200211Device() *InlineResponse200211Device`

NewInlineResponse200211Device instantiates a new InlineResponse200211Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200211DeviceWithDefaults

`func NewInlineResponse200211DeviceWithDefaults() *InlineResponse200211Device`

NewInlineResponse200211DeviceWithDefaults instantiates a new InlineResponse200211Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200211Device) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200211Device) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200211Device) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200211Device) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200211Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200211Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200211Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200211Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200211Device) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200211Device) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200211Device) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200211Device) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200211Device) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200211Device) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200211Device) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200211Device) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200211Device) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200211Device) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200211Device) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200211Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *InlineResponse200211Device) GetLastReportedAt() string`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200211Device) GetLastReportedAtOk() (*string, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200211Device) SetLastReportedAt(v string)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *InlineResponse200211Device) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200211Device) GetClients() InlineResponse200211DeviceClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200211Device) GetClientsOk() (*InlineResponse200211DeviceClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200211Device) SetClients(v InlineResponse200211DeviceClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200211Device) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200211Device) GetUplinks() []InlineResponse200211DeviceUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200211Device) GetUplinksOk() (*[]InlineResponse200211DeviceUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200211Device) SetUplinks(v []InlineResponse200211DeviceUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200211Device) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


