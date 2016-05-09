window.onload = function() {
	var bgC = new RGBA(255, 255, 255, 0);

	var champList = document.getElementById("championList").getElementsByTagName("tr");
	var found = document.getElementById("infoFoundDisplay");
	var tableFound = document.getElementById("infoFoundDisplayTable");
	var notFound = document.getElementById("infoNotFoundDisplay");
	var enemy = document.getElementById("enemyName");
	var lane = document.getElementById("role");

	var sorryMsg = "Sorry not enough people play against <b>" + enemy.innerText + "</b> in the <b>" + lane.innerText + "</b> lane.";

	if (champList.length == 1) {
		found.classList.add("hidden");
		tableFound.classList.add("hidden");
		notFound.classList.remove("hidden");
		enemy.innerHTML = sorryMsg;
		lane.innerHTML = "";
		enemy.style.marginBottom= "20px"
	} else {
		found.classList.remove("hidden");
		tableFound.classList.remove("hidden");
		notFound.classList.add("hidden");
	}

	for (var i = champList.length - 1; i > 0; i--) {
		var percent = parseInt(champList[i].children[1].innerText);
		setBG(champList[i], percent);
		if (i == 2) bgC = new RGBA(255, 245, 153, 0);
		if (i == 3) bgC = new RGBA(188, 219, 246, 0);
		if (i == 4) bgC = new RGBA(169, 239, 222, 0);
	}

	function RGBA(red, green, blue, alpha) {
		this.red = red;
		this.green = green;
		this.blue = blue;
		this.alpha = alpha;
		this.getCSS = function() {
			return "rgba(" + this.red + ", " + this.green + ", " + this.blue + ", " + this.alpha + ")";
		}
	}

	function setBG(elem, opac) {
		bgC.alpha = opac/2/100;
		elem.style.backgroundColor = bgC.getCSS();
	}
}