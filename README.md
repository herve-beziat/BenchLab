# BenchLab
 
Veille stratégique & benchmark technique — REST vs gRPC
 
Projet réalisé pour SignalWatch, une startup IoT qui collecte des données
de capteurs industriels à raison de 10 000 événements par minute.
 
## Structure du projet
 
```
BenchLab/
├── rest-service/        # Micro-service REST (Gin)
├── grpc-service/        # Micro-service gRPC (grpc-go)
│   └── proto/           # Fichiers .proto
├── internal/
│   └── store/           # Package partagé — connexion PostgreSQL
├── benchmark/
│   ├── scripts/         # Scripts k6 / ghz
│   └── results/         # Résultats bruts JSON + analyse
├── docs/                # Rapport, présentation, roadmaps
├── .env.example         # Template des variables d'environnement
└── README.md
```
 
## Prérequis
 
- Go 1.26.3+
- PostgreSQL 16+
- k6 v2.0.0+ (benchmark REST)
- ghz v0.121.0+ (benchmark gRPC)
 
**Installation de k6 (Linux) :**
```bash
sudo gpg --no-default-keyring \
  --keyring /usr/share/keyrings/k6-archive-keyring.gpg \
  --keyserver hkp://keyserver.ubuntu.com:80 \
  --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
 
echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" \
  | sudo tee /etc/apt/sources.list.d/k6.list
 
sudo apt-get update && sudo apt-get install k6
```
 
**Installation de ghz :**
```bash
go install github.com/bojand/ghz/cmd/ghz@latest
```
 
---
 
## Installation
 
```bash
# Cloner le projet
git clone https://github.com/herve-beziat/BenchLab.git
cd BenchLab
 
# Copier et remplir les variables d'environnement
cp .env.example .env
```
 
Remplir le fichier `.env` avec les informations de connexion PostgreSQL :
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=benchlab
```
 
Appliquer la migration pour créer la table `sensors` :
```bash
psql -U postgres -d benchlab -f internal/store/migrations/001_create_sensors_table.sql
```
 
---
 
## Lancer les services
 
Les deux services doivent être démarrés dans deux terminaux séparés,
**depuis la racine du projet** :
 
```bash
# Terminal 1 — Service REST (port 8080)
go run ./rest-service/.
 
# Terminal 2 — Service gRPC (port 50051)
go run ./grpc-service/.
```
 
Vérifier que les services sont opérationnels :
```bash
# REST
curl http://localhost:8080/health
 
# gRPC
grpcurl -plaintext localhost:50051 list
```
 
---
 
## Lancer les benchmarks
 
### Prérequis
 
Insérer un capteur de référence en base et noter l'UUID retourné :
```bash
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
 
### Scénario A — Lecture unitaire (1000 req, 10 VU)
 
```bash
# REST
SENSOR_ID=<uuid> k6 run \
  --out json=benchmark/results/scenario-a-rest.json \
  benchmark/scripts/scenario-a/rest.js
 
# gRPC
SENSOR_ID=<uuid> bash benchmark/scripts/scenario-a/grpc.sh
```
 
### Scénario B — Écriture (500 req, 5 VU)
 
```bash
# REST
k6 run \
  --out json=benchmark/results/scenario-b-rest.json \
  benchmark/scripts/scenario-b/rest.js
 
# gRPC
bash benchmark/scripts/scenario-b/grpc.sh
```
 
### Scénario C — Charge progressive (10 → 100 VU)
 
```bash
# REST
SENSOR_ID=<uuid> k6 run \
  --out json=benchmark/results/scenario-c-rest.json \
  benchmark/scripts/scenario-c/rest.js
 
# gRPC (3 paliers de 30s : 10, 50, 100 VU)
SENSOR_ID=<uuid> bash benchmark/scripts/scenario-c/grpc.sh
```
 
> Les résultats sont exportés en JSON dans `benchmark/results/`.
> Consulter `benchmark/results/ANALYSIS.md` pour la synthèse des métriques.
 
---
 
## Résultats
 
| Critère | REST | gRPC |
|---------|------|------|
| Latence p50 (lecture) | 6.41ms | 4.59ms |
| Latence p95 (lecture) | 58.8ms | 41.9ms |
| Throughput (lecture) | 548 req/s | 723 req/s |
| Taille payload | ~371 octets | ~100 octets (Protobuf) |
 
Analyse complète disponible dans `benchmark/results/ANALYSIS.md`.
 