# InlineResponse2008Downlink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Indicates whether the port is in &#39;trunk&#39; or &#39;access&#39; mode | [optional] 
**Sgt** | Pointer to [**DevicesSerialApplianceInterfacesPortsUpdateDownlinkSgt**](DevicesSerialApplianceInterfacesPortsUpdateDownlinkSgt.md) |  | [optional] 
**Access** | Pointer to [**DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess**](DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess.md) |  | [optional] 
**Trunk** | Pointer to [**DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk**](DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk.md) |  | [optional] 

## Methods

### NewInlineResponse2008Downlink

`func NewInlineResponse2008Downlink() *InlineResponse2008Downlink`

NewInlineResponse2008Downlink instantiates a new InlineResponse2008Downlink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008DownlinkWithDefaults

`func NewInlineResponse2008DownlinkWithDefaults() *InlineResponse2008Downlink`

NewInlineResponse2008DownlinkWithDefaults instantiates a new InlineResponse2008Downlink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse2008Downlink) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse2008Downlink) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse2008Downlink) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse2008Downlink) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse2008Downlink) GetSgt() DevicesSerialApplianceInterfacesPortsUpdateDownlinkSgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse2008Downlink) GetSgtOk() (*DevicesSerialApplianceInterfacesPortsUpdateDownlinkSgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse2008Downlink) SetSgt(v DevicesSerialApplianceInterfacesPortsUpdateDownlinkSgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse2008Downlink) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetAccess

`func (o *InlineResponse2008Downlink) GetAccess() DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse2008Downlink) GetAccessOk() (*DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse2008Downlink) SetAccess(v DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse2008Downlink) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTrunk

`func (o *InlineResponse2008Downlink) GetTrunk() DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk`

GetTrunk returns the Trunk field if non-nil, zero value otherwise.

### GetTrunkOk

`func (o *InlineResponse2008Downlink) GetTrunkOk() (*DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk, bool)`

GetTrunkOk returns a tuple with the Trunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunk

`func (o *InlineResponse2008Downlink) SetTrunk(v DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk)`

SetTrunk sets Trunk field to given value.

### HasTrunk

`func (o *InlineResponse2008Downlink) HasTrunk() bool`

HasTrunk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


