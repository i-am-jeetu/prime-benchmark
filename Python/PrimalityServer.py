from flask import Flask,request
import logging
# from miller import MillerRabin
logger = logging.getLogger('werkzeug' )
logging.basicConfig(level=logging.CRITICAL)
logger.propagate = False
logger.disabled = True

app = Flask(__name__)

def checkComposite( n ,a , d, s) :
	x= pow(a,d,n)
	# print(n,a,d,s,x)
	if(x == 1 ) or (x == n-1) :
		return False
	for r in range(1,s):
		x = (x * x)%n
		if x == n-1 :
			return False
	return True


def MillerRabin(n) :
	if n<2 :
		return False
	r = 0
	d = n-1
	while(d&1) ==0 :
		d = d//2
		r = r+1
	a = [2,3,5,7,11,13,17,19,23,31,37]

	for val in a:
		# print(val)
		if n == val:
			return True
		if(checkComposite(n, val, d ,r)):
			return False
	
	return True



@app.route('/prime', methods=['POST'])
def query():
	data = request.get_json(force = True)
	return str(MillerRabin(data["value"]))
	

if __name__ == "__main__":
	# app.debug = True
	app.run(port=8002)