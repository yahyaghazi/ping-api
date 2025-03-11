# API Ping Minimaliste

## Description
Une API minimaliste qui retourne les headers HTTP d'une requête GET sur `/ping`.  
Toutes les autres requêtes retournent une **404**.

## Lancement

1. Compiler et exécuter :
````sh
   go run main.go
````

## Requêtes

### GET /ping

#### Requête avec un header invalide

````sh
curl -X GET http://localhost:8080/ping
````

#### Résultat

````json
{
  "User-Agent": ["curl/7.79.1"],
  "Accept": ["*/*"]
}
````