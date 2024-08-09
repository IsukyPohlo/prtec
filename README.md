# prtec

1. Si la base de datos no puede aceptar solicitudes en paralelo entonces la siguiente solución seria trabajar con un orquestador de queries usando cualquier sistema de queues como SQS.:
      1. Tener un solo programa que sea por el cual todos los queries dirigidos a la DB pasen por.
      2. Para garantizar que todos los queries se completen se utiliza cualquier sistema de queues como SQS, esto para saber si existe el request todavia en la queue o si falla.
      3. Además de eso, habría que reestructurar el como se leen los datos, ya que seria un problema tener que esperar a las operaciones de escritura para poder hacer las operaciones de lectura que usualmente son mas importantes para mostrar información a un usuario o para empezar cualquier proceso.

4. Arquitecture:
      1. Los Client Apps son aquellos que van a mandar requests al API Gateway y son los usuarios que van a interactuar con los microservicios indirectamente, están divididos en Web y en Mobile ya que usualmente la plataforma de PC y la de celulares requiere de distintos desarrollos y funcionan de formas distintas.
      2. API Gateway es la abstracción que redirecciona los request a un microservicio en especifico. Usualmente es proporcionado por el proveedor de nube y sirve para redirigir trafico al API Gateway y de ahí a cada microservicio y así tener una buena organización del API Rest.
      3. Los microservicios son los programas que reciben le request y realizan la lógica del negocio. Se conectan cada uno con su propia DB y algunos como Catalog y Shopping Cart envian mensajes al Broker, mientras que discount y ordering recibe mensajes del broker. Pueden ser desarrollados en distintos lenguajes o usar distintos protocolos pero deben ser compatibles con el API GW y el Broker.
      4. Message Broker envía y recibe los eventos de los microservicios, puede funcionar con Queues, Eventos, Mensajes, etc. Pero básicamente envía y recibe información a distintos microservicios.

5. Los microservicios nos permiten no tener todos nuestros  "huevos en la misma canasta" si por ejemplo por algún problema de la plataforma o de código, uno de nuestros servicios se cae, si tenemos cada parte en su propio servicio, nos permite granularizar cada función y ser mas resilientes a los errores. Además los desarrolladores les toma menos tiempo familiarizarse con un microservicio que es considerablemente mas pequeño en código y complejidad que con todo el sistema que podría ser varias veces su tamaño. Como si dividiéramos un libro en capítulos y cada uno encuadernarlo por separado.
