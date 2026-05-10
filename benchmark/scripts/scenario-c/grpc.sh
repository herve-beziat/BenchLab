#!/bin/bash
 
# Scénario C — Charge progressive gRPC
#
# Objectif : observer le comportement sous charge croissante sur GetSensor
#
# Stratégie : 3 paliers successifs de 30s chacun
#   - Palier 1 : 10 VU pendant 30s
#   - Palier 2 : 50 VU pendant 30s
#   - Palier 3 : 100 VU pendant 30s
#
# Note : ghz v0.121.0 a un bug avec --concurrency-schedule=step + --duration
# combinés. On utilise donc 3 appels distincts, ce qui donne des résultats
# plus lisibles par palier pour l'analyse.
#
# Lancer le script :
#   SENSOR_ID=<uuid> bash benchmark/scripts/scenario-c/grpc.sh
#
# Prérequis :
#   - Le service gRPC doit tourner sur localhost:50051
#   - Lancer depuis la racine du projet BenchLab/
 
if [ -z "$SENSOR_ID" ]; then
  echo "Erreur : SENSOR_ID est requis"
  echo "Usage : SENSOR_ID=<uuid> bash benchmark/scripts/scenario-c/grpc.sh"
  exit 1
fi
 
PROTO_PATH="$PWD/grpc-service/proto/sensor.proto"
RESULTS_DIR="$PWD/benchmark/results"
 
echo "Lancement du scénario C gRPC — charge progressive"
echo "SENSOR_ID : $SENSOR_ID"
echo "Durée estimée : ~90 secondes"
echo ""
 
# Palier 1 — 10 VU pendant 30s
echo "Palier 1 : 10 VU pendant 30s..."
ghz \
  --proto "$PROTO_PATH" \
  --call sensor.SensorService.GetSensor \
  --data "{\"id\": \"$SENSOR_ID\"}" \
  --concurrency 10 \
  --duration 30s \
  --insecure \
  --output "$RESULTS_DIR/scenario-c-grpc-10vu.json" \
  --format json \
  localhost:50051
echo "Palier 1 terminé."
echo ""
 
# Palier 2 — 50 VU pendant 30s
echo "Palier 2 : 50 VU pendant 30s..."
ghz \
  --proto "$PROTO_PATH" \
  --call sensor.SensorService.GetSensor \
  --data "{\"id\": \"$SENSOR_ID\"}" \
  --concurrency 50 \
  --duration 30s \
  --insecure \
  --output "$RESULTS_DIR/scenario-c-grpc-50vu.json" \
  --format json \
  localhost:50051
echo "Palier 2 terminé."
echo ""
 
# Palier 3 — 100 VU pendant 30s
echo "Palier 3 : 100 VU pendant 30s..."
ghz \
  --proto "$PROTO_PATH" \
  --call sensor.SensorService.GetSensor \
  --data "{\"id\": \"$SENSOR_ID\"}" \
  --concurrency 100 \
  --duration 30s \
  --insecure \
  --output "$RESULTS_DIR/scenario-c-grpc-100vu.json" \
  --format json \
  localhost:50051
echo "Palier 3 terminé."
echo ""
 
echo "Scénario C gRPC terminé. Résultats dans :"
echo "  $RESULTS_DIR/scenario-c-grpc-10vu.json"
echo "  $RESULTS_DIR/scenario-c-grpc-50vu.json"
echo "  $RESULTS_DIR/scenario-c-grpc-100vu.json"