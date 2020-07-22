const express = require('express')
const bodyParser = require('body-parser')
const app = express()
app.use(bodyParser.json())

function modMultiply(a, b, mod )  {
    let res = 0
	a = a % mod
	while (b > 0) {
		if (b%2 == 1) {
			res = (res + a) % mod
		}
		a = (2 * a) % mod
		b /= 2
	}
	return res % mod
}


function binpower(base, e ,mod) {
	// base = Number.un
	let result = 1
	base%=mod
	while(e) {
		if(e&1) {
			result = result* base% mod;
			
			// result = modMultiply(result, base, mod);
		}
		base =base* base% mod
		// base = modMultiply(base, base , mod)
		e = e /2;
	}
	return result;
}


function check_composite(n, a,d,s) {
	let x = binpower(a,d,n)
	// console.log(n,a,d,s,x)
	if ((x === 1) || (x === n-1)) {
		return false;
	}

	let rs;
	for(rs=1; rs<s;rs++) {
		// console.log/(rs)
		x = (x*x)%n
		// x= modMultiply(x,x, n)
		if(x == n-1) {
			return false;
		}
	}
	return true;
}


function MillerRabin (n) {
	if (n<2) {
		return false;
	}
	// console.log(Number.MAX_VALUE, Number.MAX_SAFE_INTEGER)
	let r = 0;
	let d = n-1;
	while (d%2==0 ) {
		d >>=1;
		r++;
	}
	// console.log(n, r, d)
	let tests = [2,3,5,7,11,13,17,19,23,29,31,37]
	for (a of tests) {
		// console.log("A",a)
		if (n==a) {
			// console.log("N was a")
			return true;
		}
		if (check_composite(n,a,d,r)){
			
			// console.log("Checking Comp", n,a,d,r)
			return false;
		}
	}
	return true;
}



app.post('/prime' , (req, res) => {
	
	// res.send("HI")
	// console.log(binpower(500000003,2,1000000007))
	// res.send("0")
	// return
	// console.log("Value = " + req.body["value"])
	// res.send(req.body)
	if(MillerRabin(Number(req.body.value))) {
		res.send("True")
	} else {
		res.send("False")
	}
});


app.listen(8003, () => console.log("Listening to Port 8001"))