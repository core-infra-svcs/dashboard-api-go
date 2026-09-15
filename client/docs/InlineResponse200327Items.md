# InlineResponse200327Items

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
**TunnelAdmin** | [**InlineResponse200327TunnelAdmin**](InlineResponse200327TunnelAdmin.md) |  | 
**TunnelSchedule** | [**InlineResponse200327TunnelSchedule**](InlineResponse200327TunnelSchedule.md) |  | 
**Interfaces** | **[]string** | Interface IP addresses of the device | 
**Url** | **string** | Link to AP summary page | 
**Network** | [**InlineResponse200327Network**](InlineResponse200327Network.md) |  | 
**Counts** | [**InlineResponse200327Counts**](InlineResponse200327Counts.md) |  | 
**CampusGateways** | [**[]InlineResponse200327CampusGateways**](InlineResponse200327CampusGateways.md) | MCG information | 

## Methods

### NewInlineResponse200327Items

`func NewInlineResponse200327Items(name string, serial string, mac string, model string, uptime int32, status string, tunnelStatus string, tunnelAdmin InlineResponse200327TunnelAdmin, tunnelSchedule InlineResponse200327TunnelSchedule, interfaces []string, url string, network InlineResponse200327Network, counts InlineResponse200327Counts, campusGateways []InlineResponse200327CampusGateways, ) *InlineResponse200327Items`

NewInlineResponse200327Items instantiates a new InlineResponse200327Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327ItemsWithDefaults

`func NewInlineResponse200327ItemsWithDefaults() *InlineResponse200327Items`

NewInlineResponse200327ItemsWithDefaults instantiates a new InlineResponse200327Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200327Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200327Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200327Items) SetName(v string)`

SetName sets Name field to given value.


### GetSerial

`func (o *InlineResponse200327Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200327Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200327Items) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200327Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200327Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200327Items) SetMac(v string)`

SetMac sets Mac field to given value.


### GetModel

`func (o *InlineResponse200327Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200327Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200327Items) SetModel(v string)`

SetModel sets Model field to given value.


### GetUptime

`func (o *InlineResponse200327Items) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *InlineResponse200327Items) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *InlineResponse200327Items) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetStatus

`func (o *InlineResponse200327Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200327Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200327Items) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTunnelStatus

`func (o *InlineResponse200327Items) GetTunnelStatus() string`

GetTunnelStatus returns the TunnelStatus field if non-nil, zero value otherwise.

### GetTunnelStatusOk

`func (o *InlineResponse200327Items) GetTunnelStatusOk() (*string, bool)`

GetTunnelStatusOk returns a tuple with the TunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelStatus

`func (o *InlineResponse200327Items) SetTunnelStatus(v string)`

SetTunnelStatus sets TunnelStatus field to given value.


### GetTunnelAdmin

`func (o *InlineResponse200327Items) GetTunnelAdmin() InlineResponse200327TunnelAdmin`

GetTunnelAdmin returns the TunnelAdmin field if non-nil, zero value otherwise.

### GetTunnelAdminOk

`func (o *InlineResponse200327Items) GetTunnelAdminOk() (*InlineResponse200327TunnelAdmin, bool)`

GetTunnelAdminOk returns a tuple with the TunnelAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelAdmin

`func (o *InlineResponse200327Items) SetTunnelAdmin(v InlineResponse200327TunnelAdmin)`

SetTunnelAdmin sets TunnelAdmin field to given value.


### GetTunnelSchedule

`func (o *InlineResponse200327Items) GetTunnelSchedule() InlineResponse200327TunnelSchedule`

GetTunnelSchedule returns the TunnelSchedule field if non-nil, zero value otherwise.

### GetTunnelScheduleOk

`func (o *InlineResponse200327Items) GetTunnelScheduleOk() (*InlineResponse200327TunnelSchedule, bool)`

GetTunnelScheduleOk returns a tuple with the TunnelSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelSchedule

`func (o *InlineResponse200327Items) SetTunnelSchedule(v InlineResponse200327TunnelSchedule)`

SetTunnelSchedule sets TunnelSchedule field to given value.


### GetInterfaces

`func (o *InlineResponse200327Items) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200327Items) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200327Items) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.


### GetUrl

`func (o *InlineResponse200327Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200327Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200327Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNetwork

`func (o *InlineResponse200327Items) GetNetwork() InlineResponse200327Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200327Items) GetNetworkOk() (*InlineResponse200327Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200327Items) SetNetwork(v InlineResponse200327Network)`

SetNetwork sets Network field to given value.


### GetCounts

`func (o *InlineResponse200327Items) GetCounts() InlineResponse200327Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200327Items) GetCountsOk() (*InlineResponse200327Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200327Items) SetCounts(v InlineResponse200327Counts)`

SetCounts sets Counts field to given value.


### GetCampusGateways

`func (o *InlineResponse200327Items) GetCampusGateways() []InlineResponse200327CampusGateways`

GetCampusGateways returns the CampusGateways field if non-nil, zero value otherwise.

### GetCampusGatewaysOk

`func (o *InlineResponse200327Items) GetCampusGatewaysOk() (*[]InlineResponse200327CampusGateways, bool)`

GetCampusGatewaysOk returns a tuple with the CampusGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampusGateways

`func (o *InlineResponse200327Items) SetCampusGateways(v []InlineResponse200327CampusGateways)`

SetCampusGateways sets CampusGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


