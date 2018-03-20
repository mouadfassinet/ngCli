Un CLI simple pour baculer à une branche existante sur tous les Repositories Front à la fois.

[![asciicast](https://asciinema.org/a/JEmKeDph9INzNo2waLLELgHoH.png)](https://asciinema.org/a/JEmKeDph9INzNo2waLLELgHoH)

# Installation
Télécharger et installer le binaire Go https://golang.org/dl

## Sur GNU/Linux
```shell
sudo tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz
```


Se situer à la racine du répertoire IadNgIntranet, copier le fichier `main.go`

## Sur GNU/Linux
Lancer la commande
```shell
export PATH=$PATH:/usr/local/go/bin
```

Installer les dépendances
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
