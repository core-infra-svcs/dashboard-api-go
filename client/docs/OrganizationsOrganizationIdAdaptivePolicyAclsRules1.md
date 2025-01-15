# OrganizationsOrganizationIdAdaptivePolicyAclsRules1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule. | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39; or &#39;any&#39;). | 
**SrcPort** | Pointer to **string** | Source port. Must be in the format of single port: &#39;1&#39;, port list: &#39;1,2&#39; or port range: &#39;1-10&#39;, and in the range of 1-65535, or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 
**DstPort** | Pointer to **string** | Destination port. Must be in the format of single port: &#39;1&#39;, port list: &#39;1,2&#39; or port range: &#39;1-10&#39;, and in the range of 1-65535, or &#39;any&#39;. Default is &#39;any&#39;. | [optional] 
**Log** | Pointer to **bool** | If enabled, when this rule is hit an entry will be logged to the event log  | [optional] 
**TcpEstablished** | Pointer to **bool** | If enabled, means TCP connection with this node must be established.  | [optional] 

## Methods

### NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1

`func NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1(policy string, protocol string, ) *OrganizationsOrganizationIdAdaptivePolicyAclsRules1`

NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1 instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1WithDefaults

`func NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1WithDefaults() *OrganizationsOrganizationIdAdaptivePolicyAclsRules1`

NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1WithDefaults instantiates a new OrganizationsOrganizationIdAdaptivePolicyAclsRules1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSrcPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetDstPort() string`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetDstPortOk() (*string, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetDstPort(v string)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetLog

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetLog(v bool)`

SetLog sets Log field to given value.

### HasLog

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetTcpEstablished

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetTcpEstablished() bool`

GetTcpEstablished returns the TcpEstablished field if non-nil, zero value otherwise.

### GetTcpEstablishedOk

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) GetTcpEstablishedOk() (*bool, bool)`

GetTcpEstablishedOk returns a tuple with the TcpEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpEstablished

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) SetTcpEstablished(v bool)`

SetTcpEstablished sets TcpEstablished field to given value.

### HasTcpEstablished

`func (o *OrganizationsOrganizationIdAdaptivePolicyAclsRules1) HasTcpEstablished() bool`

HasTcpEstablished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


