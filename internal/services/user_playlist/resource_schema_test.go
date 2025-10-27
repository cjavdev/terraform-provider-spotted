// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/user_playlist"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestUserPlaylistModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user_playlist.UserPlaylistModel)(nil)
	schema := user_playlist.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
