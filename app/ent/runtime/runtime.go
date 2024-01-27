// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/endot1231/ec-backend/ent/admins"
	"github.com/endot1231/ec-backend/ent/productbrands"
	"github.com/endot1231/ec-backend/ent/productcategories"
	"github.com/endot1231/ec-backend/ent/productcolors"
	"github.com/endot1231/ec-backend/ent/products"
	"github.com/endot1231/ec-backend/ent/productsizes"
	"github.com/endot1231/ec-backend/ent/productstock"
	"github.com/endot1231/ec-backend/ent/reviews"
	"github.com/endot1231/ec-backend/ent/schema"
	"github.com/endot1231/ec-backend/ent/shops"
	"github.com/endot1231/ec-backend/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminsFields := schema.Admins{}.Fields()
	_ = adminsFields
	// adminsDescEmailVerified is the schema descriptor for email_verified field.
	adminsDescEmailVerified := adminsFields[2].Descriptor()
	// admins.DefaultEmailVerified holds the default value on creation for the email_verified field.
	admins.DefaultEmailVerified = adminsDescEmailVerified.Default.(time.Time)
	// adminsDescCreatedAt is the schema descriptor for created_at field.
	adminsDescCreatedAt := adminsFields[5].Descriptor()
	// admins.DefaultCreatedAt holds the default value on creation for the created_at field.
	admins.DefaultCreatedAt = adminsDescCreatedAt.Default.(time.Time)
	// adminsDescUpdatedAt is the schema descriptor for updated_at field.
	adminsDescUpdatedAt := adminsFields[6].Descriptor()
	// admins.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admins.DefaultUpdatedAt = adminsDescUpdatedAt.Default.(time.Time)
	// adminsDescDeletedAt is the schema descriptor for deleted_at field.
	adminsDescDeletedAt := adminsFields[7].Descriptor()
	// admins.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	admins.DefaultDeletedAt = adminsDescDeletedAt.Default.(time.Time)
	productbrandsFields := schema.ProductBrands{}.Fields()
	_ = productbrandsFields
	// productbrandsDescCreatedAt is the schema descriptor for created_at field.
	productbrandsDescCreatedAt := productbrandsFields[1].Descriptor()
	// productbrands.DefaultCreatedAt holds the default value on creation for the created_at field.
	productbrands.DefaultCreatedAt = productbrandsDescCreatedAt.Default.(time.Time)
	// productbrandsDescUpdatedAt is the schema descriptor for updated_at field.
	productbrandsDescUpdatedAt := productbrandsFields[2].Descriptor()
	// productbrands.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productbrands.DefaultUpdatedAt = productbrandsDescUpdatedAt.Default.(time.Time)
	// productbrandsDescDeletedAt is the schema descriptor for deleted_at field.
	productbrandsDescDeletedAt := productbrandsFields[3].Descriptor()
	// productbrands.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	productbrands.DefaultDeletedAt = productbrandsDescDeletedAt.Default.(time.Time)
	productcategoriesFields := schema.ProductCategories{}.Fields()
	_ = productcategoriesFields
	// productcategoriesDescCreatedAt is the schema descriptor for created_at field.
	productcategoriesDescCreatedAt := productcategoriesFields[1].Descriptor()
	// productcategories.DefaultCreatedAt holds the default value on creation for the created_at field.
	productcategories.DefaultCreatedAt = productcategoriesDescCreatedAt.Default.(time.Time)
	// productcategoriesDescUpdatedAt is the schema descriptor for updated_at field.
	productcategoriesDescUpdatedAt := productcategoriesFields[2].Descriptor()
	// productcategories.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productcategories.DefaultUpdatedAt = productcategoriesDescUpdatedAt.Default.(time.Time)
	// productcategoriesDescDeletedAt is the schema descriptor for deleted_at field.
	productcategoriesDescDeletedAt := productcategoriesFields[3].Descriptor()
	// productcategories.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	productcategories.DefaultDeletedAt = productcategoriesDescDeletedAt.Default.(time.Time)
	productcolorsFields := schema.ProductColors{}.Fields()
	_ = productcolorsFields
	// productcolorsDescCreatedAt is the schema descriptor for created_at field.
	productcolorsDescCreatedAt := productcolorsFields[1].Descriptor()
	// productcolors.DefaultCreatedAt holds the default value on creation for the created_at field.
	productcolors.DefaultCreatedAt = productcolorsDescCreatedAt.Default.(time.Time)
	// productcolorsDescUpdatedAt is the schema descriptor for updated_at field.
	productcolorsDescUpdatedAt := productcolorsFields[2].Descriptor()
	// productcolors.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productcolors.DefaultUpdatedAt = productcolorsDescUpdatedAt.Default.(time.Time)
	// productcolorsDescDeletedAt is the schema descriptor for deleted_at field.
	productcolorsDescDeletedAt := productcolorsFields[3].Descriptor()
	// productcolors.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	productcolors.DefaultDeletedAt = productcolorsDescDeletedAt.Default.(time.Time)
	productsizesFields := schema.ProductSizes{}.Fields()
	_ = productsizesFields
	// productsizesDescCreatedAt is the schema descriptor for created_at field.
	productsizesDescCreatedAt := productsizesFields[2].Descriptor()
	// productsizes.DefaultCreatedAt holds the default value on creation for the created_at field.
	productsizes.DefaultCreatedAt = productsizesDescCreatedAt.Default.(time.Time)
	// productsizesDescUpdatedAt is the schema descriptor for updated_at field.
	productsizesDescUpdatedAt := productsizesFields[3].Descriptor()
	// productsizes.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productsizes.DefaultUpdatedAt = productsizesDescUpdatedAt.Default.(time.Time)
	// productsizesDescDeletedAt is the schema descriptor for deleted_at field.
	productsizesDescDeletedAt := productsizesFields[4].Descriptor()
	// productsizes.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	productsizes.DefaultDeletedAt = productsizesDescDeletedAt.Default.(time.Time)
	productstockFields := schema.ProductStock{}.Fields()
	_ = productstockFields
	// productstockDescPrice is the schema descriptor for price field.
	productstockDescPrice := productstockFields[1].Descriptor()
	// productstock.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	productstock.PriceValidator = productstockDescPrice.Validators[0].(func(int) error)
	// productstockDescStockQuantity is the schema descriptor for stock_quantity field.
	productstockDescStockQuantity := productstockFields[2].Descriptor()
	// productstock.StockQuantityValidator is a validator for the "stock_quantity" field. It is called by the builders before save.
	productstock.StockQuantityValidator = productstockDescStockQuantity.Validators[0].(func(int) error)
	// productstockDescCreatedAt is the schema descriptor for created_at field.
	productstockDescCreatedAt := productstockFields[3].Descriptor()
	// productstock.DefaultCreatedAt holds the default value on creation for the created_at field.
	productstock.DefaultCreatedAt = productstockDescCreatedAt.Default.(time.Time)
	// productstockDescUpdatedAt is the schema descriptor for updated_at field.
	productstockDescUpdatedAt := productstockFields[4].Descriptor()
	// productstock.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productstock.DefaultUpdatedAt = productstockDescUpdatedAt.Default.(time.Time)
	productsFields := schema.Products{}.Fields()
	_ = productsFields
	// productsDescCreatedAt is the schema descriptor for created_at field.
	productsDescCreatedAt := productsFields[2].Descriptor()
	// products.DefaultCreatedAt holds the default value on creation for the created_at field.
	products.DefaultCreatedAt = productsDescCreatedAt.Default.(time.Time)
	// productsDescUpdatedAt is the schema descriptor for updated_at field.
	productsDescUpdatedAt := productsFields[3].Descriptor()
	// products.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	products.DefaultUpdatedAt = productsDescUpdatedAt.Default.(time.Time)
	// productsDescDeletedAt is the schema descriptor for deleted_at field.
	productsDescDeletedAt := productsFields[4].Descriptor()
	// products.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	products.DefaultDeletedAt = productsDescDeletedAt.Default.(time.Time)
	reviewsFields := schema.Reviews{}.Fields()
	_ = reviewsFields
	// reviewsDescCreatedAt is the schema descriptor for created_at field.
	reviewsDescCreatedAt := reviewsFields[2].Descriptor()
	// reviews.DefaultCreatedAt holds the default value on creation for the created_at field.
	reviews.DefaultCreatedAt = reviewsDescCreatedAt.Default.(time.Time)
	// reviewsDescUpdatedAt is the schema descriptor for updated_at field.
	reviewsDescUpdatedAt := reviewsFields[3].Descriptor()
	// reviews.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reviews.DefaultUpdatedAt = reviewsDescUpdatedAt.Default.(time.Time)
	// reviewsDescDeletedAt is the schema descriptor for deleted_at field.
	reviewsDescDeletedAt := reviewsFields[4].Descriptor()
	// reviews.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	reviews.DefaultDeletedAt = reviewsDescDeletedAt.Default.(time.Time)
	shopsFields := schema.Shops{}.Fields()
	_ = shopsFields
	// shopsDescCreatedAt is the schema descriptor for created_at field.
	shopsDescCreatedAt := shopsFields[2].Descriptor()
	// shops.DefaultCreatedAt holds the default value on creation for the created_at field.
	shops.DefaultCreatedAt = shopsDescCreatedAt.Default.(time.Time)
	// shopsDescUpdatedAt is the schema descriptor for updated_at field.
	shopsDescUpdatedAt := shopsFields[3].Descriptor()
	// shops.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	shops.DefaultUpdatedAt = shopsDescUpdatedAt.Default.(time.Time)
	// shopsDescDeletedAt is the schema descriptor for deleted_at field.
	shopsDescDeletedAt := shopsFields[4].Descriptor()
	// shops.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	shops.DefaultDeletedAt = shopsDescDeletedAt.Default.(time.Time)
	usersHooks := schema.Users{}.Hooks()
	users.Hooks[0] = usersHooks[0]
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescName is the schema descriptor for name field.
	usersDescName := usersFields[0].Descriptor()
	// users.NameValidator is a validator for the "name" field. It is called by the builders before save.
	users.NameValidator = usersDescName.Validators[0].(func(string) error)
	// usersDescEmail is the schema descriptor for email field.
	usersDescEmail := usersFields[1].Descriptor()
	// users.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	users.EmailValidator = usersDescEmail.Validators[0].(func(string) error)
	// usersDescEmailVerified is the schema descriptor for email_verified field.
	usersDescEmailVerified := usersFields[2].Descriptor()
	// users.DefaultEmailVerified holds the default value on creation for the email_verified field.
	users.DefaultEmailVerified = usersDescEmailVerified.Default.(time.Time)
	// usersDescCreatedAt is the schema descriptor for created_at field.
	usersDescCreatedAt := usersFields[5].Descriptor()
	// users.DefaultCreatedAt holds the default value on creation for the created_at field.
	users.DefaultCreatedAt = usersDescCreatedAt.Default.(time.Time)
	// usersDescUpdatedAt is the schema descriptor for updated_at field.
	usersDescUpdatedAt := usersFields[6].Descriptor()
	// users.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	users.DefaultUpdatedAt = usersDescUpdatedAt.Default.(time.Time)
	// usersDescDeletedAt is the schema descriptor for deleted_at field.
	usersDescDeletedAt := usersFields[7].Descriptor()
	// users.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	users.DefaultDeletedAt = usersDescDeletedAt.Default.(time.Time)
}

const (
	Version = "v0.12.3"                                         // Version of ent codegen.
	Sum     = "h1:N5lO2EOrHpCH5HYfiMOCHYbo+oh5M8GjT0/cx5x6xkk=" // Sum of ent codegen.
)
