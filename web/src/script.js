function session_start(){
	sessionStorage.patient_id = "";
	var bs = {patients: []};
	sessionStorage.doctor = "";
	sessionStorage.patient = JSON.stringify(bs);;
}

function create_rx() {
	
	var indications = JSON.stringify({
      "dpw": document.getElementById("dpw").value,
      "tpd": document.getElementById("tpd").value,
      "tod": [56,54]
    });
	
	var jsonVariable3 = {};
	
	var jsonVariable2 = {};
	jsonVariable2['daysperweek'] = document.getElementById("dpw").value;
	jsonVariable2['timesperday'] = document.getElementById("tpd").value;
	jsonVariable2['time'] = [56,54];
	
	jsonVariable3[document.getElementById("medication").value] = jsonVariable2;
	
	var jsonVariable = {};
	jsonVariable[document.getElementById("medication").value] = document.getElementById("quantity").value;
	
	var info = JSON.stringify({
      "medicines": jsonVariable3,
      "remaining": jsonVariable,
      "doctor": "cc7f1487-555c-46af-87dd-6c57f467406c"
	});
	
	var pre_id = "";
	
	xhr = new XMLHttpRequest();
		var url = "http://ezpillzz.azurewebsites.net/prescription";
		xhr.open("POST", url, true);
		xhr.setRequestHeader("Content-type", "application/json");
		xhr.onreadystatechange = function () { 
			if (xhr.readyState == 4 && xhr.status == 200) {
				pre_id = xhr.responseText;
				var json = JSON.parse(xhr.responseText);
				console.log(pre_id.substr(7,36));
			}
		}
		xhr.send(info);
		
		var qrcode = new QRCode("qrcode");
			
			qrcode.makeCode(pre_id.substr(7,36));
			console.log(elText);

}

function set_doctor(doctor){
	sessionStorage.doctor_id = doctor;
}

function set_session(name, values){
	sessionStorage.setItem(name, values);
}

function append_patients(values){
	var data = values;
	data2 = JSON.parse(sessionStorage.getItem('patient'));
	data2['patients'].push(data);
	console.log(data2);
	sessionStorage.setItem('patient', JSON.stringify(data2));
}

function get_session(name){
	console.log(JSON.parse(sessionStorage.getItem('patient')));
	return sessionStorage.getItem(name);
}

function createNode(element) {
      return document.createElement(element);
}

function append(parent, el) {
    return parent.appendChild(el);
}

function retreat(){
		$(".sick").animate({top: '0px'});
		console.log("hoooo");
}
	  
function json_load(url){
	  return new Promise((resolve, reject) => {
		fetch(url)
		  .then(function(response) {
			return resolve(response.json());
		  })
		  .catch(function(error) {
			console.log(error);
			return reject(error);
		});  
	  })
  }
  
 async function doctor_load(){
	  
	  //const url = 'https://randomuser.me/api/?results=10';
	  const url = 'http://ezpillzz.azurewebsites.net/doctor/cc7f1487-555c-46af-87dd-6c57f467406c';
	  
	  json_load(url).then( (resp) => {
		return resp.patients.map(function(patient) {
			console.log(resp);
			append_patients(patient);
		});
	  }
	  
	  )
  }

 async function patient_load(){
	  
	  //const url = 'https://randomuser.me/api/?results=10';
	  
	  
	  const resp = JSON.parse(get_session("patient"));
	  y=0;
	  while(resp['patients'].length > y){
		  //console.log(resp['patients'].length);
		  
		  const url = 'http://ezpillzz.azurewebsites.net/patient/' + resp['patients'][y];
		  console.log(url);

	  json_load(url).then( (resp) => {
		  console.log(resp['prescriptions']);
		  ///////////////////////////////////////
		  for (var key in resp['prescriptions']) {
			  //console.log(resp['prescriptions'][key]);
			json_load('http://ezpillzz.azurewebsites.net/prescription/' + resp['prescriptions'][key]).then( (prescription) => {
			
			
				console.log('http://ezpillzz.azurewebsites.net/prescription/' + key);
				console.log(prescription);
				prescription_load(prescription, resp.name);
			});	
			}
			
		///////////////////////////////////////////

		  let img = createNode('img'),
			  p = createNode('p'),
			  a = createNode('a');
			  //details = createNode('p');
		  //img.src = patient.picture.medium;
		  img.src = '../../UI Icons/Web App/ICON_Female.png';
		  img.style.height = "30px";
		  img.style.float = "left";
		  
		  a.innerHTML = '<div class=\"d-flex w-100 justify-content-between"><h5 class="mb-1">' + `${resp.name}` + '</h5><small>3 days ago</small></div><p class="mb-1">' + `${resp.phone}` + '</p><small>' + `${resp.age}` + ' , ' + `${resp.gender}` + '</small>' ;
		  a.classList.add('list-group-item');
		  a.classList.add('list-group-item-action');
		  a.classList.add('flex-column');
		  a.classList.add('align-items-start');
		  a.setAttribute("id", "sp " + `${resp.id}`);
		  a.style.float = "left";
		  
		  append(document.getElementById('patients'), a);
	  }
	  
	  );y++;}
  }
  
 async function prescription_load(prescription, x){
	  
	  //const url = 'http://meth.azurewebsites.net/patient/' + resp['patients'][i];
	  
	  //json_load(url).then( (resp) => {
		//return resp.results.map(function(prescription) {
	for (var key in prescription['medicines']) {
		  let div = createNode('div'),
			  img = createNode('img'),
			  h5 = createNode('h5');
			  ul = createNode('ul');

		  img.src = "..\\..\\UI Icons\\Web App\\pill_white.png";
		  img.style.height = "30px";
		  img.style.float = "left";
		  img.style.marginLeft = "10px";
		  
		  h5.style.float = "left";
		  h5.style.color = "white";
		  h5.innerHTML = "&nbsp Rx- " + `${x}`;

		  ul.classList.add("list-group");
		  ul.classList.add("list-group-flush");
		  
		  append(div, img);
		  append(div, h5);
		  
		  med_load(prescription, ul).then( (finall) => {
			  ul = finall;
			  append(div, ul);
		  append(document.getElementById('infoo'), div);
		  })
		  .catch(function(error) {
			console.log(error);
			return error;
		}); 
		  
	}
		  
		//});
	  //}
	  
	  //)
  }
  
  
 async function med_load(prescription, ul){
	  
	  //const url = 'http://meth.azurewebsites.net/patient/' + resp['patients'][i];
	  var size = Object.size(prescription['medicines']);
	  console.log(prescription['medicines']);

		for (var key in prescription['medicines']) {
				console.log(key);
		  let span = createNode('span'),
			  li = createNode('li');

		  
		  li.innerHTML = `${key}` + " Left:" + `${prescription.remaining[key]}` + " ";
		  li.classList.add('list-group-item');
		  if(`${prescription.remaining[key]}` < 5){
			  span.innerHTML = "Refill";
			  span.classList.add('badge');
			  span.classList.add('badge-danger');
		  }
		  append(li, span);
		  append(ul, li);
		//});
	  //}
	  //i++;
	  }
	  return ul;
	  //)
  }
  
Object.size = function(obj) {
    var size = 0, key;
    for (key in obj) {
        if (obj.hasOwnProperty(key)) size++;
    }
    return size;
};

// Get the size of an object
//var size = Object.size(myArray);
  