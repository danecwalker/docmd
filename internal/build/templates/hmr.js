(function() {
  let eventSource = new EventSource('/__hmr');

  eventSource.onmessage = function(event) {
    let data = JSON.parse(event.data);

    if (data.type === 'reload') {
      window.location.reload();
    } else if (data.type === 'update') {
      let url = data.url;
      let module = data.module;

      fetch(url)
        .then(response => response.text())
        .then(code => {
          let script = document.createElement('script');
          script.type = 'module';
          script.text = code;
          document.head.appendChild(script);

          let oldModule = window[module];
          let newModule = window[module];

          if (oldModule && oldModule.__dispose) {
            oldModule.__dispose();
          }

          if (newModule && newModule.__init) {
            newModule.__init();
          }
        });
    }
  }
})()