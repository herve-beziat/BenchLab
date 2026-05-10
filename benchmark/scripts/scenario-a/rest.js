/**
 * Scénario A — Lecture unitaire REST
 *
 * Objectif : mesurer la latence brute d'un GET /sensors/:id
 *
 * Paramètres :
 *   - 1000 requêtes au total
 *   - 10 VU (connexions concurrentes)
 *
 * Lancer le script :
 *   SENSOR_ID=<uuid> k6 run --out json=../../results/scenario-a-rest.json scenario-a-rest.js
 *
 * Prérequis :
 *   - Le service REST doit tourner sur localhost:8080
 *   - Remplacer <uuid> par l'ID d'un capteur existant en base
 */
 
import http from "k6/http";
import { check } from "k6";
 
// Options du scénario
export const options = {
  scenarios: {
    lecture_unitaire: {
      executor: "shared-iterations", // répartit les 1000 requêtes entre les 10 VU
      vus: 10,
      iterations: 1000,
    },
  },
  // Seuils d'alerte — le benchmark échoue si ces valeurs sont dépassées
  thresholds: {
    http_req_duration: ["p(95)<500"], // p95 doit rester sous 500ms
    http_req_failed: ["rate<0.01"],   // moins de 1% d'erreurs
  },
};
 
// URL de base du service REST
const BASE_URL = "http://localhost:8080";
 
// L'UUID du capteur est passé en variable d'environnement
// Exemple : SENSOR_ID=abc-123 k6 run scenario-a-rest.js
const SENSOR_ID = __ENV.SENSOR_ID;
 
export default function () {
  // Requête GET /sensors/:id
  const res = http.get(`${BASE_URL}/sensors/${SENSOR_ID}`);
 
  // Vérifications sur la réponse
  check(res, {
    "status est 200": (r) => r.status === 200,
    "le body contient un id": (r) => JSON.parse(r.body).id !== undefined,
  });
}