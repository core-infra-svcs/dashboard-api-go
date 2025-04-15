# InlineResponse200339

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the channel utilization interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the channel utilization interval. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork**](OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand**](OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewInlineResponse200339

`func NewInlineResponse200339() *InlineResponse200339`

NewInlineResponse200339 instantiates a new InlineResponse200339 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200339WithDefaults

`func NewInlineResponse200339WithDefaults() *InlineResponse200339`

NewInlineResponse200339WithDefaults instantiates a new InlineResponse200339 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200339) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200339) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200339) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200339) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200339) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200339) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200339) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200339) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200339) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200339) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200339) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200339) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *InlineResponse200339) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *InlineResponse200339) GetByBandOk() (*[]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *InlineResponse200339) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *InlineResponse200339) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


