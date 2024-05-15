# OrganizationsOrganizationIdAssuranceAlertsScopeDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL of affected device | [optional] 
**Name** | Pointer to **string** | Name of affected device | [optional] 
**Order** | Pointer to **int32** | Order of affected device in array | [optional] 
**ProductType** | Pointer to **string** | Type of affected device | [optional] 
**Serial** | Pointer to **string** | Serial of affected device | [optional] 
**Mac** | Pointer to **string** | MAC address of affected device | [optional] 
**Imei** | Pointer to **string** | IMEI of affected device | [optional] 
**Lldp** | Pointer to [**OrganizationsOrganizationIdAssuranceAlertsScopeLldp**](OrganizationsOrganizationIdAssuranceAlertsScopeLldp.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdAssuranceAlertsScopeDevices

`func NewOrganizationsOrganizationIdAssuranceAlertsScopeDevices() *OrganizationsOrganizationIdAssuranceAlertsScopeDevices`

NewOrganizationsOrganizationIdAssuranceAlertsScopeDevices instantiates a new OrganizationsOrganizationIdAssuranceAlertsScopeDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAssuranceAlertsScopeDevicesWithDefaults

`func NewOrganizationsOrganizationIdAssuranceAlertsScopeDevicesWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsScopeDevices`

NewOrganizationsOrganizationIdAssuranceAlertsScopeDevicesWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsScopeDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProductType

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSerial

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetImei

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetLldp

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetLldp() OrganizationsOrganizationIdAssuranceAlertsScopeLldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) GetLldpOk() (*OrganizationsOrganizationIdAssuranceAlertsScopeLldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) SetLldp(v OrganizationsOrganizationIdAssuranceAlertsScopeLldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *OrganizationsOrganizationIdAssuranceAlertsScopeDevices) HasLldp() bool`

HasLldp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


