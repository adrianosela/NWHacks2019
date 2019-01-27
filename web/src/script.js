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

$( "#patients" )
  .mouseover(function() {
    //$( ".sick" ).animate( {top: "-55px;"} );
    console.log("fuuuuck");
  })
  .mouseout(function() {
    $( "p:first", this ).text( "mouse out" );
  });


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

  async function patient_load(){
	  
	  const ul = document.getElementById('authors');
	  const url = 'https://randomuser.me/api/?results=10';
	  
	  json_load(url).then( (resp) => {
		return resp.results.map(function(patient) {
		  let img = createNode('img'),
			  p = createNode('p'),
			  a = createNode('a');
			  //details = createNode('p');
		  //img.src = patient.picture.medium;
		  img.src = '../../UI Icons/Web App/ICON_Female.png';
		  img.style.height = "30px";
		  img.style.float = "left";
		  
		  a.innerHTML = '<div class=\"d-flex w-100 justify-content-between"><h5 class="mb-1">' + `${patient.name.first} ${patient.name.last}` + '</h5><small>3 days ago</small></div><p class="mb-1">Donec id elit non mi porta gravida at eget metus. Maecenas sed diam eget risus varius blandit.</p><small>Donec id elit non mi porta.</small>' ;
		  a.classList.add('list-group-item');
		  a.classList.add('list-group-item-action');
		  a.classList.add('flex-column');
		  a.classList.add('align-items-start');
		  a.setAttribute("id", "single-patient");
		  a.href = "patient.html";
		  //a.style.width = "70%";
		  a.style.float = "left";
		  
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