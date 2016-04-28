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

/* autocomplete */
var champions = ["Aatrox","Ahri","Akali","Alistar","Amumu","Anivia","Annie","Ashe","Aurelionsol","Azir","Bard","Blitzcrank","Brand","Braum","Caitlyn","Cassiopeia","Chogath","Corki","Darius","Diana","Draven","Drmundo","Ekko","Elise","Evelynn","Ezreal","Fiddlesticks","Fiora","Fizz","Galio","Gangplank","Garen","Gnar","Gragas","Graves","Hecarim","Heimerdinger","Illaoi","Irelia","Janna","Jarvaniv","Jax","Jayce","Jhin","Jinx","Kalista","Karma","Karthus","Kassadin","Katarina","Kayle","Kennen","Khazix","Kindred","Kogmaw","Leblanc","Leesin","Leona","Lissandra","Lucian","Lulu","Lux","Malphite","Malzahar","Maokai","Master yi","Missfortune","Monkeyking","Mordekaiser","Morgana","Nami","Nasus","Nautilus","Nidalee","Nocturne","Nunu","Olaf","Orianna","Pantheon","Poppy","Quinn","Rammus","Reksai","Renekton","Rengar","Riven","Rumble","Ryze","Sejuani","Shaco","Shen","Shyvana","Singed","Sion","Sivir","Skarner","Sona","Soraka","Swain","Syndra","Tahmkench","Talon","Taric","Teemo","Thresh","Tristana","Trundle","Tryndamere","Twistedfate","Twitch","Udyr","Urgot","Varus","Vayne","Veigar","Velkoz","Vi","Viktor","Vladimir","Volibear","Warwick","Wukong","Xerath","Xinzhao","Yasuo","Yorick","Zac","Zed","Ziggs","Zilean","Zyra"];

