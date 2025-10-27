// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_feature

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AudioFeatureDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
				Required:    true,
			},
			"acousticness": schema.Float64Attribute{
				Description: "A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0 represents high confidence the track is acoustic.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"analysis_url": schema.StringAttribute{
				Description: "A URL to access the full audio analysis of this track. An access token is required to access this data.",
				Computed:    true,
			},
			"danceability": schema.Float64Attribute{
				Description: "Danceability describes how suitable a track is for dancing based on a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is most danceable.",
				Computed:    true,
			},
			"duration_ms": schema.Int64Attribute{
				Description: "The duration of the track in milliseconds.",
				Computed:    true,
			},
			"energy": schema.Float64Attribute{
				Description: "Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity. Typically, energetic tracks feel fast, loud, and noisy. For example, death metal has high energy, while a Bach prelude scores low on the scale. Perceptual features contributing to this attribute include dynamic range, perceived loudness, timbre, onset rate, and general entropy.",
				Computed:    true,
			},
			"instrumentalness": schema.Float64Attribute{
				Description: `Predicts whether a track contains no vocals. "Ooh" and "aah" sounds are treated as instrumental in this context. Rap or spoken word tracks are clearly "vocal". The closer the instrumentalness value is to 1.0, the greater likelihood the track contains no vocal content. Values above 0.5 are intended to represent instrumental tracks, but confidence is higher as the value approaches 1.0.`,
				Computed:    true,
			},
			"key": schema.Int64Attribute{
				Description: "The key the track is in. Integers map to pitches using standard [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(-1, 11),
				},
			},
			"liveness": schema.Float64Attribute{
				Description: "Detects the presence of an audience in the recording. Higher liveness values represent an increased probability that the track was performed live. A value above 0.8 provides strong likelihood that the track is live.",
				Computed:    true,
			},
			"loudness": schema.Float64Attribute{
				Description: "The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.",
				Computed:    true,
			},
			"mode": schema.Int64Attribute{
				Description: "Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.",
				Computed:    true,
			},
			"speechiness": schema.Float64Attribute{
				Description: "Speechiness detects the presence of spoken words in a track. The more exclusively speech-like the recording (e.g. talk show, audio book, poetry), the closer to 1.0 the attribute value. Values above 0.66 describe tracks that are probably made entirely of spoken words. Values between 0.33 and 0.66 describe tracks that may contain both music and speech, either in sections or layered, including such cases as rap music. Values below 0.33 most likely represent music and other non-speech-like tracks.",
				Computed:    true,
			},
			"tempo": schema.Float64Attribute{
				Description: "The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.",
				Computed:    true,
			},
			"time_signature": schema.Int64Attribute{
				Description: `An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".`,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(3, 7),
				},
			},
			"track_href": schema.StringAttribute{
				Description: "A link to the Web API endpoint providing full details of the track.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The object type.\nAvailable values: \"audio_features\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("audio_features"),
				},
			},
			"uri": schema.StringAttribute{
				Description: "The Spotify URI for the track.",
				Computed:    true,
			},
			"valence": schema.Float64Attribute{
				Description: "A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track. Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric), while tracks with low valence sound more negative (e.g. sad, depressed, angry).",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
		},
	}
}

func (d *AudioFeatureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AudioFeatureDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
