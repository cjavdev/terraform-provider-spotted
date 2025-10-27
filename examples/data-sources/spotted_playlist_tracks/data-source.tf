data "spotted_playlist_tracks" "example_playlist_tracks" {
  playlist_id = "3cEYpjA9oz9GiPac4AsH4n"
  additional_types = "additional_types"
  fields = "items(added_by.id,track(name,href,album(name,href)))"
  market = "ES"
}
