// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_episode_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/me_episode"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestMeEpisodesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_episode.MeEpisodesDataSourceModel)(nil)
	schema := me_episode.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
