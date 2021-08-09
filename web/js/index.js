function pick_tool(elem) {
	let nav = document.getElementById("nav")
	nav.classList.remove("col-md-3")
	nav.classList.add("col-md-1")

	let content = document.getElementById("content")
	content.classList.remove("col-md-9")
	content.classList.add("col-md-11")

	elem.classList.add("picked")
	elem.onclick = (event) => {
		remove_tool(event.currentTarget)
	}

	event = new Event('pick')
	elem.dispatchEvent(event)
}

function remove_tool(elem) {
	let nav = document.getElementById("nav")
	nav.classList.remove("col-md-1")
	nav.classList.add("col-md-3")

	let content = document.getElementById("content")
	content.classList.remove("col-md-11")
	content.classList.add("col-md-9")

	elem.classList.remove("picked")
	elem.onclick = (event) => {
		pick_tool(event.currentTarget)
	}

	event = new Event('depick')
	elem.dispatchEvent(event)
}

function init() {
	let elems = document.getElementsByClassName("node")
	for(i = 0;i < elems.length;i++) {
		elems[i].onclick = (event) => {
			event.target.onclick = null
			pick_tool(event.currentTarget)
		}
	}

	document.getElementById("item_red").addEventListener('pick', () => {
		html = document.getElementById("storage_item_red").innerHTML
		document.getElementById("content").innerHTML = html
	}, false)

	document.getElementById("item_red").addEventListener('depick', () => {
		content = document.getElementById("content")
		html = content.innerHTML
		document.getElementById("storage_item_red").innerHTML = html
		content.innerHTML = ""
	}, false)
}	