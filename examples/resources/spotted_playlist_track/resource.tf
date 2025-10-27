resource "spotted_playlist_track" "example_playlist_track" {
  playlist_id = "3cEYpjA9oz9GiPac4AsH4n"
  insert_before = 3
  range_length = 2
  range_start = 1
  snapshot_id = "snapshot_id"
  uris = ["string"]
}
