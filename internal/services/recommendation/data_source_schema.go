// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package recommendation

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*RecommendationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"market": schema.StringAttribute{
				Description: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).",
				Optional:    true,
			},
			"max_acousticness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"max_danceability": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"max_duration_ms": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
			},
			"max_energy": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"max_instrumentalness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"max_key": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 11),
				},
			},
			"max_liveness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"max_loudness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
			},
			"max_mode": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1),
				},
			},
			"max_popularity": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"max_speechiness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"max_tempo": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
			},
			"max_time_signature": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
			},
			"max_valence": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `max_instrumentalness=0.35` would filter out most tracks that are likely to be instrumental.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_acousticness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_danceability": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_duration_ms": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
			},
			"min_energy": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_instrumentalness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_key": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 11),
				},
			},
			"min_liveness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_loudness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
			},
			"min_mode": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1),
				},
			},
			"min_popularity": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"min_speechiness": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"min_tempo": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
			},
			"min_time_signature": schema.Int64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtMost(11),
				},
			},
			"min_valence": schema.Float64Attribute{
				Description: "For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, `min_tempo=140` would restrict results to only those tracks with a tempo of greater than 140 beats per minute.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"seed_artists": schema.StringAttribute{
				Description: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for seed artists.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_genres` and `seed_tracks` are not set_.",
				Optional:    true,
			},
			"seed_genres": schema.StringAttribute{
				Description: "A comma separated list of any genres in the set of [available genre seeds](/documentation/web-api/reference/get-recommendation-genres). Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_tracks` are not set_.",
				Optional:    true,
			},
			"seed_tracks": schema.StringAttribute{
				Description: "A comma separated list of [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for a seed track.  Up to 5 seed values may be provided in any combination of `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists` and `seed_genres` are not set_.",
				Optional:    true,
			},
			"target_acousticness": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"target_danceability": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"target_duration_ms": schema.Int64Attribute{
				Description: "Target duration of the track (ms)",
				Optional:    true,
			},
			"target_energy": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"target_instrumentalness": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"target_key": schema.Int64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 11),
				},
			},
			"target_liveness": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"target_loudness": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
			},
			"target_mode": schema.Int64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1),
				},
			},
			"target_popularity": schema.Int64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"target_speechiness": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"target_tempo": schema.Float64Attribute{
				Description: "Target tempo (BPM)",
				Optional:    true,
			},
			"target_time_signature": schema.Int64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
			},
			"target_valence": schema.Float64Attribute{
				Description: "For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request `target_energy=0.6` and `target_danceability=0.8`. All target values will be weighed equally in ranking results.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"limit": schema.Int64Attribute{
				Description: "The target size of the list of recommended tracks. For seeds with unusually small pools or when highly restrictive filtering is applied, it may be impossible to generate the requested number of recommended tracks. Debugging information for such cases is available in the response. Default: 20\\. Minimum: 1\\. Maximum: 100.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"published": schema.BoolAttribute{
				Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
				Computed:    true,
			},
			"seeds": schema.ListNestedAttribute{
				Description: "An array of recommendation seed objects.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[RecommendationSeedsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id used to select this seed. This will be the same as the string used in the `seed_artists`, `seed_tracks` or `seed_genres` parameter.",
							Computed:    true,
						},
						"after_filtering_size": schema.Int64Attribute{
							Description: "The number of tracks available after min\\_\\* and max\\_\\* filters have been applied.",
							Computed:    true,
						},
						"after_relinking_size": schema.Int64Attribute{
							Description: "The number of tracks available after relinking for regional availability.",
							Computed:    true,
						},
						"href": schema.StringAttribute{
							Description: "A link to the full track or artist data for this seed. For tracks this will be a link to a Track Object. For artists a link to an Artist Object. For genre seeds, this value will be `null`.",
							Computed:    true,
						},
						"initial_pool_size": schema.Int64Attribute{
							Description: "The number of recommended tracks available for this seed.",
							Computed:    true,
						},
						"published": schema.BoolAttribute{
							Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The entity type of this seed. One of `artist`, `track` or `genre`.",
							Computed:    true,
						},
					},
				},
			},
			"tracks": schema.ListNestedAttribute{
				Description: "An array of track objects ordered according to the parameters supplied.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[RecommendationTracksDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
							Computed:    true,
						},
						"album": schema.SingleNestedAttribute{
							Description: "The album on which the track appears. The album object includes a link in `href` to full information about the album.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RecommendationTracksAlbumDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the album.",
									Computed:    true,
								},
								"album_type": schema.StringAttribute{
									Description: "The type of the album.\nAvailable values: \"album\", \"single\", \"compilation\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"album",
											"single",
											"compilation",
										),
									},
								},
								"artists": schema.ListNestedAttribute{
									Description: "The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[RecommendationTracksAlbumArtistsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
											"external_urls": schema.SingleNestedAttribute{
												Description: "Known external URLs for this artist.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[RecommendationTracksAlbumArtistsExternalURLsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"published": schema.BoolAttribute{
														Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
														Computed:    true,
													},
													"spotify": schema.StringAttribute{
														Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
														Computed:    true,
													},
												},
											},
											"href": schema.StringAttribute{
												Description: "A link to the Web API endpoint providing full details of the artist.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the artist.",
												Computed:    true,
											},
											"published": schema.BoolAttribute{
												Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The object type.\nAvailable values: \"artist\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("artist"),
												},
											},
											"uri": schema.StringAttribute{
												Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
										},
									},
								},
								"available_markets": schema.ListAttribute{
									Description: "The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[RecommendationTracksAlbumExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"published": schema.BoolAttribute{
											Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
											Computed:    true,
										},
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the album.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the album in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[RecommendationTracksAlbumImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
											"published": schema.BoolAttribute{
												Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
												Computed:    true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Description: "The name of the album. In case of an album takedown, the value may be an empty string.",
									Computed:    true,
								},
								"release_date": schema.StringAttribute{
									Description: "The date the album was first released.",
									Computed:    true,
								},
								"release_date_precision": schema.StringAttribute{
									Description: "The precision with which `release_date` value is known.\nAvailable values: \"year\", \"month\", \"day\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"year",
											"month",
											"day",
										),
									},
								},
								"total_tracks": schema.Int64Attribute{
									Description: "The number of tracks in the album.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"album\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("album"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the album.",
									Computed:    true,
								},
								"published": schema.BoolAttribute{
									Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
									Computed:    true,
								},
								"restrictions": schema.SingleNestedAttribute{
									Description: "Included in the response when a content restriction is applied.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[RecommendationTracksAlbumRestrictionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"published": schema.BoolAttribute{
											Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
											Computed:    true,
										},
										"reason": schema.StringAttribute{
											Description: "The reason for the restriction. Albums may be restricted if the content is not available in a given market, to the user's subscription type, or when the user's account is set to not play explicit content.\nAdditional reasons may be added in the future.\nAvailable values: \"market\", \"product\", \"explicit\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"market",
													"product",
													"explicit",
												),
											},
										},
									},
								},
							},
						},
						"artists": schema.ListNestedAttribute{
							Description: "The artists who performed the track. Each artist object includes a link in `href` to more detailed information about the artist.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[RecommendationTracksArtistsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
										Computed:    true,
									},
									"external_urls": schema.SingleNestedAttribute{
										Description: "Known external URLs for this artist.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[RecommendationTracksArtistsExternalURLsDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"published": schema.BoolAttribute{
												Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
												Computed:    true,
											},
											"spotify": schema.StringAttribute{
												Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
												Computed:    true,
											},
										},
									},
									"href": schema.StringAttribute{
										Description: "A link to the Web API endpoint providing full details of the artist.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "The name of the artist.",
										Computed:    true,
									},
									"published": schema.BoolAttribute{
										Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
										Computed:    true,
									},
									"type": schema.StringAttribute{
										Description: "The object type.\nAvailable values: \"artist\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("artist"),
										},
									},
									"uri": schema.StringAttribute{
										Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
										Computed:    true,
									},
								},
							},
						},
						"available_markets": schema.ListAttribute{
							Description: "A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"disc_number": schema.Int64Attribute{
							Description: "The disc number (usually `1` unless the album consists of more than one disc).",
							Computed:    true,
						},
						"duration_ms": schema.Int64Attribute{
							Description: "The track length in milliseconds.",
							Computed:    true,
						},
						"explicit": schema.BoolAttribute{
							Description: "Whether or not the track has explicit lyrics ( `true` = yes it does; `false` = no it does not OR unknown).",
							Computed:    true,
						},
						"external_ids": schema.SingleNestedAttribute{
							Description: "Known external IDs for the track.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RecommendationTracksExternalIDsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ean": schema.StringAttribute{
									Description: "[International Article Number](http://en.wikipedia.org/wiki/International_Article_Number_%28EAN%29)",
									Computed:    true,
								},
								"isrc": schema.StringAttribute{
									Description: "[International Standard Recording Code](http://en.wikipedia.org/wiki/International_Standard_Recording_Code)",
									Computed:    true,
								},
								"published": schema.BoolAttribute{
									Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
									Computed:    true,
								},
								"upc": schema.StringAttribute{
									Description: "[Universal Product Code](http://en.wikipedia.org/wiki/Universal_Product_Code)",
									Computed:    true,
								},
							},
						},
						"external_urls": schema.SingleNestedAttribute{
							Description: "Known external URLs for this track.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RecommendationTracksExternalURLsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"published": schema.BoolAttribute{
									Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
									Computed:    true,
								},
								"spotify": schema.StringAttribute{
									Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
									Computed:    true,
								},
							},
						},
						"href": schema.StringAttribute{
							Description: "A link to the Web API endpoint providing full details of the track.",
							Computed:    true,
						},
						"is_local": schema.BoolAttribute{
							Description: "Whether or not the track is from a local file.",
							Computed:    true,
						},
						"is_playable": schema.BoolAttribute{
							Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If `true`, the track is playable in the given market. Otherwise `false`.",
							Computed:    true,
						},
						"linked_from": schema.SingleNestedAttribute{
							Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied, and the requested track has been replaced with different track. The track in the `linked_from` object contains information about the originally requested track.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RecommendationTracksLinkedFromDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this track.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[RecommendationTracksLinkedFromExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"published": schema.BoolAttribute{
											Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
											Computed:    true,
										},
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the track.",
									Computed:    true,
								},
								"published": schema.BoolAttribute{
									Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: `The object type: "track".`,
									Computed:    true,
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
									Computed:    true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the track.",
							Computed:    true,
						},
						"popularity": schema.Int64Attribute{
							Description: "The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.<br/>The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.<br/>Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._",
							Computed:    true,
						},
						"preview_url": schema.StringAttribute{
							Description:        "A link to a 30 second preview (MP3 format) of the track. Can be `null`",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
						},
						"published": schema.BoolAttribute{
							Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
							Computed:    true,
						},
						"restrictions": schema.SingleNestedAttribute{
							Description: "Included in the response when a content restriction is applied.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RecommendationTracksRestrictionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"published": schema.BoolAttribute{
									Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
									Computed:    true,
								},
								"reason": schema.StringAttribute{
									Description: "The reason for the restriction. Supported values:\n- `market` - The content item is not available in the given market.\n- `product` - The content item is not available for the user's subscription type.\n- `explicit` - The content item is explicit and the user's account is set to not play explicit content.\n\nAdditional reasons may be added in the future.\n**Note**: If you use this field, make sure that your application safely handles unknown values.",
									Computed:    true,
								},
							},
						},
						"track_number": schema.Int64Attribute{
							Description: "The number of the track. If an album has several discs, the track number is the number on the specified disc.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The object type: \"track\".\nAvailable values: \"track\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("track"),
							},
						},
						"uri": schema.StringAttribute{
							Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *RecommendationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *RecommendationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
