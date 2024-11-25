# SO1_proyecto2_2S_2023
SO1_proyecto2_2S_2023

# Proyecto 2
### Sistemas Operativos 1
### Kevin Golwer Enrique Ruiz Barbales 201603009

<img src="https://cunsac.usac.edu.gt/wp-content/uploads/2022/03/USAC_lineal_azul.png"  style="height: 300px; width:900px;"/>
SO1_proyecto2_2S_2023


### Introduccion

en el siguiente proyecto tuvo la finalidad de complementar y poner en practica los conocimientos adquiridos en el curso de so1 , el presente proyecto
esta desarrollado con diferentes tecnologias y lenguajes , algunos de los lenguajes usados en este proyecto fueron go , python , javascript , yaml entre otros ,
tambien entre las tecnologias usadas estan kubernetes , docker , cloudrun , gcp , el proyecto cuenta con diferentes nodos conectados dentro de un cluster de 
kubernetes , donde recibe informacion a traves de un ingress el cual divide el trafico a una api de python y otra de go, la cual trabaja con la comunicacion a
traves de grcp, este a su vez guarda la informacion en una bd externa alojada en cloud sql de google , tambien parte de la informacion se guarda en una base 
de datos en redis que es en memoria , esta es consultada a traves de un servicio de nodejs el cual se expone a traves de un load balancer y luego mostrado en 
un front montado en cloud run 



## LOCUST


esta es el inicio de la aplicacion sirve para enviar trafico a la app , a traves de un archivo json como se muestra en la imagen


<img src="https://locust.io/static/img/screenshot_2.13.1.png"  style="height: 400px; width:700px;"/>



## LOAD BALANCER

este es el encargado de poder conectar nuestra aplicacion al exterior , nos asigna una ip y con ella podremos conectarnos a nuestro proyecto en kubernetes


<img src="https://miro.medium.com/v2/resize:fit:1400/1*531E99qa7V2nDeIWdiTOxg.png"  style="height: 400px; width:400;"/>



## INGRESS

este es un servicio de kubernetes que nos permite rutear un puerto de nuestro pod, y exponerlo a traves del puerto 80 , para nuetro caso este ingress dividira
el trafico hacia un pod que contiene un cliente de go y a otro que contiene una api en python 



<img src="https://gateway-api.sigs.k8s.io/images/simple-split.png"  style="height: 200px; width:900;"/>



## POD PYTHON

este pod es una api que se encarga de guardar la informacion en este caso el json en una base de datos no relacional que seria REDIS y luego en otra base de datos
relacional que seria en MYSQL alojada en CLOUD SQL

<img src="https://images.squarespace-cdn.com/content/v1/5644323de4b07810c0b6db7b/1497657996408-PT55M6ZS0BHJLEEQJ0A4/data.PNG?format=1000w"  style="height: 400px; width:400;"/>



## POD GRCP

este es un pod que tiene como finalidad el guardar unicamente la informacion en el CLOUD MYSQL , ademas de eso cuenta con 2 contenedores , estos contenedores 
fucionan con la tecnologia grcp que permite comunicarse a traves de procedures almacenados en el codigo directamente , el primer contenedor es el cliente el cual recibe
la informacion del splitter , esta informacion se envia al servidor grcp donde luego se pasa a insertar a la base de datos en mysql


<img src="https://miro.medium.com/v2/resize:fit:1200/1*O5z3vKFFW4vY5zuFuhBjFw.png"  style="height: 400px; width:400;"/>



## POD REDIS 

este pod unicamente almacena la mitad de la informacion recibida por el ingress , para luego ser consultada por una api de nodejs en tiempo real ,
esta informacion esta en memoria por lo que si el contenedor o el pod se reinicia toda la informacion se perdera 

<img src="https://media.geeksforgeeks.org/wp-content/uploads/20230315145951/redis(2).png"  style="height: 500px; width:400;"/>



## POD NODEJS

