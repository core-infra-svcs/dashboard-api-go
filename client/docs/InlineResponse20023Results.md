# InlineResponse20023Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | The port for which the test was performed. | [optional] 
**Status** | Pointer to **string** | The current status of the port. If the cable test is still being performed on the port, \&quot;in-progress\&quot; is used. If an error occurred during the cable test, \&quot;error\&quot; is used and the error property will be populated. | [optional] 
**SpeedMbps** | Pointer to **int32** | Speed in Mbps.  A speed of 0 indicates the port is down or the port speed is automatic. | [optional] 
**Error** | Pointer to **string** | If an error occurred during the cable test, the error message will be populated here. | [optional] 
**Pairs** | Pointer to [**[]InlineResponse20023Pairs**](InlineResponse20023Pairs.md) | Results for each twisted pair within the cable. | [optional] 

## Methods

### NewInlineResponse20023Results

`func NewInlineResponse20023Results() *InlineResponse20023Results`

NewInlineResponse20023Results instantiates a new InlineResponse20023Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20023ResultsWithDefaults

`func NewInlineResponse20023ResultsWithDefaults() *InlineResponse20023Results`

NewInlineResponse20023ResultsWithDefaults instantiates a new InlineResponse20023Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *InlineResponse20023Results) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20023Results) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20023Results) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20023Results) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20023Results) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20023Results) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20023Results) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20023Results) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpeedMbps

`func (o *InlineResponse20023Results) GetSpeedMbps() int32`

GetSpeedMbps returns the SpeedMbps field if non-nil, zero value otherwise.

### GetSpeedMbpsOk

`func (o *InlineResponse20023Results) GetSpeedMbpsOk() (*int32, bool)`

GetSpeedMbpsOk returns a tuple with the SpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMbps

`func (o *InlineResponse20023Results) SetSpeedMbps(v int32)`

SetSpeedMbps sets SpeedMbps field to given value.

### HasSpeedMbps

`func (o *InlineResponse20023Results) HasSpeedMbps() bool`

HasSpeedMbps returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20023Results) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20023Results) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20023Results) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20023Results) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPairs

`func (o *InlineResponse20023Results) GetPairs() []InlineResponse20023Pairs`

GetPairs returns the Pairs field if non-nil, zero value otherwise.

### GetPairsOk

`func (o *InlineResponse20023Results) GetPairsOk() (*[]InlineResponse20023Pairs, bool)`

GetPairsOk returns a tuple with the Pairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairs

`func (o *InlineResponse20023Results) SetPairs(v []InlineResponse20023Pairs)`

SetPairs sets Pairs field to given value.

### HasPairs

`func (o *InlineResponse20023Results) HasPairs() bool`

HasPairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


