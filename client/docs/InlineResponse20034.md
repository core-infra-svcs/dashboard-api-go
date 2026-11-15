# InlineResponse20034

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CableTestId** | Pointer to **string** | Id of the cable test request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your cable test request. | [optional] 
**Request** | Pointer to [**InlineResponse2012Request**](InlineResponse2012Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the cable test request. | [optional] 
**Results** | Pointer to [**[]InlineResponse20034Results**](InlineResponse20034Results.md) | Results of the cable test request, one for each requested port. | [optional] 
**Error** | Pointer to **string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20034

`func NewInlineResponse20034() *InlineResponse20034`

NewInlineResponse20034 instantiates a new InlineResponse20034 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20034WithDefaults

`func NewInlineResponse20034WithDefaults() *InlineResponse20034`

NewInlineResponse20034WithDefaults instantiates a new InlineResponse20034 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCableTestId

`func (o *InlineResponse20034) GetCableTestId() string`

GetCableTestId returns the CableTestId field if non-nil, zero value otherwise.

### GetCableTestIdOk

`func (o *InlineResponse20034) GetCableTestIdOk() (*string, bool)`

GetCableTestIdOk returns a tuple with the CableTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableTestId

`func (o *InlineResponse20034) SetCableTestId(v string)`

SetCableTestId sets CableTestId field to given value.

### HasCableTestId

`func (o *InlineResponse20034) HasCableTestId() bool`

HasCableTestId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20034) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20034) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20034) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20034) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20034) GetRequest() InlineResponse2012Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20034) GetRequestOk() (*InlineResponse2012Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20034) SetRequest(v InlineResponse2012Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20034) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20034) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20034) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20034) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20034) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse20034) GetResults() []InlineResponse20034Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse20034) GetResultsOk() (*[]InlineResponse20034Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse20034) SetResults(v []InlineResponse20034Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse20034) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20034) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20034) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20034) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20034) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


