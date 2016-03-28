#Downloading a file

##Preferred method: using **alt=media**

To download files, you make an authorized HTTP `GET` request to the file's [resource URL]() and include the query parameter `alt=media`. For example:


```sh
GET https://www.googleapis.com/drive/v2/files/0B9jNhSvVjoIVM3dKcGRKRmVIOVU?alt=media
Authorization: Bearer ya29.AHESVbXTUv5mHMo3RYfmS1YJonjzzdTOFZwvyOAUVhrs
```


```javascript
// might not be applicable
//Downloading the file requires the user to have at least read access. Additionally, you app must be authorized with a scope that allows reading of file content.
```
Compared to downloading with `downloadUrl` (describe below), `alt=media` offers the advantage of predictable, long-lived, client-contructable download URLs.


##Alternate method: using **downloadURL**

To download file this way, you make two authorized HTTP `GET` requests:

1. First, retrieve the appropriate download URL provided in the file's metadata.
2. Then, retrieve the actual file content (or link) using the download URL.

Download a file by retrieving its download URL and the retrieving the file itself. To find the download URL, fetch a file's metadata with [files.get]() or the [files.list]() methods. The response contains the link in the downloadURL field. For example:

```json
{
 "kind": "drive#file",
 "id": "0B0iTh2ZTRkNjI1ZjUtOGIwZS00YzhiLWJmMWMtNWI4YzY3NzQyNWQ2",
 "downloadUrl": "https://doc-04-c1-docs.googleusercontent.com/docs/securesc/ivearmirmg66&e=download&gd=true",
 ...
}
```
Once you have fetched the URL, you can use it to download the file content by making an authorized GET request to it. For example:

```sh
GET https://doc-04-c1-docs.googleusercontent.com/docs/securesc/ivearmirmg66&e=download&gd=true
Authorization: Bearer ya29.AHESVbXTUv5mHMo3RYfmS1YJonjzzdTOFZwvyOAUVhrs
```

The `downloadURL` field is a short-lived value and is typically only good for x hours (configurable). For this reason, it's usually necessary to issue both requests around the same time. The `downloadURL` is not predictable by the client and must be fetched from the server metadata.

**Note: Personally, this is much more secure and suitable for AMbank's case as the URL shall be randomly generated.**


##Partial Download

Partial download involves downloading only a specified portion of a file. You can specify the portion of the file you want to download by using a byte range with the Range header. For example:

```sh
Range: bytes=500-999
```
