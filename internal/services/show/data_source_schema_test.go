// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package show_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/show"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestShowDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*show.ShowDataSourceModel)(nil)
	schema := show.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
