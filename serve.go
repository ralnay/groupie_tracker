package piscine

import (
    "html/template"
    "net/http"
    "strconv"
    "strings"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/artist/")

    if r.URL.Path == "/" {
        // Serve the main page with a list of artists
        artists := ArtistData()
        tmpl, err := template.ParseFiles("template/mainpage.html")
        if err != nil {
            http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
            return
        }
        if err := tmpl.Execute(w, artists); err != nil {
            http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
            return
        }
    } else if path != "" {
        // Convert the artist ID (from string to uint)
        artistID, err := strconv.ParseUint(path, 10, 32)
        if err != nil {
            http.Error(w, "400 - Invalid Artist ID", http.StatusBadRequest)
            return
        }

        artists := ArtistData()
        var selectedArtist Artist
        for _, artist := range artists {
            if artist.Id == uint(artistID) {
                selectedArtist = artist
                break
            }
        }

        // Handle artist not found
        if selectedArtist.Name == "" {
            http.Error(w, "404 - Artist Not Found", http.StatusNotFound)
            return
        }

        // Load relation data (concerts and dates)
        relations := RelationData()

        // Prepare the data to pass to the template
        artistData := Data{
            A: selectedArtist,
            R: relations[selectedArtist.Id-1], // Assuming artist Id matches relation index
        }

        tmpl, err := template.ParseFiles("template/artistpage.html")
        if err != nil {
            http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
            return
        }
        if err := tmpl.Execute(w, artistData); err != nil {
            http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
        }
    } else {
        http.Error(w, "404 - Page Not Found", http.StatusNotFound)
    }
}
