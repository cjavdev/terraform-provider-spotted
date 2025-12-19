// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*UserPlaylistResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"user_id": schema.StringAttribute{
				Description:   "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The name for the new playlist, for example `\"Your Coolest Playlist\"`. This name does not need to be unique; a user may have several playlists with the same name.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"collaborative": schema.BoolAttribute{
				Description:   "Defaults to `false`. If `true` the playlist will be collaborative. _**Note**: to create a collaborative playlist you must also set `public` to `false`. To create collaborative playlists you must have granted `playlist-modify-private` and `playlist-modify-public` [scopes](/documentation/web-api/concepts/scopes/#list-of-scopes)._",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description:   "value for playlist description as displayed in Spotify Clients and in the Web API.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"public": schema.BoolAttribute{
				Description:   "Defaults to `true`. The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private. To be able to create private playlists, the user must have granted the `playlist-modify-private` [scope](/documentation/web-api/concepts/scopes/#list-of-scopes). For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"href": schema.StringAttribute{
				Description: "A link to the Web API endpoint providing full details of the playlist.",
				Computed:    true,
			},
			"published": schema.BoolAttribute{
				Description: "The playlist's public/private status (if it is added to the user's profile): `true` the playlist is public, `false` the playlist is private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
				Computed:    true,
			},
			"snapshot_id": schema.StringAttribute{
				Description: "The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: `The object type: "playlist"`,
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.",
				Computed:    true,
			},
			"external_urls": schema.SingleNestedAttribute{
				Description: "Known external URLs for this playlist.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[UserPlaylistExternalURLsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"spotify": schema.StringAttribute{
						Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
						Computed:    true,
					},
				},
			},
			"followers": schema.SingleNestedAttribute{
				Description: "Information about the followers of the playlist.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[UserPlaylistFollowersModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "This will always be set to null, as the Web API does not support it at the moment.",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of followers.",
						Computed:    true,
					},
				},
			},
			"images": schema.ListNestedAttribute{
				Description: "Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order. See [Working with Playlists](/documentation/web-api/concepts/playlists). _**Note**: If returned, the source URL for the image (`url`) is temporary and will expire in less than a day._",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[UserPlaylistImagesModel](ctx),
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
			"owner": schema.SingleNestedAttribute{
				Description: "The user who owns the playlist",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[UserPlaylistOwnerModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
						Computed:    true,
					},
					"external_urls": schema.SingleNestedAttribute{
						Description: "Known public external URLs for this user.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[UserPlaylistOwnerExternalURLsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"spotify": schema.StringAttribute{
								Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
								Computed:    true,
							},
						},
					},
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint for this user.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The object type.\nAvailable values: \"user\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("user"),
						},
					},
					"uri": schema.StringAttribute{
						Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
						Computed:    true,
					},
					"display_name": schema.StringAttribute{
						Description: "The name displayed on the user's profile. `null` if not available.",
						Computed:    true,
					},
				},
			},
			"tracks": schema.SingleNestedAttribute{
				Description: "The tracks of the playlist.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
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
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[UserPlaylistTracksItemsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"added_at": schema.StringAttribute{
									Description: "The date and time the track or episode was added. _**Note**: some very old playlists may return `null` in this field._",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"added_by": schema.SingleNestedAttribute{
									Description: "The Spotify user who added the track or episode. _**Note**: some very old playlists may return `null` in this field._",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsAddedByModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
											Computed:    true,
										},
										"external_urls": schema.SingleNestedAttribute{
											Description: "Known public external URLs for this user.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsAddedByExternalURLsModel](ctx),
											Attributes: map[string]schema.Attribute{
												"spotify": schema.StringAttribute{
													Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
													Computed:    true,
												},
											},
										},
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint for this user.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: "The object type.\nAvailable values: \"user\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("user"),
											},
										},
										"uri": schema.StringAttribute{
											Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
											Computed:    true,
										},
									},
								},
								"is_local": schema.BoolAttribute{
									Description: "Whether this track or episode is a [local file](/documentation/web-api/concepts/playlists/#local-files) or not.",
									Computed:    true,
								},
								"track": schema.SingleNestedAttribute{
									Description: "Information about the track or episode.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
											Computed:    true,
										},
										"album": schema.SingleNestedAttribute{
											Description: "The album on which the track appears. The album object includes a link in `href` to full information about the album.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackAlbumModel](ctx),
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
													CustomType:  customfield.NewNestedObjectListType[UserPlaylistTracksItemsTrackAlbumArtistsModel](ctx),
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
																Computed:    true,
															},
															"external_urls": schema.SingleNestedAttribute{
																Description: "Known external URLs for this artist.",
																Computed:    true,
																CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackAlbumArtistsExternalURLsModel](ctx),
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
													CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackAlbumExternalURLsModel](ctx),
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
													CustomType:  customfield.NewNestedObjectListType[UserPlaylistTracksItemsTrackAlbumImagesModel](ctx),
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
													CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackAlbumRestrictionsModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[UserPlaylistTracksItemsTrackArtistsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
														Computed:    true,
													},
													"external_urls": schema.SingleNestedAttribute{
														Description: "Known external URLs for this artist.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackArtistsExternalURLsModel](ctx),
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
										"external_ids": schema.SingleNestedAttribute{
											Description: "Known external IDs for the track.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackExternalIDsModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackExternalURLsModel](ctx),
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
											Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If `true`, the track is playable in the given market. Otherwise `false`.",
											Computed:    true,
										},
										"linked_from": schema.SingleNestedAttribute{
											Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied, and the requested track has been replaced with different track. The track in the `linked_from` object contains information about the originally requested track.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackLinkedFromModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
													Computed:    true,
												},
												"external_urls": schema.SingleNestedAttribute{
													Description: "Known external URLs for this track.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackLinkedFromExternalURLsModel](ctx),
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
										"popularity": schema.Int64Attribute{
											Description: "The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.<br/>The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.<br/>Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._",
											Computed:    true,
										},
										"preview_url": schema.StringAttribute{
											Description:        "A link to a 30 second preview (MP3 format) of the track. Can be `null`",
											Computed:           true,
											DeprecationMessage: "This attribute is deprecated.",
										},
										"restrictions": schema.SingleNestedAttribute{
											Description: "Included in the response when a content restriction is applied.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackRestrictionsModel](ctx),
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
											Description: "The object type: \"track\".\nAvailable values: \"track\", \"episode\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("track", "episode"),
											},
										},
										"uri": schema.StringAttribute{
											Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
											Computed:    true,
										},
										"audio_preview_url": schema.StringAttribute{
											Description:        "A URL to a 30 second preview (MP3 format) of the episode. `null` if not available.",
											Computed:           true,
											DeprecationMessage: "This attribute is deprecated.",
										},
										"description": schema.StringAttribute{
											Description: "A description of the episode. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
											Computed:    true,
										},
										"html_description": schema.StringAttribute{
											Description: "A description of the episode. This field may contain HTML tags.",
											Computed:    true,
										},
										"images": schema.ListNestedAttribute{
											Description: "The cover art for the episode in various sizes, widest first.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[UserPlaylistTracksItemsTrackImagesModel](ctx),
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
										"is_externally_hosted": schema.BoolAttribute{
											Description: "True if the episode is hosted outside of Spotify's CDN.",
											Computed:    true,
										},
										"languages": schema.ListAttribute{
											Description: "A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"release_date": schema.StringAttribute{
											Description: "The date the episode was first released, for example `\"1981-12-15\"`. Depending on the precision, it might be shown as `\"1981\"` or `\"1981-12\"`.",
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
										"show": schema.SingleNestedAttribute{
											Description: "The show on which the episode belongs.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackShowModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the show.",
													Computed:    true,
												},
												"available_markets": schema.ListAttribute{
													Description: "A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
												"copyrights": schema.ListNestedAttribute{
													Description: "The copyright statements of the show.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectListType[UserPlaylistTracksItemsTrackShowCopyrightsModel](ctx),
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
												"description": schema.StringAttribute{
													Description: "A description of the show. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
													Computed:    true,
												},
												"explicit": schema.BoolAttribute{
													Description: "Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown).",
													Computed:    true,
												},
												"external_urls": schema.SingleNestedAttribute{
													Description: "External URLs for this show.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackShowExternalURLsModel](ctx),
													Attributes: map[string]schema.Attribute{
														"spotify": schema.StringAttribute{
															Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
															Computed:    true,
														},
													},
												},
												"href": schema.StringAttribute{
													Description: "A link to the Web API endpoint providing full details of the show.",
													Computed:    true,
												},
												"html_description": schema.StringAttribute{
													Description: "A description of the show. This field may contain HTML tags.",
													Computed:    true,
												},
												"images": schema.ListNestedAttribute{
													Description: "The cover art for the show in various sizes, widest first.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectListType[UserPlaylistTracksItemsTrackShowImagesModel](ctx),
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
												"is_externally_hosted": schema.BoolAttribute{
													Description: "True if all of the shows episodes are hosted outside of Spotify's CDN. This field might be `null` in some cases.",
													Computed:    true,
												},
												"languages": schema.ListAttribute{
													Description: "A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.",
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
												"media_type": schema.StringAttribute{
													Description: "The media type of the show.",
													Computed:    true,
												},
												"name": schema.StringAttribute{
													Description: "The name of the episode.",
													Computed:    true,
												},
												"publisher": schema.StringAttribute{
													Description: "The publisher of the show.",
													Computed:    true,
												},
												"total_episodes": schema.Int64Attribute{
													Description: "The total number of episodes in the show.",
													Computed:    true,
												},
												"type": schema.StringAttribute{
													Description: "The object type.\nAvailable values: \"show\".",
													Computed:    true,
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive("show"),
													},
												},
												"uri": schema.StringAttribute{
													Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the show.",
													Computed:    true,
												},
											},
										},
										"language": schema.StringAttribute{
											Description:        "The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the `languages` field instead.",
											Computed:           true,
											DeprecationMessage: "This attribute is deprecated.",
										},
										"resume_point": schema.SingleNestedAttribute{
											Description: "The user's most recent position in the episode. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[UserPlaylistTracksItemsTrackResumePointModel](ctx),
											Attributes: map[string]schema.Attribute{
												"fully_played": schema.BoolAttribute{
													Description: "Whether or not the episode has been fully played by the user.",
													Computed:    true,
												},
												"resume_position_ms": schema.Int64Attribute{
													Description: "The user's most recent position in the episode in milliseconds.",
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
			},
		},
	}
}

func (r *UserPlaylistResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserPlaylistResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
