window.onload = function() {
	var champList = document.getElementById("championList").getElementsByTagName("tr");
	for (var i = champList.length - 1; i > 0; i--) {
		var percent = parseInt(champList[i].children[1].innerText);
		setBG(champList[i], percent);
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
		var bgC = new RGBA(255, 255, 255, 0);
		bgC.alpha = opac/2/100;
		elem.style.backgroundColor = bgC.getCSS();
	}
}