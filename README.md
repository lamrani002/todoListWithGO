************ Environement Technique **************

  Back end : Api avec golang, echo, sqlite3

  Data base : faudra avoir sqlite3 ou le telecharger ici : https://www.sqlite.org/download.html
  
  Front end : Vue js et Bootsrap sont importé en CDN, donc pas besoin de les téléchargera mois que si vous n'avez pas d'internet

**********Execution du projet ******************
Une fois tout ces vérification sont faites.

executer la commande go build todo.go une fois que c'est fait.

lancer le serveur en executant ./todo (si vous voulez changer le port c'est dans le fichier )

le fichier contient egalement la fonction pour la migration de la base de données 


**********Architecture du Projet ***************

architecture MVC 

Models : En créant l'ensemble des requetes SQL besoin pour notre api 

Handlers : Ce qui permet de comuniquer avec le serveur en se basant sur les models pour manipuler les fonctions

Systeme Routage: qui permet d'acceder aux liens et recupérer les données grace au Handlers

Public / index.html  qui contient la partie HTML/CSS ainsi que le javascript (Vue JS) => un seul fichier pour se repérer facilement







