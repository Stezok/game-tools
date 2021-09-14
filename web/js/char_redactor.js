let characters = []
let glob_char_id = 0
function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}

function htmlToElement(html) {
    var template = document.createElement('template');
    html = html.trim();
    template.innerHTML = html;
    return template.content.firstChild;
}

async function request_characters() {
	response = await fetch("/characters", {
		method: "GET",
	})

	characters = await response.json()
	for(i = 0;i < characters.length;i++) {
		characters[i].uniq_id = glob_char_id++
		characters[i].deleted = false
	}
}

function make_char_node(char) {
	return `
		<div class="node" id="char${char.uniq_id}" onclick="pick_char_node(${char.uniq_id})">
			<span>${char.ID} ${char.Description}</span>
		</div>`
}

function pick_char_node(id) {
	let char = characters[id]

	let oldPicked = document.querySelector("#char-list .node.picked")
	if(oldPicked != null)
		oldPicked.classList.remove("picked")
	let elem = document.getElementById(`char${id}`)
	if(elem != null) {
		elem.classList.add("picked")
	}
	document.querySelector(`#char-red-container .redactor`).style.display = "flex"
	
	let remove = document.getElementById("remove-char")
	remove.style.display = "block"
	remove.dataset.id = id

	char_ID = document.getElementById("char-ID")
	char_ID.value = char.ID
	char_ID.oninput = () => {
		characters[id].ID = char_ID.value
		document.querySelector(`#char${id} span`).innerHTML = `${characters[id].ID} ${characters[id].Description}`
	}

	document.getElementById("char-Name").innerHTML = characters[id].Description
	char_Description = document.getElementById("char-Description")
	char_Description.value = char.Description
	char_Description.oninput = () => {
		characters[id].Description = char_Description.value
		document.getElementById("char-Name").innerHTML = characters[id].Description
		document.querySelector(`#char${id} span`).innerHTML = `${characters[id].ID} ${characters[id].Description}`
	}
	char_ModelCategory = document.getElementById("char-ModelCategory")
	char_ModelCategory.value = char.ModelCategory
	char_ModelCategory.oninput = () => {
		characters[id].ModelCategory = char_ModelCategory.value
	}
	char_ModelType = document.getElementById("char-ModelType")
	char_ModelType.value = char.ModelType
	char_ModelType.oninput = () => {
		characters[id].ModelType = char_ModelType.value
	}
	char_LogicType = document.getElementById("char-LogicType")
	char_LogicType.value = char.LogicType
	char_LogicType.oninput = () => {
		characters[id].LogicType = char_LogicType.value
	}
	char_FrameworkNumber = document.getElementById("char-FrameworkNumber")
	char_FrameworkNumber.value = char.FrameworkNumber
	char_FrameworkNumber.oninput = () => {
		characters[id].FrameworkNumber = char_FrameworkNumber.value
	}
	char_SuiteSerial = document.getElementById("char-SuiteSerial")
	char_SuiteSerial.value = char.SuiteSerial
	char_SuiteSerial.oninput = () => {
		characters[id].SuiteSerial = char_SuiteSerial.value
	}
	char_SuiteQuantity = document.getElementById("char-SuiteQuantity")
	char_SuiteQuantity.value = char.SuiteQuantity
	char_SuiteQuantity.oninput = () => {
		characters[id].SuiteQuantity = char_SuiteQuantity.value
	}
	char_PartFACE = document.getElementById("char-PartFACE")
	char_PartFACE.value = char.PartFACE
	char_PartFACE.oninput = () => {
		characters[id].PartFACE = char_PartFACE.value
	}
	char_PartHeadApparel = document.getElementById("char-PartHeadApparel")
	char_PartHeadApparel.value = char.PartHeadApparel
	char_PartHeadApparel.oninput = () => {
		characters[id].PartHeadApparel = char_PartHeadApparel.value
	}
	char_PartArmor = document.getElementById("char-PartArmor")
	char_PartArmor.value = char.PartArmor
	char_PartArmor.oninput = () => {
		characters[id].PartArmor = char_PartArmor.value
	}
	char_PartGloves = document.getElementById("char-PartGloves")
	char_PartGloves.value = char.PartGloves
	char_PartGloves.oninput = () => {
		characters[id].PartGloves = char_PartGloves.value
	}
	char_PartShoes = document.getElementById("char-PartShoes")
	char_PartShoes.value = char.PartShoes
	char_PartShoes.oninput = () => {
		characters[id].PartShoes = char_PartShoes.value
	}
	char_PartWeapon = document.getElementById("char-PartWeapon")
	char_PartWeapon.value = char.PartWeapon
	char_PartWeapon.oninput = () => {
		characters[id].PartWeapon = char_PartWeapon.value
	}
	char_Part2ndWeap = document.getElementById("char-Part2ndWeap")
	char_Part2ndWeap.value = char.Part2ndWeap
	char_Part2ndWeap.oninput = () => {
		characters[id].Part2ndWeap = char_Part2ndWeap.value
	}
	char_PartWings = document.getElementById("char-PartWings")
	char_PartWings.value = char.PartWings
	char_PartWings.oninput = () => {
		characters[id].PartWings = char_PartWings.value
	}
	char_FeffID = document.getElementById("char-FeffID")
	char_FeffID.value = char.FeffID
	char_FeffID.oninput = () => {
		characters[id].FeffID = char_FeffID.value
	}
	char_EeffID = document.getElementById("char-EeffID")
	char_EeffID.value = char.EeffID
	char_EeffID.oninput = () => {
		characters[id].EeffID = char_EeffID.value
	}
	char_SpecialEffectActionSerial = document.getElementById("char-SpecialEffectActionSerial")
	char_SpecialEffectActionSerial.value = char.SpecialEffectActionSerial
	char_SpecialEffectActionSerial.oninput = () => {
		characters[id].SpecialEffectActionSerial = char_SpecialEffectActionSerial.value
	}
	char_Shadow = document.getElementById("char-Shadow")
	char_Shadow.value = char.Shadow
	char_Shadow.oninput = () => {
		characters[id].Shadow = char_Shadow.value
	}
	char_ActionID = document.getElementById("char-ActionID")
	char_ActionID.value = char.ActionID
	char_ActionID.oninput = () => {
		characters[id].ActionID = char_ActionID.value
	}
	char_Transparency = document.getElementById("char-Transparency")
	char_Transparency.value = char.Transparency
	char_Transparency.oninput = () => {
		characters[id].Transparency = char_Transparency.value
	}
	char_MovingSoundEffect = document.getElementById("char-MovingSoundEffect")
	char_MovingSoundEffect.value = char.MovingSoundEffect
	char_MovingSoundEffect.oninput = () => {
		characters[id].MovingSoundEffect = char_MovingSoundEffect.value
	}
	char_BreathingSoundEffect = document.getElementById("char-BreathingSoundEffect")
	char_BreathingSoundEffect.value = char.BreathingSoundEffect
	char_BreathingSoundEffect.oninput = () => {
		characters[id].BreathingSoundEffect = char_BreathingSoundEffect.value
	}
	char_DeathSoundEffect = document.getElementById("char-DeathSoundEffect")
	char_DeathSoundEffect.value = char.DeathSoundEffect
	char_DeathSoundEffect.oninput = () => {
		characters[id].DeathSoundEffect = char_DeathSoundEffect.value
	}
	char_CanItBeControlled = document.getElementById("char-CanItBeControlled")
	char_CanItBeControlled.value = char.CanItBeControlled
	char_CanItBeControlled.oninput = () => {
		characters[id].CanItBeControlled = char_CanItBeControlled.value
	}
	char_AreaLimitation = document.getElementById("char-AreaLimitation")
	char_AreaLimitation.value = char.AreaLimitation
	char_AreaLimitation.oninput = () => {
		characters[id].AreaLimitation = char_AreaLimitation.value
	}
	char_AltitudeExcursion = document.getElementById("char-AltitudeExcursion")
	char_AltitudeExcursion.value = char.AltitudeExcursion
	char_AltitudeExcursion.oninput = () => {
		characters[id].AltitudeExcursion = char_AltitudeExcursion.value
	}
	char_ItemsThatCanBeEquipped = document.getElementById("char-ItemsThatCanBeEquipped")
	char_ItemsThatCanBeEquipped.value = char.ItemsThatCanBeEquipped
	char_ItemsThatCanBeEquipped.oninput = () => {
		characters[id].ItemsThatCanBeEquipped = char_ItemsThatCanBeEquipped.value
	}
	char_Length = document.getElementById("char-Length")
	char_Length.value = char.Length
	char_Length.oninput = () => {
		characters[id].Length = char_Length.value
	}
	char_Width = document.getElementById("char-Width")
	char_Width.value = char.Width
	char_Width.oninput = () => {
		characters[id].Width = char_Width.value
	}
	char_Height = document.getElementById("char-Height")
	char_Height.value = char.Height
	char_Height.oninput = () => {
		characters[id].Height = char_Height.value
	}
	char_CollisionRange = document.getElementById("char-CollisionRange")
	char_CollisionRange.value = char.CollisionRange
	char_CollisionRange.oninput = () => {
		characters[id].CollisionRange = char_CollisionRange.value
	}
	char_Birth = document.getElementById("char-Birth")
	char_Birth.value = char.Birth
	char_Birth.oninput = () => {
		characters[id].Birth = char_Birth.value
	}
	char_Dead = document.getElementById("char-Dead")
	char_Dead.value = char.Dead
	char_Dead.oninput = () => {
		characters[id].Dead = char_Dead.value
	}
	char_BirthEffect = document.getElementById("char-BirthEffect")
	char_BirthEffect.value = char.BirthEffect
	char_BirthEffect.oninput = () => {
		characters[id].BirthEffect = char_BirthEffect.value
	}
	char_DeathEffect = document.getElementById("char-DeathEffect")
	char_DeathEffect.value = char.DeathEffect
	char_DeathEffect.oninput = () => {
		characters[id].DeathEffect = char_DeathEffect.value
	}
	char_HibernateAction = document.getElementById("char-HibernateAction")
	char_HibernateAction.value = char.HibernateAction
	char_HibernateAction.oninput = () => {
		characters[id].HibernateAction = char_HibernateAction.value
	}
	char_DeathInstantAction = document.getElementById("char-DeathInstantAction")
	char_DeathInstantAction.value = char.DeathInstantAction
	char_DeathInstantAction.oninput = () => {
		characters[id].DeathInstantAction = char_DeathInstantAction.value
	}
	char_RemainingHPEffectDisplay = document.getElementById("char-RemainingHPEffectDisplay")
	char_RemainingHPEffectDisplay.value = char.RemainingHPEffectDisplay
	char_RemainingHPEffectDisplay.oninput = () => {
		characters[id].RemainingHPEffectDisplay = char_RemainingHPEffectDisplay.value
	}
	char_AttackCanBeSwerve = document.getElementById("char-AttackCanBeSwerve")
	char_AttackCanBeSwerve.value = char.AttackCanBeSwerve
	char_AttackCanBeSwerve.oninput = () => {
		characters[id].AttackCanBeSwerve = char_AttackCanBeSwerve.value
	}
	char_ConfirmToUseBlownAway = document.getElementById("char-ConfirmToUseBlownAway")
	char_ConfirmToUseBlownAway.value = char.ConfirmToUseBlownAway
	char_ConfirmToUseBlownAway.oninput = () => {
		characters[id].ConfirmToUseBlownAway = char_ConfirmToUseBlownAway.value
	}
	char_Script = document.getElementById("char-Script")
	char_Script.value = char.Script
	char_Script.oninput = () => {
		characters[id].Script = char_Script.value
	}
	char_WeaponUsed = document.getElementById("char-WeaponUsed")
	char_WeaponUsed.value = char.WeaponUsed
	char_WeaponUsed.oninput = () => {
		characters[id].WeaponUsed = char_WeaponUsed.value
	}
	char_SkillID = document.getElementById("char-SkillID")
	char_SkillID.value = char.SkillID
	char_SkillID.oninput = () => {
		characters[id].SkillID = char_SkillID.value
	}
	char_MonsterSkillRate = document.getElementById("char-MonsterSkillRate")
	char_MonsterSkillRate.value = char.MonsterSkillRate
	char_MonsterSkillRate.oninput = () => {
		characters[id].MonsterSkillRate = char_MonsterSkillRate.value
	}
	char_ItemIDToDrop = document.getElementById("char-ItemIDToDrop")
	char_ItemIDToDrop.value = char.ItemIDToDrop
	char_ItemIDToDrop.oninput = () => {
		characters[id].ItemIDToDrop = char_ItemIDToDrop.value
	}
	char_ItemDroprate = document.getElementById("char-ItemDroprate")
	char_ItemDroprate.value = char.ItemDroprate
	char_ItemDroprate.oninput = () => {
		characters[id].ItemDroprate = char_ItemDroprate.value
	}
	char_QuantityLimit = document.getElementById("char-QuantityLimit")
	char_QuantityLimit.value = char.QuantityLimit
	char_QuantityLimit.oninput = () => {
		characters[id].QuantityLimit = char_QuantityLimit.value
	}
	char_FatalityRate = document.getElementById("char-FatalityRate")
	char_FatalityRate.value = char.FatalityRate
	char_FatalityRate.oninput = () => {
		characters[id].FatalityRate = char_FatalityRate.value
	}
	char_PrefixLevel = document.getElementById("char-PrefixLevel")
	char_PrefixLevel.value = char.PrefixLevel
	char_PrefixLevel.oninput = () => {
		characters[id].PrefixLevel = char_PrefixLevel.value
	}
	char_QuestDropID = document.getElementById("char-QuestDropID")
	char_QuestDropID.value = char.QuestDropID
	char_QuestDropID.oninput = () => {
		characters[id].QuestDropID = char_QuestDropID.value
	}
	char_Rate = document.getElementById("char-Rate")
	char_Rate.value = char.Rate
	char_Rate.oninput = () => {
		characters[id].Rate = char_Rate.value
	}
	char_AI = document.getElementById("char-AI")
	char_AI.value = char.AI
	char_AI.oninput = () => {
		characters[id].AI = char_AI.value
	}
	char_Turning = document.getElementById("char-Turning")
	char_Turning.value = char.Turning
	char_Turning.oninput = () => {
		characters[id].Turning = char_Turning.value
	}
	char_Vision = document.getElementById("char-Vision")
	char_Vision.value = char.Vision
	char_Vision.oninput = () => {
		characters[id].Vision = char_Vision.value
	}
	char_Noise = document.getElementById("char-Noise")
	char_Noise.value = char.Noise
	char_Noise.oninput = () => {
		characters[id].Noise = char_Noise.value
	}
	char_GetExp = document.getElementById("char-GetExp")
	char_GetExp.value = char.GetExp
	char_GetExp.oninput = () => {
		characters[id].GetExp = char_GetExp.value
	}
	char_Light = document.getElementById("char-Light")
	char_Light.value = char.Light
	char_Light.oninput = () => {
		characters[id].Light = char_Light.value
	}
	char_MobExp = document.getElementById("char-MobExp")
	char_MobExp.value = char.MobExp
	char_MobExp.oninput = () => {
		characters[id].MobExp = char_MobExp.value
	}
	char_MonsterLVL = document.getElementById("char-MonsterLVL")
	char_MonsterLVL.value = char.MonsterLVL
	char_MonsterLVL.oninput = () => {
		characters[id].MonsterLVL = char_MonsterLVL.value
	}
	char_MaxHP = document.getElementById("char-MaxHP")
	char_MaxHP.value = char.MaxHP
	char_MaxHP.oninput = () => {
		characters[id].MaxHP = char_MaxHP.value
	}
	char_MinHP = document.getElementById("char-MinHP")
	char_MinHP.value = char.MinHP
	char_MinHP.oninput = () => {
		characters[id].MinHP = char_MinHP.value
	}
	char_MaxSP = document.getElementById("char-MaxSP")
	char_MaxSP.value = char.MaxSP
	char_MaxSP.oninput = () => {
		characters[id].MaxSP = char_MaxSP.value
	}
	char_MinSp = document.getElementById("char-MinSp")
	char_MinSp.value = char.MinSp
	char_MinSp.oninput = () => {
		characters[id].MinSp = char_MinSp.value
	}
	char_MinAtk = document.getElementById("char-MinAtk")
	char_MinAtk.value = char.MinAtk
	char_MinAtk.oninput = () => {
		characters[id].MinAtk = char_MinAtk.value
	}
	char_MaxAtk = document.getElementById("char-MaxAtk")
	char_MaxAtk.value = char.MaxAtk
	char_MaxAtk.oninput = () => {
		characters[id].MaxAtk = char_MaxAtk.value
	}
	char_PhysicalResist = document.getElementById("char-PhysicalResist")
	char_PhysicalResist.value = char.PhysicalResist
	char_PhysicalResist.oninput = () => {
		characters[id].PhysicalResist = char_PhysicalResist.value
	}
	char_Defense = document.getElementById("char-Defense")
	char_Defense.value = char.Defense
	char_Defense.oninput = () => {
		characters[id].Defense = char_Defense.value
	}
	char_HitRate = document.getElementById("char-HitRate")
	char_HitRate.value = char.HitRate
	char_HitRate.oninput = () => {
		characters[id].HitRate = char_HitRate.value
	}
	char_Flee = document.getElementById("char-Flee")
	char_Flee.value = char.Flee
	char_Flee.oninput = () => {
		characters[id].Flee = char_Flee.value
	}
	char_MonsterCriticalRate = document.getElementById("char-MonsterCriticalRate")
	char_MonsterCriticalRate.value = char.MonsterCriticalRate
	char_MonsterCriticalRate.oninput = () => {
		characters[id].MonsterCriticalRate = char_MonsterCriticalRate.value
	}
	char_MF = document.getElementById("char-MF")
	char_MF.value = char.MF
	char_MF.oninput = () => {
		characters[id].MF = char_MF.value
	}
	char_HREC = document.getElementById("char-HREC")
	char_HREC.value = char.HREC
	char_HREC.oninput = () => {
		characters[id].HREC = char_HREC.value
	}
	char_SREC = document.getElementById("char-SREC")
	char_SREC.value = char.SREC
	char_SREC.oninput = () => {
		characters[id].SREC = char_SREC.value
	}
	char_ASPDofMonster = document.getElementById("char-ASPDofMonster")
	char_ASPDofMonster.value = char.ASPDofMonster
	char_ASPDofMonster.oninput = () => {
		characters[id].ASPDofMonster = char_ASPDofMonster.value
	}
	char_AreaOfDetection = document.getElementById("char-AreaOfDetection")
	char_AreaOfDetection.value = char.AreaOfDetection
	char_AreaOfDetection.oninput = () => {
		characters[id].AreaOfDetection = char_AreaOfDetection.value
	}
	char_AfterDetectionChaseArea = document.getElementById("char-AfterDetectionChaseArea")
	char_AfterDetectionChaseArea.value = char.AfterDetectionChaseArea
	char_AfterDetectionChaseArea.oninput = () => {
		characters[id].AfterDetectionChaseArea = char_AfterDetectionChaseArea.value
	}
	char_MSPD = document.getElementById("char-MSPD")
	char_MSPD.value = char.MSPD
	char_MSPD.oninput = () => {
		characters[id].MSPD = char_MSPD.value
	}
	char_Col = document.getElementById("char-Col")
	char_Col.value = char.Col
	char_Col.oninput = () => {
		characters[id].Col = char_Col.value
	}
	char_Str = document.getElementById("char-Str")
	char_Str.value = char.Str
	char_Str.oninput = () => {
		characters[id].Str = char_Str.value
	}
	char_Agi = document.getElementById("char-Agi")
	char_Agi.value = char.Agi
	char_Agi.oninput = () => {
		characters[id].Agi = char_Agi.value
	}
	char_Dex = document.getElementById("char-Dex")
	char_Dex.value = char.Dex
	char_Dex.oninput = () => {
		characters[id].Dex = char_Dex.value
	}
	char_Con = document.getElementById("char-Con")
	char_Con.value = char.Con
	char_Con.oninput = () => {
		characters[id].Con = char_Con.value
	}
	char_Sta = document.getElementById("char-Sta")
	char_Sta.value = char.Sta
	char_Sta.oninput = () => {
		characters[id].Sta = char_Sta.value
	}
	char_Luk = document.getElementById("char-Luk")
	char_Luk.value = char.Luk
	char_Luk.oninput = () => {
		characters[id].Luk = char_Luk.value
	}
	char_LeftRad = document.getElementById("char-LeftRad")
	char_LeftRad.value = char.LeftRad
	char_LeftRad.oninput = () => {
		characters[id].LeftRad = char_LeftRad.value
	}
	char_Guild = document.getElementById("char-Guild")
	char_Guild.value = char.Guild
	char_Guild.oninput = () => {
		characters[id].Guild = char_Guild.value
	}
	char_Title = document.getElementById("char-Title")
	char_Title.value = char.Title
	char_Title.oninput = () => {
		characters[id].Title = char_Title.value
	}
	char_Job = document.getElementById("char-Job")
	char_Job.value = char.Job
	char_Job.oninput = () => {
		characters[id].Job = char_Job.value
	}
	char_ExpGainedFromKill = document.getElementById("char-ExpGainedFromKill")
	char_ExpGainedFromKill.value = char.ExpGainedFromKill
	char_ExpGainedFromKill.oninput = () => {
		characters[id].ExpGainedFromKill = char_ExpGainedFromKill.value
	}
	char_Nexp = document.getElementById("char-Nexp")
	char_Nexp.value = char.Nexp
	char_Nexp.oninput = () => {
		characters[id].Nexp = char_Nexp.value
	}
	char_Fame = document.getElementById("char-Fame")
	char_Fame.value = char.Fame
	char_Fame.oninput = () => {
		characters[id].Fame = char_Fame.value
	}
	char_Ap = document.getElementById("char-Ap")
	char_Ap.value = char.Ap
	char_Ap.oninput = () => {
		characters[id].Ap = char_Ap.value
	}
	char_Tp = document.getElementById("char-Tp")
	char_Tp.value = char.Tp
	char_Tp.oninput = () => {
		characters[id].Tp = char_Tp.value
	}
	char_Gd = document.getElementById("char-Gd")
	char_Gd.value = char.Gd
	char_Gd.oninput = () => {
		characters[id].Gd = char_Gd.value
	}
	char_Spri = document.getElementById("char-Spri")
	char_Spri.value = char.Spri
	char_Spri.oninput = () => {
		characters[id].Spri = char_Spri.value
	}
	char_Stor = document.getElementById("char-Stor")
	char_Stor.value = char.Stor
	char_Stor.oninput = () => {
		characters[id].Stor = char_Stor.value
	}
	char_MxSail = document.getElementById("char-MxSail")
	char_MxSail.value = char.MxSail
	char_MxSail.oninput = () => {
		characters[id].MxSail = char_MxSail.value
	}
	char_Sail = document.getElementById("char-Sail")
	char_Sail.value = char.Sail
	char_Sail.oninput = () => {
		characters[id].Sail = char_Sail.value
	}
	char_Stasa = document.getElementById("char-Stasa")
	char_Stasa.value = char.Stasa
	char_Stasa.oninput = () => {
		characters[id].Stasa = char_Stasa.value
	}
	char_Scsm = document.getElementById("char-Scsm")
	char_Scsm.value = char.Scsm
	char_Scsm.oninput = () => {
		characters[id].Scsm = char_Scsm.value
	}
	char_Tstr = document.getElementById("char-Tstr")
	char_Tstr.value = char.Tstr
	char_Tstr.oninput = () => {
		characters[id].Tstr = char_Tstr.value
	}
	char_Tagi = document.getElementById("char-Tagi")
	char_Tagi.value = char.Tagi
	char_Tagi.oninput = () => {
		characters[id].Tagi = char_Tagi.value
	}
	char_Tdex = document.getElementById("char-Tdex")
	char_Tdex.value = char.Tdex
	char_Tdex.oninput = () => {
		characters[id].Tdex = char_Tdex.value
	}
	char_Tcon = document.getElementById("char-Tcon")
	char_Tcon.value = char.Tcon
	char_Tcon.oninput = () => {
		characters[id].Tcon = char_Tcon.value
	}
	char_Tsta = document.getElementById("char-Tsta")
	char_Tsta.value = char.Tsta
	char_Tsta.oninput = () => {
		characters[id].Tsta = char_Tsta.value
	}
	char_Tluk = document.getElementById("char-Tluk")
	char_Tluk.value = char.Tluk
	char_Tluk.oninput = () => {
		characters[id].Tluk = char_Tluk.value
	}
	char_Tmxhp = document.getElementById("char-Tmxhp")
	char_Tmxhp.value = char.Tmxhp
	char_Tmxhp.oninput = () => {
		characters[id].Tmxhp = char_Tmxhp.value
	}
	char_Tmxsp = document.getElementById("char-Tmxsp")
	char_Tmxsp.value = char.Tmxsp
	char_Tmxsp.oninput = () => {
		characters[id].Tmxsp = char_Tmxsp.value
	}
	char_Tatk = document.getElementById("char-Tatk")
	char_Tatk.value = char.Tatk
	char_Tatk.oninput = () => {
		characters[id].Tatk = char_Tatk.value
	}
	char_Tdef = document.getElementById("char-Tdef")
	char_Tdef.value = char.Tdef
	char_Tdef.oninput = () => {
		characters[id].Tdef = char_Tdef.value
	}
	char_Thit = document.getElementById("char-Thit")
	char_Thit.value = char.Thit
	char_Thit.oninput = () => {
		characters[id].Thit = char_Thit.value
	}
	char_Tflee = document.getElementById("char-Tflee")
	char_Tflee.value = char.Tflee
	char_Tflee.oninput = () => {
		characters[id].Tflee = char_Tflee.value
	}
	char_Tmf = document.getElementById("char-Tmf")
	char_Tmf.value = char.Tmf
	char_Tmf.oninput = () => {
		characters[id].Tmf = char_Tmf.value
	}
	char_Tcrt = document.getElementById("char-Tcrt")
	char_Tcrt.value = char.Tcrt
	char_Tcrt.oninput = () => {
		characters[id].Tcrt = char_Tcrt.value
	}
	char_Threc = document.getElementById("char-Threc")
	char_Threc.value = char.Threc
	char_Threc.oninput = () => {
		characters[id].Threc = char_Threc.value
	}
	char_Tsrec = document.getElementById("char-Tsrec")
	char_Tsrec.value = char.Tsrec
	char_Tsrec.oninput = () => {
		characters[id].Tsrec = char_Tsrec.value
	}
	char_Taspd = document.getElementById("char-Taspd")
	char_Taspd.value = char.Taspd
	char_Taspd.oninput = () => {
		characters[id].Taspd = char_Taspd.value
	}
	char_Tadis = document.getElementById("char-Tadis")
	char_Tadis.value = char.Tadis
	char_Tadis.oninput = () => {
		characters[id].Tadis = char_Tadis.value
	}
	char_Tspd = document.getElementById("char-Tspd")
	char_Tspd.value = char.Tspd
	char_Tspd.oninput = () => {
		characters[id].Tspd = char_Tspd.value
	}
	char_Tspri = document.getElementById("char-Tspri")
	char_Tspri.value = char.Tspri
	char_Tspri.oninput = () => {
		characters[id].Tspri = char_Tspri.value
	}
	char_Tscsm = document.getElementById("char-Tscsm")
	char_Tscsm.value = char.Tscsm
	char_Tscsm.oninput = () => {
		characters[id].Tscsm = char_Tscsm.value
	}
}

