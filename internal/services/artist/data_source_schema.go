// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package artist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ArtistDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the artist.",
				Required:    true,
			},
			"href": schema.StringAttribute{
				Description: "A link to the Web API endpoint providing full details of the artist.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the artist.",
				Computed:    true,
			},
			"popularity": schema.Int64Attribute{
				Description: "The popularity of the artist. The value will be between 0 and 100, with 100 being the most popular. The artist's popularity is calculated from the popularity of all the artist's tracks.",
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
			"genres": schema.ListAttribute{
				Description: "A list of the genres the artist is associated with. If not yet classified, the array is empty.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"external_urls": schema.SingleNestedAttribute{
				Description: "Known external URLs for this artist.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ArtistExternalURLsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"spotify": schema.StringAttribute{
						Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
						Computed:    true,
					},
				},
			},
			"followers": schema.SingleNestedAttribute{
				Description: "Information about the followers of the artist.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ArtistFollowersDataSourceModel](ctx),
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
				Description: "Images of the artist in various sizes, widest first.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ArtistImagesDataSourceModel](ctx),
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
		},
	}
}

func (d *ArtistDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ArtistDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
