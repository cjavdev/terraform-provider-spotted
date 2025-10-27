data "spotted_search" "example_search" {
  q = "remaster%20track:Doxy%20artist:Miles%20Davis"
  type = ["album"]
  include_external = "audio"
  limit = 10
  market = "ES"
  offset = 5
}
