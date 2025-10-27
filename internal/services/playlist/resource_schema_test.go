// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/playlist"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestPlaylistModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*playlist.PlaylistModel)(nil)
	schema := playlist.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
