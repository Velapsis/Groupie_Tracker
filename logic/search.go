package groupie

import (
	"strconv"
	"strings"
)

// SearchArtistsWithFilters filtre la liste des artistes selon les critères spécifiés
// Prend en entrée la liste complète des artistes, une requête de recherche et une map de filtres
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

// matchesFilters vérifie si un artiste correspond aux critères de recherche et aux filtres
func matchesFilters(artist Artist, query string, filters map[string]string) bool {
	// Recherche textuelle dans le nom, les membres et les lieux de concert
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

	// Filtre sur la plage de dates de création
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

	// Filtre sur la plage de dates du premier album
	if min := filters["albumDateMin"]; min != "" {
		albumDate := strings.Split(artist.FirstAlbum, ".")[2]
		albumYear, _ := strconv.Atoi(albumDate)
		minYear, _ := strconv.Atoi(min)
		if albumYear < minYear {
			return false
		}
	}
	if max := filters["albumDateMax"]; max != "" {
		albumDate := strings.Split(artist.FirstAlbum, ".")[2]
		albumYear, _ := strconv.Atoi(albumDate)
		maxYear, _ := strconv.Atoi(max)
		if albumYear > maxYear {
			return false
		}
	}

	// Filtre sur les lieux de concert
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

	// Filtre sur le nombre de membres
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