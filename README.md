Un CLI simple pour baculer à une branche existante sur tous les Repositories Front à la fois.

[![asciicast](https://asciinema.org/a/JEmKeDph9INzNo2waLLELgHoH.png)](https://asciinema.org/a/JEmKeDph9INzNo2waLLELgHoH)

# Installation
Télécharger et installer le binaire Go https://golang.org/dl

Se situer à la racine du réprtoire IadNgIntranet, copier le fichier `main.go` et installer les dépendances
```shell
go get gopkg.in/kyokomi/emoji.v1
```

Lancer la commande
```shell
go run main.go -branch=nomDeLaBranche -except=production-people-css
```

## Flags
* `-branch` - La branche où basculer.
* `-except` - (Optionnel) les Repositories à ignorer.
