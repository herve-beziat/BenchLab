# Scripts de Benchmark — BenchLab
 
Ce dossier contient les scripts de benchmark pour comparer les performances
du service REST (port 8080) et du service gRPC (port 50051).
 
---
 
## Outils utilisés
 
### k6 — Benchmark REST
 
k6 est un outil de test de charge open-source. Les scripts sont écrits en
JavaScript et les résultats sont exportables en JSON.
Il calcule automatiquement les métriques p50, p95, p99.
 
**Installation (Linux / Linux Mint) :**
```bash
sudo gpg -k
sudo gpg --no-default-keyring \
  --keyring /usr/share/keyrings/k6-archive-keyring.gpg \
  --keyserver hkp://keyserver.ubuntu.com:80 \
  --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
 
echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] \
  https://dl.k6.io/deb stable main" \
  | sudo tee /etc/apt/sources.list.d/k6.list
 
sudo apt-get update && sudo apt-get install k6
```
 
**Vérification :**
```bash
k6 version
```
 
---
 
### ghz — Benchmark gRPC
 
ghz est l'équivalent de k6 pour gRPC. Il envoie des requêtes protobuf
directement au service et génère les métriques p50, p95, p99 en JSON.
 
**Installation (Linux / Linux Mint) :**
```bash
# Remplace X.X.X par la dernière version disponible sur :
# https://github.com/bojand/ghz/releases
 
wget https://github.com/bojand/ghz/releases/download/vX.X.X/ghz-linux-x86_64.tar.gz
tar -xzf ghz-linux-x86_64.tar.gz
sudo mv ghz /usr/local/bin/
rm ghz-linux-x86_64.tar.gz
```
 
**Vérification :**
```bash
ghz --version
```
 
---
 
## Prérequis avant de lancer les benchmarks
 
1. Copier et remplir le fichier d'environnement :
```bash
cp .env.example .env
```
 
2. Démarrer les deux services dans deux terminaux séparés :
```bash
# Terminal 1 — REST (port 8080)
cd rest-service && go run .
 
# Terminal 2 — gRPC (port 50051)
cd grpc-service && go run .
```
 
3. Insérer un capteur de référence en base (nécessaire pour les scénarios
   de lecture) et noter l'UUID retourné :
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
  }' | jq .
```
 
L'UUID retourné dans le champ `id` sera utilisé dans les scénarios A et C.
 
---
 
## Scénarios disponibles
 
| Fichier                | Scénario | Protocole | Description                        |
|------------------------|----------|-----------|------------------------------------|
| `scenario-a-rest.js`   | A        | REST      | Lecture unitaire — 1000 req, 10 VU |
| `scenario-a-grpc.sh`   | A        | gRPC      | Lecture unitaire — 1000 req, 10 VU |
| `scenario-b-rest.js`   | B        | REST      | Écriture — 500 req, 5 VU           |
| `scenario-b-grpc.sh`   | B        | gRPC      | Écriture — 500 req, 5 VU           |
| `scenario-c-rest.js`   | C        | REST      | Charge progressive 10 → 100 VU     |
| `scenario-c-grpc.sh`   | C        | gRPC      | Charge progressive 10 → 100 VU     |
 
---
 
## Résultats
 
Les résultats bruts sont exportés en JSON dans `benchmark/results/`.
Voir ce dossier pour les sorties de chaque scénario.