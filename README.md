# mcre
Panton Multi-Content Repository Service (Panton MCRE), provides developers and IT teams with secure, durable, highly-scalable cloud storage.


# Restful-API

## 1. Files: get

Gets a file's metadata or content by ID.

### Request

##### HTTP request
```sh
GET https://api.panton.com/mcre/v1/files/:fileId
```
##### Parameters

Parameter name | Value | Description
--- | --- | --- |
fileId | string | The ID of the file.


##### Request Body
Do not supply a request body with this method

### Response
If successful, this method returns a [Files resource]() in the response body.
