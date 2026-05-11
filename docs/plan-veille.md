# Plan de veille — BenchLab
 
## 1. Contexte et objectif
 
Ce document présente l'organisation de la veille technologique et réglementaire
conduite dans le cadre du projet BenchLab pour SignalWatch. L'objectif est de
produire une recommandation argumentée sur le choix du protocole de communication
inter-services (REST vs gRPC) en s'appuyant sur des données mesurées et des
sources fiables.
 
---
 
## 2. Répartition thématique de la veille
 
Le projet étant conduit en autonomie, la veille a été organisée par thèmes
successifs sur une semaine, en alternant recherche documentaire et implémentation
pratique.
 
| Thème | Sources principales | Livrable associé |
|-------|---------------------|-----------------|
| Protocoles REST & gRPC | grpc.io, protobuf.dev, Google Cloud Blog | Sections B.a et B.b du rapport |
| GraphQL & Message Brokers | graphql.org, kafka.apache.org, rabbitmq.com | Sections B.c et B.d du rapport |
| Outils de benchmark | k6.io/docs, ghz.sh/docs | Scripts benchmark |
| Éco-conception numérique | RGESN, greenit.fr | Section D du rapport |
| RGPD & réglementation IoT | cnil.fr, legifrance.gouv.fr, NIS2 | Section E du rapport |
 
---
 
## 3. Sources consultées par thème
 
### Protocoles & Architecture
 
- **grpc.io** — documentation officielle gRPC et blog engineering
- **protobuf.dev** — documentation Protocol Buffers
- **Google Cloud Blog** — article de lancement gRPC 1.0 (2016)
- **Google Developers Blog** — annonce open source gRPC (2015)
- **gin-gonic.com** — documentation Gin framework
### Benchmark & Performance
 
- **k6.io/docs** — documentation k6, options et métriques
- **ghz.sh/docs** — documentation ghz, options de concurrence
- **Medium — Ian Gorton** — "Scaling up REST versus gRPC Benchmark Tests" (2023)
- **grpc.io/docs/guides/benchmarking** — benchmarks officiels gRPC
### Éco-conception
 
- **RGESN** (Référentiel Général d'Écoconception des Services Numériques) — greenit.fr
- **CNCF** — données sur la réduction de payload Protobuf vs JSON
### Réglementation
 
- **cnil.fr** — obligations RGPD, traitement des données IoT
- **legifrance.gouv.fr** — directive NIS2 transposée en droit français
- **europa.eu** — texte officiel RGPD (Règlement 2016/679)
---
 
## 4. Critères de fiabilité des sources
 
Pour valider la pertinence d'une source, trois critères ont été appliqués :
 
**Origine** — priorité aux documentations officielles (grpc.io, protobuf.dev,
cnil.fr) et aux auteurs identifiés avec une expertise reconnue (ingénieurs Google,
CNCF).
 
**Date** — privilégier les sources postérieures à 2020 pour les aspects techniques,
les textes officiels pour la réglementation. Toute source antérieure à 2020 a été
recoupée avec une source plus récente.
 
**Cohérence avec la pratique** — toute affirmation théorique qui contredisait les
mesures obtenues lors du benchmark a été signalée et analysée dans le rapport
(cf. section C.c).
 
---
 
## 5. Planning de la semaine
 
| Jour | Activité |
|------|----------|
| Lundi | Lecture documentation REST, gRPC, Protobuf — mise en place de l'environnement Go |
| Mardi | Implémentation service REST (Gin) + package store PostgreSQL — veille benchmark tools |
| Mercredi | Implémentation service gRPC — définition du .proto — veille GraphQL et Kafka |
| Jeudi | Scripts de benchmark k6 et ghz — exécution des 3 scénarios — mesures CPU/RAM |
| Vendredi | Analyse des résultats — veille RGPD et éco-conception — rédaction du rapport |
 
---
 
## 6. Format de restitution
 
Projet conduit en autonomie — pas de point quotidien en groupe. La restitution
prend la forme :
 
- D'un rapport de veille structuré (8-10 pages)
- D'un support de présentation (11 slides)
- D'un dépôt GitHub documenté avec résultats bruts et scripts reproductibles
---
 
## 7. RACI simplifié
 
Projet individuel — les rôles sont portés par une seule personne. Le RACI reflète
néanmoins la répartition des responsabilités par livrable.
 
| Livrable | Responsible | Accountable | Consulted | Informed |
|----------|-------------|-------------|-----------|----------|
| Service REST (Gin) | Hervé Béziat | Hervé Béziat | Documentation Gin | Jury |
| Service gRPC (grpc-go) | Hervé Béziat | Hervé Béziat | Documentation gRPC | Jury |
| Scripts benchmark (k6 + ghz) | Hervé Béziat | Hervé Béziat | Documentation k6, ghz | Jury |
| Résultats JSON + analyse | Hervé Béziat | Hervé Béziat | Benchmarks publiés | Jury |
| Rapport de veille | Hervé Béziat | Hervé Béziat | Sources documentaires | Jury |
| Support de présentation | Hervé Béziat | Hervé Béziat | — | Jury |
| Plan de veille + RACI | Hervé Béziat | Hervé Béziat | — | Jury |
 
---
 
## 8. Lexique
 
**RACI** — matrice de responsabilités : Responsible (réalise), Accountable
(valide), Consulted (consulté), Informed (informé).
 
**Veille technologique** — processus de surveillance de l'environnement technique
pour anticiper les évolutions et prendre des décisions éclairées.
 
**Benchmark** — mesure comparative des performances de deux systèmes dans des
conditions identiques et reproductibles.
