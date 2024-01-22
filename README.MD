# Groupie Trackers Geolocalisation

## Introduction

Groupie Trackers Géolocalisation est un projet visant à utiliser une API donnée pour manipuler les données qu'elle contient et créer un site affichant ces informations. Le projet intègre une barre de recherche permettant aux utilisateurs de rechercher des informations selon divers critères, notamment l'artiste, le nom, les membres, les emplacements, la date du premier album et la date de création. De plus, une fonctionnalité de géolocalisation est mise en œuvre pour afficher les emplacements sur une carte à l'aide d'une API de MapQuest. Les principaux contributeurs à ce projet sont Cherif Nzale et Mbaye Diao.

# avec le Search Bar il est possible de rechercher par rapport a 
- artist
- name
- members
- locations
- first album date
- creation date.

#  avec geocalisation il est possible d'afficher les localisation dans une map grace a une api provenant de mapQuest
Les principaux contributeurs à ce projet sont Cherif Nzale et Mbaye Diao.

## API Structure

L'API fournie est composée de quatre parties :

1. **Artists :** Contient des informations sur des groupes et artistes tels que leur nom, image, année de début d'activité, date de leur premier album et les membres.
   
2. **Locations :** Contient des informations sur les emplacements de leurs concerts à venir et/ou passés.

3. **Dates :** Contient des informations sur les dates de leurs concerts à venir et/ou passés.

4. **Relation :** Établit le lien entre les autres parties (Artists, Dates, Locations).

## Fonctionnalités

Le projet consiste à construire un site convivial où les informations des groupes peuvent être affichées à travers différentes visualisations de données.

Le projet se concentre également sur la création d'événements/actions et leur visualisation. Un exemple d'événement/action est un appel client au serveur (client-serveur) qui doit déclencher une action. Cette action doit communiquer avec le serveur pour recevoir des informations ([request-response](https://en.wikipedia.org/wiki/Request%E2%80%93response)).

## Instructions

- Le backend doit être écrit en Go.
- Le site et le serveur ne doivent jamais planter.
- Toutes les pages doivent fonctionner correctement, et vous devez prendre en charge toutes les erreurs.
- Le code doit respecter les bonnes pratiques.
- Il est recommandé d'avoir des fichiers de test pour les tests unitaires.

## Packages Autorisés

Seuls les packages standard de Go sont autorisés.

## Utilisation

Vous pouvez voir un exemple d'API RESTful [ici](https://groupietrackers.herokuapp.com/api/artists).

Ce projet vous aidera à apprendre sur :

- La manipulation et le stockage des données.
- Les fichiers et le format JSON.
- HTML.
- Manipulation of Maps API
- Geolocation, geocoding, etc
- La création et l'affichage d'événements.
- Le client-serveur.

## Auteur
 - Cherif Nzale
 - Mbaye Diao 