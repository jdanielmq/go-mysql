# go-mysql

# Resumen
En esta sección del curso de Golang, exploramos cómo trabajar con una base de datos MySQL para la persistencia de datos en nuestras aplicaciones. A lo largo de esta sección, aprendimos los fundamentos necesarios para interactuar con una base de datos MySQL desde una aplicación Go.

Comenzamos instalando y configurando MySQL en nuestro sistema para prepararnos para trabajar con bases de datos. Luego, creamos una base de datos utilizando comandos SQL básicos y establecimos una conexión entre nuestra aplicación Go y la base de datos MySQL para poder realizar operaciones de lectura y escritura.

Utilizamos la biblioteca godotenv para cargar variables de entorno desde un archivo .env, lo que nos permitió gestionar de manera segura y conveniente la configuración de la conexión a la base de datos.

Implementamos la funcionalidad para listar todos los contactos almacenados en la base de datos y aprendimos a recuperar un contacto específico utilizando su identificador único. También exploramos cómo registrar un nuevo contacto, actualizar la información de un contacto existente y eliminar un contacto de la base de datos cuando fuera necesario.

Finalmente, concluimos la sección construyendo un menú de navegación interactivo utilizando el paquete bufio, que permitió al usuario realizar todas las operaciones CRUD de manera intuitiva.

Al finalizar esta sección, estamos bien equipados para trabajar con bases de datos MySQL en nuestros proyectos Go y gestionar eficazmente la persistencia de datos.