cat > README.md << 'EOF'
# 🎮 PUISSANCE4

> Une version moderne du célèbre jeu **Puissance 4**, développée avec **Go**, **HTML** et **CSS**.  
> Deux joueurs s’affrontent pour aligner quatre jetons de la même couleur !

---

## ✨ Fonctionnalités

- 🧩 Jeu à deux joueurs (rouge contre jaune)  
- ⚙️ Logique du jeu écrite en **Go**  
- 🎨 Interface simple et élégante en **HTML / CSS**  
- ⚡ Rapide, fluide et responsive  
- 💬 Messages d’état (victoire, égalité, tour du joueur)

---

## 🚀 Lancer le projet

### 1. Cloner le dépôt
\`\`\`bash
git clone https://github.com/nikita-ynov/PUISSANCE4.git
cd PUISSANCE4
\`\`\`

### 2. Exécuter le serveur Go
\`\`\`bash
go run main.go
\`\`\`

### 3. Ouvrir le jeu
Ouvre ton navigateur et rends-toi sur :  
👉 [http://localhost:8080](http://localhost:8080)

---

## 🧠 Règles du jeu

1. Les joueurs jouent à tour de rôle.  
2. À chaque tour, un joueur place un jeton dans une colonne.  
3. Le jeton tombe jusqu’à la première case libre de la colonne.  
4. Le premier joueur à aligner **4 jetons** (horizontalement, verticalement ou en diagonale) gagne la partie.  
5. Si la grille est pleine sans vainqueur, c’est **une égalité**.

---

## 🛠️ Structure du projet

\`\`\`
PUISSANCE4/
├── src/              # Fichiers sources du jeu
├── static/           # Ressources front-end (HTML, CSS)
├── main.go           # Point d’entrée du serveur Go
└── README.md         # Ce fichier :)
\`\`\`

---

## 🧑‍💻 À propos du projet

Ce projet a été réalisé dans un but **pédagogique** — pour s’exercer à la logique de jeu, à la programmation en **Go** et à la création d’interfaces web simples.  
Le code est open-source et librement réutilisable. 🌐

---

## 📸 Aperçu du jeu

Voici un aperçu du rendu du site :

![Aperçu du jeu](./static/screenshot.png)

---

## 📜 Licence

Ce projet est distribué sous la licence **MIT** — tu peux l’utiliser, le modifier et le partager librement.

---

### ❤️ Développé avec passion par [@nikita-ynov](https://github.com/nikita-ynov) et [@AM3X-svg](https://github.com/AM3X-svg)
EOF
