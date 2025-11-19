// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_image_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/playlist_image"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestPlaylistImageModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*playlist_image.PlaylistImageModel)(nil)
	schema := playlist_image.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
