# InlineObject233

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogEvent** | **string** | The type of log event this is recording, e.g. download or opening a banner | 
**Timestamp** | **int32** | A JavaScript UTC datetime stamp for when the even occurred | 
**TargetOS** | Pointer to **string** | The name of the onboarding distro being downloaded | [optional] 
**Request** | Pointer to **string** | Used to describe if this event was the result of a redirect. E.g. a query param if an info banner is being used | [optional] 

## Methods

### NewInlineObject233

`func NewInlineObject233(logEvent string, timestamp int32, ) *InlineObject233`

NewInlineObject233 instantiates a new InlineObject233 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject233WithDefaults

`func NewInlineObject233WithDefaults() *InlineObject233`

NewInlineObject233WithDefaults instantiates a new InlineObject233 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogEvent

`func (o *InlineObject233) GetLogEvent() string`

GetLogEvent returns the LogEvent field if non-nil, zero value otherwise.

### GetLogEventOk

`func (o *InlineObject233) GetLogEventOk() (*string, bool)`

GetLogEventOk returns a tuple with the LogEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEvent

`func (o *InlineObject233) SetLogEvent(v string)`

SetLogEvent sets LogEvent field to given value.


### GetTimestamp

`func (o *InlineObject233) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineObject233) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineObject233) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTargetOS

`func (o *InlineObject233) GetTargetOS() string`

GetTargetOS returns the TargetOS field if non-nil, zero value otherwise.

### GetTargetOSOk

`func (o *InlineObject233) GetTargetOSOk() (*string, bool)`

GetTargetOSOk returns a tuple with the TargetOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOS

`func (o *InlineObject233) SetTargetOS(v string)`

SetTargetOS sets TargetOS field to given value.

### HasTargetOS

`func (o *InlineObject233) HasTargetOS() bool`

HasTargetOS returns a boolean if a field has been set.

### GetRequest

`func (o *InlineObject233) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineObject233) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineObject233) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineObject233) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


