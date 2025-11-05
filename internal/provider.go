// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-terraform/internal/services/album"
	"github.com/stainless-sdks/spotted-terraform/internal/services/artist"
	"github.com/stainless-sdks/spotted-terraform/internal/services/audio_analysis"
	"github.com/stainless-sdks/spotted-terraform/internal/services/audio_feature"
	"github.com/stainless-sdks/spotted-terraform/internal/services/audiobook"
	"github.com/stainless-sdks/spotted-terraform/internal/services/browse_category"
	"github.com/stainless-sdks/spotted-terraform/internal/services/chapter"
	"github.com/stainless-sdks/spotted-terraform/internal/services/episode"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_album"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_audiobook"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_episode"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_player_queue"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_playlist"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_show"
	"github.com/stainless-sdks/spotted-terraform/internal/services/me_track"
	"github.com/stainless-sdks/spotted-terraform/internal/services/playlist"
	"github.com/stainless-sdks/spotted-terraform/internal/services/playlist_image"
	"github.com/stainless-sdks/spotted-terraform/internal/services/playlist_track"
	"github.com/stainless-sdks/spotted-terraform/internal/services/recommendation"
	"github.com/stainless-sdks/spotted-terraform/internal/services/search"
	"github.com/stainless-sdks/spotted-terraform/internal/services/show"
	"github.com/stainless-sdks/spotted-terraform/internal/services/track"
	"github.com/stainless-sdks/spotted-terraform/internal/services/user_playlist"
)

var _ provider.ProviderWithConfigValidators = (*SpottedProvider)(nil)

// SpottedProvider defines the provider implementation.
type SpottedProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SpottedProviderModel describes the provider data model.
type SpottedProviderModel struct {
	BaseURL      types.String `tfsdk:"base_url" json:"base_url,optional"`
	ClientID     types.String `tfsdk:"client_id" json:"client_id,optional"`
	ClientSecret types.String `tfsdk:"client_secret" json:"client_secret,optional"`
}

func (p *SpottedProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "spotted"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"client_id": schema.StringAttribute{
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *SpottedProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *SpottedProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data SpottedProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("SPOTTED_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.ClientID.IsNull() && !data.ClientID.IsUnknown() {
		opts = append(opts, option.WithClientID(data.ClientID.ValueString()))
	} else if o, ok := os.LookupEnv("SPOTIFY_CLIENT_ID"); ok {
		opts = append(opts, option.WithClientID(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing client_id value",
			"The client_id field is required. Set it in provider configuration or via the \"SPOTIFY_CLIENT_ID\" environment variable.",
		)
		return
	}

	if !data.ClientSecret.IsNull() && !data.ClientSecret.IsUnknown() {
		opts = append(opts, option.WithClientSecret(data.ClientSecret.ValueString()))
	} else if o, ok := os.LookupEnv("SPOTIFY_CLIENT_SECRET"); ok {
		opts = append(opts, option.WithClientSecret(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing client_secret value",
			"The client_secret field is required. Set it in provider configuration or via the \"SPOTIFY_CLIENT_SECRET\" environment variable.",
		)
		return
	}

	client := spotted.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *SpottedProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *SpottedProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		playlist.NewResource,
		playlist_track.NewResource,
		playlist_image.NewResource,
		user_playlist.NewResource,
	}
}

func (p *SpottedProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		album.NewAlbumDataSource,
		artist.NewArtistDataSource,
		show.NewShowDataSource,
		episode.NewEpisodeDataSource,
		audiobook.NewAudiobookDataSource,
		me.NewMeDataSource,
		me_audiobook.NewMeAudiobooksDataSource,
		me_playlist.NewMePlaylistsDataSource,
		me_album.NewMeAlbumsDataSource,
		me_track.NewMeTracksDataSource,
		me_episode.NewMeEpisodesDataSource,
		me_show.NewMeShowsDataSource,
		me_player_queue.NewMePlayerQueueDataSource,
		chapter.NewChapterDataSource,
		track.NewTrackDataSource,
		search.NewSearchDataSource,
		playlist.NewPlaylistDataSource,
		playlist_track.NewPlaylistTracksDataSource,
		user_playlist.NewUserPlaylistsDataSource,
		browse_category.NewBrowseCategoryDataSource,
		audio_feature.NewAudioFeatureDataSource,
		audio_analysis.NewAudioAnalysisDataSource,
		recommendation.NewRecommendationDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SpottedProvider{
			version: version,
		}
	}
}
