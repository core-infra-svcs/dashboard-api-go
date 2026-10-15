# InlineResponse200290ByType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Packet type. 3 packet types exist: broadcast, unicast, and multicast. \&quot;broadcast unicast multicast\&quot; is the sum of all 3 packet types. Other values denote error or size conditions. | [optional] 
**Total** | Pointer to **int32** | The total count of sent and received packets during the timespan. | [optional] 
**Sent** | Pointer to **int32** | The total count of packets sent during the timespan. | [optional] 
**Recv** | Pointer to **int32** | The total count of packets received during the timespan. | [optional] 
**Rates** | Pointer to [**InlineResponse200290Rates**](InlineResponse200290Rates.md) |  | [optional] 

## Methods

### NewInlineResponse200290ByType

`func NewInlineResponse200290ByType() *InlineResponse200290ByType`

NewInlineResponse200290ByType instantiates a new InlineResponse200290ByType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200290ByTypeWithDefaults

`func NewInlineResponse200290ByTypeWithDefaults() *InlineResponse200290ByType`

NewInlineResponse200290ByTypeWithDefaults instantiates a new InlineResponse200290ByType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200290ByType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200290ByType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200290ByType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200290ByType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200290ByType) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200290ByType) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200290ByType) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200290ByType) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSent

`func (o *InlineResponse200290ByType) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *InlineResponse200290ByType) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *InlineResponse200290ByType) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *InlineResponse200290ByType) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse200290ByType) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse200290ByType) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse200290ByType) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse200290ByType) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetRates

`func (o *InlineResponse200290ByType) GetRates() InlineResponse200290Rates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *InlineResponse200290ByType) GetRatesOk() (*InlineResponse200290Rates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *InlineResponse200290ByType) SetRates(v InlineResponse200290Rates)`

SetRates sets Rates field to given value.

### HasRates

`func (o *InlineResponse200290ByType) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


