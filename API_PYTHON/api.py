# Using flask to make an api 
# import necessary libraries and functions 
from flask import Flask, jsonify, request 
import os
import redis
from flask_cors import CORS, cross_origin
import mysql.connector


# creating a Flask app 
app = Flask(__name__) 
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'




@app.route('/', methods = ['GET', 'POST']) 
@cross_origin()
def home(): 
	if(request.method == 'GET'): 
		print ('GET')
		return jsonify({'data': "OK_py"}) 
	else: 

		print ('POST')
		
		
		r = redis.Redis(host=os.environ['HOST'], port=6379, decode_responses=True)

		content_type = request.headers.get('Content-Type')
		if(content_type=='application/json'):
			json=request.json


			#PARTE DE REDIS INSERTA BD EN MEMORIA
			#sirve para activar la lectura grupal activa
			r.publish("data",str(json))
			#para guardar toda la data en la bd seria en una lista 
			r.sadd("data",str(json))
			#smember data   --obtiene todos los datos no repetidos
			#scard data --obtiene conut de datos 

			#print(json)
			#print(json['carne'])

			
			#PARTE PARA MYSQL INSERTA EN BD EXTERNA
			mydb = mysql.connector.connect(
			host="34.121.157.97",
			user="root",
			password="asdf1234.,.,",
			database="tarea7"
			)


			try:
				mycursor = mydb.cursor()
				sql = "INSERT INTO datos(carnet,nombre,curso,nota,semestre,year)values(%s,%s,%s,%s,%s,%s)"
				val = (json['carne'], json['nombre'],json['curso'] ,json['nota'] ,json['semestre']  , json['year'] )

				print (json['carne'])
				print (json['nombre'])
				print (json['curso'])
				print (json['nota'])
				print (json['semestre'])
				print (json['year'])


				mycursor.execute(sql, val)
				mydb.commit()
				mydb.close()

			except:
				print("No se pudo conectar a la bd") 

			print (json)
			return json
		else:
			return 'Content-Type no soportado'
		


# driver function 
if __name__ == '__main__': 
	from waitress import serve
	print ('Inicio API python9 , PUERTO:'+os.environ['PUERTO'])
	serve(app, host="0.0.0.0", port=os.environ['PUERTO'])

	#app.run(debug = True) 
