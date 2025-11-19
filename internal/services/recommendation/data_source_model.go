// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package recommendation

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RecommendationDataSourceModel struct {
	Market                 types.String                                                      `tfsdk:"market" query:"market,optional"`
	MaxAcousticness        types.Float64                                                     `tfsdk:"max_acousticness" query:"max_acousticness,optional"`
	MaxDanceability        types.Float64                                                     `tfsdk:"max_danceability" query:"max_danceability,optional"`
	MaxDurationMs          types.Int64                                                       `tfsdk:"max_duration_ms" query:"max_duration_ms,optional"`
	MaxEnergy              types.Float64                                                     `tfsdk:"max_energy" query:"max_energy,optional"`
	MaxInstrumentalness    types.Float64                                                     `tfsdk:"max_instrumentalness" query:"max_instrumentalness,optional"`
	MaxKey                 types.Int64                                                       `tfsdk:"max_key" query:"max_key,optional"`
	MaxLiveness            types.Float64                                                     `tfsdk:"max_liveness" query:"max_liveness,optional"`
	MaxLoudness            types.Float64                                                     `tfsdk:"max_loudness" query:"max_loudness,optional"`
	MaxMode                types.Int64                                                       `tfsdk:"max_mode" query:"max_mode,optional"`
	MaxPopularity          types.Int64                                                       `tfsdk:"max_popularity" query:"max_popularity,optional"`
	MaxSpeechiness         types.Float64                                                     `tfsdk:"max_speechiness" query:"max_speechiness,optional"`
	MaxTempo               types.Float64                                                     `tfsdk:"max_tempo" query:"max_tempo,optional"`
	MaxTimeSignature       types.Int64                                                       `tfsdk:"max_time_signature" query:"max_time_signature,optional"`
	MaxValence             types.Float64                                                     `tfsdk:"max_valence" query:"max_valence,optional"`
	MinAcousticness        types.Float64                                                     `tfsdk:"min_acousticness" query:"min_acousticness,optional"`
	MinDanceability        types.Float64                                                     `tfsdk:"min_danceability" query:"min_danceability,optional"`
	MinDurationMs          types.Int64                                                       `tfsdk:"min_duration_ms" query:"min_duration_ms,optional"`
	MinEnergy              types.Float64                                                     `tfsdk:"min_energy" query:"min_energy,optional"`
	MinInstrumentalness    types.Float64                                                     `tfsdk:"min_instrumentalness" query:"min_instrumentalness,optional"`
	MinKey                 types.Int64                                                       `tfsdk:"min_key" query:"min_key,optional"`
	MinLiveness            types.Float64                                                     `tfsdk:"min_liveness" query:"min_liveness,optional"`
	MinLoudness            types.Float64                                                     `tfsdk:"min_loudness" query:"min_loudness,optional"`
	MinMode                types.Int64                                                       `tfsdk:"min_mode" query:"min_mode,optional"`
	MinPopularity          types.Int64                                                       `tfsdk:"min_popularity" query:"min_popularity,optional"`
	MinSpeechiness         types.Float64                                                     `tfsdk:"min_speechiness" query:"min_speechiness,optional"`
	MinTempo               types.Float64                                                     `tfsdk:"min_tempo" query:"min_tempo,optional"`
	MinTimeSignature       types.Int64                                                       `tfsdk:"min_time_signature" query:"min_time_signature,optional"`
	MinValence             types.Float64                                                     `tfsdk:"min_valence" query:"min_valence,optional"`
	SeedArtists            types.String                                                      `tfsdk:"seed_artists" query:"seed_artists,optional"`
	SeedGenres             types.String                                                      `tfsdk:"seed_genres" query:"seed_genres,optional"`
	SeedTracks             types.String                                                      `tfsdk:"seed_tracks" query:"seed_tracks,optional"`
	TargetAcousticness     types.Float64                                                     `tfsdk:"target_acousticness" query:"target_acousticness,optional"`
	TargetDanceability     types.Float64                                                     `tfsdk:"target_danceability" query:"target_danceability,optional"`
	TargetDurationMs       types.Int64                                                       `tfsdk:"target_duration_ms" query:"target_duration_ms,optional"`
	TargetEnergy           types.Float64                                                     `tfsdk:"target_energy" query:"target_energy,optional"`
	TargetInstrumentalness types.Float64                                                     `tfsdk:"target_instrumentalness" query:"target_instrumentalness,optional"`
	TargetKey              types.Int64                                                       `tfsdk:"target_key" query:"target_key,optional"`
	TargetLiveness         types.Float64                                                     `tfsdk:"target_liveness" query:"target_liveness,optional"`
	TargetLoudness         types.Float64                                                     `tfsdk:"target_loudness" query:"target_loudness,optional"`
	TargetMode             types.Int64                                                       `tfsdk:"target_mode" query:"target_mode,optional"`
	TargetPopularity       types.Int64                                                       `tfsdk:"target_popularity" query:"target_popularity,optional"`
	TargetSpeechiness      types.Float64                                                     `tfsdk:"target_speechiness" query:"target_speechiness,optional"`
	TargetTempo            types.Float64                                                     `tfsdk:"target_tempo" query:"target_tempo,optional"`
	TargetTimeSignature    types.Int64                                                       `tfsdk:"target_time_signature" query:"target_time_signature,optional"`
	TargetValence          types.Float64                                                     `tfsdk:"target_valence" query:"target_valence,optional"`
	Limit                  types.Int64                                                       `tfsdk:"limit" query:"limit,computed_optional"`
	Seeds                  customfield.NestedObjectList[RecommendationSeedsDataSourceModel]  `tfsdk:"seeds" json:"seeds,computed"`
	Tracks                 customfield.NestedObjectList[RecommendationTracksDataSourceModel] `tfsdk:"tracks" json:"tracks,computed"`
}

