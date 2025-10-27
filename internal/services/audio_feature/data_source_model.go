// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_feature

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AudioFeatureDataSourceModel struct {
	ID               types.String  `tfsdk:"id" path:"id,required"`
	Acousticness     types.Float64 `tfsdk:"acousticness" json:"acousticness,computed"`
	AnalysisURL      types.String  `tfsdk:"analysis_url" json:"analysis_url,computed"`
	Danceability     types.Float64 `tfsdk:"danceability" json:"danceability,computed"`
	DurationMs       types.Int64   `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Energy           types.Float64 `tfsdk:"energy" json:"energy,computed"`
	Instrumentalness types.Float64 `tfsdk:"instrumentalness" json:"instrumentalness,computed"`
	Key              types.Int64   `tfsdk:"key" json:"key,computed"`
	Liveness         types.Float64 `tfsdk:"liveness" json:"liveness,computed"`
	Loudness         types.Float64 `tfsdk:"loudness" json:"loudness,computed"`
	Mode             types.Int64   `tfsdk:"mode" json:"mode,computed"`
	Speechiness      types.Float64 `tfsdk:"speechiness" json:"speechiness,computed"`
	Tempo            types.Float64 `tfsdk:"tempo" json:"tempo,computed"`
	TimeSignature    types.Int64   `tfsdk:"time_signature" json:"time_signature,computed"`
	TrackHref        types.String  `tfsdk:"track_href" json:"track_href,computed"`
	Type             types.String  `tfsdk:"type" json:"type,computed"`
	Uri              types.String  `tfsdk:"uri" json:"uri,computed"`
	Valence          types.Float64 `tfsdk:"valence" json:"valence,computed"`
}
