# Music Microservices Application
This project encompasses the development of three microservices written in Go: Tracks, Search, and CoolTown. Each microservice serves a distinct purpose within the domain of music management and retrieval.

Tracks Microservice: This microservice facilitates the creation, listing, reading, and deletion of music tracks. It accepts requests to create tracks with associated audio files, lists existing tracks, reads specific tracks, and deletes tracks from the database.

Search Microservice: Designed akin to Hum-to-Search functionality, this microservice offers music recognition services. It accepts audio fragments, recognises them, and returns corresponding track IDs if successful.

CoolTown Microservice: Similar to Bixby, this microservice enables device integration for music retrieval. It accepts audio fragments as input and returns full music tracks based on the provided fragments.

Each microservice is deployed on a specific port, with the Tracks service listening on port 3000, the Search service on port 3001, and the CoolTown service on port 3002. The project emphasises modularity, scalability, and efficient music management through microservices architecture.

## File Structure
- `music/`
  - `music/tracks`
    - `music/tracks/repository`
      - `music/tracks/repository/model.go`
    - `music/tracks/resources`
      - `music/tracks/resources.go`
  - `music/search`
    - `music/search/repository`
      - `music/search/repository/model.go`
    - `music/search/resources`
      - `music/search/resources/resources.go`
  - `music/cooltown`
    - `music/cooltown/repository`
      - `music/cooltown/repository/model.go`
    - `music/cooltown/resources`
      - `music/cooltown/resources/resources.go`
- `LICENSE`
- `README.md`

## License
This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for the full license text.