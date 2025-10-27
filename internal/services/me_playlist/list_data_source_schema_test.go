// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_playlist_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/me_playlist"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestMePlaylistsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_playlist.MePlaylistsDataSourceModel)(nil)
	schema := me_playlist.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
