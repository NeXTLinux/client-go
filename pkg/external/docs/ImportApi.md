# \ImportApi

All URIs are relative to _http://localhost_

| Method                                                    | HTTP request            | Description                                                                                                                 |
| --------------------------------------------------------- | ----------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| [**ImportImageArchive**](ImportApi.md#ImportImageArchive) | **Post** /import/images | Import an nextlinux image tar.gz archive file. This is a deprecated API replaced by the \&quot;/imports/images\&quot; route |

## ImportImageArchive

> []NextlinuxImage ImportImageArchive(ctx, archiveFile)

Import an nextlinux image tar.gz archive file. This is a deprecated API replaced by the \"/imports/images\" route

### Required Parameters

| Name            | Type                           | Description                                                                 | Notes |
| --------------- | ------------------------------ | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context**            | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveFile** | **\*os.File\*\*\***os.File\*\* | nextlinux image tar archive.                                                |

### Return type

[**[]NextlinuxImage**](NextlinuxImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
