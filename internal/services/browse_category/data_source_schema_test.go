// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browse_category_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/browse_category"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestBrowseCategoryDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*browse_category.BrowseCategoryDataSourceModel)(nil)
	schema := browse_category.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
