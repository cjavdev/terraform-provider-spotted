// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_track_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/playlist_track"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestPlaylistTrackModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*playlist_track.PlaylistTrackModel)(nil)
	schema := playlist_track.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
