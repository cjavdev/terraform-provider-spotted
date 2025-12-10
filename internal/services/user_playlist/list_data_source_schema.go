// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UserPlaylistsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user_id": schema.StringAttribute{
				Description: "The user's [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids).",
				Required:    true,
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
				Description: "The index of the first playlist to return. Default:\n0 (the first object). Maximum offset: 100.000\\. Use with `limit` to get the\nnext set of playlists.",
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
				CustomType:  customfield.NewNestedObjectListType[UserPlaylistsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.",
							Computed:    true,
						},
						"collaborative": schema.BoolAttribute{
							Description: "`true` if the owner allows other users to modify the playlist.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "The playlist description. _Only returned for modified, verified playlists, otherwise_ `null`.",
							Computed:    true,
						},
						"external_urls": schema.SingleNestedAttribute{
							Description: "Known external URLs for this playlist.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[UserPlaylistsExternalURLsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"spotify": schema.StringAttribute{
									Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
									Computed:    true,
								},
							},
						},
						"href": schema.StringAttribute{
							Description: "A link to the Web API endpoint providing full details of the playlist.",
							Computed:    true,
						},
						"images": schema.ListNestedAttribute{
							Description: "Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order. See [Working with Playlists](/documentation/web-api/concepts/playlists). _**Note**: If returned, the source URL for the image (`url`) is temporary and will expire in less than a day._",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[UserPlaylistsImagesDataSourceModel](ctx),
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
							Description: "The name of the playlist.",
							Computed:    true,
						},
						"owner": schema.SingleNestedAttribute{
							Description: "The user who owns the playlist",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[UserPlaylistsOwnerDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known public external URLs for this user.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[UserPlaylistsOwnerExternalURLsDataSourceModel](ctx),
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
						"published": schema.BoolAttribute{
							Description: "The playlist's public/private status (if it is added to the user's profile): `true` the playlist is public, `false` the playlist is private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
							Computed:    true,
						},
						"snapshot_id": schema.StringAttribute{
							Description: "The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version",
							Computed:    true,
						},
						"tracks": schema.SingleNestedAttribute{
							Description: "A collection containing a link ( `href` ) to the Web API endpoint where full details of the playlist's tracks can be retrieved, along with the `total` number of tracks in the playlist. Note, a track object may be `null`. This can happen if a track is no longer available.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[UserPlaylistsTracksDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint where full details of the playlist's tracks can be retrieved.",
									Computed:    true,
								},
								"total": schema.Int64Attribute{
									Description: "Number of tracks in the playlist.",
									Computed:    true,
								},
							},
						},
						"type": schema.StringAttribute{
							Description: `The object type: "playlist"`,
							Computed:    true,
						},
						"uri": schema.StringAttribute{
							Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UserPlaylistsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *UserPlaylistsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
