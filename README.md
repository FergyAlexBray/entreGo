# EntreGo

## Contributions

Afin de contribuer au projet, vous aurez besoin de compiler celui-ci et de le lancer.
Nous allons aborder cela sur les parties suivantes.

### Compiler le projet

La compilation du projet se fait avec la commande `go`:

Il faut d'abord être sur d'avoir celle-ci installée sur notre ordinateur.
Si cela n'est pas le cas, vous pouvez donc l'installer en suivant les étapes suivantes: [go.dev](https://go.dev/doc/install)

Compilation du projet:

```bash
go build -o entrego main.go
```

Cette commande rajoutera un éxécutable à la racine appelé: `entrego`

### Lancer le projet

Une fois nous avons compilé le projet, nous pourrons alors le lancer de la façon suivante:

```bash
./entrego <fichier_de_configuration>
```

Vous pouvez trouver certains examples de fichiers de configuration dans le dossier `/map_examples`

## Organisation

Afin d'organiser le projet de la meilleure façon, nous avons decidé de diviser plusieurs parties.

Les parties étaient les suivantes:
- Fergus Bray
    - Parsing
    - Error handling
    - Linking all the pieces together
- Charles Verchères De Mateos
    - Path finding algorithm
    - State display
    - C4model component diagram
- Tomàs Forné Cappeau
    - Game core logic
    - Game loop and structure

En ce qui concerne la structure des fichiers, chaque partie et element à son propre fichier.
Dans chaque fichier, il y a les fonctions qui lui correspondent afin de séparer au mieux les parties du code.

Nous n'avons pas utilisé de packages externes pour ce projet.

## Strategie utilisée

Pour accomplir la mission de vider l'entrepot. Nous avons utilisé la strategie suivante:

- Ranger le poids des colis par ordre croissant
- Assigner les colis par ordre croissant aux transpalettes
- Faire avancer les transpalettes vers les colis
- Récupérer le colis lorsque le transpalette est à côté de sa cible
- Une fois récupérer, se diriger vers un camion libre
- Faire avancer les transpalettes vers le camion ou colis si toujours pas récupéré
- Déposer les colis dans le camion lorsqu'il est à côté de celui-ci
- Si le camion est plein, ou presque plein. Le faire partir pour le vider
- Faire avancer les transpalettes, même si le camion est partit. Et attendre devant la station de recharge.
- Une fois le colis déposé, mais il n'y a plus de colis a déplacer, et d'autres transpalettes n'ont pas fini leur mission. Nous faisons reculer les transpalettes vers leurs position de début
- Si tous les transpalettes ont déposé les colis, et il n'y a plus de colis a récupérer, nous finissons le jeux par une victoire.

L'algoritme de recherche de chemin utilisé dans ce projet, est le Breath-first search.
