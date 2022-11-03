from asyncio import events
from distutils.log import debug
from random import randrange   # para generar números aleatorios
import json
from locust import HttpUser, task,events ,between # importamos la clase HttpUser y la funcion task para crear las tareas

debug = False # para ver los mensajes de debug
def print_debug(msg):
    if debug:
        print(msg)

@events.init_command_line_parser.add_listener
def _(parser):
    parser.add_argument("--my-argument", type=int , env_var="LOCUST_MY_ARGUMENT", default="", help="It's working")

@events.test_start.add_listener
def _(environment, **kw):
    print(f"Custom argument supplied: {environment.parsed_options.my_argument}")



class Lectura():
    def __init__(self):
        self.array = [] # array de lecturas

    def pickRandom(self):
        #return self.array
        length = len(self.array)

        if (length > 0):
            random_index = randrange(0,length-1) if length > 1 else 0
            return self.array.pop(random_index)
        else:
            print('>> Lectura: No hay lecturas para leer en archivo json')
            return None

    def load(self):
        print(">> Lectura: Cargando archivo json: ")
        try:
            with open('data.json') as data_file:
                self.array = json.loads(data_file.read())
                print_debug('>> Lectura: Lecturas cargadas: ' + str(len(self.array)))
        except Exception as e:
            print('>> Lectura: Error al cargar archivo json: ', e)
            return False

class MensajeDeTrafico(HttpUser):
    wait_time = between(1, 2.5) # tiempo de espera entre cada tarea
    lectura = Lectura() # instancia de la clase Lectura
    lectura.load() # cargamos el archivo json

    def on_start(self):
        print(">> MensajeDeTrafico: Iniciando envío de trafico")

    @task
    def my_task(self):
        print(f"my_argument={self.environment.parsed_options.my_argument}")

    @task
    def PostMessage(self):
        random_data = self.lectura.pickRandom()
        
        if (random_data is not None and self.environment.parsed_options.my_argument > 0):
           # data_to_send = json.dumps(random_data)
            print_debug('>> MensajeDeTrafico: Enviando datos: ' + str(random_data))
            
            if self.environment.parsed_options.my_argument > 1:
                self.client.post("/input", json=random_data)
                self.environment.parsed_options.my_argument -= 1
            print("VALOR DE CONTADOR: ", self.environment.parsed_options.my_argument)           
        else:
            print('>> MensajeDeTrafico: No hay datos para enviar')
            #cuando no hay datos para enviar, se detiene el test
            events.request_failure.fire(request_type="PostMessage", name="PostMessage", response_time=0, exception=Exception("No hay datos para enviar"))



    