# InlineResponse200206UplinkSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failback** | Pointer to [**InlineResponse200206UplinkSelectionFailback**](InlineResponse200206UplinkSelectionFailback.md) |  | [optional] 
**Candidates** | Pointer to **string** | &#39;all&#39; lets devices try any potential interface for uplink. &#39;designated&#39; enables configuration of candidates via the Routing &amp; DHCP page. | [optional] 

## Methods

### NewInlineResponse200206UplinkSelection

`func NewInlineResponse200206UplinkSelection() *InlineResponse200206UplinkSelection`

NewInlineResponse200206UplinkSelection instantiates a new InlineResponse200206UplinkSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200206UplinkSelectionWithDefaults

`func NewInlineResponse200206UplinkSelectionWithDefaults() *InlineResponse200206UplinkSelection`

NewInlineResponse200206UplinkSelectionWithDefaults instantiates a new InlineResponse200206UplinkSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailback

`func (o *InlineResponse200206UplinkSelection) GetFailback() InlineResponse200206UplinkSelectionFailback`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *InlineResponse200206UplinkSelection) GetFailbackOk() (*InlineResponse200206UplinkSelectionFailback, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *InlineResponse200206UplinkSelection) SetFailback(v InlineResponse200206UplinkSelectionFailback)`

SetFailback sets Failback field to given value.

### HasFailback

`func (o *InlineResponse200206UplinkSelection) HasFailback() bool`

HasFailback returns a boolean if a field has been set.

### GetCandidates

`func (o *InlineResponse200206UplinkSelection) GetCandidates() string`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *InlineResponse200206UplinkSelection) GetCandidatesOk() (*string, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *InlineResponse200206UplinkSelection) SetCandidates(v string)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *InlineResponse200206UplinkSelection) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


