# InlineResponse20043

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | ID of the route details request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this URL to check the status of your route details request. | [optional] 
**Request** | Pointer to [**InlineResponse20111Request**](InlineResponse20111Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the route details request. | [optional] 
**Results** | Pointer to [**InlineResponse20043Results**](InlineResponse20043Results.md) |  | [optional] 
**Errors** | Pointer to **[]string** | Any errors encountered while retrieving route details. | [optional] 

## Methods

### NewInlineResponse20043

`func NewInlineResponse20043() *InlineResponse20043`

NewInlineResponse20043 instantiates a new InlineResponse20043 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043WithDefaults

`func NewInlineResponse20043WithDefaults() *InlineResponse20043`

NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *InlineResponse20043) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse20043) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse20043) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse20043) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20043) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20043) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20043) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20043) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20043) GetRequest() InlineResponse20111Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20043) GetRequestOk() (*InlineResponse20111Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20043) SetRequest(v InlineResponse20111Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20043) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20043) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20043) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20043) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20043) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse20043) GetResults() InlineResponse20043Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse20043) GetResultsOk() (*InlineResponse20043Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse20043) SetResults(v InlineResponse20043Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse20043) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20043) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20043) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20043) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20043) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


