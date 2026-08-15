# InlineResponse200323Items

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
**TunnelAdmin** | [**InlineResponse200323TunnelAdmin**](InlineResponse200323TunnelAdmin.md) |  | 
**TunnelSchedule** | [**InlineResponse200323TunnelSchedule**](InlineResponse200323TunnelSchedule.md) |  | 
**Interfaces** | **[]string** | Interface IP addresses of the device | 
**Url** | **string** | Link to AP summary page | 
**Network** | [**InlineResponse200323Network**](InlineResponse200323Network.md) |  | 
**Counts** | [**InlineResponse200323Counts**](InlineResponse200323Counts.md) |  | 
**CampusGateways** | [**[]InlineResponse200323CampusGateways**](InlineResponse200323CampusGateways.md) | MCG information | 

## Methods

### NewInlineResponse200323Items

`func NewInlineResponse200323Items(name string, serial string, mac string, model string, uptime int32, status string, tunnelStatus string, tunnelAdmin InlineResponse200323TunnelAdmin, tunnelSchedule InlineResponse200323TunnelSchedule, interfaces []string, url string, network InlineResponse200323Network, counts InlineResponse200323Counts, campusGateways []InlineResponse200323CampusGateways, ) *InlineResponse200323Items`

NewInlineResponse200323Items instantiates a new InlineResponse200323Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200323ItemsWithDefaults

`func NewInlineResponse200323ItemsWithDefaults() *InlineResponse200323Items`

NewInlineResponse200323ItemsWithDefaults instantiates a new InlineResponse200323Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200323Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200323Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200323Items) SetName(v string)`

SetName sets Name field to given value.


### GetSerial

`func (o *InlineResponse200323Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200323Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200323Items) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200323Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200323Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200323Items) SetMac(v string)`

SetMac sets Mac field to given value.


### GetModel

`func (o *InlineResponse200323Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200323Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200323Items) SetModel(v string)`

SetModel sets Model field to given value.


### GetUptime

`func (o *InlineResponse200323Items) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *InlineResponse200323Items) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *InlineResponse200323Items) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetStatus

`func (o *InlineResponse200323Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200323Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200323Items) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTunnelStatus

`func (o *InlineResponse200323Items) GetTunnelStatus() string`

GetTunnelStatus returns the TunnelStatus field if non-nil, zero value otherwise.

### GetTunnelStatusOk

`func (o *InlineResponse200323Items) GetTunnelStatusOk() (*string, bool)`

GetTunnelStatusOk returns a tuple with the TunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelStatus

`func (o *InlineResponse200323Items) SetTunnelStatus(v string)`

SetTunnelStatus sets TunnelStatus field to given value.


### GetTunnelAdmin

`func (o *InlineResponse200323Items) GetTunnelAdmin() InlineResponse200323TunnelAdmin`

GetTunnelAdmin returns the TunnelAdmin field if non-nil, zero value otherwise.

### GetTunnelAdminOk

`func (o *InlineResponse200323Items) GetTunnelAdminOk() (*InlineResponse200323TunnelAdmin, bool)`

GetTunnelAdminOk returns a tuple with the TunnelAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelAdmin

`func (o *InlineResponse200323Items) SetTunnelAdmin(v InlineResponse200323TunnelAdmin)`

SetTunnelAdmin sets TunnelAdmin field to given value.


### GetTunnelSchedule

`func (o *InlineResponse200323Items) GetTunnelSchedule() InlineResponse200323TunnelSchedule`

GetTunnelSchedule returns the TunnelSchedule field if non-nil, zero value otherwise.

### GetTunnelScheduleOk

`func (o *InlineResponse200323Items) GetTunnelScheduleOk() (*InlineResponse200323TunnelSchedule, bool)`

GetTunnelScheduleOk returns a tuple with the TunnelSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelSchedule

`func (o *InlineResponse200323Items) SetTunnelSchedule(v InlineResponse200323TunnelSchedule)`

SetTunnelSchedule sets TunnelSchedule field to given value.


### GetInterfaces

`func (o *InlineResponse200323Items) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200323Items) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200323Items) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.


### GetUrl

`func (o *InlineResponse200323Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200323Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200323Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200323Items) GetNetwork() InlineResponse200323Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200323Items) GetNetworkOk() (*InlineResponse200323Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200323Items) SetNetwork(v InlineResponse200323Network)`

SetNetwork sets Network field to given value.


### GetCounts

`func (o *InlineResponse200323Items) GetCounts() InlineResponse200323Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200323Items) GetCountsOk() (*InlineResponse200323Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200323Items) SetCounts(v InlineResponse200323Counts)`

SetCounts sets Counts field to given value.


### GetCampusGateways

`func (o *InlineResponse200323Items) GetCampusGateways() []InlineResponse200323CampusGateways`

GetCampusGateways returns the CampusGateways field if non-nil, zero value otherwise.

### GetCampusGatewaysOk

`func (o *InlineResponse200323Items) GetCampusGatewaysOk() (*[]InlineResponse200323CampusGateways, bool)`

GetCampusGatewaysOk returns a tuple with the CampusGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampusGateways

`func (o *InlineResponse200323Items) SetCampusGateways(v []InlineResponse200323CampusGateways)`

SetCampusGateways sets CampusGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


