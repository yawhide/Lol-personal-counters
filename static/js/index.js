var storage;
try {
  storage = JSON.parse(localStorage['lol-personal-counter']);
} catch (e) {
  storage = {};
}
var $form = document.querySelector('form');
var $regionBtn = document.querySelector('.btn.btn-default');
var $regionList = document.querySelector('.dropdown-menu');

// analytics
var $role = document.querySelector('.role');
var $summonerName = document.querySelector('.summonerName');
var $championGGLinks = document.querySelectorAll('table tbody tr th a');
var $yawhideLink = document.querySelector('p a');

if ($form) {
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
      storage.region = $regionBtn.innerText.trim();
      storage.role = $form.role.value;
    } else if (localStorage['lol-personal-counter']) {
      delete storage.summonerName;
      delete storage.region;
      delete storage.role;
    }
    localStorage['lol-personal-counter'] = JSON.stringify(storage);
    fetch(urlPrefix + 'analytics/index', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        SummonerName: $form.name.value,
        Enemy       : $form.enemy.value,
        Role        : $form.role.value,
        RememberMe  : $form.save.checked,
        Referral    : document.referrer,
      })
    });
  }

  $regionBtn.onclick = function () {
    $regionList.style.display = !$regionList.style.display || $regionList.style.display === 'none' ? 'block' : 'none';
  }

  $regionList.onmousedown = function (e) {
    $form.region.value = e.target.innerText;
    $regionBtn.innerText = e.target.innerText;
    $regionList.style.display = 'none';
  }

  $regionList.onblur = function () {
    $regionList.style.display = 'none';
  }

  $regionBtn.onblur = function (e) {
    $regionList.style.display = 'none';
  }
}

for (var i = 0; i < $championGGLinks.length; i++) {
  $championGGLinks[i].onclick = function (e) {
    fetch(urlPrefix + 'analytics/matchup', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        SummonerName: $summonerName.id,
        Enemy       : e.target.innerText,
        Role        : $role.id,
        Click       : e.target.href,
      })
    });
  }
}

if ($yawhideLink) {
  $yawhideLink.onclick = function (e) {
    fetch(urlPrefix + 'analytics/external', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        Url : e.target.href,
        Page: window.location.pathname,
      })
    });
  }
}
