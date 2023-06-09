# \EventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEvent**](EventsApi.md#DeleteEvent) | **Delete** /events/{eventId} | Delete Event
[**DeleteEvents**](EventsApi.md#DeleteEvents) | **Delete** /events | Delete Events
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /events/{eventId} | Get Event
[**ListEventTypes**](EventsApi.md#ListEventTypes) | **Get** /event_types | List Event Types
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /events | List Events



## DeleteEvent

> DeleteEvent(ctx, eventId, optional)

Delete Event

Delete an event by its event ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| Event ID of the event to be deleted | 
 **optional** | ***DeleteEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNextlinuxAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEvents

> []string DeleteEvents(ctx, optional)

Delete Events

Delete all or a subset of events filtered using the optional query parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **optional.String**| Delete events that occurred before the timestamp | 
 **since** | **optional.String**| Delete events that occurred after the timestamp | 
 **level** | **optional.String**| Delete events that match the level - INFO or ERROR | 
 **xNextlinuxAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> EventResponse GetEvent(ctx, eventId, optional)

Get Event

Lookup an event by its event ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| Event ID of the event for lookup | 
 **optional** | ***GetEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNextlinuxAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventTypes

> []EventCategory ListEventTypes(ctx, )

List Event Types

Returns list of event types in the category hierarchy

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EventCategory**](EventCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> EventsList ListEvents(ctx, optional)

List Events

Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceServicename** | **optional.String**| Filter events by the originating service | 
 **sourceHostid** | **optional.String**| Filter events by the originating host ID | 
 **eventType** | **optional.String**| Filter events by a prefix match on the event type (e.g. \&quot;user.image.\&quot;) | 
 **resourceType** | **optional.String**| Filter events by the type of resource - tag, imageDigest, repository etc | 
 **resourceId** | **optional.String**| Filter events by the id of the resource | 
 **level** | **optional.String**| Filter events by the level - INFO or ERROR | 
 **since** | **optional.String**| Return events that occurred after the timestamp | 
 **before** | **optional.String**| Return events that occurred before the timestamp | 
 **page** | **optional.Int32**| Pagination controls - return the nth page of results. Defaults to first page if left empty | [default to 1]
 **limit** | **optional.Int32**| Number of events in the result set. Defaults to 100 if left empty | [default to 100]
 **xNextlinuxAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EventsList**](EventsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

