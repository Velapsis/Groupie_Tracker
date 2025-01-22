package groupie

import (
	"strconv"
	"strings"
)

func SearchArtistsWithFilters(artists []Artist, query string, filters map[string]string) []Artist {
	query = strings.ToLower(query)
	var results []Artist

	for _, artist := range artists {
		if !matchesFilters(artist, query, filters) {
			continue
		}
		results = append(results, artist)
	}

	return results
}

func matchesFilters(artist Artist, query string, filters map[string]string) bool {
	// Vérification de la recherche textuelle
	if query != "" {
		found := false
		if strings.Contains(strings.ToLower(artist.Name), query) {
			found = true
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				found = true
			}
		}
		for location := range artist.Relations {
			if strings.Contains(strings.ToLower(location), query) {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	// Filtres de dates de création
	if min := filters["creationDateMin"]; min != "" {
		minYear, _ := strconv.Atoi(min)
		if artist.CreationDate < minYear {
			return false
		}
	}
	if max := filters["creationDateMax"]; max != "" {
		maxYear, _ := strconv.Atoi(max)
		if artist.CreationDate > maxYear {
			return false
		}
	}

	// Filtres de dates de premier album
	if min := filters["albumDateMin"]; min != "" {
		albumYear, _ := strconv.Atoi(artist.FirstAlbum[:4])
		minYear, _ := strconv.Atoi(min)
		if albumYear < minYear {
			return false
		}
	}
	if max := filters["albumDateMax"]; max != "" {
		albumYear, _ := strconv.Atoi(artist.FirstAlbum[:4])
		maxYear, _ := strconv.Atoi(max)
		if albumYear > maxYear {
			return false
		}
	}

	// Filtre de localisation
	if location := filters["location"]; location != "" {
		found := false
		for loc := range artist.Relations {
			if strings.Contains(strings.ToLower(loc), strings.ToLower(location)) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Filtre de nombre de membres
	if members := filters["members"]; members != "" {
		memberCounts := strings.Split(members, ",")
		found := false
		numMembers := len(artist.Members)
		for _, count := range memberCounts {
			if count == "5+" && numMembers >= 5 {
				found = true
				break
			}
			countNum, _ := strconv.Atoi(count)
			if countNum == numMembers {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
