# InlineResponse200162

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Qos Rule id | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the incoming packet. A null value will match any VLAN. | [optional] 
**Protocol** | Pointer to **string** | The protocol of the incoming packet. Can be one of \&quot;ANY\&quot;, \&quot;TCP\&quot; or \&quot;UDP\&quot;. Default value is \&quot;ANY\&quot; | [optional] 
**SrcPort** | Pointer to **int32** | The source port of the incoming packet. Applicable only if protocol is TCP or UDP. | [optional] 
**SrcPortRange** | Pointer to **string** | The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80 | [optional] 
**DstPort** | Pointer to **int32** | The destination port of the incoming packet. Applicable only if protocol is TCP or UDP. | [optional] 
**DstPortRange** | Pointer to **string** | The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80 | [optional] 
**Dscp** | Pointer to **int32** | DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0 | [optional] 

## Methods

### NewInlineResponse200162

`func NewInlineResponse200162() *InlineResponse200162`

NewInlineResponse200162 instantiates a new InlineResponse200162 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200162WithDefaults

`func NewInlineResponse200162WithDefaults() *InlineResponse200162`

NewInlineResponse200162WithDefaults instantiates a new InlineResponse200162 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200162) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200162) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200162) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200162) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200162) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200162) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200162) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200162) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse200162) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200162) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200162) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200162) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineResponse200162) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineResponse200162) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineResponse200162) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineResponse200162) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcPortRange

`func (o *InlineResponse200162) GetSrcPortRange() string`

GetSrcPortRange returns the SrcPortRange field if non-nil, zero value otherwise.

### GetSrcPortRangeOk

`func (o *InlineResponse200162) GetSrcPortRangeOk() (*string, bool)`

GetSrcPortRangeOk returns a tuple with the SrcPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortRange

`func (o *InlineResponse200162) SetSrcPortRange(v string)`

SetSrcPortRange sets SrcPortRange field to given value.

### HasSrcPortRange

`func (o *InlineResponse200162) HasSrcPortRange() bool`

HasSrcPortRange returns a boolean if a field has been set.

### GetDstPort

`func (o *InlineResponse200162) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *InlineResponse200162) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *InlineResponse200162) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *InlineResponse200162) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetDstPortRange

`func (o *InlineResponse200162) GetDstPortRange() string`

GetDstPortRange returns the DstPortRange field if non-nil, zero value otherwise.

### GetDstPortRangeOk

`func (o *InlineResponse200162) GetDstPortRangeOk() (*string, bool)`

GetDstPortRangeOk returns a tuple with the DstPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortRange

`func (o *InlineResponse200162) SetDstPortRange(v string)`

SetDstPortRange sets DstPortRange field to given value.

### HasDstPortRange

`func (o *InlineResponse200162) HasDstPortRange() bool`

HasDstPortRange returns a boolean if a field has been set.

### GetDscp

`func (o *InlineResponse200162) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *InlineResponse200162) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *InlineResponse200162) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *InlineResponse200162) HasDscp() bool`

HasDscp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


