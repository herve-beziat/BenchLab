# Conditions de test — BenchLab
 
Ce document décrit l'environnement exact dans lequel les benchmarks ont été
exécutés. Il est indispensable pour reproduire les résultats ou interpréter
les écarts avec d'autres machines.
 
---
 
## Machine utilisée
 
| Paramètre | Valeur |
|-----------|--------|
| OS | Linux Mint 22 |
| CPU | Intel Core i7-10750H @ 2.60GHz (12 cœurs) |
| Mémoire RAM | 16 GB |
| Stockage | SSD NVMe |
 
---
 
## Stack technique
 
| Composant | Version |
|-----------|---------|
| Go | 1.26.3 |
| Gin | v1.12.0 |
| grpc-go | v1.81.0 |
| PostgreSQL | 16 (local) |
| k6 | v2.0.0-rc1 |
| ghz | v0.121.0 |
 
---
 
## Configuration des services
 
- Les deux services tournent en local sur la même machine
- Aucune limite de ressources appliquée (pas de cgroups, pas de Docker)
- Base de données PostgreSQL locale, pas de réseau entre services et BDD
- Les services sont lancés avec `go run` (pas de binaire optimisé)
 
> Note : utiliser `go run` au lieu d'un binaire compilé introduit un léger
> overhead au démarrage mais n'impacte pas les performances en régime établi.
 
---
 
## Conditions d'exécution
 
- Les benchmarks REST et gRPC sont lancés **séparément** (jamais en parallèle)
- Un capteur de référence est inséré en base avant les scénarios de lecture
- Aucun autre processus intensif ne tourne pendant les tests
- Les services sont redémarrés entre chaque scénario pour éviter les effets
  de cache ou de connexions résiduelles
 
---
 
## Commandes exactes pour reproduire
 
### Prérequis
 
```bash
# 1. Cloner le projet
git clone https://github.com/herve-beziat/BenchLab.git
cd BenchLab
 
# 2. Configurer l'environnement
cp .env.example .env
# Remplir les variables DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
 
# 3. Appliquer la migration
psql -U postgres -d benchlab -f internal/store/migrations/001_create_sensors_table.sql
 
# 4. Démarrer les services (deux terminaux séparés)
go run ./rest-service/.   # port 8080
go run ./grpc-service/.   # port 50051
 
# 5. Insérer un capteur de référence et noter l'UUID retourné
curl -s -X POST http://localhost:8080/sensors \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Turbine-A3-Temp",
    "type": "TEMPERATURE",
    "location": "Bâtiment C - Salle 12",
    "unit": "°C",
    "status": "ACTIVE",
    "last_value": 72.4
  }'
```
 
### Scénario A — Lecture unitaire
 
```bash
# REST (1000 req, 10 VU)
SENSOR_ID=<uuid> k6 run \
  --out json=benchmark/results/scenario-a-rest.json \
  benchmark/scripts/scenario-a/rest.js
 
# gRPC (1000 req, 10 VU)
SENSOR_ID=<uuid> bash benchmark/scripts/scenario-a/grpc.sh
```
 
### Scénario B — Écriture
 
```bash
# REST (500 req, 5 VU)
k6 run \
  --out json=benchmark/results/scenario-b-rest.json \
  benchmark/scripts/scenario-b/rest.js
 
# gRPC (500 req, 5 VU)
bash benchmark/scripts/scenario-b/grpc.sh
```
 
### Scénario C — Charge progressive
 
```bash
# REST (10 → 100 VU sur 90s)
SENSOR_ID=<uuid> k6 run \
  --out json=benchmark/results/scenario-c-rest.json \
  benchmark/scripts/scenario-c/rest.js
 
# gRPC (3 paliers : 10, 50, 100 VU × 30s)
SENSOR_ID=<uuid> bash benchmark/scripts/scenario-c/grpc.sh
```
 
---
 
## Note sur le scénario C gRPC
 
ghz v0.121.0 présente un bug avec la combinaison `--concurrency-schedule=step`
et `--duration` (panic: send on closed channel). Les paliers ont donc été
exécutés en 3 appels distincts, produisant 3 fichiers JSON séparés :
 
- `scenario-c-grpc-10vu.json` — palier à 10 VU
- `scenario-c-grpc-50vu.json` — palier à 50 VU
- `scenario-c-grpc-100vu.json` — palier à 100 VU
 
Les résultats restent comparables avec les données k6 palier par palier.