window.onload = function() {
	var bgC = new RGBA(255, 255, 255, 0);

	var champList = document.getElementById("championList").getElementsByTagName("tr");

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
