# InlineResponse200347Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | AP name | 
**Serial** | **string** | AP serial | 
**Mac** | **string** | Mac address | 
**Model** | **string** | MR model | 
**Uptime** | **int32** | Device uptime in seconds | 
**Status** | **string** | Device status | 
**TunnelStatus** | **string** | Cumulative tunnel status across all campus gateways in the cluster | 
**TunnelAdmin** | [**InlineResponse200347TunnelAdmin**](InlineResponse200347TunnelAdmin.md) |  | 
**TunnelSchedule** | [**InlineResponse200347TunnelSchedule**](InlineResponse200347TunnelSchedule.md) |  | 
**Interfaces** | **[]string** | Interface IP addresses of the device | 
**Url** | **string** | Link to AP summary page | 
**Network** | [**InlineResponse200347Network**](InlineResponse200347Network.md) |  | 
**Counts** | [**InlineResponse200347Counts**](InlineResponse200347Counts.md) |  | 
**CampusGateways** | [**[]InlineResponse200347CampusGateways**](InlineResponse200347CampusGateways.md) | MCG information | 

## Methods

### NewInlineResponse200347Items

`func NewInlineResponse200347Items(name string, serial string, mac string, model string, uptime int32, status string, tunnelStatus string, tunnelAdmin InlineResponse200347TunnelAdmin, tunnelSchedule InlineResponse200347TunnelSchedule, interfaces []string, url string, network InlineResponse200347Network, counts InlineResponse200347Counts, campusGateways []InlineResponse200347CampusGateways, ) *InlineResponse200347Items`

NewInlineResponse200347Items instantiates a new InlineResponse200347Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200347ItemsWithDefaults

`func NewInlineResponse200347ItemsWithDefaults() *InlineResponse200347Items`

NewInlineResponse200347ItemsWithDefaults instantiates a new InlineResponse200347Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200347Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200347Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200347Items) SetName(v string)`

SetName sets Name field to given value.


### GetSerial

`func (o *InlineResponse200347Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200347Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200347Items) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200347Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200347Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200347Items) SetMac(v string)`

SetMac sets Mac field to given value.


### GetModel

`func (o *InlineResponse200347Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200347Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200347Items) SetModel(v string)`

SetModel sets Model field to given value.


### GetUptime

`func (o *InlineResponse200347Items) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *InlineResponse200347Items) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *InlineResponse200347Items) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetStatus

`func (o *InlineResponse200347Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200347Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200347Items) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTunnelStatus

`func (o *InlineResponse200347Items) GetTunnelStatus() string`

GetTunnelStatus returns the TunnelStatus field if non-nil, zero value otherwise.

### GetTunnelStatusOk

`func (o *InlineResponse200347Items) GetTunnelStatusOk() (*string, bool)`

GetTunnelStatusOk returns a tuple with the TunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelStatus

`func (o *InlineResponse200347Items) SetTunnelStatus(v string)`

SetTunnelStatus sets TunnelStatus field to given value.


### GetTunnelAdmin

`func (o *InlineResponse200347Items) GetTunnelAdmin() InlineResponse200347TunnelAdmin`

GetTunnelAdmin returns the TunnelAdmin field if non-nil, zero value otherwise.

### GetTunnelAdminOk

`func (o *InlineResponse200347Items) GetTunnelAdminOk() (*InlineResponse200347TunnelAdmin, bool)`

GetTunnelAdminOk returns a tuple with the TunnelAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelAdmin

`func (o *InlineResponse200347Items) SetTunnelAdmin(v InlineResponse200347TunnelAdmin)`

SetTunnelAdmin sets TunnelAdmin field to given value.


### GetTunnelSchedule

`func (o *InlineResponse200347Items) GetTunnelSchedule() InlineResponse200347TunnelSchedule`

GetTunnelSchedule returns the TunnelSchedule field if non-nil, zero value otherwise.

### GetTunnelScheduleOk

`func (o *InlineResponse200347Items) GetTunnelScheduleOk() (*InlineResponse200347TunnelSchedule, bool)`

GetTunnelScheduleOk returns a tuple with the TunnelSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelSchedule

`func (o *InlineResponse200347Items) SetTunnelSchedule(v InlineResponse200347TunnelSchedule)`

SetTunnelSchedule sets TunnelSchedule field to given value.


### GetInterfaces

`func (o *InlineResponse200347Items) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200347Items) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200347Items) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.


### GetUrl

`func (o *InlineResponse200347Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200347Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200347Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200347Items) GetNetwork() InlineResponse200347Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200347Items) GetNetworkOk() (*InlineResponse200347Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200347Items) SetNetwork(v InlineResponse200347Network)`

SetNetwork sets Network field to given value.


### GetCounts

`func (o *InlineResponse200347Items) GetCounts() InlineResponse200347Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200347Items) GetCountsOk() (*InlineResponse200347Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200347Items) SetCounts(v InlineResponse200347Counts)`

SetCounts sets Counts field to given value.


### GetCampusGateways

`func (o *InlineResponse200347Items) GetCampusGateways() []InlineResponse200347CampusGateways`

GetCampusGateways returns the CampusGateways field if non-nil, zero value otherwise.

### GetCampusGatewaysOk

`func (o *InlineResponse200347Items) GetCampusGatewaysOk() (*[]InlineResponse200347CampusGateways, bool)`

GetCampusGatewaysOk returns a tuple with the CampusGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampusGateways

`func (o *InlineResponse200347Items) SetCampusGateways(v []InlineResponse200347CampusGateways)`

SetCampusGateways sets CampusGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


