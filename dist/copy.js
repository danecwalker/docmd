// find all pre tags
document.onload = (function() {
  var pre = document.querySelectorAll('pre.chroma');
  // loop through all pre tags
  for (var i = 0; i < pre.length; i++) {
    console.log('pre', pre[i]);
    let thispre = pre[i];
    // create a new div
    var btn = document.createElement('button');
    btn.classList.add('copy');
    btn.innerHTML = '📋';

    btn.addEventListener('click', function() {
      var text = thispre.firstChild.textContent;
      navigator.clipboard.writeText(text).then(function() {
        console.log('Copied to clipboard');
      }, function(err) {
        console.error('Failed to copy to clipboard', err);
      });
    });

    // append the button to the pre tag
    pre[i].appendChild(btn);
  }
})()