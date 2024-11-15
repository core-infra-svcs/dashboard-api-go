# InlineResponse2013Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **string** | IP address or FQDN to ping | [optional] 
**Count** | Pointer to **int32** | Number of pings to send. [1..5], default 5 | [optional] 
**SourceInterface** | Pointer to **string** | The specific L3 Interface IP address to ping through | [optional] 

## Methods

### NewInlineResponse2013Request

`func NewInlineResponse2013Request() *InlineResponse2013Request`

NewInlineResponse2013Request instantiates a new InlineResponse2013Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2013RequestWithDefaults

`func NewInlineResponse2013RequestWithDefaults() *InlineResponse2013Request`

NewInlineResponse2013RequestWithDefaults instantiates a new InlineResponse2013Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *InlineResponse2013Request) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineResponse2013Request) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineResponse2013Request) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InlineResponse2013Request) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse2013Request) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse2013Request) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse2013Request) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse2013Request) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSourceInterface

`func (o *InlineResponse2013Request) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *InlineResponse2013Request) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *InlineResponse2013Request) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *InlineResponse2013Request) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


