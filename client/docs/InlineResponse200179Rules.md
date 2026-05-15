# InlineResponse200179Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | Pointer to **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | [optional] 
**IpVersion** | Pointer to **string** | IP address version | [optional] 
**Protocol** | Pointer to **string** | The type of protocol | [optional] 
**SrcCidr** | Pointer to **string** | Source IP address (in IP or CIDR notation) | [optional] 
**SrcPort** | Pointer to **string** | Source port | [optional] 
**DstCidr** | Pointer to **string** | Destination IP address (in IP or CIDR notation) | [optional] 
**DstPort** | Pointer to **string** | Destination port | [optional] 
**Vlan** | Pointer to **string** | ncoming traffic VLAN | [optional] 

## Methods

### NewInlineResponse200179Rules

`func NewInlineResponse200179Rules() *InlineResponse200179Rules`

NewInlineResponse200179Rules instantiates a new InlineResponse200179Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200179RulesWithDefaults

`func NewInlineResponse200179RulesWithDefaults() *InlineResponse200179Rules`

NewInlineResponse200179RulesWithDefaults instantiates a new InlineResponse200179Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineResponse200179Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse200179Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse200179Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse200179Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200179Rules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200179Rules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200179Rules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200179Rules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse200179Rules) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse200179Rules) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse200179Rules) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse200179Rules) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse200179Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200179Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200179Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200179Rules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcCidr

`func (o *InlineResponse200179Rules) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *InlineResponse200179Rules) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *InlineResponse200179Rules) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *InlineResponse200179Rules) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineResponse200179Rules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineResponse200179Rules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineResponse200179Rules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineResponse200179Rules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstCidr

`func (o *InlineResponse200179Rules) GetDstCidr() string`

GetDstCidr returns the DstCidr field if non-nil, zero value otherwise.

### GetDstCidrOk

`func (o *InlineResponse200179Rules) GetDstCidrOk() (*string, bool)`

GetDstCidrOk returns a tuple with the DstCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstCidr

`func (o *InlineResponse200179Rules) SetDstCidr(v string)`

SetDstCidr sets DstCidr field to given value.

### HasDstCidr

`func (o *InlineResponse200179Rules) HasDstCidr() bool`

HasDstCidr returns a boolean if a field has been set.

### GetDstPort

`func (o *InlineResponse200179Rules) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *InlineResponse200179Rules) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *InlineResponse200179Rules) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *InlineResponse200179Rules) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200179Rules) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200179Rules) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200179Rules) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200179Rules) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


