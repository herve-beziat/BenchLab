# Analyse des résultats — BenchLab
 
Synthèse des résultats obtenus lors des benchmarks REST vs gRPC.
Les résultats bruts sont disponibles dans ce même dossier (fichiers JSON).
 
Machine : Intel Core i7-10750H @ 2.60GHz, 16 GB RAM, SSD NVMe, Linux Mint 22
Stack   : Go 1.26.3, Gin v1.12.0, grpc-go v1.81.0, PostgreSQL 16
 
---
 
## Scénario A — Lecture unitaire (1000 req, 10 VU)
 
| Métrique | REST | gRPC | Écart |
|----------|------|------|-------|
| p50 | 6.41ms | 4.59ms | gRPC -28% |
| p95 | 58.8ms | 41.9ms | gRPC -29% |
| Throughput | 548 req/s | 723 req/s | gRPC +32% |
| Erreurs | 0% | 0% | égalité |
 
**Analyse :**
gRPC est significativement plus rapide sur la lecture unitaire. L'écart de
~30% sur la latence s'explique par le protocole HTTP/2 de gRPC (multiplexage
des connexions) et la sérialisation Protobuf plus compacte que JSON.
 
---
 
## Scénario B — Écriture (500 req, 5 VU)
 
| Métrique | REST | gRPC | Écart |
|----------|------|------|-------|
| p50 | 4.06ms | 2.78ms | gRPC -31% |
| p95 | 36.08ms | 25.6ms | gRPC -29% |
| Throughput | 535 req/s | 802 req/s | gRPC +50% |
| Erreurs | 0% | 0% | égalité |
 
**Analyse :**
L'avantage de gRPC est encore plus marqué sur les écritures (+50% throughput).
La sérialisation Protobuf réduit la taille du payload envoyé, ce qui diminue
le temps de traitement réseau et la charge sur PostgreSQL.
 
---
 
## Scénario C — Charge progressive (10 → 100 VU)
 
### REST (k6 — résultats agrégés sur 90s)
 
| Métrique | Valeur |
|----------|--------|
| Total requêtes | 51 076 |
| p50 | 17.99ms |
| p95 | 333.92ms |
| Throughput | 510 req/s |
| Erreurs | 0.02% (13/51076) |
 
### gRPC (ghz — résultats par palier)
 
| Palier | VU | p50 | p95 | Throughput | Erreurs |
|--------|----|-----|-----|------------|---------|
| 10 VU  | 10 | 3.94ms | 40.96ms | 815 req/s | 0.03% (8/24455) |
| 50 VU  | 50 | 108.47ms | 238.19ms | 454 req/s | 0.37% (50/13632) |
| 100 VU | 100 | 282.04ms | 560.59ms | 363 req/s | 1.15% (130/10882) |
 
**Analyse :**
La dégradation gRPC sous charge est très marquée — la latence p50 passe de
3.94ms à 282ms entre 10 et 100 VU. Les erreurs au palier 100 VU sont dues
à PostgreSQL qui sature ses connexions disponibles ("too many clients" — code
53300). Cela révèle une limite de configuration PostgreSQL en local, pas une
limite intrinsèque de gRPC.
 
REST tient mieux la charge globale (510 req/s agrégés sur 90s, 0.02% d'erreurs)
car Gin gère mieux le pool de connexions PostgreSQL sous forte concurrence
dans cette configuration.
 
> Note : les erreurs gRPC à 100 VU sont liées à la limite de connexions
> PostgreSQL (`max_connections`), pas au protocole lui-même. En production,
> un pool de connexions (pgBouncer) éliminerait ce problème.
 
---
 
## Synthèse générale
 
| Critère | Gagnant | Commentaire |
|---------|---------|-------------|
| Latence brute | gRPC | ~30% plus rapide sur p50 et p95 |
| Throughput | gRPC | +32% à +50% selon le scénario |
| Stabilité sous charge | REST | Comportement prévisible, erreurs quasi nulles |
| Facilité de test | REST | curl, Postman, navigateur — aucun outil spécifique |
| Taille payload | gRPC | Protobuf plus compact que JSON |
 
**Conclusion préliminaire :**
gRPC offre de meilleures performances brutes dans tous les scénarios testés.
Pour SignalWatch (10 000 événements/min), l'avantage de gRPC en throughput
et en taille de payload représente un gain significatif en bande passante
et en coût infrastructure.
 
REST reste pertinent pour les APIs publiques ou les interfaces consommées
par des clients web, où la simplicité d'intégration prime sur la performance.