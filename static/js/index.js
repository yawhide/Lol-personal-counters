var storage;
try {
  storage = JSON.parse(localStorage['lol-personal-counter']);
} catch (e) {
  storage = {};
}
var $form = document.querySelector('form');
// var $saveUsername = document.querySelector('input[type="checkbox"]');

if (storage.summonerName) {
  // $saveUsername.value = 'checked'
  $form.name.value = storage.summonerName;
  $form.save.checked = true;
  $form.enemy.focus();
  console.info('Username saved:', storage.summonerName);
}

$form.onsubmit = function (e) {
  if ($form.save.checked) {
    storage.summonerName = $form.name.value;
  } else if (localStorage['lol-personal-counter']) {
    delete storage.summonerName;
  }
  localStorage['lol-personal-counter'] = JSON.stringify(storage);
}
