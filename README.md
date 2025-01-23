# KAIZEN - Groupie Tracker

Un site web élégant avec une esthétique japonaise zen qui affiche et explore les données d'artistes musicaux via l'API Groupie Tracker.

![Theme](static/img/fuji.gif)

## Fonctionnalités

- Recherche dynamique en temps réel
- Filtres avancés :
  - Nombre de membres
  - Date de création
  - Date du premier album
  - Localisation des concerts
- Visualisation des concerts sur une carte interactive
- Interface responsive avec animations sakura
- Thème japonais zen avec palette rose pâle/noir

## Technologies

- Backend : Go
- Frontend : HTML/CSS/JavaScript
- Cartes : Leaflet.js
- API : [https://groupietrackers.herokuapp.com/api]

## Installation

```bash
# Cloner le repository
git clone [votre-repo-url]

# Aller dans le dossier
cd groupie-tracker

# Lancer le serveur
go run main.go
```

Le site sera accessible sur `http://localhost:8000`

## Structure

```
groupie-tracker/
├── logic/
│   ├── api.go        # Gestion API
│   ├── search.go     # Logique de recherche
│   └── struct.go     # Structures de données
├── static/
│   ├── css/          # Styles
│   ├── js/           # Scripts
│   └── img/          # Images
├── templates/        # Pages HTML
└── server/          # Serveur web
```

## API Endpoints

- `/` : Page d'accueil
- `/index` : Liste des artistes
- `/artist?id=X` : Détails d'un artiste
- `/search` : Recherche et filtrage

## Développé par Kaizen

- [Romain VICENTE](https://github.com/Velapsis)
- [Mahan MIR](https://github.com/Nothypaa)

Projet réalisé dans le cadre du cursus Ynov Montpellier.