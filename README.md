# API Vote

Une API REST pour : 
- créer un nouveau de ballot de Vote avec la méthode _/new_ballot_
- voter avec la méthode _/vote_
- récupérer le résultat des votes avec la méthode _/result_

1. Possibilité de lancer le serveur et de faire les commandes dans le navigateur avec l'executable dans le dossier launch-rsagt

2. Possibilité de lancer tous les agents de votes en même temps pour tester toutes les méthodes avec l'executable dans le dossier launch-all-rest-agents

## Utilisation locale

**Récupération du projet**

```bash
  go mod init projet_test
  go get gitlab.utc.fr/langemat/ia04@install 

```
Les fichiers se trouvent dans le GOPATH dans le dossier pkg/mod/gitlab.utc.fr/langemat

**Démarrer le serveur**

```bash
  cd cmd/launch-rsagt
  ./launch-rsagt.exe
```
Des tests peuvent ensuite être effectués par API REST sur localhost:8080

**Démarrer le serveur & les agents de votes**

```bash
  cd ..
  cd launch-all-rest-agents
  ./launch-all-rest-agents.exe
```
Ce fichier contient des tests pour toutes les méthodes de vote implémentées. L'utilisateur doit renseigner le nombre d'agents votants et le nombre de candidats (chaque candidat est associé à un entier entre 1 et n). La deadline, les voter-ids et le tiebreak pour les ballots sont générés automatiquement, ainsi que les préférences des votants. La deadline est créée afin de permettre d'obtenir un résultat rapidement.

## Méthodes de vote implémentées

- Majorité (rule:"majority"),
- Borda (rule:"borda"),
- Approval (rule:"approval"), options à renseigner sous forme d'un int[] de longueur 1,
- Condorcet (rule:"condorcet"). Si pas de gagnant de Condorcet, la valeur renvoyée est 0 car il s'agit de la 0-value du type entier. Or cela ne correspond à aucun candidat existant,
- Copeland (rule:"copeland")
- STV (rule:"stv"). Pour chaque tour on applique la majorité. En cas d'égalité, le candidat avec le plus petit indice est choisi car la fonction STV_SWF ne prend pas de tie-break en argument. STV_SWF renvoie un count tel que plus une alternative progresse dans les tours, plus son score est élévé.

Les factories TieBreakFactory, SWFFactory et SCFFactory sont dans `/comsoc/utils.go` et les tests dans `/comsoc/go_tests.go`.

## Authors

- Mathilde Lange
- Solenn Lenoir

