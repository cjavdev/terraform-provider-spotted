// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browse_category_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/browse_category"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestBrowseCategoryDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*browse_category.BrowseCategoryDataSourceModel)(nil)
	schema := browse_category.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
