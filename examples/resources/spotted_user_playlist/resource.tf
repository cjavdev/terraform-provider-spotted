resource "spotted_user_playlist" "example_user_playlist" {
  user_id = "smedjan"
  name = "New Playlist"
  collaborative = true
  description = "New playlist description"
  published = true
}
