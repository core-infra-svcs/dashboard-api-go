# InlineResponse20026

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingId** | Pointer to **string** | Id to check the status of your ping request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your ping request. | [optional] 
**Request** | Pointer to [**InlineResponse2015Request**](InlineResponse2015Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the ping request. | [optional] 
**Results** | Pointer to [**InlineResponse20026Results**](InlineResponse20026Results.md) |  | [optional] 

## Methods

### NewInlineResponse20026

`func NewInlineResponse20026() *InlineResponse20026`

NewInlineResponse20026 instantiates a new InlineResponse20026 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026WithDefaults

`func NewInlineResponse20026WithDefaults() *InlineResponse20026`

NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingId

`func (o *InlineResponse20026) GetPingId() string`

GetPingId returns the PingId field if non-nil, zero value otherwise.

### GetPingIdOk

`func (o *InlineResponse20026) GetPingIdOk() (*string, bool)`

GetPingIdOk returns a tuple with the PingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingId

`func (o *InlineResponse20026) SetPingId(v string)`

SetPingId sets PingId field to given value.

### HasPingId

`func (o *InlineResponse20026) HasPingId() bool`

HasPingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20026) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20026) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20026) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20026) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20026) GetRequest() InlineResponse2015Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20026) GetRequestOk() (*InlineResponse2015Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20026) SetRequest(v InlineResponse2015Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20026) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20026) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20026) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20026) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20026) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse20026) GetResults() InlineResponse20026Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse20026) GetResultsOk() (*InlineResponse20026Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse20026) SetResults(v InlineResponse20026Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse20026) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


