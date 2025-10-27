// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package search_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/search"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestSearchDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*search.SearchDataSourceModel)(nil)
	schema := search.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
