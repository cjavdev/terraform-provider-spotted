// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_audiobook_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/me_audiobook"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestMeAudiobooksDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_audiobook.MeAudiobooksDataSourceModel)(nil)
	schema := me_audiobook.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
