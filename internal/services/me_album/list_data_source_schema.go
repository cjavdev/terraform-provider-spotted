// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_album

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*MeAlbumsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"market": schema.StringAttribute{
				Description: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).",
				Optional:    true,
			},
			"limit": schema.Int64Attribute{
				Description: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 50),
				},
			},
			"offset": schema.Int64Attribute{
				Description: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.",
				Computed:    true,
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MeAlbumsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"added_at": schema.StringAttribute{
							Description: "The date and time the album was saved\nTimestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.\nIf the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"album": schema.SingleNestedAttribute{
							Description: "Information about the album.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumDataSourceModel](ctx),
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
								"available_markets": schema.ListAttribute{
									Description: "The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumExternalURLsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[MeAlbumsAlbumImagesDataSourceModel](ctx),
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
								"artists": schema.ListNestedAttribute{
									Description: "The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeAlbumsAlbumArtistsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
											"external_urls": schema.SingleNestedAttribute{
												Description: "Known external URLs for this artist.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumArtistsExternalURLsDataSourceModel](ctx),
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
								"copyrights": schema.ListNestedAttribute{
									Description: "The copyright statements of the album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[MeAlbumsAlbumCopyrightsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"text": schema.StringAttribute{
												Description: "The copyright text for this content.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The type of copyright: `C` = the copyright, `P` = the sound recording (performance) copyright.",
												Computed:    true,
											},
										},
									},
								},
								"external_ids": schema.SingleNestedAttribute{
									Description: "Known external IDs for the album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumExternalIDsDataSourceModel](ctx),
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
								"genres": schema.ListAttribute{
									Description:        "**Deprecated** The array is always empty.",
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
									CustomType:         customfield.NewListType[types.String](ctx),
									ElementType:        types.StringType,
								},
								"label": schema.StringAttribute{
									Description: "The label associated with the album.",
									Computed:    true,
								},
								"popularity": schema.Int64Attribute{
									Description: "The popularity of the album. The value will be between 0 and 100, with 100 being the most popular.",
									Computed:    true,
								},
								"restrictions": schema.SingleNestedAttribute{
									Description: "Included in the response when a content restriction is applied.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumRestrictionsDataSourceModel](ctx),
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
								"tracks": schema.SingleNestedAttribute{
									Description: "The tracks of the album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumTracksDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint returning the full result of the request",
											Computed:    true,
										},
										"items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[MeAlbumsAlbumTracksItemsDataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
														Computed:    true,
													},
													"artists": schema.ListNestedAttribute{
														Description: "The artists who performed the track. Each artist object includes a link in `href` to more detailed information about the artist.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectListType[MeAlbumsAlbumTracksItemsArtistsDataSourceModel](ctx),
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"id": schema.StringAttribute{
																	Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
																	Computed:    true,
																},
																"external_urls": schema.SingleNestedAttribute{
																	Description: "Known external URLs for this artist.",
																	Computed:    true,
																	CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumTracksItemsArtistsExternalURLsDataSourceModel](ctx),
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
													"external_urls": schema.SingleNestedAttribute{
														Description: "External URLs for this track.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumTracksItemsExternalURLsDataSourceModel](ctx),
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
													"is_local": schema.BoolAttribute{
														Description: "Whether or not the track is from a local file.",
														Computed:    true,
													},
													"is_playable": schema.BoolAttribute{
														Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking/) is applied. If `true`, the track is playable in the given market. Otherwise `false`.",
														Computed:    true,
													},
													"linked_from": schema.SingleNestedAttribute{
														Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking/) is applied and is only part of the response if the track linking, in fact, exists. The requested track has been replaced with a different track. The track in the `linked_from` object contains information about the originally requested track.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumTracksItemsLinkedFromDataSourceModel](ctx),
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
																Computed:    true,
															},
															"external_urls": schema.SingleNestedAttribute{
																Description: "Known external URLs for this track.",
																Computed:    true,
																CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumTracksItemsLinkedFromExternalURLsDataSourceModel](ctx),
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
													"name": schema.StringAttribute{
														Description: "The name of the track.",
														Computed:    true,
													},
													"preview_url": schema.StringAttribute{
														Description:        "A URL to a 30 second preview (MP3 format) of the track.",
														Computed:           true,
														DeprecationMessage: "This attribute is deprecated.",
													},
													"restrictions": schema.SingleNestedAttribute{
														Description: "Included in the response when a content restriction is applied.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[MeAlbumsAlbumTracksItemsRestrictionsDataSourceModel](ctx),
														Attributes: map[string]schema.Attribute{
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
														Description: `The object type: "track".`,
														Computed:    true,
													},
													"uri": schema.StringAttribute{
														Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
														Computed:    true,
													},
												},
											},
										},
										"limit": schema.Int64Attribute{
											Description: "The maximum number of items in the response (as set in the query or by default).",
											Computed:    true,
										},
										"next": schema.StringAttribute{
											Description: "URL to the next page of items. ( `null` if none)",
											Computed:    true,
										},
										"offset": schema.Int64Attribute{
											Description: "The offset of the items returned (as set in the query or by default)",
											Computed:    true,
										},
										"previous": schema.StringAttribute{
											Description: "URL to the previous page of items. ( `null` if none)",
											Computed:    true,
										},
										"total": schema.Int64Attribute{
											Description: "The total number of items available to return.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *MeAlbumsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MeAlbumsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
