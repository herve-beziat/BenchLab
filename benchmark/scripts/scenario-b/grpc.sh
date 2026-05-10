#!/bin/bash
 
# Scénario B — Écriture gRPC
#
# Objectif : mesurer le coût d'un CreateSensor
#
# Paramètres :
#   - 500 requêtes au total
#   - 5 connexions concurrentes
#
# Lancer le script :
#   bash benchmark/scripts/scenario-b/grpc.sh
#
# Prérequis :
#   - Le service gRPC doit tourner sur localhost:50051
#   - Lancer depuis la racine du projet BenchLab/
 
# Chemins absolus construits depuis le dossier courant (racine du projet)
PROTO_PATH="$PWD/grpc-service/proto/sensor.proto"
OUTPUT_FILE="$PWD/benchmark/results/scenario-b-grpc.json"
 
echo "Lancement du scénario B gRPC..."
echo "Résultat  : $OUTPUT_FILE"
echo ""
 
# Payload identique pour REST et gRPC — même capteur créé des deux côtés
ghz \
  --proto "$PROTO_PATH" \
  --call sensor.SensorService.CreateSensor \
  --data '{
    "name": "Bench-Sensor-Write",
    "type": 1,
    "location": "Zone de test",
    "unit": "bar",
    "status": 0,
    "last_value": 1.013
  }' \
  --total 500 \
  --concurrency 5 \
  --insecure \
  --output "$OUTPUT_FILE" \
  --format json \
  localhost:50051
 
echo ""
echo "Scénario B gRPC terminé. Résultats dans : $OUTPUT_FILE"