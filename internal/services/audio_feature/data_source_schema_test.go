// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_feature_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/spotted-terraform/internal/services/audio_feature"
	"github.com/stainless-sdks/spotted-terraform/internal/test_helpers"
)

func TestAudioFeatureDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audio_feature.AudioFeatureDataSourceModel)(nil)
	schema := audio_feature.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
