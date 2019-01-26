function createNode(element) {
      return document.createElement(element);
  }

  function append(parent, el) {
    return parent.appendChild(el);
  }
  
  function json_load(url){
	  fetch(url)
		  .then(function(response) {
			return response.json();
		  })
		  .catch(function(error) {
			console.log(error);
		});
  }

  function patient_load(){
	  const ul = document.getElementById('authors');
	  const url = 'https://randomuser.me/api/?results=10';
	  authors = json_load(url);
	  console.log(authors);
		return authors.map(function(patient) {
		  let li = createNode('li'),
			  img = createNode('img'),
			  span = createNode('span');
			  //details = createNode('p');
		  img.src = patient.picture.medium;
		  span.innerHTML = `${patient.name.first} ${patient.name.last}`;
		  //details.innerHTML = `${patient.name.first} ${patient.name.last}`;
		  append(li, img);
		  append(li, span);
		  append(document.getElementById('authors'), li);
		});
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