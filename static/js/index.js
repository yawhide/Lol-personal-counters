var storage;
try {
  storage = JSON.parse(localStorage['lol-personal-counter']);
} catch (e) {
  storage = {};
}
var $form = document.querySelector('form');
var $regionBtn = document.querySelector('.btn.btn-default');
var $regionList = document.querySelector('.dropdown-menu');

if (storage.summonerName) {
  $form.name.value = storage.summonerName;
  $regionBtn.innerText = storage.region || 'NA';
  $form.region.value = $regionBtn.innerText;
  $form.role.value = storage.role || 'Top';
  $form.save.checked = true;
  $form.enemy.focus();
  console.info('Username saved:', storage.summonerName);
}

$form.onsubmit = function (e) {
  if ($form.save.checked) {
    storage.summonerName = $form.name.value;
    storage.region = $regionBtn.innerText;
    storage.role = $form.role.value;
  } else if (localStorage['lol-personal-counter']) {
    delete storage.summonerName;
    delete storage.region;
    delete storage.role;
  }
  localStorage['lol-personal-counter'] = JSON.stringify(storage);
}

$regionBtn.onclick = function () {
  $regionList.style.display = !$regionList.style.display || $regionList.style.display === 'none' ? 'block' : 'none';
}

$regionBtn.onblur = function () {
  $regionList.style.display = 'none';
}

$regionList.onclick = function (e) {
  $form.region.value = e.target.innerText;
  $regionBtn.innerText = e.target.innerText;
  $regionList.style.display = 'none';
}
