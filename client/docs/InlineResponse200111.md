# InlineResponse200111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportingEnabled** | Pointer to **bool** | Boolean indicating whether NetFlow traffic reporting is enabled (true) or disabled (false). | [optional] 
**CollectorIp** | Pointer to **string** | The IPv4 address of the NetFlow collector. | [optional] 
**CollectorPort** | Pointer to **int32** | The port that the NetFlow collector will be listening on. | [optional] 
**EtaEnabled** | Pointer to **bool** | Boolean indicating whether Encrypted Traffic Analytics is enabled (true) or disabled (false). | [optional] 
**EtaDstPort** | Pointer to **int32** | The port that the Encrypted Traffic Analytics collector will be listening on. | [optional] 

## Methods

### NewInlineResponse200111

`func NewInlineResponse200111() *InlineResponse200111`

NewInlineResponse200111 instantiates a new InlineResponse200111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111WithDefaults

`func NewInlineResponse200111WithDefaults() *InlineResponse200111`

NewInlineResponse200111WithDefaults instantiates a new InlineResponse200111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportingEnabled

`func (o *InlineResponse200111) GetReportingEnabled() bool`

GetReportingEnabled returns the ReportingEnabled field if non-nil, zero value otherwise.

### GetReportingEnabledOk

`func (o *InlineResponse200111) GetReportingEnabledOk() (*bool, bool)`

GetReportingEnabledOk returns a tuple with the ReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingEnabled

`func (o *InlineResponse200111) SetReportingEnabled(v bool)`

SetReportingEnabled sets ReportingEnabled field to given value.

### HasReportingEnabled

`func (o *InlineResponse200111) HasReportingEnabled() bool`

HasReportingEnabled returns a boolean if a field has been set.

### GetCollectorIp

`func (o *InlineResponse200111) GetCollectorIp() string`

GetCollectorIp returns the CollectorIp field if non-nil, zero value otherwise.

### GetCollectorIpOk

`func (o *InlineResponse200111) GetCollectorIpOk() (*string, bool)`

GetCollectorIpOk returns a tuple with the CollectorIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorIp

`func (o *InlineResponse200111) SetCollectorIp(v string)`

SetCollectorIp sets CollectorIp field to given value.

### HasCollectorIp

`func (o *InlineResponse200111) HasCollectorIp() bool`

HasCollectorIp returns a boolean if a field has been set.

### GetCollectorPort

`func (o *InlineResponse200111) GetCollectorPort() int32`

GetCollectorPort returns the CollectorPort field if non-nil, zero value otherwise.

### GetCollectorPortOk

`func (o *InlineResponse200111) GetCollectorPortOk() (*int32, bool)`

GetCollectorPortOk returns a tuple with the CollectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorPort

`func (o *InlineResponse200111) SetCollectorPort(v int32)`

SetCollectorPort sets CollectorPort field to given value.

### HasCollectorPort

`func (o *InlineResponse200111) HasCollectorPort() bool`

HasCollectorPort returns a boolean if a field has been set.

### GetEtaEnabled

`func (o *InlineResponse200111) GetEtaEnabled() bool`

GetEtaEnabled returns the EtaEnabled field if non-nil, zero value otherwise.

### GetEtaEnabledOk

`func (o *InlineResponse200111) GetEtaEnabledOk() (*bool, bool)`

GetEtaEnabledOk returns a tuple with the EtaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtaEnabled

`func (o *InlineResponse200111) SetEtaEnabled(v bool)`

SetEtaEnabled sets EtaEnabled field to given value.

### HasEtaEnabled

`func (o *InlineResponse200111) HasEtaEnabled() bool`

HasEtaEnabled returns a boolean if a field has been set.

### GetEtaDstPort

`func (o *InlineResponse200111) GetEtaDstPort() int32`

GetEtaDstPort returns the EtaDstPort field if non-nil, zero value otherwise.

### GetEtaDstPortOk

`func (o *InlineResponse200111) GetEtaDstPortOk() (*int32, bool)`

GetEtaDstPortOk returns a tuple with the EtaDstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtaDstPort

`func (o *InlineResponse200111) SetEtaDstPort(v int32)`

SetEtaDstPort sets EtaDstPort field to given value.

### HasEtaDstPort

`func (o *InlineResponse200111) HasEtaDstPort() bool`

HasEtaDstPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


