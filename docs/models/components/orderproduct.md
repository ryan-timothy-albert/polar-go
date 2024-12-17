# OrderProduct


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `CreatedAt`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Creation timestamp of the object.                        |
| `ModifiedAt`                                             | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Last modification timestamp of the object.               |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | The ID of the product.                                   |
| `Name`                                                   | *string*                                                 | :heavy_check_mark:                                       | The name of the product.                                 |
| `Description`                                            | *string*                                                 | :heavy_check_mark:                                       | The description of the product.                          |
| `IsRecurring`                                            | *bool*                                                   | :heavy_check_mark:                                       | Whether the product is a subscription tier.              |
| `IsArchived`                                             | *bool*                                                   | :heavy_check_mark:                                       | Whether the product is archived and no longer available. |
| `OrganizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | The ID of the organization owning the product.           |