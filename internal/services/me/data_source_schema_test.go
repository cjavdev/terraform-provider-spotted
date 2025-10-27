// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/me"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestMeDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me.MeDataSourceModel)(nil)
	schema := me.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
