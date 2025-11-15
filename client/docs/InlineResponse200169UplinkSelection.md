# InlineResponse200169UplinkSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failback** | Pointer to [**InlineResponse200169UplinkSelectionFailback**](InlineResponse200169UplinkSelectionFailback.md) |  | [optional] 
**Candidates** | Pointer to **string** | &#39;all&#39; lets devices try any potential interface for uplink. &#39;designated&#39; enables configuration of candidates via the Routing &amp; DHCP page. | [optional] 

## Methods

### NewInlineResponse200169UplinkSelection

`func NewInlineResponse200169UplinkSelection() *InlineResponse200169UplinkSelection`

NewInlineResponse200169UplinkSelection instantiates a new InlineResponse200169UplinkSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200169UplinkSelectionWithDefaults

`func NewInlineResponse200169UplinkSelectionWithDefaults() *InlineResponse200169UplinkSelection`

NewInlineResponse200169UplinkSelectionWithDefaults instantiates a new InlineResponse200169UplinkSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailback

`func (o *InlineResponse200169UplinkSelection) GetFailback() InlineResponse200169UplinkSelectionFailback`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *InlineResponse200169UplinkSelection) GetFailbackOk() (*InlineResponse200169UplinkSelectionFailback, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *InlineResponse200169UplinkSelection) SetFailback(v InlineResponse200169UplinkSelectionFailback)`

SetFailback sets Failback field to given value.

### HasFailback

`func (o *InlineResponse200169UplinkSelection) HasFailback() bool`

HasFailback returns a boolean if a field has been set.

### GetCandidates

`func (o *InlineResponse200169UplinkSelection) GetCandidates() string`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *InlineResponse200169UplinkSelection) GetCandidatesOk() (*string, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *InlineResponse200169UplinkSelection) SetCandidates(v string)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *InlineResponse200169UplinkSelection) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


