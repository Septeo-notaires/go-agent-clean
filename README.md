
# go-agent-clean

Un outil en Go permettant de nettoyer les agents Azure DevOps présents sur une machine. Compatible avec Linux et Windows.

## 📌 Fonctionnalités
- Nettoie les dossiers de builds des agents
- Nettoie les fichiers temporaires associés aux agents.

## ⚡ Installation

Assurez-vous d'avoir [Go](https://go.dev/dl/) installé sur votre machine.

Sur Linux :
-
```sh
git clone https://github.com/votre-repo/go-agent-clean.git
cd go-agent-clean
go build -o go-agent-clean
```

Sur Windows
-

```sh
go build -o go-agent-clean.exe
```

### Installation depuis une release

Télécharger l'archive associer depuis l'onglet release dans github.

## 🚀 Utilisation

L'outil utilise un fichier de configuration au format TOML.
Créez un fichier `config.toml` dans le même dossier que l'exécutable avec le contenu suivant :

#### Mono-agent
```toml
[[agents]]
name="agent_name"
service="service_name"
path='agent_path' ## Spécifier le dossier _work de l'agent
```

#### Multi-agent
```toml
[[agents]]
name="agent1_name"
service="service1_name"
path='agent1_path' ## Spécifier le dossier _work de l'agent

[[agents]]
name="agent2_name"
service="service2_name"
path='agent2_path' ## Spécifier le dossier _work de l'agent

...
```

Ensuite, exécutez l'outil :

Sur Linux :
```sh
./go-agent-clean
```

Sur Windows :
```sh
go-agent-clean.exe
```

## ⚙️ Fonctionnement

go-agent-clean analyse le répertoire spécifié dans le fichier `config.toml` pour identifier les agents Azure DevOps présents sur la machine. Il procède ensuite aux étapes suivantes :

1. Vérification des agents actifs.
2. Arrêt des agents
3. Suppression des fichiers du dossier work.
5. Démarrage des agents.

Le processus garantit qu'aucun agent en cours d'utilisation ne sera supprimé.

## 🏗️ Contribution
Les contributions sont les bienvenues ! Merci de créer une issue ou une pull request.
