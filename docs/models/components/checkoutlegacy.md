# CheckoutLegacy

A checkout session.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ID`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | The ID of the checkout.                                                  |
| `URL`                                                                    | **string*                                                                | :heavy_minus_sign:                                                       | URL the customer should be redirected to complete the purchase.          |
| `CustomerEmail`                                                          | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `CustomerName`                                                           | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `Product`                                                                | [components.CheckoutProduct](../../models/components/checkoutproduct.md) | :heavy_check_mark:                                                       | Product data for a checkout session.                                     |
| `ProductPrice`                                                           | [components.ProductPrice](../../models/components/productprice.md)       | :heavy_check_mark:                                                       | N/A                                                                      |