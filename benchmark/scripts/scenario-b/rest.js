/**
 * Scénario B — Écriture REST
 *
 * Objectif : mesurer le coût d'un POST /sensors
 *
 * Paramètres :
 *   - 500 requêtes au total
 *   - 5 VU (connexions concurrentes)
 *
 * Lancer le script :
 *   k6 run --out json=../../../benchmark/results/scenario-b-rest.json rest.js
 *
 * Prérequis :
 *   - Le service REST doit tourner sur localhost:8080
 */
 
import http from "k6/http";
import { check } from "k6";
 
export const options = {
  scenarios: {
    ecriture: {
      executor: "shared-iterations",
      vus: 5,
      iterations: 500,
    },
  },
  thresholds: {
    http_req_duration: ["p(95)<1000"], // écriture tolère un peu plus de latence
    http_req_failed: ["rate<0.01"],
  },
};
 
const BASE_URL = "http://localhost:8080";
 
// Payload identique pour REST et gRPC — même capteur créé des deux côtés
const PAYLOAD = JSON.stringify({
  name: "Bench-Sensor-Write",
  type: "PRESSURE",
  location: "Zone de test",
  unit: "bar",
  status: "ACTIVE",
  last_value: 1.013,
});
 
export default function () {
  const res = http.post(`${BASE_URL}/sensors`, PAYLOAD, {
    headers: { "Content-Type": "application/json" },
  });
 
  check(res, {
    "status est 201": (r) => r.status === 201,
    "le body contient un id": (r) => JSON.parse(r.body).id !== undefined,
  });
}