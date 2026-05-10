#!/bin/bash
 
# Scénario A — Lecture unitaire gRPC
#
# Objectif : mesurer la latence brute d'un GetSensor
#
# Paramètres :
#   - 1000 requêtes au total
#   - 10 connexions concurrentes
#
# Lancer le script :
#   SENSOR_ID=<uuid> bash benchmark/scripts/scenario-a/grpc.sh
#
# Prérequis :
#   - Le service gRPC doit tourner sur localhost:50051
#   - Remplacer <uuid> par l'ID d'un capteur existant en base
#   - Lancer depuis la racine du projet BenchLab/
 
if [ -z "$SENSOR_ID" ]; then
  echo "Erreur : SENSOR_ID est requis"
  echo "Usage : SENSOR_ID=<uuid> bash benchmark/scripts/scenario-a/grpc.sh"
  exit 1
fi
 
# Chemins absolus construits depuis le dossier courant (racine du projet)
# Ce script doit toujours être lancé depuis la racine de BenchLab/
PROTO_PATH="$PWD/grpc-service/proto/sensor.proto"
OUTPUT_FILE="$PWD/benchmark/results/scenario-a-grpc.json"
 
echo "Lancement du scénario A gRPC..."
echo "SENSOR_ID : $SENSOR_ID"
echo "Résultat  : $OUTPUT_FILE"
echo ""
 
ghz \
  --proto "$PROTO_PATH" \
  --call sensor.SensorService.GetSensor \
  --data "{\"id\": \"$SENSOR_ID\"}" \
  --total 1000 \
  --concurrency 10 \
  --insecure \
  --output "$OUTPUT_FILE" \
  --format json \
  localhost:50051
 
echo ""
echo "Scénario A gRPC terminé. Résultats dans : $OUTPUT_FILE"