// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_player_queue_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-spotted/internal/services/me_player_queue"
	"github.com/cjavdev/terraform-provider-spotted/internal/test_helpers"
)

func TestMePlayerQueueDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*me_player_queue.MePlayerQueueDataSourceModel)(nil)
	schema := me_player_queue.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