async function update_char_list() {
	let list = document.getElementById("char-list")
	list.innerHTML = ""
	for(i = 0;i < Math.min(400, characters.length);i++) {
		let char = characters[i]
		if(char.deleted) {
			continue
		}
		
		list.insertAdjacentElement('beforeend', htmlToElement(make_char_node(char)))
	}
}

async function update_char_with_search(search) {
	let inserted = 0
	let list = document.getElementById("char-list")
	list.innerHTML = ""
	for(i = 0;i < characters.length && inserted < 200;i++) {
		if(characters[i].deleted) {
			continue
		}
		let name = `${characters[i].ID} ${characters[i].Description}`
		if(name.includes(search)) {		
			inserted++
			list.insertAdjacentElement('beforeend', htmlToElement(make_char_node(characters[i])))
		}
	}
}

async function init_char() {
	document.getElementById("char-search-bar").addEventListener('input', () => {
		search = document.getElementById("char-search-bar").value 
		if(search == "") {
			update_char_list()
		} else {
			update_char_with_search(search)
		}
	})

	document.getElementById("save-all-char").addEventListener('click', () => {
		data = []
		for(i = 0;i < characters.length;i++) {
			if(!characters[i].deleted) {
				data.push(characters[i])
			}
		}
		console.log(data)
		const resp = fetch("/characters", {
			method: "POST",
			body: JSON.stringify(data)
		})
		console.log(resp)
	})

	document.getElementById("add-char").addEventListener('click', () => {
		let char = {}
		for(key in characters[0]) {
			char[key] = "0"
		}
		char.uniq_id = glob_char_id++
		char.deleted = false
		char.Description = "New Monster"
		characters.push(char)
		document.getElementById("char-list").insertAdjacentElement('afterbegin', htmlToElement(make_char_node(char)))
		pick_char_node(char.uniq_id)
	})

	document.getElementById("remove-char").addEventListener('click', (event) => {
		let id = event.target.dataset.id
		console.log(id)
		characters[id].deleted = true

		event.target.style.display = "none"
		document.querySelector("#char-red-container .redactor").style.display = "none"
	
		let search = document.getElementById("char-search-bar").value
		if(search == "") {
			update_char_list()
		} else {
			update_char_with_search(search)
		}

		show_char_delete_cancle(id)
	})

	document.querySelector("#cancle-char-action div").onclick = cancle_char_action

    await request_characters()
    update_char_list()
}


let char_cancle_id = -1
function cancle_char_action() {
	characters[char_cancle_id].deleted = false
	let search = document.getElementById("char-search-bar").value
	if(search == "") {
		update_char_list()
	} else {
		update_char_with_search(search)
	}
	document.getElementById("cancle-char-action").style.display = "none"
}

let char_father_id = 0
async function show_char_delete_cancle(id) {
	let cancle = document.getElementById("cancle-char-action")
	let father = ++char_father_id
	cancle.style.display = "flex"
	char_cancle_id = id
	await sleep(10 * 1000)
	if(father == char_father_id) {
		cancle.style.display = "none"
	}
}