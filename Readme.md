# Ejemplos de uso de docker network y loadlancer con NGINX
 El repositorio cuenta con 3 ejemplos sencillos, cada uno dentro de los directorios del repositorio:
-network
-load-balancer
-bd

Cada directorio tiene dentro un archivo docker-compose.yaml que es el cual se encarga de levantar los servicios con el comando:
"docker-compose up". Al ejecutar dicho comando, docker crea las network y contenedores.

## Network
  Solo contiene dentro un archivo docker-compose.yaml donse se realizan las siguientes acciones:
  - Se crea una network llamada "red-servidores" con el driver bridge.
  - Se crean dos contenedores basados en una imagen de un servidor hecho en go, la cual esta alojada en docker-hub. Dicho servidor posee el endpoint /mensaje, el cual devuelve un mensaje simple. Los puertos de los contenedores estan mapeados en el host, por lo que es posible acceder a los servidores mediante la url : localhost:5000/mensaje y localhost:5001/mensaje

## load-balancer
 Contiene dentro un archivo docker-compose.yaml con los elementos a crear y un archivo de configuración  nginx.conf, el cual es usado como base para crear un contenedor con nginx el cual se encargara de realizar balanceo de carga entre contenedores.
 El docker-compose.yaml realiza las siguientes acciones:
 - Se crean las networks "red-servidores" y "red-servidores2", ambas utilizando el driver bridge.
 - Se crean dos contenedores, los cuales tienen como base la imagen alojada en docker-hub llamada "api-servidor", la cual es un servidor en go, el cual devuelve un mensaje, con la peculiaridad de que cada vez que se levanta el servidor, este obtiene un nombre propio, el cual adjunta al mensaje devuelto.
 - Se crea un contenedor de NGINX, basado en el archivo de configuración nginx.conf, el cual configura el balanceo de carga para los servidores. Se puede probar accediendo a la url localhost:4000/mensaje

## bd
En este directorio se realiza una reimplementación del ejemplo de load-balancer, pero se agrega la funcionalidad de conexion con base de datos.
El archivo docker-compose.yaml realiza las siguientes acciones:
- Se crean dos networks, la network "red-servidores" y la network "bd", donde se conectaran los distintos contenedores.
- Se crea un contendor con mysql para la base de datos. Dicha base de datos se crea a partir del archivo setup.sql en el directorio interno mysql. Para el build se usa el archivo Dockerfile dentro del directorio mysql.
- Se crea un servidor el cual se conecta al contendor de base de datos y se encarga de devolver un listado de nombres, el contenedor se genera con la imagen descargada de docker-hub.El codigo del servidor se adjunta en el directorio, pero se debe aclarar que **dicho codigo no se utiliza en el build de la imagen, esta es descargada de docker-hub, pero es el mismo codigo con el que se genero dicha imagen.**
- Un balanceador de carga basado en nginx, la configuración es la misma que en el ejemplo del directorio **load-balancer**.

### Tomar en cuenta la siguiente
- El codigo de los servidores puede contener bugs, malas practicas, falta de comentarios, etc.**Es solo para referencia del uso de contendores**
- ....