este pod tiene como finalidad el exponer los datos guardados en la base de datos de redis esto a traves de un loadbalancer , de lo contrario no podriamos acceder a esa informacion, tambien obtiene los datos guardados en mysql para luego mostrarlos en una interfaz de usuario hecha en react , 
es un componente importante ya que sin el no podremos visualizar nada de lo que tenemos almacenado en las base de datos , tambien funciona una parte con sockets
para mostar en tiempo real la informacion al cliente de la base de datos de redis 

<img src="https://miro.medium.com/v2/resize:fit:923/1*_9SnIxRyD5nCdqMMK4lXQQ.png"  style="height: 500px; width:400;"/>





## REACT VIEW

esta es la vista para mostrar en un dashboard sencillo toda la informacion almacenada en MYSQL y en la base de datos de REDIS , 
realiza las siguientes consultas para mostrar:


#### MYSQL (Parte estatica)
Datos Almacenados.
- Gráfica Circular de las Notas de un Curso en un semestre. (No. Aprobados y
Reprobados)

- Gráfica de Barras de Cursos con Mayor número de alumnos en un semestre
específico. (Mostrar Top 3)

- Gráfica de Barras de Alumnos con mejor Promedio (Mostrar únicamente un Top 5)


#### REDIS (Parte dinámica)
- Cantidad Total de Registros en Tiempo Real.
- Cantidad de Alumnos en un Curso y Semestre Específico


<img src="https://www.devaim.com/wp-content/uploads/2021/08/visu5.png"  style="height: 500px; width:500;"/>


