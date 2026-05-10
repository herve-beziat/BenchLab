/**
 * Scénario C — Charge progressive REST
 *
 * Objectif : observer le comportement sous charge croissante
 *            sur l'endpoint de lecture (le plus sollicité en production)
 *
 * Paramètres :
 *   - Montée de 10 à 100 VU en 3 paliers
 *   - Palier 1 : 10 VU pendant 30s
 *   - Palier 2 : 50 VU pendant 30s
 *   - Palier 3 : 100 VU pendant 30s
 *
 * Lancer le script :
 *   SENSOR_ID=<uuid> k6 run --out json=../../../benchmark/results/scenario-c-rest.json rest.js
 *
 * Prérequis :
 *   - Le service REST doit tourner sur localhost:8080
 *   - Remplacer <uuid> par l'ID d'un capteur existant en base
 */
 
import http from "k6/http";
import { check } from "k6";
 
export const options = {
  scenarios: {
    charge_progressive: {
      executor: "ramping-vus", // fait varier le nombre de VU dans le temps
      startVUs: 0,
      stages: [
        { duration: "30s", target: 10 },  // montée à 10 VU en 30s
        { duration: "30s", target: 50 },  // montée à 50 VU en 30s
        { duration: "30s", target: 100 }, // montée à 100 VU en 30s
        { duration: "10s", target: 0 },   // descente progressive (clean shutdown)
      ],
    },
  },
  thresholds: {
    http_req_duration: ["p(95)<2000"], // on tolère plus de latence sous forte charge
    http_req_failed: ["rate<0.05"],    // moins de 5% d'erreurs
  },
};
 
const BASE_URL = "http://localhost:8080";
const SENSOR_ID = __ENV.SENSOR_ID;
 
export default function () {
  const res = http.get(`${BASE_URL}/sensors/${SENSOR_ID}`);
 
  check(res, {
    "status est 200": (r) => r.status === 200,
  });
}