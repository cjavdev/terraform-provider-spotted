// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_track_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/playlist_track"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestPlaylistTracksDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*playlist_track.PlaylistTracksDataSourceModel)(nil)
	schema := playlist_track.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
