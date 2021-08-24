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

	let event = new Event('pick')
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

	let event = new Event('depick')
	elem.dispatchEvent(event)
}


let tags = new Set()
let items = []
let filtered = []
let glob_id = 0

async function request_items() {
	response = await fetch("/items", {
		method: "GET",
	})

	items = await response.json()
	for(i = 0;i < items.length;i++) {
		items[i].uniq_id = glob_id++
		items[i].deleted = false
	}
}

function pick_item_node(id) {
	console.log(items[id])

	let item = items[id]

	let oldPicked = document.querySelector("#item-list .node.picked")
	if(oldPicked != null)
		oldPicked.classList.remove("picked")
	let elem = document.getElementById(`item${id}`)
	if(elem != null) {
		elem.classList.add("picked")
	}

	document.getElementById("remove-item").style.display = "block"
	document.getElementById("remove-item").dataset.id = item.uniq_id
	
	document.getElementsByClassName("redactor")[0].style.display = "block"

	let name_label = document.getElementById("name-label")
	name_label.innerHTML = item.Name
	let item_name = document.getElementById("item-name")
	item_name.value = item.Name
	item_name.oninput = () => {
		let val = item_name.value
		name_label.innerHTML = val
		items[id].Name = val
		document.querySelector(`#item${item.uniq_id} span`).innerHTML = `${items[id].ID} ${items[id].Name}`
	}
	
	let item_image = document.getElementById("item-image")
	item_image.src = `/image/icon/${item.IconName}.png`	
	item_image_path = document.getElementById("item-image-path")
	item_image_path.value = item.IconName
	item_image_path.oninput = () => {
		let val = `/image/icon/${item_image_path.value}.png`
		item_image.src = val
		items[id].IconName = item_image_path.value
		document.querySelector(`#item${items[id].uniq_id} img`).src = val
	}

	item_id = document.getElementById("item-id")
	item_id.value = item.ID
	item_id.oninput = () => {
		let val = item_id.value
		items[id].ID = val
		document.querySelector(`#item${items[id].uniq_id} span`).innerHTML = `${items[id].ID} ${items[id].Name}`
	}

	// document.getElementById("item-type").value = item.ItemType
	// document.getElementById("ground-item-model").value = item.ModelGround
	// document.getElementById("item-description").innerHTML = item.Description
	// document.getElementById("item-script").value = item.UsageScript
	// document.getElementById("item-level").value = item.CharacterLevel
	// document.getElementById("item-stack").value = item.MaxStackSize
	// document.getElementById("item-price").value = item.Price
	// document.getElementById("item-slots").value = item.GemSockets
	// document.getElementById("item-effect").value = item.DisplayEffect
	// document.getElementById("item-effect-second").value = item.ItemUsageEffect
	
	item_type = document.getElementById("item-type")
	item_type.value = item.ItemType
	item_type.oninput = ()=> {
		items[id].ItemType = item_type.value
	}
	ground_item_model = document.getElementById("ground-item-model")
	ground_item_model.value = item.ModelGround
	ground_item_model.oninput = ()=> {
		items[id].ModelGround = ground_item_model.value
	}
	item_description = document.getElementById("item-description")
	item_description.value = item.Description
	item_description.oninput = ()=> {
		items[id].Description = item_description.value
	}
	item_script = document.getElementById("item-script")
	item_script.value = item.UsageScript
	item_script.oninput = ()=> {
		items[id].UsageScript = item_script.value
	}
	item_level = document.getElementById("item-level")
	item_level.value = item.CharacterLevel
	item_level.oninput = ()=> {
		items[id].CharacterLevel = item_level.value
	}
	item_stack = document.getElementById("item-stack")
	item_stack.value = item.MaxStackSize
	item_stack.oninput = ()=> {
		items[id].MaxStackSize = item_stack.value
	}
	item_price = document.getElementById("item-price")
	item_price.value = item.Price
	item_price.oninput = ()=> {
		items[id].Price = item_price.value
	}
	item_slots = document.getElementById("item-slots")
	item_slots.value = item.GemSockets
	item_slots.oninput = ()=> {
		items[id].GemSockets = item_slots.value
	}
	item_effect = document.getElementById("item-effect")
	item_effect.value = item.DisplayEffect
	item_effect.oninput = ()=> {
		items[id].DisplayEffect = item_effect.value
	}
	item_effect_second = document.getElementById("item-effect-second")
	item_effect_second.value = item.ItemUsageEffect
	item_effect_second.oninput = ()=> {
		items[id].ItemUsageEffect = item_effect_second.value
	}


	// document.getElementById("item-repair").checked = (item.Repairable == "1")
	// document.getElementById("item-sell").checked = (item.Tradable == "1")
	// document.getElementById("item-take").checked = (item.PickUpable == "1")
	// document.getElementById("item-throw").checked = (item.Droppable == "1")
	// document.getElementById("item-delete").checked = (item.Deletable == "1")

	item_repair = document.getElementById("item-repair")
	item_repair.checked = (item.Repairable == "1")
	item_repair.onchange = ()=> {
		if(item_repair.checked) {
			items[id].Repairable = "1"
		} else {
			items[id].Repairable = "0"
		}
	}
	item_sell = document.getElementById("item-sell")
	item_sell.checked = (item.Tradable == "1")
	item_sell.onchange = ()=> {
		if(item_sell.checked) {
			items[id].Tradable = "1"
		} else {
			items[id].Tradable = "0"
		}
	}
	item_take = document.getElementById("item-take")
	item_take.checked = (item.PickUpable == "1")
	item_take.onchange = ()=> {
		if(item_take.checked) {
			items[id].PickUpable = "1"
		} else {
			items[id].PickUpable = "0"
		}
	}
	item_throw = document.getElementById("item-throw")
	item_throw.checked = (item.Droppable == "1")
	item_throw.onchange = ()=> {
		if(item_throw.checked) {
			items[id].Droppable = "1"
		} else {
			items[id].Droppable = "0"
		}
	}
	item_delete = document.getElementById("item-delete")
	item_delete.checked = (item.Deletable == "1")
	item_delete.onchange = ()=> {
		if(item_delete.checked) {
			items[id].Deletable = "1"
		} else {
			items[id].Deletable = "0"
		}
	}

	
	let races_arr = [
		"item-lance",
		"item-carsise",
		"item-phyllis",
		"item-ami"
	]
	if(typeof item.CharacterTypes === 'undefined') {
		for(i = 0;i < 4;i++) {
			document.getElementById(races_arr[i]).checked = false
		}
	} else {
		let types = item.CharacterTypes
		if(types == "0" || types == "-1") {
			for(i = 0;i < 4;i++) {
				document.getElementById(races_arr[i]).checked = true
			}			
		} else {
			for(i = 0;i < 4;i++) {
				document.getElementById(races_arr[i]).checked = false
			}
			if(item.CharacterTypes != "") {
				races = types.split(",")
				for(i = 0;i < races.length;i++) {
					document.getElementById(races_arr[parseInt(races[i]) - 1]).checked = true
				}
			}
		}
	}
	item_lance = document.getElementById("item-lance")
	item_lance.onchange = () => {
		ft = true
		if(item_lance.checked) {
			picked = items[id].CharacterTypes.split(",")
			if(items[id].CharacterTypes == "") {
				picked = []
			}
			picked.push("1")
			picked.sort()
			if(picked.length == 4) {
				items[id].CharacterTypes = "0"
				return
			}
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(i == 0 || picked[i] != picked[i - 1]) {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
		} else {
			if(items[id].CharacterTypes == "0") {
				items[id].CharacterTypes = "1,2,3,4"
			}		
			picked = items[id].CharacterTypes.split(",")
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(picked[i] != "1") {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
			// if(items[id].CharacterTypes == "") {
			// 	items[id].CharacterTypes = "0"
			// }
		}
	}
	item_carsise = document.getElementById("item-carsise")
	item_carsise.onchange = () => {
		ft = true
		if(item_carsise.checked) {
			picked = items[id].CharacterTypes.split(",")
			if(items[id].CharacterTypes == "") {
				picked = []
			}
			picked.push("2")
			picked.sort()
			if(picked.length == 4) {
				items[id].CharacterTypes = "0"
				return
			}
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(i == 0 || picked[i] != picked[i - 1]) {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
		} else {
			if(items[id].CharacterTypes == "0") {
				items[id].CharacterTypes = "1,2,3,4"
			}		
			picked = items[id].CharacterTypes.split(",")
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(picked[i] != "2") {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
			// if(items[id].CharacterTypes == "") {
			// 	items[id].CharacterTypes = "0"
			// }
		}
	}
	item_phyllis = document.getElementById("item-phyllis")
	item_phyllis.onchange = () => {
		ft = true
		if(item_phyllis.checked) {
			picked = items[id].CharacterTypes.split(",")
			if(items[id].CharacterTypes == "") {
				picked = []
			}
			picked.push("3")
			picked.sort()
			if(picked.length == 4) {
				items[id].CharacterTypes = "0"
				return
			}
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(i == 0 || picked[i] != picked[i - 1]) {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
		} else {
			if(items[id].CharacterTypes == "0") {
				items[id].CharacterTypes = "1,2,3,4"
			}		
			picked = items[id].CharacterTypes.split(",")
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(picked[i] != "3") {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
			// if(items[id].CharacterTypes == "") {
			// 	items[id].CharacterTypes = "0"
			// }
		}
	}
	item_ami = document.getElementById("item-ami")
	item_ami.onchange = () => {
		ft = true
		if(item_ami.checked) {
			picked = items[id].CharacterTypes.split(",")
			if(items[id].CharacterTypes == "") {
				picked = []
			}
			picked.push("4")
			picked.sort()
			if(picked.length == 4) {
				items[id].CharacterTypes = "0"
				return
			}
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(i == 0 || picked[i] != picked[i - 1]) {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
		} else {
			if(items[id].CharacterTypes == "0") {
				items[id].CharacterTypes = "1,2,3,4"
			}		
			picked = items[id].CharacterTypes.split(",")
			items[id].CharacterTypes = ""
			for(i = 0;i < picked.length;i++) {
				if(picked[i] != "4") {
					if(!ft) {
						items[id].CharacterTypes += ","
					}
					items[id].CharacterTypes += picked[i]
					ft = false
				}
			}
			// if(items[id].CharacterTypes == "") {
			// 	items[id].CharacterTypes = "0"
			// }
		}
	}



	item_model_lance = document.getElementById("item-model-lance")
	item_model_lance.value = item.ModelLance
	item_model_lance.oninput = ()=> {
		items[id].ModelLance = item_model_lance.value
	}
	item_model_carsise = document.getElementById("item-model-carsise")
	item_model_carsise.value = item.ModelCarsise
	item_model_carsise.oninput = ()=> {
		items[id].ModelCarsise = item_model_carsise.value
	}
	item_model_phyllis = document.getElementById("item-model-phyllis")
	item_model_phyllis.value = item.ModelPhyllis
	item_model_phyllis.oninput = ()=> {
		items[id].ModelPhyllis = item_model_phyllis.value
	}
	item_model_ami = document.getElementById("item-model-ami")
	item_model_ami.value = item.ModelAmi
	item_model_ami.oninput = ()=> {
		items[id].ModelAmi = item_model_ami.value
	}

	let classes_arr = [
		"item-begginer",
		"item-swordmaster",
		"item-hunter",
		"item-sailor",
		"-",
		"item-healer",
		"-",
		"-",
		"item-champion",
		"item-warrior",
		"-",
		"-",
		"item-archer",
		"item-priest",
		"item-sealmaster",
		"-",
		"item-voyager",
		"-",
		"-"
	]

	valid_count = 0
	for(i = 0;i < classes_arr.length;i++) {
		if(classes_arr[i] != "-") {
			valid_count++
		}
	}
	valid_str = "1"
	for(i = 1;i < classes_arr.length;i++) {
		if(classes_arr[i] != "-") {
			valid_str += "," + (i + 1).toString()
		}
	}

	mp = new Map()
	for(i = 0;i < classes_arr.length;i++) {
		mp.set(classes_arr[i], i + 1)
	}

	let classes = item.CharacterClasses
	if(classes == "0" || classes == "-1") {
		for(i = 0;i < classes_arr.length;i++) {
			if(classes_arr[i] != "-")
				document.getElementById(classes_arr[i]).checked = true
		}
	} else {
		
		if(classes != "" && classes != undefined) {
			let classes_splt = classes.split(",")

			for(i = 0;i < classes_arr.length;i++) {
				if(classes_arr[i] != "-")
					document.getElementById(classes_arr[i]).checked = false
			}

			for(i = 0;i < classes_splt.length;i++) {
				if(classes_arr[classes_splt[i] - 1] != "-")
					document.getElementById(classes_arr[classes_splt[i] - 1]).checked = true
			}
		}
	}

	for(i = 0;i < classes_arr.length;i++) {
		if(classes_arr[i] != "-") {
			elm = document.getElementById(classes_arr[i])
			elm.onchange = (event) => {
				index = mp.get(event.target.id).toString()
				if(event.target.checked) {
					if(items[id].CharacterClasses == "") {
						items[id].CharacterClasses = index
						return
					}

					if(items[id].CharacterClasses == "0") {
						items[id].CharacterClasses = valid_str
					}

					splt = items[id].CharacterClasses.split(",")
					splt.push(index)
					if(splt.length >= valid_count) {
						items[id].CharacterClasses = "0"
						return
					}
					splt.sort()
					ft = true
					for(ii = 0;ii < splt.length;ii++) {
						if(ii == 0 || splt[ii] != splt[ii - 1]) {
							if(!ft) {
								items[id].CharacterClasses += ","
								items[id].CharacterClasses += splt[ii]
							} else {
								items[id].CharacterClasses = splt[ii]
							}
							ft = false
						}
					}
				} else {
					
					if(items[id].CharacterClasses == "") {
						return
					}
					if(items[id].CharacterClasses == "0") {
						items[id].CharacterClasses = valid_str
					}

					splt = items[id].CharacterClasses.split(",")
					items[id].CharacterClasses = ""
					splt.sort()
					ft = true
					for(ii = 0;ii < splt.length;ii++) {
						if(splt[ii] != index) {
							if(!ft) {
								items[id].CharacterClasses += ","
							}
							items[id].CharacterClasses += splt[ii]
							ft = false
						}
					}
				}
			}
		}
	}

	item_strength_add = document.getElementById("item-strength-add")
	item_strength_add.value = item.STRMinMax
	item_strength_add.oninput = ()=> {
		items[id].STRMinMax = item_strength_add.value
	}
	item_strength_mlt = document.getElementById("item-strength-mlt")
	item_strength_mlt.value = item.STR
	item_strength_mlt.oninput = ()=> {
		items[id].STR = item_strength_mlt.value
	}
	item_agility_add = document.getElementById("item-agility-add")
	item_agility_add.value = item.AGIMinMax
	item_agility_add.oninput = ()=> {
		items[id].AGIMinMax = item_agility_add.value
	}
	item_agility_mlt = document.getElementById("item-agility-mlt")
	item_agility_mlt.value = item.AGI
	item_agility_mlt.oninput = ()=> {
		items[id].AGI = item_agility_mlt.value
	}
	item_vitality_add = document.getElementById("item-vitality-add")
	item_vitality_add.value = item.CONMinMax
	item_vitality_add.oninput = ()=> {
		items[id].CONMinMax = item_vitality_add.value
	}
	item_vitality_mlt = document.getElementById("item-vitality-mlt")
	item_vitality_mlt.value = item.CON
	item_vitality_mlt.oninput = ()=> {
		items[id].CON = item_vitality_mlt.value
	}
	item_spirit_add = document.getElementById("item-spirit-add")
	item_spirit_add.value = item.SPRMinMax
	item_spirit_add.oninput = ()=> {
		items[id].SPRMinMax = item_spirit_add.value
	}
	item_spirit_mlt = document.getElementById("item-spirit-mlt")
	item_spirit_mlt.value = item.SPR
	item_spirit_mlt.oninput = ()=> {
		items[id].SPR = item_spirit_mlt.value
	}
	item_accuracy_add = document.getElementById("item-accuracy-add")
	item_accuracy_add.value = item.ACCMinMax
	item_accuracy_add.oninput = ()=> {
		items[id].ACCMinMax = item_accuracy_add.value
	}
	item_accuracy_mlt = document.getElementById("item-accuracy-mlt")
	item_accuracy_mlt.value = item.ACC
	item_accuracy_mlt.oninput = ()=> {
		items[id].ACC = item_accuracy_mlt.value
	}
	item_luck_add = document.getElementById("item-luck-add")
	item_luck_add.value = item.LUKMinMax
	item_luck_add.oninput = ()=> {
		items[id].LUKMinMax = item_luck_add.value
	}
	item_luck_mlt = document.getElementById("item-luck-mlt")
	item_luck_mlt.value = item.LUK
	item_luck_mlt.oninput = ()=> {
		items[id].LUK = item_luck_mlt.value
	}
	item_min_attack_add = document.getElementById("item-min-attack-add")
	item_min_attack_add.value = item.MinAttackAdd
	item_min_attack_add.oninput = ()=> {
		items[id].MinAttackAdd = item_min_attack_add.value
	}
	item_min_attack_mlt = document.getElementById("item-min-attack-mlt")
	item_min_attack_mlt.value = item.MinAttack
	item_min_attack_mlt.oninput = ()=> {
		items[id].MinAttack = item_min_attack_mlt.value
	}
	item_max_attack_add = document.getElementById("item-max-attack-add")
	item_max_attack_add.value = item.MaxAttackAdd
	item_max_attack_add.oninput = ()=> {
		items[id].MaxAttackAdd = item_max_attack_add.value
	}
	item_max_attack_mlt = document.getElementById("item-max-attack-mlt")
	item_max_attack_mlt.value = item.MaxAttack
	item_max_attack_mlt.oninput = ()=> {
		items[id].MaxAttack = item_max_attack_mlt.value
	}
	item_attack_speed_add = document.getElementById("item-attack-speed-add")
	item_attack_speed_add.value = item.AttackSpeedAdd
	item_attack_speed_add.oninput = ()=> {
		items[id].AttackSpeedAdd = item_attack_speed_add.value
	}
	item_attack_speed_mlt = document.getElementById("item-attack-speed-mlt")
	item_attack_speed_mlt.value = item.AttackSpeed
	item_attack_speed_mlt.oninput = ()=> {
		items[id].AttackSpeed = item_attack_speed_mlt.value
	}
	item_hit_chance_add = document.getElementById("item-hit-chance-add")
	item_hit_chance_add.value = item.HitRateAdd
	item_hit_chance_add.oninput = ()=> {
		items[id].HitRateAdd = item_hit_chance_add.value
	}
	item_hit_chance_mlt = document.getElementById("item-hit-chance-mlt")
	item_hit_chance_mlt.value = item.HitRate
	item_hit_chance_mlt.oninput = ()=> {
		items[id].HitRate = item_hit_chance_mlt.value
	}
	item_crit_chance_add = document.getElementById("item-crit-chance-add")
	item_crit_chance_add.value = item.CriticalRateAdd
	item_crit_chance_add.oninput = ()=> {
		items[id].CriticalRateAdd = item_crit_chance_add.value
	}
	item_crit_chance_mlt = document.getElementById("item-crit-chance-mlt")
	item_crit_chance_mlt.value = item.CriticalRate
	item_crit_chance_mlt.oninput = ()=> {
		items[id].CriticalRate = item_crit_chance_mlt.value
	}
	item_move_speed_add = document.getElementById("item-move-speed-add")
	item_move_speed_add.value = item.MovementSpeedAdd
	item_move_speed_add.oninput = ()=> {
		items[id].MovementSpeedAdd = item_move_speed_add.value
	}
	item_move_speed_mlt = document.getElementById("item-move-speed-mlt")
	item_move_speed_mlt.value = item.MovementSpeed
	item_move_speed_mlt.oninput = ()=> {
		items[id].MovementSpeed = item_move_speed_mlt.value
	}
	item_max_hp_add = document.getElementById("item-max-hp-add")
	item_max_hp_add.value = item.MaxHPAdd
	item_max_hp_add.oninput = ()=> {
		items[id].MaxHPAdd = item_max_hp_add.value
	}
	item_max_hp_mlt = document.getElementById("item-max-hp-mlt")
	item_max_hp_mlt.value = item.MaxHP
	item_max_hp_mlt.oninput = ()=> {
		items[id].MaxHP = item_max_hp_mlt.value
	}
	item_max_mp_add = document.getElementById("item-max-mp-add")
	item_max_mp_add.value = item.MaxSPAdd
	item_max_mp_add.oninput = ()=> {
		items[id].MaxSPAdd = item_max_mp_add.value
	}
	item_max_mp_mlt = document.getElementById("item-max-mp-mlt")
	item_max_mp_mlt.value = item.MaxSP
	item_max_mp_mlt.oninput = ()=> {
		items[id].MaxSP = item_max_mp_mlt.value
	}
	item_hp_regen_add = document.getElementById("item-hp-regen-add")
	item_hp_regen_add.value = item.HPRecoveryAdd
	item_hp_regen_add.oninput = ()=> {
		items[id].HPRecoveryAdd = item_hp_regen_add.value
	}
	item_hp_regen_mlt = document.getElementById("item-hp-regen-mlt")
	item_hp_regen_mlt.value = item.HPRecovery
	item_hp_regen_mlt.oninput = ()=> {
		items[id].HPRecovery = item_hp_regen_mlt.value
	}
	item_mp_regen_add = document.getElementById("item-mp-regen-add")
	item_mp_regen_add.value = item.SPRecoveryAdd
	item_mp_regen_add.oninput = ()=> {
		items[id].SPRecoveryAdd = item_mp_regen_add.value
	}
	item_mp_regen_mlt = document.getElementById("item-mp-regen-mlt")
	item_mp_regen_mlt.value = item.SPRecovery
	item_mp_regen_mlt.oninput = ()=> {
		items[id].SPRecovery = item_mp_regen_mlt.value
	}
	item_armour_add = document.getElementById("item-armour-add")
	item_armour_add.value = item.DefenseAdd
	item_armour_add.oninput = ()=> {
		items[id].DefenseAdd = item_armour_add.value
	}
	item_armour_mlt = document.getElementById("item-armour-mlt")
	item_armour_mlt.value = item.Defense
	item_armour_mlt.oninput = ()=> {
		items[id].Defense = item_armour_mlt.value
	}
	item_evasion_add = document.getElementById("item-evasion-add")
	item_evasion_add.value = item.DodgeAdd
	item_evasion_add.oninput = ()=> {
		items[id].DodgeAdd = item_evasion_add.value
	}
	item_evasion_mlt = document.getElementById("item-evasion-mlt")
	item_evasion_mlt.value = item.DodgeRate
	item_evasion_mlt.oninput = ()=> {
		items[id].DodgeRate = item_evasion_mlt.value
	}
	item_energy_add = document.getElementById("item-energy-add")
	item_energy_add.value = item.Energy
	item_energy_add.oninput = ()=> {
		items[id].Energy = item_energy_add.value
	}
	item_durability_add = document.getElementById("item-durability-add")
	item_durability_add.value = item.Durability
	item_durability_add.oninput = ()=> {
		items[id].Durability = item_durability_add.value
	}
	item_phys_armour_add = document.getElementById("item-phys-armour-add")
	item_phys_armour_add.value = item.PhysicalResistAdd
	item_phys_armour_add.oninput = ()=> {
		items[id].PhysicalResistAdd = item_phys_armour_add.value
	}
}

function make_node(item) {
	return `
		<div class="node" id="item${item.uniq_id}" onclick="pick_item_node(${item.uniq_id})">
			<img src="/image/icon/${item.IconName}.png">
			<span>${item.ID} ${item.Name}</span>
		</div>`
}

function htmlToElement(html) {
    var template = document.createElement('template');
    html = html.trim();
    template.innerHTML = html;
    return template.content.firstChild;
}

async function update_list() {
	let list = document.getElementById("item-list")
	list.innerHTML = ""
	for(i = 0;i < Math.min(400, items.length);i++) {
		let item = items[i]
		if(item.deleted) {
			continue
		}
		
		list.insertAdjacentElement('beforeend', htmlToElement(make_node(item)))
	}
}

async function update_with_search(search) {
	let inserted = 0
	let list = document.getElementById("item-list")
	list.innerHTML = ""
	for(i = 0;i < items.length && inserted < 200;i++) {
		if(items[i].deleted) {
			continue
		}
		let name = `${items[i].ID} ${items[i].Name}`
		if(name.includes(search)) {		
			inserted++
			list.insertAdjacentElement('beforeend', htmlToElement(make_node(items[i])))
		}
	}
}

async function init() {
	let elems = document.getElementsByClassName("node")
	for(i = 0;i < elems.length;i++) {
		elems[i].onclick = (event) => {
			event.target.onclick = null
			pick_tool(event.currentTarget)
		}
	}

	document.getElementById("item_red").addEventListener('pick', () => {
		document.getElementById("item-redactor").style.display = "block"
		// html = document.getElementById("storage_item_red").innerHTML
		// document.getElementById("content").innerHTML = html
	}, false)

	document.getElementById("item_red").addEventListener('depick', () => {
		document.getElementById("item-redactor").style.display = "none"
		// content = document.getElementById("content")
		// html = content.innerHTML
		// document.getElementById("storage_item_red").innerHTML = html
		// content.innerHTML = ""
	}, false)

	document.getElementById("item-search-bar").addEventListener('input', () => {
		patter = []
		list = document.getElementById("item-list")
		list.innerHTML = ""
		search = document.getElementById("item-search-bar").value 
		if(search == "") {
			update_list()
			return
		}
		update_with_search(search)
	})

	document.getElementById("save-all").addEventListener('click', () => {
		data = []
		for(i = 0;i < items.length;i++) {
			if(!items[i].deleted) {
				data.push(items[i])
			}
		}
		console.log(data)
		const resp = fetch("/items", {
			method: "POST",
			body: JSON.stringify(data)
		})
		console.log(resp)
	})

	document.getElementById("add-item").addEventListener('click', () => {
		let item = {}
		for(key in items[0]) {
			item[key] = "0"
		}
		item.uniq_id = glob_id++
		item.deleted = false
		item.Name = "NEW ITEM"
		items.push(item)
		document.getElementById("item-list").insertAdjacentElement('afterbegin', htmlToElement(make_node(item)))
		pick_item_node(item.uniq_id)
	})

	document.getElementById("remove-item").addEventListener('click', (event) => {
		let id = event.target.dataset.id
		console.log(id)
		items[id].deleted = true
		let search = document.getElementById("item-search-bar").value
		if(search == "") {
			update_list()
		} else {
			update_with_search(search)
		}

		event.target.style.display = "none"
		document.getElementsByClassName("redactor")[0].style.display = "none"
	})

	await request_items()
	update_list()
}	