## MANIFIESTO KUBERNETES 


 
  
     
>   
>   
>   
> 
>#DEPLOYMENT PARA LEVANTAR NODEJS API  
>apiVersion: apps/v1 
>kind: Deployment 
>metadata: 
>  name: nodejsdeployment 
>spec: 
>  selector: 
>    matchLabels: 
>      greeting: hello4 
>  replicas: 1 
>  template: 
>    metadata: 
>      labels: 
>        greeting: hello4 
>    spec: 
>      containers: 
>       name: pythoncontainer 
>        image: "kruiz9/nodejs_api:latest" 
>        imagePullPolicy: Always 
>        command: ["sh","c"] 
>        args: ["chmod u+r+x ips.sh && /ips.sh && node index.js"] 
>        env: 
>         name: "PUERTO" 
>          value: "8000" 
>         name: "HOST" 
>          value: "service3" 
>         name: "SUSCRIBE" 
>          value: "datos" 
>   
>   
>   
>   
>   
> 
>#DEPLOYMENT PARA LEVANTAR REDIS 
>apiVersion: apps/v1 
>kind: Deployment 
>metadata: 
>  name: redisdeployment 
>spec: 
>  selector: 
>    matchLabels: 
>      greeting: hello3 
>  replicas: 1 
>  template: 
>    metadata: 
>      labels: 
>        greeting: hello3 
>    spec: 
>      containers: 
>       name: rediscontainer 
>        image: "redis/redisstackserver:latest" 
> 
>       
>       
> 
>#DEPLOYMENT PARA LEVANTAR PYTHON API  
>apiVersion: apps/v1 
>kind: Deployment 
>metadata: 
>  name: pythondeployment 
>spec: 
>  selector: 
>    matchLabels: 
>      greeting: hello2 
>  replicas: 1 
>  template: 
>    metadata: 
>      labels: 
>        greeting: hello2 
>    spec: 
>      containers: 
>       name: pythoncontainer 
>        image: "kruiz9/python_api:latest" 
>        imagePullPolicy: Always 
>        command: ["sh","c"] 
>        args: ["chmod u+r+x ips.sh && /ips.sh && python3 u api.py"] 
>        env: 
>         name: "PUERTO" 
>          value: "8000" 
>         name: "HOST" 
>          value: "service3" 
>      
> 
> 
>#DEPLOYMENT PARA LEVANTAR GRCP 
>apiVersion: apps/v1 
>kind: Deployment 
>metadata: 
>  name: grcpdeployment 
>spec: 
>  selector: 
>    matchLabels: 
>      greeting: hello 
>  replicas: 1 
>  template: 
>    metadata: 
>      labels: 
>        greeting: hello 
>    spec: 
>      containers: 
>       name: grcpclientcontainer 
>        image: "kruiz9/grcp_client" 
>        imagePullPolicy: Always 
>       name: grcpservercontainer 
>        image: "kruiz9/grcp_server" 
>        imagePullPolicy: Always 
>        command: ["sh","c"] 
>        args: ["chmod u+r+x ips.sh && /ips.sh && /server"] 
>        #command: ["sh", "ips.sh"] 
> 
> 
> 
> 
>#servicio 4 para API_NODEJS 
>apiVersion: v1 
>kind: Service 
>metadata: 
>  name: service4 
>spec: 
>  type: ClusterIP 
>  selector: 
>    greeting: hello4 
>  ports: 
>   name: worldport 
>    protocol: TCP 
>    port: 40000 
>    targetPort: 8000 
>  type: LoadBalancer 
> 
> 
> 
>#servicio 3 para REDIS 
>apiVersion: v1 
>kind: Service 
>metadata: 
>  name: service3 
>spec: 
>  type: ClusterIP 
>  selector: 
>    greeting: hello3 
>  ports: 
>   name: worldport 
>    protocol: TCP 
>    port: 6379 
>    targetPort: 6379 
> 
> 
> 
> 
>#servicio 2 para PYTHON API 
>apiVersion: v1 
>kind: Service 
>metadata: 
>  name: service2 
>spec: 
>  type: ClusterIP 
>  selector: 
>    greeting: hello2 
>  ports: 
>   name: worldport 
>    protocol: TCP 
>    port: 60000 
>    targetPort: 8000 
> 
> 
> 
> 
>#servicio 1 para GRCP 
>apiVersion: v1 
>kind: Service 
>metadata: 
>  name: service1 
>spec: 
>  type: ClusterIP 
>  selector: 
>    greeting: hello 
>  ports: 
>   name: worldport 
>    protocol: TCP 
>    port: 50000 
>    targetPort: 8000 
> 
> 
> 
> 
> 
>   
> 
> #crea el pod para dividir el trafico 
>  
>apiVersion: v1 
>kind: Pod 
>metadata: 
>  name: splitter 
>  labels: 
>    app.kubernetes.io/name: splitter 
>spec: 
>  containers: 
>   name: splittercontainer 
>    image: kruiz9/splitter 
>    imagePullPolicy: Always 
>    ports: 
>     containerPort: 20000 
>    env: 
>     name: "PUERTO" 
>      value: "20000" 
>     name: "RUTA1" 
>      value: "http://service1:50000" 
>     name: "RUTA2" 
>      value: "http://service2:60000" 
>       
>       
>       
>       
> 
>#crea el servicio para conectar al divisor de trafico 
>apiVersion: v1 
>kind: Service 
>metadata: 
>  name: pservice 
>spec: 
>  type: ClusterIP 
>  selector: 
>    app.kubernetes.io/name: splitter 
>  ports: 
>   name: worldport 
>    protocol: TCP 
>    port: 10000 
>    targetPort: 20000 
>  
>  
>  
>  
>  
>  
>  
>  
>  
>  
>  
>apiVersion: networking.k8s.io/v1 
>kind: Ingress 
>metadata: 
>  name: myingress 
>  annotations: 
>    kubernetes.io/ingress.globalstaticipname: "webstaticip" 
> 
>spec: 
>  rules: 
>   http: 
>      paths: 
>       path: / 
>        pathType: Exact 
>        backend: 
>          service: 
>            name: pservice 
>            port: 
>              number: 10000 
> 
>  
>  
>  
>   
>   
>   
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
>
 
 
  
  
  

















