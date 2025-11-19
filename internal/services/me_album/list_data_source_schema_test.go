// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_album_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/me_album"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestMeAlbumsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_album.MeAlbumsDataSourceModel)(nil)
	schema := me_album.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
