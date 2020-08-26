# restApiWithGO

============== Software necesario ==============
- Version de Go: 1.14 o superior.
- Postman o equivalente.
- Paquete mux


============== Pasos para levantar el proyecto ==============
1. Descargar el modulo de mux
- go get github.com/gorilla/mux

2. Poner a correr
- go run main.go


============== Endpoints ==============
- localhost:3000 
Es el root, la podemos abrir en el explorador para confirmar que el servidor esta Up.


- localhost:3000/tasks
GET  : Si la petición es un GET, solo devolverá una lista con todas las tareas.

POST : Si se envia una peticion POST, se agregará una tarea, si la tarea tiene el "body" vacío
se agrega una tarea vacía.

=== ejemplo de lo que puede ir en el body ===
{
    "Name":    "nombre",
    "Content": "Descripcion"
}


- localhost:3000/tasks/1
GET    : Si se envia un GET, devuelve la tarea especificada.

DELETE : En caso de ser un DELETE, borra la tarea especificada.

PUT    : Con un PUT, la peticion debe llevar un body, para actualizar la tarea, sino
la tarea se actualiza con valores vacios.

=== ejemplo de lo que puede ir en el body ===
{
    "Name":    "nombre actualizado",
    "Content": "Descripcion actualizada"
}



============== Notas ==============
El modulo de compileDaemon es un paquete que reinicie el server despues hacer algun cambio.
Fue usado al momento de hacer el ejemplo para no estar bajando y subiendo el servidor, pero 
para probar el ejemplo no es necesario.

go get github.com/githubnemo/CompileDaemon <- para traer el paquete

CompileDaemon -command="nombreProyecto.exe" <- para ejecutarlo