function change_select_003() {
	var select = document.getElementById("viewPhone");
	var selectedOption = select.options[select.selectedIndex].value;

	var modelDiv = document.getElementById("model-div");
	var brandDiv = document.getElementById("brand-div");
	var buttonsDiv = document.getElementById("buttons-div");

	modelDiv.classList.add("hidden");
	brandDiv.classList.add("hidden");
	buttonsDiv.classList.add("hidden");

	if (selectedOption === "applePhone") {
		modelDiv.classList.remove("hidden");
	} else if (selectedOption === "androidPhone") {
		brandDiv.classList.remove("hidden");
		modelDiv.classList.remove("hidden");
	} else if (selectedOption === "radioPhone") {
		buttonsDiv.classList.remove("hidden");
		brandDiv.classList.remove("hidden");
		modelDiv.classList.remove("hidden");
	}
} 

function submitFormTask_003() {
	var el_viewPhone = document.getElementById("viewPhone");
	var el_brand = document.getElementById("brand").value;
	var el_model = document.getElementById("model").value;
	var el_buttons = document.getElementById("buttons").value;

	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/process_task_003_01_calc", true);
	xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
	xhr.onreadystatechange = function () {
		if (xhr.readyState === 4 && xhr.status === 200) {
			document.getElementById("output").innerHTML = xhr.responseText;
			
		}
	};

	var data = "viewPhone=" + encodeURIComponent(el_viewPhone.value)+ "&brand=" + encodeURIComponent(el_brand)+ "&model=" + encodeURIComponent(el_model)+ "&buttons=" + encodeURIComponent(el_buttons);
	xhr.send(data);
}

function showFormTask_003() {
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/process_task_003_01_show", true);
	xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
	xhr.onreadystatechange = function () {
		if (xhr.readyState === 4 && xhr.status === 200) {
			document.getElementById("output").innerHTML = xhr.responseText;
		}
	};
	var data = "show=true";
	xhr.send(data);
}

function refreshForm_003(){
  change_select_003();
  showFormTask_003();
  //location.reload();
}

//task_004
function showFormTask_004() {
	var el_viewUnitType = document.getElementById("viewUnitType");

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask004Show", true);
    xhr.setRequestHeader("Content-Type", "application/json");  // Изменение типа содержимого на application/json
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("output").innerHTML = xhr.responseText;
     }
    };

    var data = {
      viewPhone: "45454",
      brand: el_viewUnitType.value  // Значение типа UnitType "inch", переданное как JSON-объект
    };

    xhr.send(JSON.stringify(data));  // Преобразование в JSON-строку
}

function refreshForm_004(){
	showFormTask_004();
}

//task_005
function showFormTask005() {
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask005Show", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("output").innerHTML = xhr.responseText;
     }
    };

    var data = {
      listSuspects: document.getElementById("comment").value  
    };

    xhr.send(JSON.stringify(data));  
}

function refreshForm005(){
	showFormTask005();
}

function clearAreaTask005() {
	document.getElementById("comment").value = "";
}

  function insertNameIntoAreaTask005(obj) {
	document.getElementById("comment").value += obj.id+"\n";
}

//task_006
function showSortArray006() {
	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask006ShowSortArray", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("output").innerHTML = xhr.responseText;
     }
    };

    var data = {
      listInputboxSize: parseInt(document.getElementById("inputboxSize").value, 10),  
	  listInputboxMax:  parseInt(document.getElementById("inputboxMax").value, 10),
	  listInputboxAsc:  JSON.parse(document.getElementById("inputboxAsc").value)  
    };

    xhr.send(JSON.stringify(data));  
}

function showSortArrayTest006() {

	document.getElementById("outputTest").innerHTML = "Ждите... :)";

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask006ShowSortArrayTest", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("outputTest").innerHTML = xhr.responseText;
     }
    };

    var data = {
	  listInputboxTypeSort: parseInt(document.getElementById("inputboxTypeSort").value, 10),   
    };

    xhr.send(JSON.stringify(data));  
}

//task_007
function task007loadPage() {
  task007treeCreate();
  task007graphBfsAction(1);
  task007graphDfsAction(1);
}

function task007treeCreate() {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask007treeCreate", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("outputSvgTree").innerHTML = xhr.responseText;
     }
    };

    var data = {
	  listInputbox: parseInt(document.getElementById("numElementTree").value, 10),   
    };

    xhr.send(JSON.stringify(data));  
}

function task007treeAction(action) {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask007treeAction", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("outputSvgTree").innerHTML = xhr.responseText;
     }
    };

    var data = {
	  listInputbox: parseInt(document.getElementById("numElementTree").value, 10),
	  listAction:   action,
    };

    xhr.send(JSON.stringify(data));  
}

function task007graphBfsAction(action) {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask007graphBfsAction", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("outputSvgGraphBfs").innerHTML = xhr.responseText;
     }
    };

    var data = {
	  numElementGraphBfsAVal: parseInt(document.getElementById("numElementGraphBfsA").value, 10),
    numElementGraphBfsBVal: parseInt(document.getElementById("numElementGraphBfsB").value, 10),
	  actionVal:   action,
    };

    xhr.send(JSON.stringify(data));  
}

function task007graphDfsAction(action) {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask007graphDfsAction", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
       document.getElementById("outputSvgGraphDfs").innerHTML = xhr.responseText;
     }
    };

    var data = {
	  numElementGraphDfsAVal: parseInt(document.getElementById("numElementGraphDfsA").value, 10),
    numElementGraphDfsBVal: parseInt(document.getElementById("numElementGraphDfsB").value, 10),
	  actionVal:   action,
    };

    xhr.send(JSON.stringify(data));  
}


//task_008
function task008mapEqualAction(action) {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask008mapEqualAction", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200 && action == 1) {
        document.getElementById("outputFind").innerHTML = xhr.responseText;
      }
    };

    var data = {
      inputString1Val: document.getElementById("inputString1").value,
      inputString2Val: document.getElementById("inputString2").value,
      actionVal:   action,
    };

    xhr.send(JSON.stringify(data));  
}

//task_009
function task009Action(action) {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask009Action", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {

        document.getElementById("output").innerHTML = xhr.responseText;
      }
    };

    var data = {
      inputAmount: parseInt(document.getElementById("inputAmount").value, 10),
      actionVal:   action,
    };

    xhr.send(JSON.stringify(data));  
}

//task_010
function task010Action(action) {

	var xhr = new XMLHttpRequest();
    xhr.open("POST", "/processHandlerTask010Action", true);
    xhr.setRequestHeader("Content-Type", "application/json");  
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200 && action == 1) {
        document.getElementById("outputFind").innerHTML = xhr.responseText;
      }
    };

    var data = {
      inputString1Val: document.getElementById("inputString1").value,
      inputString2Val: document.getElementById("inputString2").value,
      actionVal:   action,
    };

    xhr.send(JSON.stringify(data));  
}






