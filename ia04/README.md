
# API Vote

Une API REST pour : 
- créer un nouveau de ballot de Vote
- voter 
- récupérer le résultat des votes 

1. Possibilité de lancer le serveur et de faire les commandes dans le navigateur. 
2. Possibilité de lancer tous les agents de votes en même temps pour tester toutes les méthodes.

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

- Majorité
- Borda
- Approval
- Condorcet
- Copeland
- STV

Les factories sont dans `/comsoc/utils.go` et les tests dans `/comsoc/go_tests.go`.



## Authors

- Solenn Lenoir
- Mathilde Lange

