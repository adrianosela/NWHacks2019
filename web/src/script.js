function session_start(){
	sessionStorage.doctor = "";
	sessionStorage.patient = "";
}

function set_doctor(doctor){
	sessionStorage.doctor_id = doctor;
}

function set_session(name, values){
	sessionStorage.setItem(name, values);
}

function get_session(name){
	return sessionStorage.getItem(name);
}

function createNode(element) {
      return document.createElement(element);
  }

function append(parent, el) {
    return parent.appendChild(el);
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

  async function patient_load(){
	  
	  const ul = document.getElementById('authors');
	  const url = 'https://randomuser.me/api/?results=10';
	  
	  json_load(url).then( (resp) => {
		return resp.results.map(function(patient) {
		  let li = createNode('li'),
			  p = createNode('p'),
			  a = createNode('a');
			  //details = createNode('p');
		  //img.src = patient.picture.medium;
		  a.innerHTML = '<h4>' + `${patient.name.first} ${patient.name.last}` + '</h4>';
		  a.classList.add('list-group-item');
		  a.classList.add('list-group-item-action');
		  a.href = "patient.html";
		  p.innerHTML = `${patient.location.street} ${patient.location.city}`;
		  //details.innerHTML = `${patient.name.first} ${patient.name.last}`;
		  //append(li, img);
		  append(a, p);
		  append(document.getElementById('patients'), a);
		});
	  }
	  
	  )
  }
  
	function prescription_load(){
	  const ul = document.getElementById('authors');
	  const url = 'https://randomuser.me/api/?results=10';
	  fetch(url)
	  .then((resp) => resp.json())
	  .then(function(data) {
		let authors = data.results;
		return authors.map(function(author) {
		  let li = createNode('li'),
			  img = createNode('img'),
			  span = createNode('span');
		  img.src = author.picture.medium;
		  span.innerHTML = `${author.name.first} ${author.name.last}`;
		  append(li, img);
		  append(li, span);
		  append(document.getElementById('authors'), li);
		})
	  })
	  .catch(function(error) {
		console.log(error);
	  });
  }