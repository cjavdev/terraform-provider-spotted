// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_track_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/me_track"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestMeTracksDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_track.MeTracksDataSourceModel)(nil)
	schema := me_track.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
