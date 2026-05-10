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
│   └── results/         # Résultats bruts JSON
├── docs/                # Rapport, présentation, roadmaps
├── .env.example         # Template des variables d'environnement
└── README.md
```

## Prérequis

- Go 1.26.3+
- PostgreSQL 15+
- k6 (benchmark REST)
- ghz (benchmark gRPC)

## Installation

```bash
# Cloner le projet
git clone https://github.com/herve-beziat/BenchLab.git
cd BenchLab

# Copier et remplir les variables d'environnement
cp .env.example .env
```

## Lancer les services

*Instructions à venir*

## Lancer les benchmarks

*Instructions à venir*