document.getElementById('getOpinionButton').addEventListener('click', async () => {
    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });
    const videoLink = encodeURIComponent(tab.url);
    const apiUrl = `http://localhost:8080/get_opinion?link=${videoLink}`;
    console.log(videoLink)

  function typeWriterEffect(text, speed) {
    console.log("test2")
    const opinionResult = document.getElementById('opinionResult');
    let i = 0;
    const typingInterval = setInterval(function () {
      console.log("test")
      if (i < text.length) {
        opinionResult.textContent += text.charAt(i);
        i++;
      } else {
        clearInterval(typingInterval);
      }
    }, speed);
  }

    try {
        const response = await fetch(apiUrl, {
            mode: 'cors'
        });


        if (response.ok) {
            const data = await response.json();
            const text = `Opinion: ${data.opinion}`;
            typeWriterEffect(text, 35);
          } else {
          throw new Error('Failed to fetch: ' + response.statusText);
        }}
    catch (error) {
        const text = 'Error fetching opinion.';
          typeWriterEffect(text, 100);
          console.error('Error fetching opinion:', error);
        }
});

