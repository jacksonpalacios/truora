# Truora

(api-rest) permite leer, crear, actualizar y borrar recetas de cocina. Implementa base datos mySQL, y golang como 
lenguaje de programación, dicha res api es consumida por una app (recipes-api) hecha con Vue.js

## Instalación Linux
Descargar el directorio truora, en esta carpeta encontrará dos carpetas más, (api-rest), (recipes-api) y la base de datos 
(go-api.sql), importa la base de datos en su máquina local, después de esto entra a la carpeta (api-rest) y ejecuta
el archivo main
<br>
Dentro de la carpeta api-rest, ejecuta el archivo main, así:
```
 ./main

```
En esta instancia, el servidor estará corriendo.
<br>
Dentro de la carpeta (recipe-api), ejecuta:
<br>
Ejecutar archivos de desarrollo
<br>
De este modo, el archivo será sevidor por la dirección http://localhost:8080/
```
 npm install
 npm run dev

```
<br>
O ejecutar archivos de producción:
<br>
Ponga a correr cualquier servidor de su preferencia, y copia el archivo index.html y la carpeta dist y la pega 
en el lugar correspondiente para que dicho servidor pueda leer el archivo index
