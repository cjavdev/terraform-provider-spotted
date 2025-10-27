// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package track

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*TrackDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids)\nfor the track.",
				Required:    true,
			},
			"market": schema.StringAttribute{
				Description: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).",
				Optional:    true,
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
			"available_markets": schema.ListAttribute{
				Description: "A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"album": schema.SingleNestedAttribute{
				Description: "The album on which the track appears. The album object includes a link in `href` to full information about the album.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TrackAlbumDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[TrackAlbumArtistsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[TrackAlbumArtistsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
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
						CustomType:  customfield.NewNestedObjectType[TrackAlbumExternalURLsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
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
						CustomType:  customfield.NewNestedObjectListType[TrackAlbumImagesDataSourceModel](ctx),
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
					"restrictions": schema.SingleNestedAttribute{
						Description: "Included in the response when a content restriction is applied.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[TrackAlbumRestrictionsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
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
				CustomType:  customfield.NewNestedObjectListType[TrackArtistsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
							Computed:    true,
						},
						"external_urls": schema.SingleNestedAttribute{
							Description: "Known external URLs for this artist.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[TrackArtistsExternalURLsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
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
			"external_ids": schema.SingleNestedAttribute{
				Description: "Known external IDs for the track.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TrackExternalIDsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ean": schema.StringAttribute{
						Description: "[International Article Number](http://en.wikipedia.org/wiki/International_Article_Number_%28EAN%29)",
						Computed:    true,
					},
					"isrc": schema.StringAttribute{
						Description: "[International Standard Recording Code](http://en.wikipedia.org/wiki/International_Standard_Recording_Code)",
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
				CustomType:  customfield.NewNestedObjectType[TrackExternalURLsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"spotify": schema.StringAttribute{
						Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
						Computed:    true,
					},
				},
			},
			"linked_from": schema.SingleNestedAttribute{
				Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied, and the requested track has been replaced with different track. The track in the `linked_from` object contains information about the originally requested track.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TrackLinkedFromDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
						Computed:    true,
					},
					"external_urls": schema.SingleNestedAttribute{
						Description: "Known external URLs for this track.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[TrackLinkedFromExternalURLsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
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
			"restrictions": schema.SingleNestedAttribute{
				Description: "Included in the response when a content restriction is applied.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TrackRestrictionsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"reason": schema.StringAttribute{
						Description: "The reason for the restriction. Supported values:\n- `market` - The content item is not available in the given market.\n- `product` - The content item is not available for the user's subscription type.\n- `explicit` - The content item is explicit and the user's account is set to not play explicit content.\n\nAdditional reasons may be added in the future.\n**Note**: If you use this field, make sure that your application safely handles unknown values.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *TrackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TrackDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
