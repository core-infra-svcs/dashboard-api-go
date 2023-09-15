# InlineResponse2004

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**InlineResponse2011Request**](InlineResponse2011Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Results** | Pointer to [**InlineResponse2004Results**](InlineResponse2004Results.md) |  | [optional] 

## Methods

### NewInlineResponse2004

`func NewInlineResponse2004() *InlineResponse2004`

NewInlineResponse2004 instantiates a new InlineResponse2004 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004WithDefaults

`func NewInlineResponse2004WithDefaults() *InlineResponse2004`

NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingId

`func (o *InlineResponse2004) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *InlineResponse2004) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *InlineResponse2004) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *InlineResponse2004) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse2004) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse2004) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse2004) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse2004) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse2004) GetRequest() InlineResponse2011Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse2004) GetRequestOk() (*InlineResponse2011Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse2004) SetRequest(v InlineResponse2011Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse2004) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse2004) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2004) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2004) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse2004) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse2004) GetResults() InlineResponse2004Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse2004) GetResultsOk() (*InlineResponse2004Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse2004) SetResults(v InlineResponse2004Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse2004) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


