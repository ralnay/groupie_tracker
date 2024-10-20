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
        tmpl, err := template.ParseFiles("template/homepage.html")
        if err != nil {
            http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
            return
        }
        if err := tmpl.Execute(w, nil); err != nil {
            http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
            return
        }


    } else if path == "/main/" {
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

    } else if path !="" {
        artistID, err := strconv.ParseUint(path, 10, 32)
        if err != nil {
            http.Error(w, "404 - Page Not Found", http.StatusNotFound)
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

        if selectedArtist.Name == "" {
            http.Error(w, "404 - Page Not Found", http.StatusNotFound)
            return
        }

        relations := RelationData()

        artistData := Data{
            A: selectedArtist,
            R: relations[selectedArtist.Id-1],
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
