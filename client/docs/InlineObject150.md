# InlineObject150

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | **NullableInt32** | The VLAN of the incoming packet. A null value will match any VLAN. | 
**Protocol** | Pointer to **string** | The protocol of the incoming packet. Default value is \&quot;ANY\&quot; | [optional] 
**SrcPort** | Pointer to **NullableInt32** | The source port of the incoming packet. Applicable only if protocol is TCP or UDP. | [optional] 
**SrcPortRange** | Pointer to **string** | The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. | [optional] 
**DstPort** | Pointer to **NullableInt32** | The destination port of the incoming packet. Applicable only if protocol is TCP or UDP. | [optional] 
**DstPortRange** | Pointer to **string** | The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. | [optional] 
**Dscp** | Pointer to **int32** | DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0 | [optional] 

## Methods

### NewInlineObject150

`func NewInlineObject150(vlan NullableInt32, ) *InlineObject150`

NewInlineObject150 instantiates a new InlineObject150 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject150WithDefaults

`func NewInlineObject150WithDefaults() *InlineObject150`

NewInlineObject150WithDefaults instantiates a new InlineObject150 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *InlineObject150) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject150) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject150) SetVlan(v int32)`

SetVlan sets Vlan field to given value.


### SetVlanNil

`func (o *InlineObject150) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *InlineObject150) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetProtocol

`func (o *InlineObject150) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineObject150) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineObject150) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineObject150) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineObject150) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineObject150) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineObject150) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineObject150) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### SetSrcPortNil

`func (o *InlineObject150) SetSrcPortNil(b bool)`

 SetSrcPortNil sets the value for SrcPort to be an explicit nil

### UnsetSrcPort
`func (o *InlineObject150) UnsetSrcPort()`

UnsetSrcPort ensures that no value is present for SrcPort, not even an explicit nil
### GetSrcPortRange

`func (o *InlineObject150) GetSrcPortRange() string`

GetSrcPortRange returns the SrcPortRange field if non-nil, zero value otherwise.

### GetSrcPortRangeOk

`func (o *InlineObject150) GetSrcPortRangeOk() (*string, bool)`

GetSrcPortRangeOk returns a tuple with the SrcPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortRange

`func (o *InlineObject150) SetSrcPortRange(v string)`

SetSrcPortRange sets SrcPortRange field to given value.

### HasSrcPortRange

`func (o *InlineObject150) HasSrcPortRange() bool`

HasSrcPortRange returns a boolean if a field has been set.

### GetDstPort

`func (o *InlineObject150) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *InlineObject150) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *InlineObject150) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *InlineObject150) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### SetDstPortNil

`func (o *InlineObject150) SetDstPortNil(b bool)`

 SetDstPortNil sets the value for DstPort to be an explicit nil

### UnsetDstPort
`func (o *InlineObject150) UnsetDstPort()`

UnsetDstPort ensures that no value is present for DstPort, not even an explicit nil
### GetDstPortRange

`func (o *InlineObject150) GetDstPortRange() string`

GetDstPortRange returns the DstPortRange field if non-nil, zero value otherwise.

### GetDstPortRangeOk

`func (o *InlineObject150) GetDstPortRangeOk() (*string, bool)`

GetDstPortRangeOk returns a tuple with the DstPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortRange

`func (o *InlineObject150) SetDstPortRange(v string)`

SetDstPortRange sets DstPortRange field to given value.

### HasDstPortRange

`func (o *InlineObject150) HasDstPortRange() bool`

HasDstPortRange returns a boolean if a field has been set.

### GetDscp

`func (o *InlineObject150) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *InlineObject150) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *InlineObject150) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *InlineObject150) HasDscp() bool`

HasDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


