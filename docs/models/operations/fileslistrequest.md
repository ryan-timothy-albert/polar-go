# FilesListRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `OrganizationID`                                | **string*                                       | :heavy_minus_sign:                              | N/A                                             |
| `Ids`                                           | []*string*                                      | :heavy_minus_sign:                              | List of file IDs to get.                        |
| `Page`                                          | **int64*                                        | :heavy_minus_sign:                              | Page number, defaults to 1.                     |
| `Limit`                                         | **int64*                                        | :heavy_minus_sign:                              | Size of a page, defaults to 10. Maximum is 100. |