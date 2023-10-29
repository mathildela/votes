# API Vote

Une API REST pour : 
- créer un nouveau de ballot de Vote /new_ballot
- voter /vote
- récupérer le résultat des votes /result

1. Possibilité de lancer le serveur et de faire les commandes dans le navigateur. 

dans dossier launch-rsagt

2. Possibilité de lancer tous les agents de votes en même temps pour tester toutes les méthodes.

dans dossier launch-all-rest-agents

on demande le nombre de votants et leurs prefs sont générées aléatoirement
alternative c'est toujours entre 1 et #alts
si pas de gagnant, winner:0 (voir si on change)
attente

## Utilisation locale

Récupération du projet

```bash
  go install github.com/yourusername/mypackage
```

Démarrer le server

```bash
  cd []
  ./mon_programme (main juste server)
```

Démarrer le server & les agents de votes

```bash
  cd []
  ./mon_programme 2 (all_agents)
```
## Méthodes de vote implémentées

- Majorité "majority"
- Borda "borda"
- Approval "approval" (options à renseigner)
- Condorcet "condorcet"
- Copeland "copeland"
- STV "stv"

Les factories sont dans `/comsoc/utils.go` et les tests dans `/comsoc/go_tests.go`.



## Authors

- Mathilde Lange
- Solenn Lenoir

