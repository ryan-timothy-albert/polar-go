# ProductPriceTypeFilter

Filter by product price type. `recurring` will return orders corresponding to subscriptions creations or renewals. `one_time` will return orders corresponding to one-time purchases.


## Supported Types

### ProductPriceType

```go
productPriceTypeFilter := operations.CreateProductPriceTypeFilterProductPriceType(components.ProductPriceType{/* values here */})
```

### 

```go
productPriceTypeFilter := operations.CreateProductPriceTypeFilterArrayOfProductPriceType([]components.ProductPriceType{/* values here */})
```

