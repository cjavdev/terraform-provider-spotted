// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_show_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/me_show"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestMeShowsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_show.MeShowsDataSourceModel)(nil)
	schema := me_show.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
