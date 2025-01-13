package groupie

import (
    "strconv"
    "strings"
)

func SearchArtistsWithFilters(artists []Artist, query string, filters map[string]string) []Artist {
    query = strings.ToLower(query)
    var results []map[string]string

    var artiststruct []Artist = []Artist{} 

    for _, artist := range artists {
        match := false

        
        if query != "" {
            if strings.Contains(strings.ToLower(artist.Name), query) {
                match = true
            }
            for _, member := range artist.Members {
                if strings.Contains(strings.ToLower(member), query) {
                    match = true
                }
            }
            for location := range artist.Relations {
                if strings.Contains(strings.ToLower(location), query) {
                    match = true
                }
            }
            if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
                match = true
            }
            if strings.Contains(strconv.Itoa(artist.CreationDate), query) {
                match = true
            }
        }

        
        if creationDate, ok := filters["creation_date"]; ok && creationDate != "" {
            if strconv.Itoa(artist.CreationDate) != creationDate {
                continue
            }
        }
        if firstAlbumDate, ok := filters["first_album_date"]; ok && firstAlbumDate != "" {
            if !strings.Contains(artist.FirstAlbum, firstAlbumDate) {
                continue
            }
        }
        if location, ok := filters["location"]; ok && location != "" {
            found := false
            for loc := range artist.Relations {
                if strings.Contains(strings.ToLower(loc), strings.ToLower(location)) {
                    found = true
                    break
                }
            }
            if !found {
                continue
            }
        }

        
        if match {
            results = append(results, map[string]string{
                "type":          "Artist",
                "name":          artist.Name,
                "creation_date": strconv.Itoa(artist.CreationDate),
                "first_album":   artist.FirstAlbum,
            })

            artiststruct = append(artiststruct, Artist{
                Name: artist.Name,
                CreationDate: artist.CreationDate,
                FirstAlbum: artist.FirstAlbum,
                Image: artist.Image,
            })


            
        }
    }
    return artiststruct 
}
