# InlineResponse200199

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device. | [optional] 
**MeshRoute** | Pointer to **[]string** | List of device serials that make up the mesh. | [optional] 
**LatestMeshPerformance** | Pointer to [**NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance**](NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance.md) |  | [optional] 

## Methods

### NewInlineResponse200199

`func NewInlineResponse200199() *InlineResponse200199`

NewInlineResponse200199 instantiates a new InlineResponse200199 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200199WithDefaults

`func NewInlineResponse200199WithDefaults() *InlineResponse200199`

NewInlineResponse200199WithDefaults instantiates a new InlineResponse200199 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200199) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200199) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200199) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200199) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMeshRoute

`func (o *InlineResponse200199) GetMeshRoute() []string`

GetMeshRoute returns the MeshRoute field if non-nil, zero value otherwise.

### GetMeshRouteOk

`func (o *InlineResponse200199) GetMeshRouteOk() (*[]string, bool)`

GetMeshRouteOk returns a tuple with the MeshRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshRoute

`func (o *InlineResponse200199) SetMeshRoute(v []string)`

SetMeshRoute sets MeshRoute field to given value.

### HasMeshRoute

`func (o *InlineResponse200199) HasMeshRoute() bool`

HasMeshRoute returns a boolean if a field has been set.

### GetLatestMeshPerformance

`func (o *InlineResponse200199) GetLatestMeshPerformance() NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance`

GetLatestMeshPerformance returns the LatestMeshPerformance field if non-nil, zero value otherwise.

### GetLatestMeshPerformanceOk

`func (o *InlineResponse200199) GetLatestMeshPerformanceOk() (*NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance, bool)`

GetLatestMeshPerformanceOk returns a tuple with the LatestMeshPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMeshPerformance

`func (o *InlineResponse200199) SetLatestMeshPerformance(v NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance)`

SetLatestMeshPerformance sets LatestMeshPerformance field to given value.

### HasLatestMeshPerformance

`func (o *InlineResponse200199) HasLatestMeshPerformance() bool`

HasLatestMeshPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