func (m *RecommendationDataSourceModel) toReadParams(_ context.Context) (params spotted.RecommendationGetParams, diags diag.Diagnostics) {
	params = spotted.RecommendationGetParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}
	if !m.MaxAcousticness.IsNull() {
		params.MaxAcousticness = param.NewOpt(m.MaxAcousticness.ValueFloat64())
	}
	if !m.MaxDanceability.IsNull() {
		params.MaxDanceability = param.NewOpt(m.MaxDanceability.ValueFloat64())
	}
	if !m.MaxDurationMs.IsNull() {
		params.MaxDurationMs = param.NewOpt(m.MaxDurationMs.ValueInt64())
	}
	if !m.MaxEnergy.IsNull() {
		params.MaxEnergy = param.NewOpt(m.MaxEnergy.ValueFloat64())
	}
	if !m.MaxInstrumentalness.IsNull() {
		params.MaxInstrumentalness = param.NewOpt(m.MaxInstrumentalness.ValueFloat64())
	}
	if !m.MaxKey.IsNull() {
		params.MaxKey = param.NewOpt(m.MaxKey.ValueInt64())
	}
	if !m.MaxLiveness.IsNull() {
		params.MaxLiveness = param.NewOpt(m.MaxLiveness.ValueFloat64())
	}
	if !m.MaxLoudness.IsNull() {
		params.MaxLoudness = param.NewOpt(m.MaxLoudness.ValueFloat64())
	}
	if !m.MaxMode.IsNull() {
		params.MaxMode = param.NewOpt(m.MaxMode.ValueInt64())
	}
	if !m.MaxPopularity.IsNull() {
		params.MaxPopularity = param.NewOpt(m.MaxPopularity.ValueInt64())
	}
	if !m.MaxSpeechiness.IsNull() {
		params.MaxSpeechiness = param.NewOpt(m.MaxSpeechiness.ValueFloat64())
	}
	if !m.MaxTempo.IsNull() {
		params.MaxTempo = param.NewOpt(m.MaxTempo.ValueFloat64())
	}
	if !m.MaxTimeSignature.IsNull() {
		params.MaxTimeSignature = param.NewOpt(m.MaxTimeSignature.ValueInt64())
	}
	if !m.MaxValence.IsNull() {
		params.MaxValence = param.NewOpt(m.MaxValence.ValueFloat64())
	}
	if !m.MinAcousticness.IsNull() {
		params.MinAcousticness = param.NewOpt(m.MinAcousticness.ValueFloat64())
	}
	if !m.MinDanceability.IsNull() {
		params.MinDanceability = param.NewOpt(m.MinDanceability.ValueFloat64())
	}
	if !m.MinDurationMs.IsNull() {
		params.MinDurationMs = param.NewOpt(m.MinDurationMs.ValueInt64())
	}
	if !m.MinEnergy.IsNull() {
		params.MinEnergy = param.NewOpt(m.MinEnergy.ValueFloat64())
	}
	if !m.MinInstrumentalness.IsNull() {
		params.MinInstrumentalness = param.NewOpt(m.MinInstrumentalness.ValueFloat64())
	}
	if !m.MinKey.IsNull() {
		params.MinKey = param.NewOpt(m.MinKey.ValueInt64())
	}
	if !m.MinLiveness.IsNull() {
		params.MinLiveness = param.NewOpt(m.MinLiveness.ValueFloat64())
	}
	if !m.MinLoudness.IsNull() {
		params.MinLoudness = param.NewOpt(m.MinLoudness.ValueFloat64())
	}
	if !m.MinMode.IsNull() {
		params.MinMode = param.NewOpt(m.MinMode.ValueInt64())
	}
	if !m.MinPopularity.IsNull() {
		params.MinPopularity = param.NewOpt(m.MinPopularity.ValueInt64())
	}
	if !m.MinSpeechiness.IsNull() {
		params.MinSpeechiness = param.NewOpt(m.MinSpeechiness.ValueFloat64())
	}
	if !m.MinTempo.IsNull() {
		params.MinTempo = param.NewOpt(m.MinTempo.ValueFloat64())
	}
	if !m.MinTimeSignature.IsNull() {
		params.MinTimeSignature = param.NewOpt(m.MinTimeSignature.ValueInt64())
	}
	if !m.MinValence.IsNull() {
		params.MinValence = param.NewOpt(m.MinValence.ValueFloat64())
	}
	if !m.SeedArtists.IsNull() {
		params.SeedArtists = param.NewOpt(m.SeedArtists.ValueString())
	}
	if !m.SeedGenres.IsNull() {
		params.SeedGenres = param.NewOpt(m.SeedGenres.ValueString())
	}
	if !m.SeedTracks.IsNull() {
		params.SeedTracks = param.NewOpt(m.SeedTracks.ValueString())
	}
	if !m.TargetAcousticness.IsNull() {
		params.TargetAcousticness = param.NewOpt(m.TargetAcousticness.ValueFloat64())
	}
	if !m.TargetDanceability.IsNull() {
		params.TargetDanceability = param.NewOpt(m.TargetDanceability.ValueFloat64())
	}
	if !m.TargetDurationMs.IsNull() {
		params.TargetDurationMs = param.NewOpt(m.TargetDurationMs.ValueInt64())
	}
	if !m.TargetEnergy.IsNull() {
		params.TargetEnergy = param.NewOpt(m.TargetEnergy.ValueFloat64())
	}
	if !m.TargetInstrumentalness.IsNull() {
		params.TargetInstrumentalness = param.NewOpt(m.TargetInstrumentalness.ValueFloat64())
	}
	if !m.TargetKey.IsNull() {
		params.TargetKey = param.NewOpt(m.TargetKey.ValueInt64())
	}
	if !m.TargetLiveness.IsNull() {
		params.TargetLiveness = param.NewOpt(m.TargetLiveness.ValueFloat64())
	}
	if !m.TargetLoudness.IsNull() {
		params.TargetLoudness = param.NewOpt(m.TargetLoudness.ValueFloat64())
	}
	if !m.TargetMode.IsNull() {
		params.TargetMode = param.NewOpt(m.TargetMode.ValueInt64())
	}
	if !m.TargetPopularity.IsNull() {
		params.TargetPopularity = param.NewOpt(m.TargetPopularity.ValueInt64())
	}
	if !m.TargetSpeechiness.IsNull() {
		params.TargetSpeechiness = param.NewOpt(m.TargetSpeechiness.ValueFloat64())
	}
	if !m.TargetTempo.IsNull() {
		params.TargetTempo = param.NewOpt(m.TargetTempo.ValueFloat64())
	}
	if !m.TargetTimeSignature.IsNull() {
		params.TargetTimeSignature = param.NewOpt(m.TargetTimeSignature.ValueInt64())
	}
	if !m.TargetValence.IsNull() {
		params.TargetValence = param.NewOpt(m.TargetValence.ValueFloat64())
	}

	return
}

type RecommendationSeedsDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AfterFilteringSize types.Int64  `tfsdk:"after_filtering_size" json:"afterFilteringSize,computed"`
	AfterRelinkingSize types.Int64  `tfsdk:"after_relinking_size" json:"afterRelinkingSize,computed"`
	Href               types.String `tfsdk:"href" json:"href,computed"`
	InitialPoolSize    types.Int64  `tfsdk:"initial_pool_size" json:"initialPoolSize,computed"`
	Type               types.String `tfsdk:"type" json:"type,computed"`
}

type RecommendationTracksDataSourceModel struct {
	ID               types.String                                                              `tfsdk:"id" json:"id,computed"`
	Album            customfield.NestedObject[RecommendationTracksAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists          customfield.NestedObjectList[RecommendationTracksArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets customfield.List[types.String]                                            `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber       types.Int64                                                               `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs       types.Int64                                                               `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit         types.Bool                                                                `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs      customfield.NestedObject[RecommendationTracksExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs     customfield.NestedObject[RecommendationTracksExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                              `tfsdk:"href" json:"href,computed"`
	IsLocal          types.Bool                                                                `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable       types.Bool                                                                `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom       customfield.NestedObject[RecommendationTracksLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name             types.String                                                              `tfsdk:"name" json:"name,computed"`
	Popularity       types.Int64                                                               `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL       types.String                                                              `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions     customfield.NestedObject[RecommendationTracksRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber      types.Int64                                                               `tfsdk:"track_number" json:"track_number,computed"`
	Type             types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type RecommendationTracksAlbumDataSourceModel struct {
	ID                   types.String                                                                   `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                   `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[RecommendationTracksAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                                 `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[RecommendationTracksAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                   `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[RecommendationTracksAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                   `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                   `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                   `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                    `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[RecommendationTracksAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type RecommendationTracksAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                          `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[RecommendationTracksAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                          `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                          `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                          `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                          `tfsdk:"uri" json:"uri,computed"`
}

type RecommendationTracksAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type RecommendationTracksAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type RecommendationTracksAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type RecommendationTracksAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type RecommendationTracksArtistsDataSourceModel struct {
	ID           types.String                                                                     `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[RecommendationTracksArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                     `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                     `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                     `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                     `tfsdk:"uri" json:"uri,computed"`
}

type RecommendationTracksArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type RecommendationTracksExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type RecommendationTracksExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type RecommendationTracksLinkedFromDataSourceModel struct {
	ID           types.String                                                                        `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[RecommendationTracksLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                        `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                        `tfsdk:"uri" json:"uri,computed"`
}

type RecommendationTracksLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type RecommendationTracksRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}
