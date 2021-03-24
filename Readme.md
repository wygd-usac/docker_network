El repositorio cuenta con 3 ejemplos sencillos, cada uno dentro de los directorios del repositorio:
-network
-load-balancer
-bd

Cada directorio tiene dentro un archivo docker-compose.yaml que es el cual se encarga de levantar los servicios con el comando:
"docker-compose up". Al ejecutar dicho comando, docker crea las network y contenedores.

Network:
Solo contiene dentro un archivo docker-compose.yaml donse se realizan las siguientes acciones:
-Se crea una network llamada "red-servidores" con el driver bridge.
-Se crean dos contenedores basados en una imagen de un servidor hecho en go, la cual esta alojada en docker-hub. Dicho servidor posee el endpoint /mensaje, el cual devuelve un mensaje simple. Los puertos de los contenedores estan mapeados en el host, por lo que es posible acceder a los servidores mediante la url : localhost:5000/mensaje y localhost:5001/mensaje
